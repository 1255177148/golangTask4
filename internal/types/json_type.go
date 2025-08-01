package types

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type JSONTime time.Time

const layout = "2006-01-02 15:04:05"

// MarshalJSON JSON 返回时自动格式化
func (t *JSONTime) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%s\"", time.Time(*t).Format(layout))
	return []byte(formatted), nil
}

// UnmarshalJSON 解析 JSON 时支持字符串
func (t *JSONTime) UnmarshalJSON(data []byte) error {
	parsed, err := time.ParseInLocation(`"`+layout+`"`, string(data), time.Local)
	if err != nil {
		return err
	}
	*t = JSONTime(parsed)
	return nil
}

// Value 让 GORM 存数据库时正常存储
func (t *JSONTime) Value() (driver.Value, error) {
	return time.Time(*t), nil
}

// Scan 让 GORM 读数据库时正常读取
func (t *JSONTime) Scan(value interface{}) error {
	if v, ok := value.(time.Time); ok {
		*t = JSONTime(v)
	}
	return nil
}

// Time 转换成 time.Time 使用
func (t *JSONTime) Time() time.Time {
	return time.Time(*t)
}
