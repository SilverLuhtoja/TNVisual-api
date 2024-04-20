package api

import (
	"errors"
	"net/http"
)

func SetCookieHandler(w http.ResponseWriter, r *http.Request, apiKey string) {
	cookie := http.Cookie{
		Name:     "tns_cookie",
		Value:    "ApiKey " + apiKey,
		Path:     "/",
		MaxAge:   36000,
		HttpOnly: true,
		// SameSite: http.SameSiteLaxMode,
		Secure:   true,
		SameSite: http.SameSiteNoneMode,
	}

	// Use the http.SetCookie() function to send the cookie to the client.
	// Behind the scenes this adds a `Set-Cookie` header to the response
	// containing the necessary cookie data.
	http.SetCookie(w, &cookie)
}

func GetCookieHandler(w http.ResponseWriter, r *http.Request) error {
	// Retrieve the cookie from the request using its name (which in our case is
	// "exampleCookie"). If no matching cookie is found, this will return a
	// http.ErrNoCookie error. We check for this, and return a 400 Bad Request
	// response to the client.
	_, err := r.Cookie("tns_cookie")
	if err != nil {
		return errors.New("cookie not found")
	}

	return nil
}
