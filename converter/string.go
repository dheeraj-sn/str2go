package converter

import (
	"reflect"
)

func init() {
	var stringVar string
	registerConverter(reflect.TypeOf(stringVar), StringToString)
	registerConverter(reflect.TypeOf(&stringVar), StringToStringPtr)
}

func StringToString(value string) (interface{}, error) {
	return value, nil
}

func StringToStringPtr(value string) (interface{}, error) {
	return &value, nil
}
