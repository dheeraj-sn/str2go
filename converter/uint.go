package converter

import (
	"reflect"
	"strconv"
)

func init() {
	registerConverter(reflect.TypeOf(uint(0)), StringToUint)
	registerConverter(reflect.TypeOf(uint8(0)), StringToUint8)
	registerConverter(reflect.TypeOf(uint16(0)), StringToUint16)
	registerConverter(reflect.TypeOf(uint32(0)), StringToUint32)
	registerConverter(reflect.TypeOf(uint64(0)), StringToUint64)
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
