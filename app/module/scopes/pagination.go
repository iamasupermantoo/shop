package scopes

import (
	"gofiber/utils"
	"gorm.io/gorm"
)

// Pagination 分页
type Pagination struct {
	SortBy      string `json:"sortBy"`      //	排序字段
	Descending  bool   `json:"descending"`  //	排序 真DESC 假ASC
	Page        int    `json:"page"`        //	当前页数
	RowsPerPage int    `json:"rowsPerPage"` //	每页显示条数
}

// Scopes 分页处理
func (_Pagination *Pagination) Scopes() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		tableName := db.Statement.Table

		if _Pagination == nil || _Pagination.RowsPerPage <= 0 {
			return db
		}

		//	页数不对的话
		if _Pagination.Page <= 0 {
			_Pagination.Page = 1
		}

		if _Pagination.SortBy != "" && _Pagination.Descending {
			if tableName != "" {
				tableName += "."
			}
			orderStr := tableName + utils.ToUnderlinedWords(_Pagination.SortBy) + " DESC"
			if !_Pagination.Descending {
				orderStr = tableName + utils.ToUnderlinedWords(_Pagination.SortBy) + " ASC"
			}
			db.Order(orderStr)
		}

		return db.Offset((_Pagination.Page - 1) * _Pagination.RowsPerPage).Limit(_Pagination.RowsPerPage)
	}
}
