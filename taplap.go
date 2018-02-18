package taplap

import (
	"reflect"
)

func Fill(val interface{}) {
	value := reflect.ValueOf(val)
	if value.Kind() != reflect.Struct {
		return
	}
	for fieldIndex := 0; fieldIndex < value.NumField(); fieldIndex++ {
		field := value.Field(fieldIndex)
		switch field.Kind() {
		case reflect.String:
			field.Set(reflect.ValueOf("foo bar"))
		default:
		}
	}
}
