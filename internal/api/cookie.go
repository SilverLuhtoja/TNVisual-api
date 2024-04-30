package api

import (
	"net/http"
	"strings"
)

func SetCookieHandler(w http.ResponseWriter, r *http.Request, apiKey string) {
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

func GetCookieHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("tnsCookie")
	if err != nil {
		RespondWithError(w, http.StatusUnauthorized, http.ErrNoCookie.Error())
		return
	}

	RespondWithJSON(w, http.StatusOK, strings.Split(cookie.Value, " ")[1])
}
