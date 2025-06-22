package converter

import (
	"reflect"
	"strconv"
)

func init() {
	var intVar int
	var int8Var int8
	var int16Var int16
	var int32Var int32
	var int64Var int64

	registerConverter(reflect.TypeOf(&intVar), StringToIntPtr)
	registerConverter(reflect.TypeOf(&int8Var), StringToInt8Ptr)
	registerConverter(reflect.TypeOf(&int16Var), StringToInt16Ptr)
	registerConverter(reflect.TypeOf(&int32Var), StringToInt32Ptr)
	registerConverter(reflect.TypeOf(&int64Var), StringToInt64Ptr)
}

func StringToIntPtr(value string) (interface{}, error) {
	intValue, err := strconv.Atoi(value)
	if err != nil {
		return nil, err
	}
	return &intValue, nil
}

func StringToInt8Ptr(value string) (interface{}, error) {
	intValue, err := strconv.ParseInt(value, 10, 8)
	if err != nil {
		return nil, err
	}
	return &intValue, nil
}

func StringToInt16Ptr(value string) (interface{}, error) {
	intValue, err := strconv.ParseInt(value, 10, 16)
	if err != nil {
		return nil, err
	}
	return &intValue, nil
}

func StringToInt32Ptr(value string) (interface{}, error) {
	intValue, err := strconv.ParseInt(value, 10, 32)
	if err != nil {
		return nil, err
	}
	return &intValue, nil
}

func StringToInt64Ptr(value string) (interface{}, error) {
	intValue, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return nil, err
	}
	return &intValue, nil
}
