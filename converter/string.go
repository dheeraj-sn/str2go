package converter

import (
	"reflect"

	"github.com/dheeraj-sn/str2go/globalregistry"
)

func init() {
	var stringVar string
	globalregistry.GetGlobalRegistry().Register(reflect.TypeOf(stringVar), StringToString)
	globalregistry.GetGlobalRegistry().Register(reflect.TypeOf(&stringVar), StringToStringPtr)
}

func StringToString(value string) (interface{}, error) {
	return value, nil
}

func StringToStringPtr(value string) (interface{}, error) {
	return &value, nil
}
