package scopes

import "time"

// RangeDatePicker 时间范围
type RangeDatePicker struct {
	From string `json:"from"` //	开始时间
	To   string `json:"to"`   //	结束时间
}

// ToAddDate 加一天时间
func (r *RangeDatePicker) ToAddDate() *RangeDatePicker {
	location, err := time.Parse("2006/01/02", r.To)
	if err != nil {
		return r
	}
	location = location.AddDate(0, 0, 1)
	r.To = location.Format("2006/01/02")
	return r
}

// FromAddDate 加一天时间
func (r *RangeDatePicker) FromAddDate() *RangeDatePicker {
	location, err := time.Parse("2006/01/02", r.From)
	if err != nil {
		return r
	}
	location = location.AddDate(0, 0, 1)
	r.From = location.Format("2006/01/02")
	return r
}
