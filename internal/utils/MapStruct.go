package utils

import "reflect"

// MapStruct 通用的结构体映射
// src 要复制的源结构体
// dst 要复制到的目标结构体
func MapStruct(src, dst interface{}) error {
	sVal := reflect.ValueOf(src)
	dVal := reflect.ValueOf(dst).Elem()

	if sVal.Kind() == reflect.Ptr {
		sVal = sVal.Elem()
	}

	for i := 0; i < sVal.NumField(); i++ {
		field := sVal.Type().Field(i).Name
		if dField := dVal.FieldByName(field); dField.IsValid() && dField.CanSet() {
			dField.Set(sVal.Field(i))
		}
	}
	return nil
}
