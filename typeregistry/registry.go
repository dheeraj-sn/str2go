package typeregistry

import (
	"fmt"
	"reflect"

	"github.com/dheeraj-sn/str2go/model"
)

// TypeRegistry holds all registered type converters
type TypeRegistry struct {
	converters map[reflect.Type]model.ConverterFunc
}

// NewTypeRegistry creates a new type registry with default converters
func NewTypeRegistry() *TypeRegistry {
	registry := &TypeRegistry{
		converters: make(map[reflect.Type]model.ConverterFunc),
	}
	return registry
}

// Register adds a new type converter to the registry
func (tr *TypeRegistry) Register(typeName reflect.Type, converter model.ConverterFunc) {
	tr.converters[typeName] = converter
}

func (tr *TypeRegistry) RegisterAll(converters map[reflect.Type]model.ConverterFunc) {
	for typeName, converter := range converters {
		tr.Register(typeName, converter)
	}
}

// Get retrieves a converter for the given type
func (tr *TypeRegistry) Get(typeName reflect.Type) (model.ConverterFunc, bool) {
	converter, exists := tr.converters[typeName]
	return converter, exists
}

// Convert uses the registry to convert a string to the specified type
func (tr *TypeRegistry) Convert(value string, targetType reflect.Type) (interface{}, error) {
	// Check for exact match first
	if converter, exists := tr.Get(targetType); exists {
		return converter(value)
	}

	return nil, fmt.Errorf("no converter registered for type: %s", targetType)
}

// GetSupportedTypes returns all registered types
func (tr *TypeRegistry) GetSupportedTypes() []reflect.Type {
	types := make([]reflect.Type, 0, len(tr.converters))
	for typeName := range tr.converters {
		types = append(types, typeName)
	}
	return types
}
