package auth

import (
	"net/http"
)

func SetCookieHandler(w http.ResponseWriter, apiKey string) {
	var minutes int = 15
	cookie := http.Cookie{
		Name:     "tnsCookie",
		Value:    "ApiKey " + apiKey,
		Path:     "/",
		MaxAge:   minutes * 60,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Secure:   true,
		// SameSite: http.SameSiteNoneMode,
	}

	// Use the http.SetCookie() function to send the cookie to the client.
	// Behind the scenes this adds a `Set-Cookie` header to the response
	// containing the necessary cookie data.
	http.SetCookie(w, &cookie)
}
