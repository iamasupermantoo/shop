package productsService

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"gofiber/app/models/model/productsModel"
	"gofiber/app/models/service/adminsService"
	"gofiber/app/models/service/systemsService"
	"gofiber/app/module/database"
	"gofiber/app/module/views"
	"strconv"
)

type ProductCategory struct {
	rdsConn redis.Conn
	adminId uint
}

func NewProductCategory(rdsConn redis.Conn, adminId uint) *ProductCategory {
	return &ProductCategory{rdsConn: rdsConn, adminId: adminId}
}

// GetViewsOptions 获取视图Options
func (_ProductCategory *ProductCategory) GetViewsOptions() []*views.InputOptions {
	adminService := adminsService.NewAdminUser(_ProductCategory.rdsConn, _ProductCategory.adminId)
	translateService := systemsService.NewSystemTranslate(_ProductCategory.rdsConn, _ProductCategory.adminId)

	categoryList := make([]*productsModel.Category, 0)
	database.Db.Model(&productsModel.Category{}).
		Where("admin_id IN ?", adminService.GetRedisChildrenIds()).Where("status = ?", productsModel.CategoryStatusActive).Find(&categoryList)

	data := make([]*views.InputOptions, 0)
	for _, category := range categoryList {
		data = append(data, &views.InputOptions{
			Label: translateService.GetRedisAdminTranslateLangField("zh-CN", category.Name) + "." + strconv.Itoa(int(category.AdminId)),
			Value: category.ID,
		})
	}
	return data
}

// CategoryChildrenIds 分类获取子级Ids
func (_ProductCategory *ProductCategory) CategoryChildrenIds(categoryList []*productsModel.Category, currentId uint) []int {
	data := make([]int, 0)

	categoryListTmp := make(map[uint][]*productsModel.Category)
	for _, category := range categoryList {
		if categoryListTmp[category.ParentId] == nil {
			categoryListTmp[category.ParentId] = make([]*productsModel.Category, 0)
		}
		categoryListTmp[category.ParentId] = append(categoryListTmp[category.ParentId], category)
	}

	// 递归查找子集
	findEndChildrenIds := func(map[uint][]*productsModel.Category, uint) {}
	findEndChildrenIds = func(categoryListTmp map[uint][]*productsModel.Category, currentId uint) {
		for _, categories := range categoryListTmp[currentId] {
			data = append(data, int(categories.ID))
			findEndChildrenIds(categoryListTmp, categories.ID)
		}
	}
	findEndChildrenIds(categoryListTmp, currentId)

	// 当没有子集的时候返回本身
	if len(data) == 0 {
		data = append(data, int(currentId))
	}

	return data
}

// FindAllEndClient 获取最低的子类
func (_ProductCategory *ProductCategory) FindAllEndClient(adminIds []uint) []uint {
	ids := make([]uint, 0)
	err := database.Db.Raw("SELECT c.id FROM category AS c LEFT JOIN category AS c1 ON c1.parent_id = c.id WHERE c1.id IS NULL").Where("c.id IN ?", adminIds).Scan(&ids).Error
	if err != nil {
		return nil
	}
	return ids
}

// GetEndClientShowWhere 获取最低子类的显示条件
func (_ProductCategory *ProductCategory) GetEndClientShowWhere(adminIds []uint) string {
	endClientIds := _ProductCategory.FindAllEndClient(adminIds)
	scopeRow := ""
	for i, id := range endClientIds {
		scopeRow += fmt.Sprintf("scope.row.id == %d", id)
		if i != len(endClientIds)-1 {
			scopeRow += "||"
		}
	}
	return scopeRow
}
