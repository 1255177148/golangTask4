package utils

import (
	"errors"
	"fmt"
	"github.com/1255177148/golangTask4/internal/utils/log"
	"reflect"
	"strconv"
	"time"
)

// 基础时间格式，按需增加
const timeLayout = "2006-01-02 15:04:05"

// MapStruct 支持基础类型转换的结构体映射
func MapStruct(src, dst interface{}) error {
	sVal := reflect.ValueOf(src)
	dVal := reflect.ValueOf(dst)

	// dst 必须是指针
	if dVal.Kind() != reflect.Ptr || dVal.IsNil() {
		return errors.New("dst must be a non-nil pointer")
	}
	dVal = dVal.Elem()

	if sVal.Kind() == reflect.Ptr {
		sVal = sVal.Elem()
	}

	if sVal.Kind() != reflect.Struct || dVal.Kind() != reflect.Struct {
		return errors.New("src and dst must be struct or pointer to struct")
	}

	for i := 0; i < sVal.NumField(); i++ {
		fieldName := sVal.Type().Field(i).Name
		sField := sVal.Field(i)

		dField := dVal.FieldByName(fieldName)
		if !dField.IsValid() || !dField.CanSet() {
			continue
		}

		if err := setValueWithConvert(dField, sField); err != nil {
			// 可选择返回错误，或者忽略单个字段错误
			log.ErrorF("field %s: %w", fieldName, err)
			return fmt.Errorf("field %s: %w", fieldName, err)
		}
	}
	return nil
}

// setValueWithConvert 支持部分基础类型的赋值和转换
func setValueWithConvert(dst, src reflect.Value) error {
	if !src.IsValid() {
		return errors.New("source value invalid")
	}

	// 1. 类型完全相同，直接赋值
	if src.Type() == dst.Type() {
		dst.Set(src)
		return nil
	}

	// 2. 支持数字类型转换（int, int64, float64 等）
	if isNumericKind(src.Kind()) && isNumericKind(dst.Kind()) {
		srcFloat := src.Convert(reflect.TypeOf(float64(0))).Float()
		switch dst.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			dst.SetInt(int64(srcFloat))
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			dst.SetUint(uint64(srcFloat))
		case reflect.Float32, reflect.Float64:
			dst.SetFloat(srcFloat)
		default:
			return fmt.Errorf("unsupported numeric dst kind: %v", dst.Kind())
		}
		return nil
	}

	// 3. 数字类型转字符串
	if isNumericKind(src.Kind()) && dst.Kind() == reflect.String {
		dst.SetString(fmt.Sprintf("%v", src.Interface()))
		return nil
	}

	// 4. 字符串转数字类型
	if src.Kind() == reflect.String && isNumericKind(dst.Kind()) {
		strVal := src.String()
		switch dst.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			v, err := strconv.ParseInt(strVal, 10, 64)
			if err != nil {
				return err
			}
			dst.SetInt(v)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			v, err := strconv.ParseUint(strVal, 10, 64)
			if err != nil {
				return err
			}
			dst.SetUint(v)
		case reflect.Float32, reflect.Float64:
			v, err := strconv.ParseFloat(strVal, 64)
			if err != nil {
				return err
			}
			dst.SetFloat(v)
		default:
			return fmt.Errorf("unsupported numeric dst kind: %v", dst.Kind())
		}
		return nil
	}

	// 5. time.Time 或 *time.Time  和 string 互转（格式 yyyy-mm-dd hh:mm:ss）
	if (src.Type() == reflect.TypeOf(time.Time{}) || src.Type() == reflect.TypeOf(&time.Time{})) && dst.Kind() == reflect.String {
		if src.IsNil() && src.Kind() == reflect.Ptr {
			// 如果是 *time.Time 且为 nil，就赋空字符串
			dst.SetString("")
			return nil
		}

		var t time.Time
		if src.Type() == reflect.TypeOf(&time.Time{}) {
			t = *src.Interface().(*time.Time)
		} else {
			t = src.Interface().(time.Time)
		}
		dst.SetString(t.Format(timeLayout))
		return nil
	}
	if src.Kind() == reflect.String && dst.Type() == reflect.TypeOf(time.Time{}) {
		str := src.String()
		if str == "" {
			// 如果是空字符串，跳过，不转换为时间
			return nil
		}
		t, err := time.Parse(timeLayout, src.String())
		if err != nil {
			return err
		}
		dst.Set(reflect.ValueOf(t))
		return nil
	}
	// 字符串转 *time.Time
	if src.Kind() == reflect.String && dst.Type() == reflect.TypeOf(&time.Time{}) {
		str := src.String()
		if str == "" {
			dst.Set(reflect.Zero(dst.Type())) // 赋 nil
			return nil
		}
		t, err := time.Parse(timeLayout, str)
		if err != nil {
			return err
		}
		dst.Set(reflect.ValueOf(&t))
		return nil
	}

	return fmt.Errorf("unsupported conversion: src type %v to dst type %v", src.Type(), dst.Type())
}

func isNumericKind(k reflect.Kind) bool {
	switch k {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64:
		return true
	default:
		return false
	}
}
