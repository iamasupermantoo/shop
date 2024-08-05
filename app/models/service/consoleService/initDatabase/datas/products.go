package datas

import (
	"gofiber/app/module/database"
	"os"
	"strings"
)

// InitProduct 初始化产品信息
func InitProduct(sqlUrl string) {
	dir, err := os.ReadDir(sqlUrl)
	if err != nil {
		panic(err)
	}

	tx := database.Db.Begin()
	for _, file := range dir {
		if !file.IsDir() {
			fileName := file.Name()
			if strings.LastIndex(fileName, ".sql") > -1 {
				value, err := os.ReadFile(sqlUrl + "/" + fileName)
				if err != nil {
					tx.Rollback()
					panic(err)
				}
				valueArr := strings.Split(string(value), ";\n")
				for _, v := range valueArr {
					if v != "" {
						if err = tx.Exec(v).Error; err != nil {
							tx.Rollback()
							panic(err)
						}
					}
				}
			}
		}
	}
	tx.Commit()
}
