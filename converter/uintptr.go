package converter

import (
	"reflect"
	"strconv"

	"github.com/dheeraj-sn/str2go/globalregistry"
)

func init() {
	var uintVar uint
	var uint8Var uint8
	var uint16Var uint16
	var uint32Var uint32
	var uint64Var uint64

	globalregistry.GetGlobalRegistry().Register(reflect.TypeOf(&uintVar), StringToUintPtr)
	globalregistry.GetGlobalRegistry().Register(reflect.TypeOf(&uint8Var), StringToUint8Ptr)
	globalregistry.GetGlobalRegistry().Register(reflect.TypeOf(&uint16Var), StringToUint16Ptr)
	globalregistry.GetGlobalRegistry().Register(reflect.TypeOf(&uint32Var), StringToUint32Ptr)
	globalregistry.GetGlobalRegistry().Register(reflect.TypeOf(&uint64Var), StringToUint64Ptr)
}

func StringToUintPtr(value string) (interface{}, error) {
	uintValue, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return nil, err
	}
	return &uintValue, nil
}

func StringToUint8Ptr(value string) (interface{}, error) {
	uintValue, err := strconv.ParseUint(value, 10, 8)
	if err != nil {
		return nil, err
	}
	return &uintValue, nil
}

func StringToUint16Ptr(value string) (interface{}, error) {
	uintValue, err := strconv.ParseUint(value, 10, 16)
	if err != nil {
		return nil, err
	}
	return &uintValue, nil
}

func StringToUint32Ptr(value string) (interface{}, error) {
	uintValue, err := strconv.ParseUint(value, 10, 32)
	if err != nil {
		return nil, err
	}
	return &uintValue, nil
}

func StringToUint64Ptr(value string) (interface{}, error) {
	uintValue, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return nil, err
	}
	return &uintValue, nil
}
