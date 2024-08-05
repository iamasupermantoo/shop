package views

// TableViews 数据表格视图
type TableViews struct {
	Url        string            `json:"url"`        //	请求数据地址
	Pagination *Pagination       `json:"pagination"` //	排序规则
	Search     *TableSearchViews `json:"search"`     //	查询字段列表
	Table      *TableConfigViews `json:"table"`      //	数据表格配置
}

// TableSearchViews 视图查询配置
type TableSearchViews struct {
	Params    interface{}        `json:"params"`    //	参数
	InputList []*InputAttrsViews `json:"inputList"` //	输入框列表
}

// TableConfigViews 视图表格配置
type TableConfigViews struct {
	Key       string               `json:"key"`       //	主键
	UpdateUrl string               `json:"updateUrl"` //	更新路由
	Tools     []*DialogButtonViews `json:"tools"`     //	工具栏
	Columns   []*ColumnsAttrsViews `json:"columns"`   //	字段名称
	Options   []*DialogButtonViews `json:"options"`   //	操作栏
}

// NewTableViews 创建表格视图对象
func NewTableViews(indexUrl, updateUrl string) *TableViews {
	return &TableViews{
		Url:        indexUrl,
		Pagination: &Pagination{SortBy: "id", Descending: true, Page: 1, RowsPerPage: 20},
		Search: &TableSearchViews{
			Params:    map[string]interface{}{},
			InputList: make([]*InputAttrsViews, 0),
		},
		Table: &TableConfigViews{
			Key:       "id",
			UpdateUrl: updateUrl,
			Tools:     make([]*DialogButtonViews, 0),
			Columns:   make([]*ColumnsAttrsViews, 0),
			Options:   make([]*DialogButtonViews, 0),
		},
	}
}

// SetTableKey 设置Table key
func (_TableViews *TableViews) SetTableKey(key string) *TableViews {
	_TableViews.Table.Key = key
	return _TableViews
}

// SetSearch 设置查询
func (_TableViews *TableViews) SetSearch(inputViews *InputViews) *TableViews {
	_TableViews.Search.Params, _TableViews.Search.InputList = inputViews.GetInputListInfo()
	return _TableViews
}

// SetTools 设置工具栏按钮
func (_TableViews *TableViews) SetTools(dialogBtnList ...*DialogButtonViews) *TableViews {
	for _, dialogBtn := range dialogBtnList {
		_TableViews.Table.Tools = append(_TableViews.Table.Tools, dialogBtn)
	}
	return _TableViews
}

// SetColumn 设置数据表格
func (_TableViews *TableViews) SetColumn(columnsViews *ColumnsViews) *TableViews {
	_TableViews.Table.Columns = columnsViews.GetColumnsListInfo()
	return _TableViews
}

// SetOptions 设置表格Options
func (_TableViews *TableViews) SetOptions(dialogBtnList ...*DialogButtonViews) *TableViews {
	for _, dialogBtn := range dialogBtnList {
		_TableViews.Table.Options = append(_TableViews.Table.Options, dialogBtn)
	}
	return _TableViews
}
