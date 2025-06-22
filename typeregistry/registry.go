package typeregistry

import (
	"fmt"
	"reflect"
)

// ConverterFunc represents a function that converts a string to a specific type
type ConverterFunc func(value string) (interface{}, error)

// TypeRegistry holds all registered type converters
type TypeRegistry struct {
	converters map[reflect.Type]ConverterFunc
}

// NewTypeRegistry creates a new type registry with default converters
func NewTypeRegistry() *TypeRegistry {
	registry := &TypeRegistry{
		converters: make(map[reflect.Type]ConverterFunc),
	}
	return registry
}

// Register adds a new type converter to the registry
func (tr *TypeRegistry) Register(typeName reflect.Type, converter ConverterFunc) {
	tr.converters[typeName] = converter
}

// Get retrieves a converter for the given type
func (tr *TypeRegistry) Get(typeName reflect.Type) (ConverterFunc, bool) {
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
