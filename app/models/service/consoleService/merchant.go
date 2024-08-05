package consoleService

import (
	"gofiber/app/models/model/adminsModel"
	"gofiber/app/models/model/productsModel"
	"gofiber/app/models/model/systemsModel"
	"gofiber/app/models/model/walletsModel"
	"gofiber/app/models/service/adminsService"
	"gofiber/app/models/service/productsService"
	"gofiber/app/models/service/systemsService"
	"gofiber/app/module/cache"
	"gofiber/app/module/database"
	"gofiber/utils"
	"gorm.io/gorm"
)

const (
	MerchantOperateDelete = -1
	MerchantOperateReset  = 1
	MerchantOperateSync   = 2
)

type Merchant struct {
	tables      []*merchantTable //	操作的表
	FilterTable []string         // 过滤的表名
	adminId     uint             //	当前操作管理
	operate     int              //	当前操作方法
}

func NewMerchant(adminId uint, filterTable []string) *Merchant {
	return &Merchant{
		FilterTable: filterTable,
		tables: []*merchantTable{
			{Name: "admin_setting", Model: &adminsModel.AdminSetting{}, FindFields: []string{"field"}},
			{Name: "wallet_assets", Model: &walletsModel.WalletAssets{}, FindFields: []string{"symbol", "type"}},
			{Name: "wallet_payment", Model: &walletsModel.WalletPayment{}, FindFields: []string{"name", "type", "mode"}, Association: []*Association{{ForeignKey: "assets_id", Model: &walletsModel.WalletAssets{}, FindFields: []string{"symbol", "type"}}}},
			{Name: "translate", Model: &systemsModel.Translate{}, FindFields: []string{"lang", "field"}},
			{Name: "lang", Model: &systemsModel.Lang{}, FindFields: []string{"alias", "symbol"}},
			{Name: "country", Model: &systemsModel.Country{}, FindFields: []string{"alias"}},
			{Name: "level", Model: &systemsModel.Level{}, FindFields: []string{"symbol"}},
			{Name: "article", Model: &systemsModel.Article{}, FindFields: []string{"name", "type"}},
			{Name: "menu", Model: &systemsModel.Menu{}, FindFields: []string{"route", "type", "name"}, Association: []*Association{{ForeignKey: "parent_id", Model: &systemsModel.Menu{}, FindFields: []string{"route", "type", "name"}}}},
			{Name: "category", Model: &productsModel.Category{}, FindFields: []string{"name", "type"}, Association: []*Association{{ForeignKey: "parent_id", Model: &productsModel.Category{}, FindFields: []string{"name", "type"}}}},
			{Name: "product", Model: &productsModel.Product{}, FindFields: []string{"name", "type"}, Association: []*Association{
				{ForeignKey: "category_id", Model: &productsModel.Category{}, FindFields: []string{"name", "type"}},
				{ForeignKey: "assets_id", Model: &walletsModel.WalletAssets{}, FindFields: []string{"name", "type"}},
			}, DoneFunc: func(tx *gorm.DB, oldId uint, m map[string]interface{}) error {
				var newId uint

				switch m["id"].(type) {
				case int64:
					newId = uint(m["id"].(int64))
				case int:
					newId = uint(m["id"].(int))
				}
				return productsService.NewProduct(nil).InserterProductAttrs(productsService.ProductAttrsSkuWholesale, oldId, newId)
			}},
		},
		adminId: adminId,
	}
}

// RunSync 同步设置
func (_Merchant *Merchant) RunSync() error {
	_Merchant.operate = MerchantOperateSync
	return _Merchant.run()
}

// RunRest 重置设置
func (_Merchant *Merchant) RunRest() error {
	_Merchant.operate = MerchantOperateReset
	err := _Merchant.run()
	if err != nil {
		return err
	}
	return nil
}

// Delete 删除商户设置
func (_Merchant *Merchant) Delete() error {
	_Merchant.operate = MerchantOperateDelete
	return _Merchant.run()
}

