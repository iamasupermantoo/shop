package context

import "gofiber/app/module/scopes"

// DeleteParams 删除参数
type DeleteParams struct {
	Ids []int `validate:"required" json:"ids"` //	删除ID
}

// PrimaryKeyParams 主键ID
type PrimaryKeyParams struct {
	ID uint `validate:"required" json:"id"` //	ID
}

// IndexData 列表数据
type IndexData struct {
	Items interface{} `json:"items"` //	数据列表
	Count int64       `json:"count"` //	总数
}

// IndexParams 列表参数
type IndexParams struct {
	UpdatedAt  *scopes.RangeDatePicker `json:"updatedAt"`
	CreatedAt  *scopes.RangeDatePicker `json:"createdAt"`
	Pagination *scopes.Pagination      `json:"pagination"`
}
