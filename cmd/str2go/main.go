package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/dheeraj-sn/str2go/converter"
	"github.com/dheeraj-sn/str2go/server"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "str2go",
	Short: "Convert strings to Go types",
	Long: `A powerful tool for converting strings to various Go types.
Supports basic types, slices, maps, and complex structures.`,
}

var convertCmd = &cobra.Command{
	Use:   "convert [value]",
	Short: "Convert a string value to a Go type",
	Long: `Convert a string value to the specified Go type.
Examples:
  str2go convert "42" --type int
  str2go convert '{"name":"John","age":30}' --type "Person"
  str2go convert "1,2,3,4,5" --type "[]int" --delimiter ","`,
	Args: cobra.ExactArgs(1),
	Run:  runConvert,
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the web server",
	Long:  "Start a web server to provide REST API for string conversion",
	Run:   runServe,
}

var (
	targetType string
	delimiter  string
	port       int
	pretty     bool
)

func init() {
	convertCmd.Flags().StringVarP(&targetType, "type", "t", "", "Target Go type (required)")
	convertCmd.Flags().StringVarP(&delimiter, "delimiter", "d", ",", "Delimiter for slice conversion")
	convertCmd.Flags().BoolVarP(&pretty, "pretty", "p", false, "Pretty print JSON output")
	convertCmd.MarkFlagRequired("type")

	serveCmd.Flags().IntVarP(&port, "port", "p", 8080, "Port to listen on")

	rootCmd.AddCommand(convertCmd)
	rootCmd.AddCommand(serveCmd)
}

func runConvert(cmd *cobra.Command, args []string) {
	value := args[0]

	var result interface{}
	var err error

	if delimiter != "," && targetType != "" {
		result, err = converter.ConvertWithDelimiter(value, targetType, delimiter)
	} else {
		result, err = converter.Convert(value, targetType)
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	// Output the result
	if pretty {
		jsonData, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error marshaling result: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(string(jsonData))
	} else {
		fmt.Printf("%v\n", result)
	}
}

func runServe(cmd *cobra.Command, args []string) {
	fmt.Printf("Starting server on port %d...\n", port)
	if err := server.StartServer(port); err != nil {
		fmt.Fprintf(os.Stderr, "Server error: %v\n", err)
		os.Exit(1)
	}
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
