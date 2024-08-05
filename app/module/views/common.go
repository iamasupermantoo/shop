package views

import "github.com/goccy/go-json"

const (
	SizingSmall    = "small"     //	dialog 小型宽度
	SizingMedium   = "medium"    //	dialog	中型宽度 默认值
	ButtonSizeXS   = "xs"        //	小按钮
	ButtonSizeSM   = "sm"        //	一般
	ButtonSizeMD   = "md"        //	正常
	ColorPrimary   = "primary"   //	主颜色
	ColorSecondary = "secondary" //	辅色
	ColorPositive  = "positive"  //	绿色
	ColorNegative  = "negative"  //	红色
	ColorAccent    = "accent"    //	紫色
	ColorWarning   = "warning"   //	警告色

	OperateCheckbox = "checkbox" //	选择操作
	OperateSetting  = "setting"  //	配置操作

	// InputTypeText 文本类型
	InputTypeText = 1
	// InputTypeTextArea 文本域类型
	InputTypeTextArea = 2
	// InputTypeEditor 富文本类型
	InputTypeEditor = 3
	// InputTypeNumber 数字类型
	InputTypeNumber = 4
	// InputTypePassword 密码类型
	InputTypePassword = 5
	// InputTypeSelect 选择类型
	InputTypeSelect = 10
	// InputTypeRadio 单选类型
	InputTypeRadio = 11
	// InputTypeCheckbox 多选类型
	InputTypeCheckbox = 12
	// InputTypeToggle 开关类型
	InputTypeToggle = 13
	// InputTypeFile 文件类型
	InputTypeFile = 21
	// InputTypeImage 图片类型
	InputTypeImage = 22
	// InputTypeImages 图片组类型
	InputTypeImages = 23
	// InputTypeIcon 图标
	InputTypeIcon = 24
	// InputTypeDatePicker 时间类型
	InputTypeDatePicker = 31
	// InputTypeRangeDatePicker 时间范围类型
	InputTypeRangeDatePicker = 32
	// InputTypeJson 对象类型
	InputTypeJson = 41
	// InputTypeChildren 子级对象类型
	InputTypeChildren = 42
	// InputTypeEditText 编辑文本类型
	InputTypeEditText = 51
	// InputTypeEditNumber 编辑数字类型
	InputTypeEditNumber = 52
	// InputTypeEditTextArea 编辑富文本类型
	InputTypeEditTextArea = 53
	// InputTypeEditToggle 编辑开关类型
	InputTypeEditToggle = 54
	// InputTypeTranslate 翻译类型
	InputTypeTranslate = 61
)

// InputOptions select|checkbox|toggle|radio
type InputOptions struct {
	Label string      `gorm:"column:label" json:"label"` //	名称
	Value interface{} `gorm:"column:value" json:"value"` //	值
}

// InputCheckboxOptions checkbox值
type InputCheckboxOptions struct {
	Label string      `gorm:"column:label" json:"label"` //	名称
	Value interface{} `gorm:"column:value" json:"value"` //	值
	Field string      `json:"field"`                     //	字段名称
}

// InputJsonOptions Json格式
type InputJsonOptions struct {
	Items [][]*InputAttrsViews `json:"items"` //	input组
}

// InputChildrenOptions 子级
type InputChildrenOptions struct {
	Items    [][]*InputAttrsViews `json:"items"`    //	input组
	IsCreate bool                 `json:"isCreate"` //	是否新增
	IsDelete bool                 `json:"isDelete"` //	是否删除
}

// CheckboxListOptions Table多选配置参数
type CheckboxListOptions struct {
	Operate string `json:"operate"` //	操作名称
	Name    string `json:"name"`    //	名称
	Field   string `json:"field"`   //	提交的字段名称
	Scan    string `json:"scan"`    //	提取的字段名称
}

// SettingOptions 设置参数
type SettingOptions struct {
	Operate string `json:"operate"` //	操作名称
	Name    string `json:"name"`    //	名称
	Params  string `json:"params"`  //	参数
	Type    string `json:"type"`    //  类型字段
	Value   string `json:"value"`   // 	值字段
	Input   string `json:"input"`   //	输入框
}

// Pagination 分页
type Pagination struct {
	SortBy      string `json:"sortBy"`      //	排序字段
	Descending  bool   `json:"descending"`  //	排序 真DESC 假ASC
	Page        int    `json:"page"`        //	当前页数
	RowsPerPage int    `json:"rowsPerPage"` //	每页显示条数
}

// RangeDatePicker 时间范围
type RangeDatePicker struct {
	From string `json:"from"` //	开始时间
	To   string `json:"to"`   //	结束时间
}

var InputTypeOptions = []*InputOptions{
	{Label: "文本类型", Value: InputTypeText},
	{Label: "文本域类型", Value: InputTypeTextArea},
	{Label: "富文本类型", Value: InputTypeEditor},
	{Label: "数字类型", Value: InputTypeNumber},
	{Label: "密码类型", Value: InputTypePassword},
	{Label: "选择类型", Value: InputTypeSelect},
	{Label: "单选类型", Value: InputTypeRadio},
	{Label: "多选类型", Value: InputTypeCheckbox},
	{Label: "开关类型", Value: InputTypeToggle},
	{Label: "文件类型", Value: InputTypeFile},
	{Label: "图片类型", Value: InputTypeImage},
	{Label: "图片图标", Value: InputTypeIcon},
	{Label: "图片组类型", Value: InputTypeImages},
	{Label: "时间类型", Value: InputTypeDatePicker},
	{Label: "时间范围类型", Value: InputTypeRangeDatePicker},
	{Label: "对象类型", Value: InputTypeJson},
	{Label: "子级对象类型", Value: InputTypeChildren},
	{Label: "编辑文本类型", Value: InputTypeEditText},
	{Label: "编辑数字类型", Value: InputTypeEditNumber},
	{Label: "编辑文本域类型", Value: InputTypeEditTextArea},
	{Label: "编辑开关类型", Value: InputTypeEditToggle},
	{Label: "翻译类型", Value: InputTypeTranslate},
}

// InputViewsStringToData 值字符串转对象
func InputViewsStringToData(inputType int, value string) interface{} {
	var data interface{}
	switch inputType {
	case InputTypeChildren, InputTypeJson, InputTypeImages, InputTypeCheckbox, InputTypeSelect:
		_ = json.Unmarshal([]byte(value), &data)
	default:
		data = value
	}

	return data
}
