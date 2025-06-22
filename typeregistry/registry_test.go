package typeregistry

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/dheeraj-sn/str2go/model"
)

// TestNewTypeRegistry tests the creation of a new type registry
func TestNewTypeRegistry(t *testing.T) {
	registry := NewTypeRegistry()

	if registry == nil {
		t.Fatal("NewTypeRegistry() returned nil")
	}

	if registry.converters == nil {
		t.Fatal("converters map was not initialized")
	}

	if len(registry.converters) != 0 {
		t.Fatal("new registry should have no converters")
	}
}

// TestRegister tests registering converters
func TestRegister(t *testing.T) {
	registry := NewTypeRegistry()

	// Test registering a converter
	testConverter := func(value string) (interface{}, error) {
		return value, nil
	}

	registry.Register(reflect.TypeOf(""), testConverter)

	if len(registry.converters) != 1 {
		t.Fatalf("expected 1 converter, got %d", len(registry.converters))
	}

	// Test registering another converter
	intConverter := func(value string) (interface{}, error) {
		return 42, nil
	}

	registry.Register(reflect.TypeOf(0), intConverter)

	if len(registry.converters) != 2 {
		t.Fatalf("expected 2 converters, got %d", len(registry.converters))
	}
}

// TestGet tests retrieving converters
func TestGet(t *testing.T) {
	registry := NewTypeRegistry()

	// Test getting non-existent converter
	converter, exists := registry.Get(reflect.TypeOf(""))
	if exists {
		t.Fatal("should not find converter for unregistered type")
	}
	if converter != nil {
		t.Fatal("converter should be nil for unregistered type")
	}

	// Test getting existing converter
	testConverter := func(value string) (interface{}, error) {
		return value, nil
	}

	registry.Register(reflect.TypeOf(""), testConverter)

	converter, exists = registry.Get(reflect.TypeOf(""))
	if !exists {
		t.Fatal("should find converter for registered type")
	}
	if converter == nil {
		t.Fatal("converter should not be nil for registered type")
	}

	// Test that the converter works
	result, err := converter("test")
	if err != nil {
		t.Fatalf("converter should not return error: %v", err)
	}
	if result != "test" {
		t.Fatalf("expected 'test', got %v", result)
	}
}

// TestConvert tests the Convert method
func TestConvert(t *testing.T) {
	registry := NewTypeRegistry()

	// Test converting with unregistered type
	_, err := registry.Convert("test", reflect.TypeOf(""))
	if err == nil {
		t.Fatal("should return error for unregistered type")
	}

	// Test converting with registered type
	testConverter := func(value string) (interface{}, error) {
		return "converted: " + value, nil
	}

	registry.Register(reflect.TypeOf(""), testConverter)

	result, err := registry.Convert("test", reflect.TypeOf(""))
	if err != nil {
		t.Fatalf("should not return error: %v", err)
	}

	expected := "converted: test"
	if result != expected {
		t.Fatalf("expected '%s', got '%v'", expected, result)
	}
}

// TestConvertWithError tests conversion when converter returns error
func TestConvertWithError(t *testing.T) {
	registry := NewTypeRegistry()

	errorConverter := func(value string) (interface{}, error) {
		return nil, fmt.Errorf("conversion error")
	}

	registry.Register(reflect.TypeOf(""), errorConverter)

	_, err := registry.Convert("test", reflect.TypeOf(""))
	if err == nil {
		t.Fatal("should return error from converter")
	}
	if err.Error() != "conversion error" {
		t.Fatalf("expected 'conversion error', got '%s'", err.Error())
	}
}

// TestGetSupportedTypes tests getting all supported types
func TestGetSupportedTypes(t *testing.T) {
	registry := NewTypeRegistry()

	// Test empty registry
	types := registry.GetSupportedTypes()
	if len(types) != 0 {
		t.Fatalf("expected 0 types, got %d", len(types))
	}

	// Test with registered types
	stringType := reflect.TypeOf("")
	intType := reflect.TypeOf(0)

	registry.Register(stringType, func(value string) (interface{}, error) { return value, nil })
	registry.Register(intType, func(value string) (interface{}, error) { return 0, nil })

	types = registry.GetSupportedTypes()
	if len(types) != 2 {
		t.Fatalf("expected 2 types, got %d", len(types))
	}

	// Check that both types are present
	foundString := false
	foundInt := false

	for _, t := range types {
		if t == stringType {
			foundString = true
		}
		if t == intType {
			foundInt = true
		}
	}

	if !foundString {
		t.Fatal("string type not found in supported types")
	}
	if !foundInt {
		t.Fatal("int type not found in supported types")
	}
}

