package model

// ConverterFunc represents a function that converts a string to a specific type
type ConverterFunc func(value string) (interface{}, error)
