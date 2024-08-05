package types

import (
	"database/sql/driver"
	"gofiber/utils"
)

type GormPasswordParams string

func (_GormPasswordParams GormPasswordParams) Value() (driver.Value, error) {
	if string(_GormPasswordParams) == "" {
		return "", nil
	}
	return utils.PasswordEncrypt(string(_GormPasswordParams)), nil
}
