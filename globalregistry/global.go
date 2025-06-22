package globalregistry

import (
	"reflect"
	"sync"

	"github.com/dheeraj-sn/str2go/converter"
	"github.com/dheeraj-sn/str2go/model"
	"github.com/dheeraj-sn/str2go/typeregistry"
)

var once sync.Once
var globalRegistry *typeregistry.TypeRegistry

func init() {
	getGlobalRegistry()
}

func getGlobalRegistry() *typeregistry.TypeRegistry {
	once.Do(func() {
		globalRegistry = typeregistry.NewTypeRegistry()

		// Get all registered converters and register them with the global registry
		globalRegistry.RegisterAll(converter.GetConvertorMap())
	})

	return globalRegistry
}

func GetConverter(targetType reflect.Type) (model.ConverterFunc, bool) {
	return getGlobalRegistry().Get(targetType)
}
