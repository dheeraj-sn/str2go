package converter

import (
	"reflect"
	"sync"

	"github.com/dheeraj-sn/str2go/model"
)

var (
	once         sync.Once
	convertorMap map[reflect.Type]model.ConverterFunc
)

func init() {
	getConvertorMap()
}

func getConvertorMap() map[reflect.Type]model.ConverterFunc {
	once.Do(func() {
		convertorMap = make(map[reflect.Type]model.ConverterFunc)
	})
	return convertorMap
}

func GetConvertorMap() map[reflect.Type]model.ConverterFunc {
	return getConvertorMap()
}

func registerConverter(targetType reflect.Type, converter model.ConverterFunc) {
	getConvertorMap()[targetType] = converter
}
