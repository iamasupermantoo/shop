package shopsService

import (
	"errors"
	"github.com/brianvoe/gofakeit/v6"
	"go.uber.org/zap"
	"gofiber/app/models/model/productsModel"
	"gofiber/app/models/model/shopsModel"
	"gofiber/app/models/model/systemsModel"
	"gofiber/app/models/model/usersModel"
	"gofiber/app/models/service/productsService"
	"gofiber/app/module/database"
	"gofiber/utils"
	"gorm.io/gorm"
	"strconv"
	"strings"
	"time"
)

type StoresService struct {
	adminId uint
	tx      *gorm.DB
}

// NewStoresService 商品服务模型
func NewStoresService(tx *gorm.DB, adminId uint) *StoresService {
	return &StoresService{tx: tx, adminId: adminId}
}

// InitStore 初始化店铺
func (_StoresService *StoresService) InitStore() error {
	if _StoresService.tx == nil {
		_StoresService.tx = database.Db
	}
	adminId := _StoresService.adminId
	productList := make([]*productsModel.Product, 0)
	err := database.Db.Where("admin_id = ?", adminId).
		Where("status = ?", productsModel.ProductStatusActive).
		Where("type = ?", productsModel.ProductTypeWholesale).
		Find(&productList).Error
	if err != nil || len(productList) == 0 {
		return errors.New("not find product")
	}

	productMapByCategory := make(map[uint][]*productsModel.Product)
	for _, product := range productList {
		if productMapByCategory[product.CategoryId] == nil {
			productMapByCategory[product.CategoryId] = make([]*productsModel.Product, 0)
		}
		productMapByCategory[product.CategoryId] = append(productMapByCategory[product.CategoryId], product)
	}

	for i := 0; i < 9; i++ {
		err = _StoresService.tx.Transaction(func(tx *gorm.DB) error {
			userInfo := usersModel.User{
				AdminId:     adminId,
				ParentId:    uint(gofakeit.Number(2, 5)),
				UserName:    strings.ReplaceAll(gofakeit.Name(), " ", ""),
				NickName:    strings.ReplaceAll(gofakeit.Name(), " ", ""),
				Email:       gofakeit.Email(),
				Telephone:   gofakeit.Phone(),
				Avatar:      "/assets/store/" + strconv.Itoa(i+1) + ".jpeg",
				Score:       gofakeit.Number(80, 100),
				Sex:         gofakeit.RandomInt([]int{1, 2}),
				Password:    utils.PasswordEncrypt("abc123"),
				SecurityKey: utils.PasswordEncrypt("abc123"),
				Money:       100000,
			}
			if err = tx.Create(&userInfo).Error; err != nil {
				return err
			}
			if err = tx.Create(&shopsModel.StoreSettled{
				AdminId:  userInfo.AdminId,
				UserId:   userInfo.ID,
				Name:     strings.ReplaceAll(gofakeit.Name(), " ", ""),
				Address:  gofakeit.Address().Address,
				RealName: strings.ReplaceAll(gofakeit.Name(), " ", ""),
				Logo:     "/assets/store/" + strconv.Itoa(i+1) + ".jpeg",
				Photo1:   "/assets/store/business_license_front.png",
				Photo2:   "/assets/store/business_license_reverse.png",
				Number:   gofakeit.Password(false, false, true, false, false, 20),
				Email:    gofakeit.Email(),
				Contact:  gofakeit.Phone(),
				Status:   shopsModel.StoreSettledStatusPass,
			}).Error; err != nil {
				return err
			}

			// 购买第一个会员
			currentLevelInfo := &systemsModel.Level{}
			tx.Model(currentLevelInfo).Where("admin_id = ?", adminId).Where("status = ?", systemsModel.LevelStatusActive).
				Order("symbol ASC").Limit(1).Find(currentLevelInfo)
			if currentLevelInfo.ID > 0 {
				expireTime := time.Now()
				if currentLevelInfo.Days == -1 {
					expireTime = expireTime.Add(365 * 24 * time.Hour)
				} else {
					expireTime = expireTime.Add(time.Duration(currentLevelInfo.Days) * 24 * time.Hour)
				}
				tx.Create(&usersModel.UserLevel{
					AdminId:   userInfo.AdminId,
					UserId:    userInfo.ID,
					Name:      currentLevelInfo.Name,
					Icon:      currentLevelInfo.Icon,
					Symbol:    currentLevelInfo.Symbol,
					Money:     currentLevelInfo.Money,
					ExpiredAt: expireTime.Add(time.Duration(currentLevelInfo.Days) * 24 * time.Hour),
					Increase:  currentLevelInfo.Increase,
				})
			}

			storeInfo := shopsModel.Store{
				AdminId:  userInfo.AdminId,
				UserId:   userInfo.ID,
				Logo:     "/assets/store/" + strconv.Itoa(i+1) + ".jpeg",
				Name:     gofakeit.Name(),
				Contact:  gofakeit.Phone(),
				Keywords: gofakeit.Password(true, true, true, false, false, 15),
				Address:  gofakeit.Address().Address,
				Desc:     gofakeit.HackerPhrase(),
				Rating:   utils.FloatAccuracy(gofakeit.Float64Range(3.0, 5), 2),
				Sales:    gofakeit.Number(300, 8888),
				Score:    gofakeit.Number(50, 100),
			}
			if err = tx.Create(&storeInfo).Error; err != nil {
				return err
			}

			createProductList := make([]productsModel.Product, 0)
			for _, v := range productMapByCategory {
				numb := 5
				index := i + 1*numb
				switch {

				case index < len(v): // 如果该分类的数据小于第i页第numb条，则顺序获取numb条数据
					for j := index - numb; j < index; j++ {
						createProductList = append(createProductList, *v[j])
					}
				default: // 随机获取numb条数据
					indexes := make([]int, 0)
					if len(v) < numb {
						for x, _ := range v {
							indexes = append(indexes, x)
						}
					} else {
						indexes = utils.NewRandom().IntArray(numb, 0, len(v)-1)
					}

					for j := 0; j < len(indexes); j++ {
						index = indexes[j]
						createProductList = append(createProductList, *v[index])
					}
				}
			}

			for _, v := range createProductList {
				currentId := v.ID
				v.ID = 0
				v.Type = productsModel.ProductTypeDefault
				v.Status = productsModel.ProductStatusActive
				v.StoreId = storeInfo.ID
				v.ParentId = currentId

				if err = tx.Create(&v).Error; err != nil {
					return err
				}

				if err = productsService.NewProduct(tx).InserterProductAttrs(productsService.ProductAttrsSkuWholesale, currentId, v.ID); err != nil {
					zap.L().Error("initProduct", zap.Error(err), zap.Reflect("product", v))
					return err
				}
			}
			return nil
		})
		if err != nil {
			return err
		}
	}
	return nil
}

// DropStore 删除店铺信息
func (_StoresService *StoresService) DropStore() *StoresService {
	if _StoresService.tx == nil {
		_StoresService.tx = database.Db
	}
	_StoresService.tx.Unscoped().Where("admin_id = ?", _StoresService.adminId).Delete(&usersModel.User{})
	_StoresService.tx.Unscoped().Where("admin_id = ?", _StoresService.adminId).Delete(&shopsModel.StoreSettled{})
	_StoresService.tx.Unscoped().Where("admin_id = ?", _StoresService.adminId).Delete(&usersModel.UserLevel{})
	_StoresService.tx.Unscoped().Where("admin_id = ?", _StoresService.adminId).Delete(&shopsModel.Store{})
	_StoresService.tx.Unscoped().Where("admin_id = ?", _StoresService.adminId).
		Where("store_id > ?", 0).
		Delete(&productsModel.Product{})
	return _StoresService
}
