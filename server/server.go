package server

import (
	"net/http"
	"strconv"

	"github.com/dheeraj-sn/str2go/converter"
	"github.com/gin-gonic/gin"
)

// ConvertRequest represents the request body for conversion
type ConvertRequest struct {
	Value     string `json:"value" binding:"required"`
	Type      string `json:"type" binding:"required"`
	Delimiter string `json:"delimiter,omitempty"`
}

// ConvertResponse represents the response body for conversion
type ConvertResponse struct {
	Success bool        `json:"success"`
	Result  interface{} `json:"result,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// TypesResponse represents the response for supported types
type TypesResponse struct {
	Success bool     `json:"success"`
	Types   []string `json:"types"`
}

// StartServer starts the web server on the specified port
func StartServer(port int) error {
	r := gin.Default()

	// Add CORS middleware
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// Health check endpoint
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"service": "str2go",
		})
	})

	// Get supported types
	r.GET("/types", func(c *gin.Context) {
		types := converter.GetSupportedTypes()
		c.JSON(http.StatusOK, TypesResponse{
			Success: true,
			Types:   types,
		})
	})

	// Convert endpoint
	r.POST("/convert", handleConvert)

	// Serve static files for a simple web interface
	r.Static("/static", "./static")
	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "str2go - String to Go Type Converter",
		})
	})

	// Start the server
	return r.Run(":" + strconv.Itoa(port))
}

// handleConvert handles the conversion request
func handleConvert(c *gin.Context) {
	var req ConvertRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ConvertResponse{
			Success: false,
			Error:   "Invalid request: " + err.Error(),
		})
		return
	}

	var result interface{}
	var err error

	// Use delimiter if provided and type is a slice
	if req.Delimiter != "" && len(req.Type) > 2 && req.Type[:2] == "[]" {
		result, err = converter.ConvertWithDelimiter(req.Value, req.Type, req.Delimiter)
	} else {
		result, err = converter.Convert(req.Value, req.Type)
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, ConvertResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, ConvertResponse{
		Success: true,
		Result:  result,
	})
}

// GetServerInfo returns information about the server
func GetServerInfo() map[string]interface{} {
	return map[string]interface{}{
		"name":        "str2go",
		"version":     "1.0.0",
		"description": "String to Go Type Converter API",
		"endpoints": []string{
			"GET  /health",
			"GET  /types",
			"POST /convert",
		},
	}
}
