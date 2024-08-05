package datas

import (
	"gofiber/app/models/model/shopsModel"
)

func InitShippingAddress() []*shopsModel.ShippingAddress {
	return []*shopsModel.ShippingAddress{
		{AdminId: 2, UserId: 1, Name: "隔壁老王", Contact: "muiprosperyls15@gmail.com", City: "美国佛罗里达州", Address: "美国佛罗里达州 18#3008栋", IsShow: shopsModel.ShippingAddressIsShowYes},
	}
}
