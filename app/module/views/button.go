package views

// ButtonViews 按钮
type ButtonViews struct {
	Label string `json:"label"` //	按钮名称
	Color string `json:"color"` //	颜色
	Size  string `json:"size"`  //	按钮大小
	Eval  string `json:"eval"`  //	执行字符串		在数据表格中可以使用 scope.row(当前行) scope.col(当前列)
}

// DialogButtonViews 按钮Dialog
type DialogButtonViews struct {
	ButtonViews
	Config *DialogViews `json:"config"` //	配置
}

// DialogActionsButtonsViews Dialog 提交按钮组
type DialogActionsButtonsViews struct {
	Cancel  *ButtonViews `json:"cancel"`  //	取消按钮
	Confirm *ButtonViews `json:"confirm"` //	确认按钮
}
