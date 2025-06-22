package converter

import (
	"reflect"
	"strconv"

	"github.com/dheeraj-sn/str2go/globalregistry"
)

func init() {
	globalregistry.GetGlobalRegistry().Register(reflect.TypeOf(float32(0)), StringToFloat32)
	globalregistry.GetGlobalRegistry().Register(reflect.TypeOf(float64(0)), StringToFloat64)
}

func StringToFloat32(value string) (interface{}, error) {
	return strconv.ParseFloat(value, 32)
}

func StringToFloat64(value string) (interface{}, error) {
	return strconv.ParseFloat(value, 64)
}
