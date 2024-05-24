package main

import (
	"net/http"
	"path/filepath"
	"strings"
)

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
