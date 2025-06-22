package converter

import (
	"reflect"
	"strconv"
)

func init() {
	registerConverter(reflect.TypeOf(0), StringToInt)
	registerConverter(reflect.TypeOf(int8(0)), StringToInt8)
	registerConverter(reflect.TypeOf(int16(0)), StringToInt16)
	registerConverter(reflect.TypeOf(int32(0)), StringToInt32)
	registerConverter(reflect.TypeOf(int64(0)), StringToInt64)
}

func StringToInt(value string) (interface{}, error) {
	return strconv.Atoi(value)
}

func StringToInt8(value string) (interface{}, error) {
	return strconv.ParseInt(value, 10, 8)
}

func StringToInt16(value string) (interface{}, error) {
	return strconv.ParseInt(value, 10, 16)
}

func StringToInt32(value string) (interface{}, error) {
	return strconv.ParseInt(value, 10, 32)
}

func StringToInt64(value string) (interface{}, error) {
	return strconv.ParseInt(value, 10, 64)
}
