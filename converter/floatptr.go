package converter

import (
	"reflect"
	"strconv"

	"github.com/dheeraj-sn/str2go/globalregistry"
)

func init() {
	var float32Var float32
	var float64Var float64

	globalregistry.GetGlobalRegistry().Register(reflect.TypeOf(&float32Var), StringToFloat32Ptr)
	globalregistry.GetGlobalRegistry().Register(reflect.TypeOf(&float64Var), StringToFloat64Ptr)
}

func StringToFloat32Ptr(value string) (interface{}, error) {
	floatValue, err := strconv.ParseFloat(value, 32)
	if err != nil {
		return nil, err
	}
	return &floatValue, nil
}

func StringToFloat64Ptr(value string) (interface{}, error) {
	floatValue, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return nil, err
	}
	return &floatValue, nil
}