// TestRegisterOverwrite tests that registering the same type overwrites the previous converter
func TestRegisterOverwrite(t *testing.T) {
	registry := NewTypeRegistry()

	// Register first converter
	firstConverter := func(value string) (interface{}, error) {
		return "first", nil
	}
	registry.Register(reflect.TypeOf(""), firstConverter)

	// Register second converter for same type
	secondConverter := func(value string) (interface{}, error) {
		return "second", nil
	}
	registry.Register(reflect.TypeOf(""), secondConverter)

	// Should only have one converter
	if len(registry.converters) != 1 {
		t.Fatalf("expected 1 converter, got %d", len(registry.converters))
	}

	// Should use the second converter
	result, err := registry.Convert("test", reflect.TypeOf(""))
	if err != nil {
		t.Fatalf("should not return error: %v", err)
	}
	if result != "second" {
		t.Fatalf("expected 'second', got '%v'", result)
	}
}

// TestConvertWithComplexTypes tests conversion with more complex types
func TestConvertWithComplexTypes(t *testing.T) {
	registry := NewTypeRegistry()

	// Test with slice type
	sliceType := reflect.TypeOf([]string{})
	sliceConverter := func(value string) (interface{}, error) {
		return []string{value}, nil
	}
	registry.Register(sliceType, sliceConverter)

	result, err := registry.Convert("test", sliceType)
	if err != nil {
		t.Fatalf("should not return error: %v", err)
	}

	slice, ok := result.([]string)
	if !ok {
		t.Fatal("result should be []string")
	}
	if len(slice) != 1 || slice[0] != "test" {
		t.Fatalf("expected ['test'], got %v", slice)
	}

	// Test with map type
	mapType := reflect.TypeOf(map[string]int{})
	mapConverter := func(value string) (interface{}, error) {
		return map[string]int{value: 1}, nil
	}
	registry.Register(mapType, mapConverter)

	result, err = registry.Convert("test", mapType)
	if err != nil {
		t.Fatalf("should not return error: %v", err)
	}

	m, ok := result.(map[string]int)
	if !ok {
		t.Fatal("result should be map[string]int")
	}
	if m["test"] != 1 {
		t.Fatalf("expected map[test:1], got %v", m)
	}
}

// TestConvertWithPointerTypes tests conversion with pointer types
func TestConvertWithPointerTypes(t *testing.T) {
	registry := NewTypeRegistry()

	// Test with pointer to string
	stringPtrType := reflect.TypeOf((*string)(nil))
	stringPtrConverter := func(value string) (interface{}, error) {
		return &value, nil
	}
	registry.Register(stringPtrType, stringPtrConverter)

	result, err := registry.Convert("test", stringPtrType)
	if err != nil {
		t.Fatalf("should not return error: %v", err)
	}

	strPtr, ok := result.(*string)
	if !ok {
		t.Fatal("result should be *string")
	}
	if *strPtr != "test" {
		t.Fatalf("expected 'test', got '%s'", *strPtr)
	}
}

// Benchmark tests for performance
func BenchmarkRegister(b *testing.B) {
	registry := NewTypeRegistry()
	stringType := reflect.TypeOf("")
	converter := func(value string) (interface{}, error) { return value, nil }

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		registry.Register(stringType, converter)
	}
}

func BenchmarkGet(b *testing.B) {
	registry := NewTypeRegistry()
	stringType := reflect.TypeOf("")
	converter := func(value string) (interface{}, error) { return value, nil }
	registry.Register(stringType, converter)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		registry.Get(stringType)
	}
}

func BenchmarkConvert(b *testing.B) {
	registry := NewTypeRegistry()
	stringType := reflect.TypeOf("")
	converter := func(value string) (interface{}, error) { return value, nil }
	registry.Register(stringType, converter)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		registry.Convert("test", stringType)
	}
}

