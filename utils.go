package main

import (
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

func replaceIds(userID string, lobbyId string) []byte {
	assetPath := "client/game/game.html"
	fileContents, err := embeddedFs.ReadFile(assetPath)
	if err != nil {
		log.Printf("Failed to read" + assetPath)
	}

	// replace ids in html
	// convert to string
	tmp := string(fileContents)
	tmp = strings.Replace(tmp, "{UserToken}", userID, 1)
	tmp = strings.Replace(tmp, "{LobbyToken}", lobbyId, 1)

	// convert back to bytes
	fileContents = []byte(tmp)
	return fileContents
}

func handleHtmlPath(writer http.ResponseWriter, _ *http.Request, filepath string) {
	fileContents, err := embeddedFs.ReadFile(filepath)
	if err != nil {
		log.Printf("Failed to read" + filepath)
	}

	writer.Header().Add("Content-Type", "text/html")
	_, err = writer.Write(fileContents)
	if err != nil {
		log.Printf("Failed to write" + filepath)
		return
	}
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
