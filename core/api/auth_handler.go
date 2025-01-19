package api

import (
	connect "connectrpc.com/connect"
	context "context"
	"database/sql"
	v1 "github.com/RA341/multipacman/gen/auth/v1"
	"github.com/RA341/multipacman/service"
	"net/http"
	"net/url"
)

type AuthHandler struct {
	auth *service.AuthService
}

func (a AuthHandler) Authenticate(ctx context.Context, c *connect.Request[v1.AuthRequest]) (*connect.Response[v1.AuthResponse], error) {
	//TODO implement me
	panic("implement me")
}

func (a AuthHandler) NewUser(ctx context.Context, c *connect.Request[v1.NewUserReq]) (*connect.Response[v1.NewUserRes], error) {
	//if c.Msg.Username == "" || c.Msg.Password == "" || c.Msg. == "" {
	//
	//}

	registerUser(a.auth, c)

	//TODO implement me
	panic("implement me")
}

func (a AuthHandler) Test(ctx context.Context, c *connect.Request[v1.AuthResponse]) (*connect.Response[v1.TestResponse], error) {
	//TODO implement me
	panic("implement me")
}

func registerUser(auth *service.AuthService, req *connect.Request[v1.NewUserReq]) {
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
