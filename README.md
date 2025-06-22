# str2go

A Go library for converting strings to various Go types with a flexible type registry system. This library provides a clean, extensible architecture for string-to-type conversion with support for basic types, pointer types, and custom converters.

## Features

- **Type Registry System**: Flexible registry for managing type converters
- **Basic Type Support**: Convert strings to int, uint, float, bool, string, time.Time
- **Pointer Type Support**: Convert strings to pointer types (*int, *bool, etc.)
- **Extensible Architecture**: Easy to add custom type converters
- **Thread-Safe**: Concurrent access to converters and registries
- **Comprehensive Testing**: Full test coverage with benchmarks

## Architecture

The library is organized into several key packages:

- **`converter/`**: Contains individual type converters and a global converter map
- **`typeregistry/`**: Provides a flexible type registry system
- **`globalregistry/`**: Offers a global singleton registry with all converters
- **`model/`**: Defines the core interfaces and types

## Installation

```bash
go get github.com/dheeraj-sn/str2go
```

## Usage

### Using the Global Registry

The simplest way to use the library is through the global registry:

```go
package main

import (
    "fmt"
    "reflect"
    "github.com/dheeraj-sn/str2go/globalregistry"
)

func main() {
    // Get a converter for int type
    converter, exists := globalregistry.GetConverter(reflect.TypeOf(0))
    if !exists {
        panic("int converter not found")
    }

    // Convert string to int
    result, err := converter("42")
    if err != nil {
        panic(err)
    }

    fmt.Printf("Result: %v (%T)\n", result, result)
    // Output: Result: 42 (int)
}
```

### Using Type Registry

For more control, use the type registry directly:

```go
package main

import (
    "fmt"
    "reflect"
    "github.com/dheeraj-sn/str2go/typeregistry"
    "github.com/dheeraj-sn/str2go/converter"
)

func main() {
    // Create a new registry
    registry := typeregistry.NewTypeRegistry()

    // Register all default converters
    registry.RegisterAll(converter.GetConvertorMap())

    // Convert string to float64
    result, err := registry.Convert("3.14", reflect.TypeOf(0.0))
    if err != nil {
        panic(err)
    }

    fmt.Printf("Result: %v (%T)\n", result, result)
    // Output: Result: 3.14 (float64)
}
```

### Adding Custom Converters

You can easily add custom converters:

```go
package main

import (
    "fmt"
    "reflect"
    "strings"
    "github.com/dheeraj-sn/str2go/typeregistry"
)

func main() {
    registry := typeregistry.NewTypeRegistry()

    // Custom converter for string to uppercase
    uppercaseConverter := func(value string) (interface{}, error) {
        return strings.ToUpper(value), nil
    }

    // Register the custom converter
    registry.Register(reflect.TypeOf(""), uppercaseConverter)

    // Use the converter
    result, err := registry.Convert("hello world", reflect.TypeOf(""))
    if err != nil {
        panic(err)
    }

    fmt.Printf("Result: %v\n", result)
    // Output: Result: HELLO WORLD
}
```

## Supported Types

### Basic Types
- `int`, `int8`, `int16`, `int32`, `int64`
- `uint`, `uint8`, `uint16`, `uint32`, `uint64`
- `float32`, `float64`
- `bool`
- `string`
- `time.Time`

### Pointer Types
- `*int`, `*int8`, `*int16`, `*int32`, `*int64`
- `*uint`, `*uint8`, `*uint16`, `*uint32`, `*uint64`
- `*float32`, `*float64`
- `*bool`
- `*string`
- `*time.Time`

## API Reference

### Type Registry

```go
// Create a new registry
registry := typeregistry.NewTypeRegistry()

// Register a converter
registry.Register(targetType, converterFunc)

// Register multiple converters
registry.RegisterAll(convertersMap)

// Get a converter
converter, exists := registry.Get(targetType)

// Convert a string to a type
result, err := registry.Convert(value, targetType)

// Get all supported types
types := registry.GetSupportedTypes()
```

### Global Registry

```go
// Get a converter from global registry
converter, exists := globalregistry.GetConverter(targetType)
```

### Converter Function Signature

```go
type ConverterFunc func(value string) (interface{}, error)
```

## Testing

Run the test suite:

```bash
go test ./...
```

Run tests with coverage:

```bash
go test -cover ./...
```

Run benchmarks:

```bash
go test -bench=. ./...
```

## Contributing

1. Fork the repository
2. Create a feature branch
3. Add your changes with tests
4. Ensure all tests pass
5. Submit a pull request