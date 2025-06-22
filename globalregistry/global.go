package globalregistry

import (
	"sync"

	"github.com/dheeraj-sn/str2go/typeregistry"
)

var once sync.Once
var globalRegistry *typeregistry.TypeRegistry

func init() {
	once.Do(func() {
		globalRegistry = typeregistry.NewTypeRegistry()
	})
}

func getGlobalRegistry() *typeregistry.TypeRegistry {
	once.Do(func() {
		globalRegistry = typeregistry.NewTypeRegistry()
	})
	return globalRegistry
}

func GetGlobalRegistry() *typeregistry.TypeRegistry {
	return getGlobalRegistry()
}
