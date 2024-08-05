package command

import (
	"gofiber/app/models/service/consoleService/initDatabase"
)

// InitDatabase 初始化数据库
func InitDatabase(filterTable []string) *initDatabase.Database {
	databases := &initDatabase.Database{FilterTable: filterTable}
	return databases.InitTables()
}
