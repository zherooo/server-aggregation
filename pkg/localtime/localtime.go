package localtime

import (
	"database/sql/driver"
	"fmt"
	"server-aggregation/internal/consts"
	"strings"
	"time"
)

// LocalDate 本地日期
type LocalTime struct {
	time.Time
}

// UnmarshalJSON gin bind 反射结构体
func (t *LocalTime) UnmarshalJSON(bytes []byte) (err error) {
	if string(bytes) == "null" || string(bytes) == "\"\"" {
		return
	}
	t.Time, err = time.ParseInLocation(consts.DateTimeFormat, strings.Trim(string(bytes), "\""), time.Local)
	return
}

// MarshalJSON gorm marshal 序列化结构体
func (t LocalTime) MarshalJSON() ([]byte, error) {
	var str string
	if !t.Time.IsZero() {
		str = t.Time.Format(consts.DateTimeFormat)
	}
	return []byte(fmt.Sprintf("\"%s\"", str)), nil
}

// Value LocalDate 转 time
func (t LocalTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

// Scan gorm Scan 扫描时的数据赋值
func (t *LocalTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = LocalTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
