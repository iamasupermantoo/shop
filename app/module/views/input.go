package views

import (
	"gofiber/utils"
)

type InputAttrsViews struct {
	Label    string      `json:"label"`    //	显示标题
	Field    string      `json:"field"`    //	字段名称
	Alias    string      `json:"alias"`    //	字段别名
	Type     int         `json:"type"`     //	input类型
	Readonly bool        `json:"readonly"` //	只读
	Data     interface{} `json:"data"`     //	input配置
}

type InputViews struct {
	InputList []*InputAttrsViews     `json:"inputList"` //	input组
	Params    map[string]interface{} `json:"params"`    //	input参数
}

// NewInputViews 创建Inputs
func NewInputViews() *InputViews {
	return &InputViews{
		Params:    map[string]interface{}{},
		InputList: make([]*InputAttrsViews, 0),
	}
}

// GetInputListInfo 获取参数信息
func (_InputViews *InputViews) GetInputListInfo() (interface{}, []*InputAttrsViews) {
	return _InputViews.Params, _InputViews.InputList
}

// GetInputListRows 获取Input数组
func (_InputViews *InputViews) GetInputListRows() [][]*InputAttrsViews {
	inputList := make([][]*InputAttrsViews, 0)
	for _, inputInfo := range _InputViews.InputList {
		inputList = append(inputList, []*InputAttrsViews{inputInfo})
	}

	return inputList
}

// GetInputListColumn 获取Input数组
func (_InputViews *InputViews) GetInputListColumn() [][]*InputAttrsViews {
	return [][]*InputAttrsViews{_InputViews.InputList}
}

// Text 文本
func (_InputViews *InputViews) Text(label, field string) *InputViews {
	_InputViews.SetInput(InputTypeText, label, field, nil)
	_InputViews.Params[field] = ""
	return _InputViews
}

// TextArea 文本域
func (_InputViews *InputViews) TextArea(label, field string) *InputViews {
	_InputViews.SetInput(InputTypeTextArea, label, field, nil)
	_InputViews.Params[field] = ""
	return _InputViews
}

// Editor 富文本
func (_InputViews *InputViews) Editor(label, field string) *InputViews {
	_InputViews.SetInput(InputTypeEditor, label, field, nil)
	_InputViews.Params[field] = ""
	return _InputViews
}

// Number 数字
func (_InputViews *InputViews) Number(label, field string) *InputViews {
	_InputViews.SetInput(InputTypeNumber, label, field, nil)
	_InputViews.Params[field] = nil
	return _InputViews
}

// Password 密码
func (_InputViews *InputViews) Password(label, field string) *InputViews {
	_InputViews.SetInput(InputTypePassword, label, field, nil)
	_InputViews.Params[field] = ""
	return _InputViews
}

// Select 选择框
func (_InputViews *InputViews) Select(label, field string, data []*InputOptions) *InputViews {
	if len(data) > 0 && !utils.IsNumeric(data[0].Value) {
		if _, ok := _InputViews.Params[field]; !ok {
			_InputViews.Params[field] = ""
		}
		data = append([]*InputOptions{{Label: "全部", Value: ""}}, data...)
	} else {
		if _, ok := _InputViews.Params[field]; !ok {
			_InputViews.Params[field] = 0
		}
		data = append([]*InputOptions{{Label: "全部", Value: 0}}, data...)
	}

	// 去掉重复掉值
	newData := make([]*InputOptions, 0)
	for _, datum := range data {
		isAppend := true
		for newDatumIndex, newDatum := range newData {
			if newDatum.Value == datum.Value {
				isAppend = false
				newData[newDatumIndex] = datum
			}
		}
		if isAppend {
			newData = append(newData, datum)
		}
	}

	_InputViews.SetInput(InputTypeSelect, label, field, newData)
	return _InputViews
}

// SelectDefault 设置默认值
func (_InputViews *InputViews) SelectDefault(label, field string, data []*InputOptions) *InputViews {
	if len(data) > 0 {
		_InputViews.Params[field] = data[0].Value
	}
	_InputViews.Select(label, field, data)
	return _InputViews
}

// Radio 单选框
func (_InputViews *InputViews) Radio(label, field string, data []*InputOptions) *InputViews {
	_InputViews.SetInput(InputTypeRadio, label, field, data)
	_InputViews.Params[field] = ""
	return _InputViews
}

// Checkbox 多选框
func (_InputViews *InputViews) Checkbox(label, field string, data []*InputOptions) *InputViews {
	_InputViews.SetInput(InputTypeCheckbox, label, field, data)
	_InputViews.Params[field] = ""
	return _InputViews
}

// Toggle 开关
func (_InputViews *InputViews) Toggle(label, field string, data []*InputOptions) *InputViews {
	_InputViews.SetInput(InputTypeToggle, label, field, data)
	_InputViews.Params[field] = false
	return _InputViews
}

// File 文件
func (_InputViews *InputViews) File(label, field string) *InputViews {
	_InputViews.SetInput(InputTypeFile, label, field, nil)
	_InputViews.Params[field] = ""
	return _InputViews
}

// Image 图片
func (_InputViews *InputViews) Image(label, field string) *InputViews {
	_InputViews.SetInput(InputTypeImage, label, field, nil)
	_InputViews.Params[field] = ""
	return _InputViews
}

// Images 图片组
func (_InputViews *InputViews) Images(label, field string) *InputViews {
	_InputViews.SetInput(InputTypeImages, label, field, nil)
	_InputViews.Params[field] = make([]string, 0)
	return _InputViews
}

// DatePicker 时间格式
func (_InputViews *InputViews) DatePicker(label, field string) *InputViews {
	_InputViews.SetInput(InputTypeDatePicker, label, field, nil)
	_InputViews.Params[field] = ""
	return _InputViews
}

// RangeDatePicker 时间范围
func (_InputViews *InputViews) RangeDatePicker(label, field string) *InputViews {
	_InputViews.SetInput(InputTypeRangeDatePicker, label, field, nil)
	_InputViews.Params[field] = &RangeDatePicker{}
	return _InputViews
}

// Json 结构体
func (_InputViews *InputViews) Json(label, field string, value interface{}) *InputViews {
	_InputViews.SetInput(InputTypeJson, label, field, value)
	_InputViews.Params[field] = map[string]interface{}{}
	return _InputViews
}

// Children 子级
func (_InputViews *InputViews) Children(label, field string, value interface{}) *InputViews {
	_InputViews.SetInput(InputTypeChildren, label, field, value)
	_InputViews.Params[field] = []interface{}{}
	return _InputViews
}

// SetValue 设置值
func (_InputViews *InputViews) SetValue(field string, value interface{}) *InputViews {
	_InputViews.Params[field] = value
	return _InputViews
}

// SetAlias 设置别名
func (_InputViews *InputViews) SetAlias(field string, alias string) *InputViews {
	for _, v := range _InputViews.InputList {
		if v.Field == field {
			v.Alias = alias
		}
	}
	return _InputViews
}

// SetReadonly 设置是否显示
func (_InputViews *InputViews) SetReadonly(field string) *InputViews {
	for _, v := range _InputViews.InputList {
		if v.Field == field {
			v.Readonly = true
		}
	}
	return _InputViews
}

// SetInput 设置Input
func (_InputViews *InputViews) SetInput(inputType int, label, field string, data any) *InputViews {
	_InputViews.InputList = append(_InputViews.InputList, &InputAttrsViews{
		Type:  inputType,
		Field: field,
		Label: label,
		Data:  data,
	})
	return _InputViews
}
