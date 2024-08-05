package curd

import (
	"errors"
	"fmt"
	"gofiber/app/models/service/consoleService/initDatabase"
	"gofiber/utils"
	"os"
	"strings"
)

type Generate struct {
	filePath          string             //	文件路径
	filePackageName   string             //	文件包名
	modelsPathStr     string             //	模型路径
	modelsPackageName string             // 	模型包名
	tableName         string             //	表名
	tableTitleName    string             //	驼峰表名
	table             initDatabase.Table //	当前表
}

func NewGenerate(name string, path string) *Generate {
	pathList := strings.Split(path, "/")
	tableName := name
	currentTable := new(initDatabase.Database).InitTables().GetTable(tableName)

	schema := currentTable.GetFields().String()
	schemaList := strings.Split(schema, ".")
	schemaPathList := strings.Split(schemaList[0], "/")
	modelsPathStr := strings.Join(schemaPathList[:len(schemaPathList)-1], "/")
	modelsPackageName := schemaPathList[len(schemaPathList)-1]

	//	{package}   				===>	文件包名
	//	{modelsPath}				===>	模型路径
	//	{modelsPackage}				===>	模型包名
	//	{modelsStruct}				===>	模型结构体
	//	{modelsFieldStruct}			===>	模型字段结构体
	//	{modelsFieldParams}			===>	模型对应的参数

	return &Generate{
		filePath:          "/app/" + path,
		filePackageName:   pathList[len(pathList)-1],
		modelsPathStr:     modelsPathStr,
		modelsPackageName: modelsPackageName,
		tableName:         tableName,
		tableTitleName:    schemaList[1],
		table:             currentTable,
	}
}

func (_Generate *Generate) CURD() error {

	_Generate.GenerateCreate()

	_Generate.GenerateDelete()

	_Generate.GenerateUpdate()

	_Generate.GenerateIndex()

	_Generate.GenerateViews()

	return nil
}

func (_Generate *Generate) GenerateIndex() {
	str := indexTmp
	str = strings.ReplaceAll(str, "{package}", _Generate.filePackageName)
	str = strings.ReplaceAll(str, "{modelsPath}", _Generate.modelsPathStr)
	str = strings.ReplaceAll(str, "{modelsPackage}", _Generate.modelsPackageName)
	str = strings.ReplaceAll(str, "{modelsStruct}", _Generate.tableTitleName)
	str = strings.ReplaceAll(str, "{modelsFieldStruct}", _Generate.table.GetFieldsToStruct())

	err := _Generate.save("index.go", str)
	if err == nil {
		fmt.Println("生成查询接口文件成功....")
	}
}

func (_Generate *Generate) GenerateViews() {
	str := viewsTmp
	str = strings.ReplaceAll(str, "{package}", _Generate.filePackageName)
	str = strings.ReplaceAll(str, "{modelsPath}", _Generate.modelsPathStr)
	str = strings.ReplaceAll(str, "{modelsPackage}", _Generate.modelsPackageName)
	str = strings.ReplaceAll(str, "{modelsStruct}", _Generate.tableTitleName)

	err := _Generate.save("views.go", str)
	if err == nil {
		fmt.Println("生成视图接口文件成功....")
	}
}

// GenerateCreate 生成新增文件
func (_Generate *Generate) GenerateCreate() {
	str := createTmp
	str = strings.ReplaceAll(str, "{package}", _Generate.filePackageName)
	str = strings.ReplaceAll(str, "{modelsPath}", _Generate.modelsPathStr)
	str = strings.ReplaceAll(str, "{modelsPackage}", _Generate.modelsPackageName)
	str = strings.ReplaceAll(str, "{modelsStruct}", _Generate.tableTitleName)
	str = strings.ReplaceAll(str, "{modelsFieldStruct}", _Generate.table.GetFieldsToStruct())
	str = strings.ReplaceAll(str, "{modelsFieldParams}", _Generate.table.GetFieldsToParams())

	err := _Generate.save("create.go", str)
	if err == nil {
		fmt.Println("生成新增接口文件成功....")
	}
}

// GenerateUpdate 生成更新接口
func (_Generate *Generate) GenerateUpdate() {
	str := updateTmp
	str = strings.ReplaceAll(str, "{package}", _Generate.filePackageName)
	str = strings.ReplaceAll(str, "{modelsPath}", _Generate.modelsPathStr)
	str = strings.ReplaceAll(str, "{modelsPackage}", _Generate.modelsPackageName)
	str = strings.ReplaceAll(str, "{modelsStruct}", _Generate.tableTitleName)
	str = strings.ReplaceAll(str, "{modelsFieldStruct}", _Generate.table.GetFieldsToStruct())

	err := _Generate.save("update.go", str)
	if err == nil {
		fmt.Println("生成更新接口文件成功....")
	}
}

// GenerateDelete 生成删除文件
func (_Generate *Generate) GenerateDelete() {
	str := deleteTmp
	str = strings.ReplaceAll(str, "{package}", _Generate.filePackageName)
	str = strings.ReplaceAll(str, "{modelsPath}", _Generate.modelsPathStr)
	str = strings.ReplaceAll(str, "{modelsPackage}", _Generate.modelsPackageName)
	str = strings.ReplaceAll(str, "{modelsStruct}", _Generate.tableTitleName)

	err := _Generate.save("delete.go", str)
	if err == nil {
		fmt.Println("生成删除接口文件成功....")
	}
}

func (_Generate *Generate) save(fileName string, fileContent string) error {
	dir, _ := os.Getwd()

	if !utils.PathExists(dir + _Generate.filePath) {
		utils.PathMkdirAll(dir + _Generate.filePath)
	}

	filePath := dir + _Generate.filePath + "/" + fileName
	if utils.PathExists(filePath) {
		return errors.New("【" + filePath + "】已存在...")
	}

	err := utils.FileWrite(filePath, []byte(fileContent))
	if err != nil {
		panic(err)
	}
	return nil
}
