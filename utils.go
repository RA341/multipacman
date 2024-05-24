package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// app related

func getServerPort() string {
	var port string
	if len(os.Args) > 1 {
		port = os.Args[1]
		fmt.Println(port)
	} else {
		port = "5000" // Default port if not specified
	}
	fmt.Println("Server started at " + port)
	return port
}

// File related

// FileExists checks if a file exists and is not a directory
func FileExists(filename string) bool {
	if _, err := os.Stat(filename); errors.Is(err, os.ErrNotExist) {
		log.Printf("File does not exist")
		return false
	}
	return true
}

func handleHtmlPath(writer http.ResponseWriter, request *http.Request, filepath string) {
	if !FileExists(filepath) {
		http.NotFound(writer, request)
		return
	}
	http.ServeFile(writer, request, filepath)
}

func DetectContentType(assetPath string, fileContents []byte) string {
	fileExt := strings.ToLower(filepath.Ext(assetPath))
	var contentType string

	switch fileExt {
	case ".html":
		contentType = "text/html"
	case ".css":
		contentType = "text/css"
	case ".js":
		contentType = "application/javascript"
	case ".png":
		contentType = "image/png"
	case ".jpg", ".jpeg":
		contentType = "image/jpeg"
	case ".gif":
		contentType = "image/gif"
	case ".svg":
		contentType = "image/svg+xml"
	case ".json":
		contentType = "application/json"
	case ".pdf":
		contentType = "application/pdf"
	default:
		contentType = http.DetectContentType(fileContents)
	}
	return contentType
}