// TestRegisterAll tests registering multiple converters at once
func TestRegisterAll(t *testing.T) {
	registry := NewTypeRegistry()

	// Create a map of converters to register
	converters := map[reflect.Type]model.ConverterFunc{
		reflect.TypeOf(""): func(value string) (interface{}, error) {
			return "string: " + value, nil
		},
		reflect.TypeOf(0): func(value string) (interface{}, error) {
			return 42, nil
		},
		reflect.TypeOf(0.0): func(value string) (interface{}, error) {
			return 3.14, nil
		},
	}

	// Register all converters at once
	registry.RegisterAll(converters)

	// Verify all converters were registered
	if len(registry.converters) != 3 {
		t.Fatalf("expected 3 converters, got %d", len(registry.converters))
	}

	// Test each registered converter
	stringType := reflect.TypeOf("")
	intType := reflect.TypeOf(0)
	floatType := reflect.TypeOf(0.0)

	// Test string converter
	result, err := registry.Convert("test", stringType)
	if err != nil {
		t.Fatalf("string converter should not return error: %v", err)
	}
	if result != "string: test" {
		t.Fatalf("expected 'string: test', got '%v'", result)
	}

	// Test int converter
	result, err = registry.Convert("test", intType)
	if err != nil {
		t.Fatalf("int converter should not return error: %v", err)
	}
	if result != 42 {
		t.Fatalf("expected 42, got %v", result)
	}

	// Test float converter
	result, err = registry.Convert("test", floatType)
	if err != nil {
		t.Fatalf("float converter should not return error: %v", err)
	}
	if result != 3.14 {
		t.Fatalf("expected 3.14, got %v", result)
	}
}

// TestRegisterAllOverwrite tests that RegisterAll overwrites existing converters
func TestRegisterAllOverwrite(t *testing.T) {
	registry := NewTypeRegistry()

	// Register initial converter
	initialConverter := func(value string) (interface{}, error) {
		return "initial", nil
	}
	registry.Register(reflect.TypeOf(""), initialConverter)

	// Verify initial registration
	if len(registry.converters) != 1 {
		t.Fatalf("expected 1 converter after initial registration, got %d", len(registry.converters))
	}

	// Create converters map with same type but different converter
	converters := map[reflect.Type]model.ConverterFunc{
		reflect.TypeOf(""): func(value string) (interface{}, error) {
			return "overwritten", nil
		},
		reflect.TypeOf(0): func(value string) (interface{}, error) {
			return 100, nil
		},
	}

	// Register all converters (should overwrite existing and add new)
	registry.RegisterAll(converters)

	// Verify we have 2 converters now
	if len(registry.converters) != 2 {
		t.Fatalf("expected 2 converters after RegisterAll, got %d", len(registry.converters))
	}

	// Test that the string converter was overwritten
	result, err := registry.Convert("test", reflect.TypeOf(""))
	if err != nil {
		t.Fatalf("converter should not return error: %v", err)
	}
	if result != "overwritten" {
		t.Fatalf("expected 'overwritten', got '%v'", result)
	}

	// Test that the new int converter works
	result, err = registry.Convert("test", reflect.TypeOf(0))
	if err != nil {
		t.Fatalf("int converter should not return error: %v", err)
	}
	if result != 100 {
		t.Fatalf("expected 100, got %v", result)
	}
}

// TestRegisterAllEmptyMap tests RegisterAll with an empty map
func TestRegisterAllEmptyMap(t *testing.T) {
	registry := NewTypeRegistry()

	// Register some initial converters
	registry.Register(reflect.TypeOf(""), func(value string) (interface{}, error) { return value, nil })
	registry.Register(reflect.TypeOf(0), func(value string) (interface{}, error) { return 0, nil })

	initialCount := len(registry.converters)

	// Register empty map
	registry.RegisterAll(map[reflect.Type]model.ConverterFunc{})

	// Should not change the count
	if len(registry.converters) != initialCount {
		t.Fatalf("expected %d converters after empty RegisterAll, got %d", initialCount, len(registry.converters))
	}
}
