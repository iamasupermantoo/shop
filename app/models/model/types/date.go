package types

import (
	"database/sql/driver"
	"time"
)

type GormTimeParams string

func (_GormTimeParams GormTimeParams) Value() (driver.Value, error) {
	if string(_GormTimeParams) == "" {
		return time.Now(), nil
	}

	var currentTime time.Time
	var err error
	if len(string(_GormTimeParams)) == 19 {
		currentTime, err = time.ParseInLocation("2006/01/02 15:04:05", string(_GormTimeParams), time.Local)
	} else {
		currentTime, err = time.ParseInLocation("2006/01/02", string(_GormTimeParams), time.Local)
	}
	if err != nil {
		return time.Now(), nil
	}
	return currentTime, nil
}
