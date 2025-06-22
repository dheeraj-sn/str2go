# str2go

A powerful Go library and CLI tool for converting strings to any Go type with support for basic types, complex structures, slices, and maps.

## Features

- **Basic Type Conversion**: Convert strings to int, float, bool, time.Time, etc.
- **Complex Type Support**: Convert strings to structs, slices, maps
- **Flexible Input Formats**: Support for JSON, YAML, and custom formats
- **CLI Interface**: Command-line tool for quick conversions
- **Web API**: RESTful API for programmatic access
- **Extensible**: Easy to add custom type converters
- **Comprehensive Testing**: Full test coverage

## Installation

```bash
go install github.com/dheeraj-sn/str2go/cmd/str2go@latest
```

## Usage

### CLI Usage

```bash
# Convert string to int
str2go convert "42" --type int

# Convert JSON string to struct
str2go convert '{"name":"John","age":30}' --type "Person"

# Convert string to slice
str2go convert "1,2,3,4,5" --type "[]int" --delimiter ","

# Start web server
str2go serve --port 8080
```

### Library Usage

```go
package main

import (
    "fmt"
    "github.com/dheeraj-sn/str2go/converter"
)

func main() {
    // Convert string to int
    result, err := converter.Convert("42", "int")
    if err != nil {
        panic(err)
    }
    fmt.Printf("Result: %v (%T)\n", result, result)

    // Convert JSON to struct
    type Person struct {
        Name string `json:"name"`
        Age  int    `json:"age"`
    }
    
    jsonStr := `{"name":"John","age":30}`
    person, err := converter.Convert(jsonStr, "Person")
    if err != nil {
        panic(err)
    }
    fmt.Printf("Person: %+v\n", person)
}
```

### Web API

Start the server:
```bash
str2go serve --port 8080
```

API Endpoints:
- `POST /convert` - Convert string to specified type
- `GET /types` - List supported types

Example API call:
```bash
curl -X POST http://localhost:8080/convert \
  -H "Content-Type: application/json" \
  -d '{"value": "42", "type": "int"}'
```

## Supported Types

### Basic Types
- `int`, `int8`, `int16`, `int32`, `int64`
- `uint`, `uint8`, `uint16`, `uint32`, `uint64`
- `float32`, `float64`
- `bool`
- `string`
- `time.Time`
- `[]byte`

### Complex Types
- Structs (from JSON/YAML)
- Slices (`[]int`, `[]string`, etc.)
- Maps (`map[string]int`, etc.)
- Pointers (`*int`, `*string`, etc.)