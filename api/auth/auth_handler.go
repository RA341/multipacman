package auth

import (
	"database/sql"
	"github.com/go-chi/chi/v5"
	"net/http"
	"net/url"
)

func SetupAuthRouter(db *sql.DB) *chi.Mux {
	r := chi.NewRouter()
	r.Post("/register", func(writer http.ResponseWriter, request *http.Request) {
		registerUser(db, writer, request)
	})
	r.Post("/login", func(writer http.ResponseWriter, request *http.Request) {
		loginUser(db, writer, request)
	})
	return r
}

func registerUser(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")
	passwordVerify := r.FormValue("passwordVerify")

	if username == "" || password == "" || passwordVerify == "" {
		redirectUrl := generateRedirectUrl("/register", "one or inputs empty")
		http.Redirect(w, r, redirectUrl, http.StatusSeeOther)
		return
	}

	// Ensure that the password & passwordVerify match
	if password != passwordVerify {
		redirectUrl := generateRedirectUrl("/register", "Passwords do not match")
		http.Redirect(w, r, redirectUrl, http.StatusSeeOther)
		return
	}

	res := createUser(db, username, password, "")

	if !res {
		redirectUrl := generateRedirectUrl("/register", "Failed to create user, try again")
		http.Redirect(w, r, redirectUrl, http.StatusSeeOther)
		return
	}

	// Redirect to login.
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func generateRedirectUrl(baseUrl string, message string) string {
	params := url.Values{}
	params.Set("error", message)
	return baseUrl + "?" + params.Encode()
}

func loginUser(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user := r.FormValue("username")
	pass := r.FormValue("password")

	userData := retrieveUser(db, user)
	if userData.Username == "" || userData.Password == "" {
		redirectUrl := generateRedirectUrl("/login", "Could not find user")
		http.Redirect(w, r, redirectUrl, http.StatusSeeOther)
		return
	}
	// Verify username and password
	if !checkPassword(pass, userData.Password) {
		redirectUrl := generateRedirectUrl("/login", "Invalid Username or Password")
		http.Redirect(w, r, redirectUrl, http.StatusSeeOther)
		return
	}

	// Generate auth token and Insert hashed auth token to DB
	token := createAuthToken(25)
	res := updateUserAuthToken(db, hashString(token), userData.ID)
	if !res {
		redirectUrl := generateRedirectUrl("/login", "Something went wrong when signing in try again")
		http.Redirect(w, r, redirectUrl, http.StatusSeeOther)
		return
	}

	// Set cookie
	expires := 3 * 3600 // 3 hours in seconds
	cookieOptions := &http.Cookie{
		Name:     "auth",
		Value:    token,
		MaxAge:   expires,
		HttpOnly: true,
		Path:     "/",
	}
	http.SetCookie(w, cookieOptions)

	// Redirect to lobby
	http.Redirect(w, r, "/lobby", http.StatusSeeOther)
}
