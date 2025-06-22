package converter

import (
	"reflect"
	"sync"
	"testing"
)

func TestGetConvertorMap(t *testing.T) {
	// Reset the map for testing
	once = sync.Once{}
	convertorMap = nil

	// Test that getConvertorMap returns a non-nil map
	result := getConvertorMap()
	if result == nil {
		t.Error("getConvertorMap should return a non-nil map")
	}

	// Test that subsequent calls return the same map by mutating and checking
	result2 := getConvertorMap()
	type testType struct{}
	key := reflect.TypeOf(testType{})
	val := func(string) (interface{}, error) { return "ok", nil }
	result[key] = val
	if result2[key] == nil {
		t.Error("getConvertorMap should return the same map instance on subsequent calls (mutation not visible)")
	}
}

func TestGetConvertorMapPublic(t *testing.T) {
	// Reset the map for testing
	once = sync.Once{}
	convertorMap = nil

	// Test the public GetConvertorMap function
	result := GetConvertorMap()
	if result == nil {
		t.Error("GetConvertorMap should return a non-nil map")
	}

	// Test that it returns the same map as the private function by mutating and checking
	expected := getConvertorMap()
	type testType struct{}
	key := reflect.TypeOf(testType{})
	val := func(string) (interface{}, error) { return "ok", nil }
	result[key] = val
	if expected[key] == nil {
		t.Error("GetConvertorMap should return the same map as getConvertorMap (mutation not visible)")
	}
}

func TestRegisterConverter(t *testing.T) {
	// Reset the map for testing
	once = sync.Once{}
	convertorMap = nil

	// Create a test converter function
	testConverter := func(value string) (interface{}, error) {
		return "test", nil
	}

	// Test registering a converter
	testType := reflect.TypeOf("")
	registerConverter(testType, testConverter)

	// Verify the converter was registered
	converterMap := getConvertorMap()
	if converterMap[testType] == nil {
		t.Error("registerConverter should add the converter to the map")
	}

	// Test that the registered converter works
	result, err := converterMap[testType]("test")
	if err != nil {
		t.Errorf("Registered converter should work without error: %v", err)
	}
	if result != "test" {
		t.Errorf("Expected 'test', got %v", result)
	}
}

func TestConverterMapConcurrency(t *testing.T) {
	// Reset the map for testing
	once = sync.Once{}
	convertorMap = nil

	// Test that multiple goroutines can safely access the map
	done := make(chan bool, 10)
	for i := 0; i < 10; i++ {
		go func() {
			GetConvertorMap()
			done <- true
		}()
	}

	// Wait for all goroutines to complete
	for i := 0; i < 10; i++ {
		<-done
	}

	// Verify the map is still accessible
	result := GetConvertorMap()
	if result == nil {
		t.Error("Map should still be accessible after concurrent access")
	}
}
