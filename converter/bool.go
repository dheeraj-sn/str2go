package converter

import (
	"reflect"
	"strconv"
)

func init() {
	var boolVar bool
	registerConverter(reflect.TypeOf(false), StringToBool)
	registerConverter(reflect.TypeOf(&boolVar), StringToBoolPtr)
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
