package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"gofiber/app/module/cache"
	"gofiber/console/command"
	"gofiber/console/curd"
)

func main() {
	console := &cobra.Command{
		Use:   "app",
		Short: "基础版本命令行工具",
	}

	// 初始化数据库
	// go run main.go init
	console.AddCommand(&cobra.Command{
		Use:   "init",
		Short: "数据库初始化 false",
		Run: func(cmd *cobra.Command, args []string) {
			command.InitDatabase(args).ExecInit()

			// 清空缓存数据
			rdsConn := cache.Rds.Get()
			defer rdsConn.Close()
			_, _ = rdsConn.Do("flushall")
		},
	})

	// 数据库文件生成 CURD
	// go run main.go curd admin_menu admin/controllers/admins/menu
	// admin_menu 				表名			admin_menu[数据库表名]
	// admin/controllers/admins/menu			路径			其中menu[包名]				固定是在 ./app/
	console.AddCommand(&cobra.Command{
		Use:   "curd",
		Short: "数据表生成增删改查 table path",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 2 {
				fmt.Println("缺少参数, 第一个参数为表名, 第二个参数为路径")
				return
			}

			// 新增数据
			_ = curd.NewGenerate(args[0], args[1]).CURD()
		},
	})

	_ = console.Execute()
}
