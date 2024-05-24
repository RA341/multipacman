package login

import "net/http"

// Define your route handlers
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Home Page"))
	if err != nil {
		return
	}
}
