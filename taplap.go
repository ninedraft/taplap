package taplap

import (
	"reflect"
	"strings"
)

func Fill(val interface{}) {
	value := reflect.ValueOf(val).Elem()
	if value.Kind() != reflect.Struct {
		return
	}
	switch value.Kind() {
	case reflect.Struct:
		fillStruct(&value)
	default:
	}
}

func fillStruct(value *reflect.Value) {
	tt := value.Type()
	for fieldIndex := 0; fieldIndex < value.NumField(); fieldIndex++ {
		field := value.Field(fieldIndex)
		name := tt.Field(fieldIndex).Name
		if field.IsValid() && field.CanSet() {
			switch field.Kind() {
			case reflect.String:
				matchString(name, &field)
			case reflect.Struct:
				fillStruct(&field)
			case reflect.Bool:
				field.SetBool(true)
			default:
			}
		}
	}
}

func matchString(name string, field *reflect.Value) {
	name = strings.ToLower(name)
	match := func(name string, elems ...string) bool {
		for _, elem := range elems {
			if strings.Contains(name, elem) {
				return true
			}
		}
		return false
	}
	switch {
	case match(name, "uuid"):
		field.SetString("uuid")
	case match(name, "username", "firstname", "secondname", "name"):
		field.SetString("merlin")
	default:
		field.SetString("foobar")
	}
}
