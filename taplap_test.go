package taplap

import (
	"reflect"
	"testing"
)

func TestTaplap(test *testing.T) {
	type TestData struct {
		A    string
		B    string
		C    bool
		UUID string
		Name string
		F    int64
		Meta struct {
			Lambda string
		}
	}
	data := TestData{}
	Fill(&data)
	if notInitialized(data) {
		test.Fatalf("not initialized values")
	}
}

func notInitialized(val interface{}) bool {
	value := reflect.ValueOf(val)
	return reflect.DeepEqual(
		reflect.Zero(value.Type()).Interface(),
		value.Interface())
}
