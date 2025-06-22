package main

import (
	"fmt"
	"log"

	"github.com/dheeraj-sn/str2go/converter"
)

func main() {
	fmt.Println("str2go Library Usage Examples")
	fmt.Println("=============================")

	// Example 1: Basic type conversions
	fmt.Println("\n1. Basic Type Conversions:")

	// String to int
	if result, err := converter.Convert("42", "int"); err == nil {
		fmt.Printf("  \"42\" -> int: %v (%T)\n", result, result)
	}

	// String to float
	if result, err := converter.Convert("3.14159", "float64"); err == nil {
		fmt.Printf("  \"3.14159\" -> float64: %v (%T)\n", result, result)
	}

	// String to bool
	if result, err := converter.Convert("true", "bool"); err == nil {
		fmt.Printf("  \"true\" -> bool: %v (%T)\n", result, result)
	}

	// String to time
	if result, err := converter.Convert("2023-01-02T15:04:05Z", "time.Time"); err == nil {
		fmt.Printf("  \"2023-01-02T15:04:05Z\" -> time.Time: %v (%T)\n", result, result)
	}

	// Example 2: Slice conversions
	fmt.Println("\n2. Slice Conversions:")

	// JSON array to slice
	if result, err := converter.Convert(`["apple", "banana", "cherry"]`, "[]string"); err == nil {
		fmt.Printf("  JSON array -> []string: %v (%T)\n", result, result)
	}

	// Comma-separated values to slice
	if result, err := converter.ConvertWithDelimiter("1,2,3,4,5", "[]int", ","); err == nil {
		fmt.Printf("  CSV -> []int: %v (%T)\n", result, result)
	}

	// Custom delimiter
	if result, err := converter.ConvertWithDelimiter("red;green;blue", "[]string", ";"); err == nil {
		fmt.Printf("  Custom delimiter -> []string: %v (%T)\n", result, result)
	}

	// Example 3: Map conversions
	fmt.Println("\n3. Map Conversions:")

	// JSON object to map
	if result, err := converter.Convert(`{"name": "John", "age": 30}`, "map[string]int"); err == nil {
		fmt.Printf("  JSON object -> map[string]int: %v (%T)\n", result, result)
	}

	// Example 4: Pointer conversions
	fmt.Println("\n4. Pointer Conversions:")

	// String to pointer
	if result, err := converter.Convert("42", "*int"); err == nil {
		fmt.Printf("  \"42\" -> *int: %v (%T)\n", result, result)
	}

	// Example 5: Error handling
	fmt.Println("\n5. Error Handling:")

	if result, err := converter.Convert("not_a_number", "int"); err != nil {
		fmt.Printf("  Error converting \"not_a_number\" to int: %v\n", err)
	} else {
		fmt.Printf("  Result: %v\n", result)
	}

	// Example 6: Get supported types
	fmt.Println("\n6. Supported Types:")
	types := converter.GetSupportedTypes()
	fmt.Printf("  Total supported types: %d\n", len(types))
	fmt.Printf("  First 5 types: %v\n", types[:5])

	// Example 7: Complex JSON to struct-like conversion
	fmt.Println("\n7. Complex JSON Conversion:")

	jsonData := `{
		"name": "Alice",
		"age": 25,
		"active": true,
		"tags": ["developer", "golang"],
		"metadata": {
			"department": "engineering",
			"level": "senior"
		}
	}`

	if result, err := converter.Convert(jsonData, "Person"); err == nil {
		fmt.Printf("  Complex JSON -> struct-like: %v (%T)\n", result, result)
	} else {
		log.Printf("Error: %v", err)
	}

	fmt.Println("\nExamples completed!")
}
