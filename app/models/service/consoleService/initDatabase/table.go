package initDatabase

import (
	"gorm.io/gorm/schema"
	"sync"
)

// Table 表名
type Table struct {
	Name    string        //	表名
	Model   interface{}   //	模型
	Data    []interface{} //	初始数据
	Comment string        //	注释
}

// GetFields 获取所有字段
func (_Table *Table) GetFields() *schema.Schema {
	m, _ := schema.Parse(_Table.Model, &sync.Map{}, schema.NamingStrategy{})
	return m
}

// GetFieldsToStruct 字段转结构体
func (_Table *Table) GetFieldsToStruct() string {
	fields := _Table.GetFields()
	s := "\n"
	for _, field := range fields.Fields {
		if field.Comment != "" {
			s += "\t" + field.StructField.Name + "\t" + field.FieldType.String() + "\t// " + field.Comment + "\n"
		}
	}
	return s
}

func (_Table *Table) GetFieldsToParams() string {
	fields := _Table.GetFields()
	s := "\n"
	for _, field := range fields.Fields {
		if field.Comment != "" {
			s += "\t\t" + field.StructField.Name + ": params." + field.StructField.Name + ",\n"
		}
	}
	s += "\t"
	return s
}
