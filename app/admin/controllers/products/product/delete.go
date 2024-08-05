package product

import (
	"errors"
	"gofiber/app/config"
	"gofiber/app/models/model/productsModel"
	"gofiber/app/models/model/types"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
	"gofiber/utils"
	"gorm.io/gorm"
	"math"
	"os"
	"strings"
)

// Delete 删除接口
func Delete(ctx *context.CustomCtx, params *context.DeleteParams) error {
	err := PermanentDelete(params.Ids, ctx.GetAdminChildIds())
	if err != nil {
		return ctx.ErrorJson(err.Error())
	}

	return ctx.SuccessJsonOK()
}

// PermanentDelete 永久化删除产品
func PermanentDelete(ids []int, adminIds []uint) error {
	err := database.Db.Transaction(func(tx *gorm.DB) error {
		err := tx.Unscoped().Where("id IN ?", ids).
			Where("admin_id IN ?", adminIds).
			Delete(&productsModel.Product{}).Error
		if err != nil {
			return err
		}

		err = tx.Unscoped().Where("product_id IN ?", ids).
			Delete(&productsModel.ProductAttrsSku{}).Error
		if err != nil {
			return err
		}

		productKeyIds := make([]int, 0)
		err = tx.Model(&productsModel.ProductAttrsKey{}).Unscoped().Select("id").
			Where("product_id IN ?", ids).
			Scan(&productKeyIds).Error

		err = tx.Unscoped().Where("id IN ?", productKeyIds).
			Delete(&productsModel.ProductAttrsKey{}).Error
		if err != nil {
			return err
		}

		err = tx.Unscoped().Where("key_id IN ?", productKeyIds).Delete(&productsModel.ProductAttrsVal{}).Error
		if err != nil {
			return err
		}

		return nil
	})
	return err
}

// removeImageByIds 同步Id 移动产品图片
func removeImageByIds(ids []int, adminIds []uint) {
	productList := make([]*productsModel.Product, 0)
	err := database.Db.Where("admin_id IN ?", adminIds).
		Where("id IN ?", ids).
		Find(&productList).Error
	if err != nil {
		return
	}

	cfg := config.Conf
	for _, p := range productList {
		for _, image := range p.Images {
			index := strings.LastIndex(image, "/")
			fileName := image[index+1:]
			_ = removeFile(cfg.FileRoot+image, cfg.FileRoot+"/crawling/product_old/"+fileName)
		}
	}
}

// removeFile 移动文件
func removeFile(sourcePath, targetPath string) error {
	sourceDir := strings.TrimSuffix(sourcePath, "/"+strings.Split(sourcePath, "/")[len(strings.Split(sourcePath, "/"))-1])
	if !utils.PathExists(sourceDir) {
		return errors.New("path not exists")
	}

	targetDir := strings.TrimSuffix(targetPath, "/"+strings.Split(targetPath, "/")[len(strings.Split(targetPath, "/"))-1])
	if !utils.PathExists(targetDir) {
		utils.PathMkdirAll(targetDir)
	}

	err := os.Rename(sourcePath, targetPath)
	if os.IsExist(err) {
		err = os.Remove(sourcePath)
		if err != nil {
			return err
		}
	}
	return err
}

// findNotFindProductImage 查找没有找到产品图片的url路径
func findNotFindProductImage(settingAdminId uint) {
	productList := make([]*productsModel.Product, 0)
	err := database.Db.Where("admin_id = ?", settingAdminId).Where("type = ?", productsModel.ProductTypeWholesale).Find(&productList).Error
	if err != nil {
		return
	}

	cfg := config.Conf
	for _, p := range productList {
		for _, image := range p.Images {
			if !utils.PathExists(cfg.FileRoot + image) {
				index := strings.LastIndex(image, "/")
				fileName := image[index+1:]
				_ = removeFile(cfg.FileRoot+"/crawling/product_old/"+fileName, cfg.FileRoot+image)
			}
		}
	}

}

// findRepeatIds 查找全部重复的产品Id
func findRepeatIds(settingAdminId uint) []int {
	productList := make([]*productsModel.Product, 0)
	err := database.Db.Where("admin_id = ?", settingAdminId).Where("type = ?", productsModel.ProductTypeWholesale).Find(&productList).Error
	if err != nil {
		return nil
	}

	deleteIds := make([]int, 0)
	productListMap := make(map[string]int)
	for _, p := range productList {
		key := p.Name
		if _, ok := productListMap[key]; ok {
			deleteIds = append(deleteIds, int(p.ID))
		}
		productListMap[key] = int(p.ID)
	}
	return deleteIds
}

// UpdateDiscount 修改折扣为负数的产品
func UpdateDiscount(settingAdminId uint) {
	productList := make([]*productsModel.Product, 0)
	err := database.Db.Where("admin_id = ?", settingAdminId).Where("type = ?", productsModel.ProductTypeWholesale).Find(&productList).Error
	if err != nil {
		return
	}
	for _, v := range productList {
		if v.Discount < 0 {
			v.Discount = math.Abs(v.Discount)
			database.Db.Updates(&productsModel.Product{GormModel: types.GormModel{ID: v.ID}, Discount: v.Discount})
		}
	}

}
