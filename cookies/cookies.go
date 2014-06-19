package cookies

import (
	"github.com/gorilla/securecookie"
	"gothere/config"
	"log"
	"net/http"
	"time"
)

func generateCookie() *securecookie.SecureCookie {
	// Generates SecureCookie type object and returns a pointer to it.
	// It is used to Encode/Decode plain data to/from a cookie.

	S := securecookie.New([]byte(config.Config.Secret1), []byte(config.Config.Secret2))
	return S
}

var cookiesChan = make(chan *securecookie.SecureCookie, config.Config.NrOfCookies)

func init() {
	for i := 0; i < config.Config.NrOfCookies; i++ {
		cookiesChan <- generateCookie()
	}
}

func getCookie() *securecookie.SecureCookie {
	return <-cookiesChan
}

func recycleCookie(c *securecookie.SecureCookie) {
	cookiesChan <- c
}

func GenerateSessionId(username string) string {
	// Given a SecureCookie object and a plain-text string encodes
	// it to a new cookie and returns SessionID to be used in
	// HTTP response.

	secCookie := getCookie()
	defer recycleCookie(secCookie)

	encoded_cookie, err := secCookie.Encode("sessionid", username)
	if err != nil {
		// Input can't affect this func -> bad secureCookie.
		log.Fatal(err)
	}
	return string(encoded_cookie)
}

func UsernameFromCookie(cookie string) string {
	// Given a SecureCookie object and a SessionID from request
	// decodes it and returns plain-text data from it (username).

	secCookie := getCookie()
	defer recycleCookie(secCookie)

	var username string
	secCookie.Decode("sessionid", cookie, &username)
	return username
}

func GetCookieVal(r *http.Request, name string) string {
	// From a given http.Request gets a cookie value.
	// Returns a string (empty one if there is no such cookie).

	c, err := r.Cookie(name)
	if err != nil {
		return ""
	}

	return c.Value
}

func SetCookieVal(w http.ResponseWriter, name string, value string, path string) {
	var C http.Cookie
	C.Name = name
	C.Value = value
	C.Path = path
	C.Expires = time.Now().AddDate(0, 1, 0)
	http.SetCookie(w, &C)
}
func SetSessionId(w http.ResponseWriter, sessionid string, remember bool) {
	// Writes a new (rewrites) cookie to set a sessionid.
	// If remember is false, cookies expire field is not set
	// and cookie will expire when the session ends.
	// Otherwise it exipres in 2 weeks.

	var C http.Cookie
	C.Name = "sessionid"
	C.Value = sessionid
	C.Path = "/"
	if remember {
		C.Expires = time.Now().AddDate(0, 0, 14)
	}
	http.SetCookie(w, &C)
}

func DeleteSessionId(w http.ResponseWriter) {
	// Resets sessionId.
	// When user logs out, his session id is set to
	// none, and thus Gorila will not be able to decode
	// the username of the user.

	var C http.Cookie
	C.Name = "sessionid"
	C.Value = "none"
	C.Path = "/"
	http.SetCookie(w, &C)
}
