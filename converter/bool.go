package converter

import (
	"reflect"
	"strconv"

	"github.com/dheeraj-sn/str2go/globalregistry"
)

func init() {
	var boolVar bool
	globalregistry.GetGlobalRegistry().Register(reflect.TypeOf(false), StringToBool)
	globalregistry.GetGlobalRegistry().Register(reflect.TypeOf(&boolVar), StringToBoolPtr)
}

func StringToBool(value string) (interface{}, error) {
	return strconv.ParseBool(value)
}

func StringToBoolPtr(value string) (interface{}, error) {
	boolValue, err := strconv.ParseBool(value)
	if err != nil {
		return nil, err
	}
	return &boolValue, nil
}
