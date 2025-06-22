package converter

import (
	"reflect"
	"strconv"

	"github.com/dheeraj-sn/str2go/globalregistry"
)

func init() {
	globalregistry.GetGlobalRegistry().Register(reflect.TypeOf(uint(0)), StringToUint)
	globalregistry.GetGlobalRegistry().Register(reflect.TypeOf(uint8(0)), StringToUint8)
	globalregistry.GetGlobalRegistry().Register(reflect.TypeOf(uint16(0)), StringToUint16)
	globalregistry.GetGlobalRegistry().Register(reflect.TypeOf(uint32(0)), StringToUint32)
	globalregistry.GetGlobalRegistry().Register(reflect.TypeOf(uint64(0)), StringToUint64)
}

func StringToUint(value string) (interface{}, error) {
	return strconv.ParseUint(value, 10, 64)
}

func StringToUint8(value string) (interface{}, error) {
	return strconv.ParseUint(value, 10, 8)
}

func StringToUint16(value string) (interface{}, error) {
	return strconv.ParseUint(value, 10, 16)
}

func StringToUint32(value string) (interface{}, error) {
	return strconv.ParseUint(value, 10, 32)
}

func StringToUint64(value string) (interface{}, error) {
	return strconv.ParseUint(value, 10, 64)
}
