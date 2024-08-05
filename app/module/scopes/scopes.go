package scopes

import (
	"gofiber/app/module/database"
	"gofiber/utils"
	"gorm.io/gorm"
	"time"
)

type GormWhere struct {
	query any
	args  []any
}

type Scopes struct {
	gormWhere []*GormWhere
}

func NewScopes() *Scopes {
	return &Scopes{}
}

func (_Scopes *Scopes) Eq(field string, val any) *Scopes {
	if !utils.IsEmpty(val) {
		_Scopes.gormWhere = append(_Scopes.gormWhere, &GormWhere{query: field + " = ?", args: []any{val}})
	}
	return _Scopes
}

// Like 操作
func (_Scopes *Scopes) Like(field string, val string) *Scopes {
	if val != "" && val != "%" && val != "%%" {
		_Scopes.gormWhere = append(_Scopes.gormWhere, &GormWhere{query: field + " LIKE ?", args: []any{val}})
	}
	return _Scopes
}

// In 操作
func (_Scopes *Scopes) In(field string, val any) *Scopes {
	if !utils.IsEmpty(val) {
		_Scopes.gormWhere = append(_Scopes.gormWhere, &GormWhere{query: field + " IN ?", args: []any{val}})
	}
	return _Scopes
}

// FindModeIn 查询模型
func (_Scopes *Scopes) FindModeIn(foreignKey string, model any, field string, query any, args ...any) *Scopes {
	if !utils.IsEmpty(args[0]) {
		data := make([]any, 0)
		database.Db.Model(model).Select(field).Where(query, args...).Find(&data)
		if len(data) == 0 {
			data = append(data, "0")
		}
		_Scopes.gormWhere = append(_Scopes.gormWhere, &GormWhere{query: foreignKey + " IN ?", args: []any{data}})
	}

	return _Scopes
}

// Between 范围方法
func (_Scopes *Scopes) Between(field string, datePicker *RangeDatePicker) *Scopes {
	if datePicker != nil && datePicker.From != "" && datePicker.To != "" {
		var staTime, endTime time.Time
		if len(datePicker.From) == 19 {
			staTime, _ = time.ParseInLocation("2006/01/02 15:04:05", datePicker.From, time.Local)
		} else {
			staTime, _ = time.ParseInLocation("2006/01/02", datePicker.From, time.Local)
		}

		if len(datePicker.To) == 19 {
			endTime, _ = time.ParseInLocation("2006/01/02 15:04:05", datePicker.To, time.Local)
		} else {
			endTime, _ = time.ParseInLocation("2006/01/02", datePicker.To, time.Local)
		}

		_Scopes.gormWhere = append(_Scopes.gormWhere, &GormWhere{query: field + " BETWEEN ? AND ?", args: []any{staTime, endTime}})
	}
	return _Scopes
}

func (_Scopes *Scopes) Scopes() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		for _, where := range _Scopes.gormWhere {
			db.Where(where.query, where.args...)
		}
		return db
	}
}
