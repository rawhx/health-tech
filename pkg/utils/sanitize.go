package utils

import (
	"reflect"

	"github.com/microcosm-cc/bluemonday"
)

func SanitizePlain(input string) string {
	p := bluemonday.StrictPolicy()
	return p.Sanitize(input)
}

func SanitizeStruct(v any) {
	val := reflect.ValueOf(v)

	if val.Kind() != reflect.Ptr || val.IsNil() {
		return
	}

	elem := val.Elem()
	if elem.Kind() != reflect.Struct {
		return
	}

	for i := 0; i < elem.NumField(); i++ {
		field := elem.Field(i)
		if !field.CanSet() {
			continue
		}

		switch field.Kind() {
		case reflect.String:
			field.SetString(SanitizePlain(field.String()))

		case reflect.Struct:
			SanitizeStruct(field.Addr().Interface())

		case reflect.Ptr:
			if !field.IsNil() && field.Elem().Kind() == reflect.Struct {
				SanitizeStruct(field.Interface())
			}
		}
	}
}