// run 执行重置同步删除的方法
func (_Merchant *Merchant) run() error {
	return database.Db.Transaction(func(tx *gorm.DB) error {
		for _, table := range _Merchant.tables {
			// 过滤需要过滤的数据
			if utils.ArrayStringIndexOf(_Merchant.FilterTable, table.Name) > -1 {
				continue
			}

			superData := make([]map[string]interface{}, 0)
			model := tx.Model(table.Model).Where("admin_id = ?", adminsModel.SuperAdminId)
			if table.Where != "" {
				model.Where(table.Where)
			}
			model.Find(&superData)

			switch _Merchant.operate {
			case MerchantOperateReset:
				// 重置数据
				tx.Unscoped().Where("admin_id = ?", _Merchant.adminId).Delete(table.Model)
				for _, datum := range superData {
					currentId := datum["id"].(uint)
					datum["id"] = 0
					datum["admin_id"] = _Merchant.adminId

					// 查找上级ID
					for _, association := range table.Association {
						datum[association.ForeignKey] = _Merchant.getParent(tx, association, datum[association.ForeignKey].(uint))
					}
					tx.Model(table.Model).Create(&datum)

					// 执行插入后回调方法
					if table.DoneFunc != nil {
						_ = table.DoneFunc(tx, currentId, datum)
					}
				}

			case MerchantOperateSync:
				// 同步操作
				merchantData := make([]map[string]interface{}, 0)
				tx.Model(table.Model).Where("admin_id = ?", _Merchant.adminId).Find(&merchantData)
				for _, datum := range superData {
					if !_Merchant.matchFields(merchantData, datum, table.FindFields) {
						currentId := datum["id"].(uint)
						datum["id"] = 0
						datum["admin_id"] = _Merchant.adminId

						// 查找上级ID
						for _, association := range table.Association {
							datum[association.ForeignKey] = _Merchant.getParent(tx, association, datum[association.ForeignKey].(uint))
						}
						tx.Model(table.Model).Create(&datum)

						// 执行插入后回调方法
						if table.DoneFunc != nil {
							_ = table.DoneFunc(tx, currentId, datum)
						}
					}
				}

			case MerchantOperateDelete:
				// 删除数据
				tx.Where("admin_id = ?", _Merchant.adminId).Delete(table.Model)
			}
		}
		// 删除缓存内容
		rdsConn := cache.Rds.Get()
		defer rdsConn.Close()

		langService := systemsService.NewSystemLang(rdsConn, _Merchant.adminId)
		countryService := systemsService.NewSystemCountry(rdsConn, _Merchant.adminId)
		translateService := systemsService.NewSystemTranslate(rdsConn, _Merchant.adminId)
		adminSettingService := adminsService.NewAdminSetting(rdsConn, _Merchant.adminId)
		adminMenuService := adminsService.NewAdminMenu(rdsConn, _Merchant.adminId)
		adminService := adminsService.NewAdminUser(rdsConn, _Merchant.adminId)
		userMenuService := systemsService.NewSystemMenu(rdsConn, _Merchant.adminId)

		// 删除商户语言缓存
		langService.DelRedisAdminLangList()
		// 删除商户国家缓存
		countryService.DelRedisAdminCountryList()
		// 删除商户配置缓存
		adminSettingService.DelRedisAdminSetting()
		// 删除商户管理菜单缓存
		adminMenuService.DelRedisAdminMenuList()
		// 删除商户下级管理缓存
		adminService.DelRedisChildrenIds()
		// 删除商户前台菜单
		userMenuService.DelRedisSystemMenuList()
		// 删除商户翻译包缓存
		translateService.DelRedisAdminTranslate()

		return nil
	})
}

// getParent 获取上级Id
func (_Merchant *Merchant) getParent(tx *gorm.DB, table *Association, parentId uint) interface{} {
	if parentId == 0 {
		return 0
	}

	superData := map[string]interface{}{}
	currentData := map[string]interface{}{}
	var model *gorm.DB
	tx.Model(table.Model).Where("id = ?", parentId).Find(&superData)
	model = tx.Model(table.Model).Where("admin_id = ?", _Merchant.adminId)

	for _, field := range table.FindFields {
		model.Where(field+" = ?", superData[field])
	}
	model.Find(&currentData)
	return currentData["id"]
}

// matchFields 对比新旧数据是否相同
func (_Merchant *Merchant) matchFields(merchantData []map[string]interface{}, datum map[string]interface{}, FindFields []string) bool {
	for _, merchantDatum := range merchantData {
		isExist := true
		for _, field := range FindFields {
			if merchantDatum[field] != datum[field] {
				isExist = false
				break
			}
		}
		if isExist {
			return isExist
		}
	}
	return false
}

// Association 关联模型数据
type Association struct {
	Model      interface{} // 关联模型
	ForeignKey string      // 外键
	FindFields []string    // 查找匹配字段
}

// merchantTable 同步删除重置操作模型数据
type merchantTable struct {
	Name        string                                                               // 表明
	Model       interface{}                                                          // 模型
	Where       string                                                               // 条件
	FindFields  []string                                                             // 查找匹配字段
	Association []*Association                                                       // 关联表
	DoneFunc    func(tx *gorm.DB, currentId uint, data map[string]interface{}) error // 执行完后回调
}
