package utils

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"time"
)

type LocalTime time.Time

func (t *LocalTime) MarshalJSON() ([]byte, error) {
	tTime := time.Time(*t)
	return []byte(fmt.Sprintf("\"%v\"", tTime.Format("2006-01-02 15:04:05"))), nil
}

// Scan 实现 Scanner 接口
func (lt *LocalTime) Scan(value interface{}) error {
	if value == nil {
		*lt = LocalTime(time.Time{})
		return nil
	}

	if t, ok := value.(time.Time); ok {
		*lt = LocalTime(t)
		return nil
	}

	return errors.New("invalid data type for LocalTime")
}

// Value 实现 Valuer 接口
func (lt LocalTime) Value() (driver.Value, error) {
	t := time.Time(lt)
	return t, nil
}
