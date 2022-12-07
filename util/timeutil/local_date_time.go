package timeutil

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type LocalDateTime struct {
	time.Time
}

// MarshalJSON 重写 MarshaJSON 方法，在此方法中实现自定义格式的转换；
func (t *LocalDateTime) MarshalJSON() ([]byte, error) {
	output := fmt.Sprintf("\"%s\"", t.Format("2006-01-02 15:04:05"))
	return []byte(output), nil
}

// Value 实现 Value 方法，写入数据库时会调用该方法将自定义时间类型转换并写入数据库；
func (t *LocalDateTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

// Scan 实现 Scan 方法，读取数据库时会调用该方法将时间数据转换成自定义时间类型；
func (t *LocalDateTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = LocalDateTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to datetime", v)
}
