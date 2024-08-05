package views

// DialogViews 弹窗
type DialogViews struct {
	Id         string                     `json:"id"`         //	ID
	Url        string                     `json:"url"`        //	提交地址
	Title      string                     `json:"title"`      //	标题
	Small      string                     `json:"small"`      //	副标题
	Content    string                     `json:"content"`    //	内容
	Sizing     string                     `json:"sizing"`     //	大小
	FullWidth  bool                       `json:"fullWidth"`  //	满屏宽度
	FullHeight bool                       `json:"fullHeight"` //	满屏高度
	Params     interface{}                `json:"params"`     //	参数
	InputList  []*InputAttrsViews         `json:"inputList"`  //	Input列表
	Buttons    *DialogActionsButtonsViews `json:"buttons"`    //	按钮组
}

func NewDialogViews(id, submitUrl, title string) *DialogViews {
	return &DialogViews{
		Id:        id,
		Url:       submitUrl,
		Title:     title,
		Sizing:    SizingMedium,
		Params:    map[string]string{},
		InputList: make([]*InputAttrsViews, 0),
		Buttons: &DialogActionsButtonsViews{
			Cancel: &ButtonViews{
				Label: "取消",
			},
			Confirm: &ButtonViews{
				Label: "提交",
			},
		},
	}
}

// SetSmall 设置副标题
func (_DialogViews *DialogViews) SetSmall(small string) *DialogViews {
	_DialogViews.Small = small
	return _DialogViews
}

// SetContent 设置内容
func (_DialogViews *DialogViews) SetContent(content string) *DialogViews {
	_DialogViews.Content = content
	return _DialogViews
}

// SetSizingSmall 设置大小
func (_DialogViews *DialogViews) SetSizingSmall() *DialogViews {
	_DialogViews.Sizing = SizingSmall
	return _DialogViews
}

// SetFullWidth 设置满屏宽度
func (_DialogViews *DialogViews) SetFullWidth() *DialogViews {
	_DialogViews.FullWidth = true
	return _DialogViews
}

// SetFullHeight 设置满屏高度
func (_DialogViews *DialogViews) SetFullHeight() *DialogViews {
	_DialogViews.FullHeight = true
	return _DialogViews
}

// SetInputViews 设置InputViews
func (_DialogViews *DialogViews) SetInputViews(inputViews *InputViews) *DialogViews {
	params, inputList := inputViews.GetInputListInfo()
	_DialogViews.SetParams(params)
	_DialogViews.SetInputList(inputList)
	return _DialogViews
}

// SetParams 设置参数
func (_DialogViews *DialogViews) SetParams(params interface{}) *DialogViews {
	_DialogViews.Params = params
	return _DialogViews
}

// SetInputList 甚至input输入框
func (_DialogViews *DialogViews) SetInputList(inputList []*InputAttrsViews) *DialogViews {
	_DialogViews.InputList = inputList
	return _DialogViews
}

// SetCancelButton 设置取消按钮
func (_DialogViews *DialogViews) SetCancelButton(button *ButtonViews) *DialogViews {
	_DialogViews.Buttons.Cancel = button
	return _DialogViews
}

// SetConfirmButton 设置提交按钮
func (_DialogViews *DialogViews) SetConfirmButton(button *ButtonViews) *DialogViews {
	_DialogViews.Buttons.Confirm = button
	return _DialogViews
}
