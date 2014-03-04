package cookies

import (
    "github.com/gorilla/securecookie"
    "log"
    "gothere/config"
)

func generateCookie() (*securecookie.SecureCookie) {
    // Generates SecureCookie type object and returns a pointer to it.
    // It is used to Encode/Decode plain data to/from a cookie.

    S := securecookie.New([]byte(config.Secret1), []byte(config.Secret1))
    return S
}

var cookiesChan = make(chan *securecookie.SecureCookie, config.NrOfCookies)

func init() {
    for i := 0; i < config.NrOfCookies; i++ {
        cookiesChan <- generateCookie()
    }
    log.Printf("Generating %d secure cookies.\n", config.NrOfCookies)
}

func getCookie() (*securecookie.SecureCookie) {
    return <- cookiesChan
}

func recycleCookie(c *securecookie.SecureCookie) {
    cookiesChan <- c
}

func GenerateSessionId(username string) (string) {
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

func UsernameFromCookie(cookie string) (string) {
    // Given a SecureCookie object and a SessionID from request
    // decodes it and returns plain-text data from it (username).

    secCookie := getCookie()
    defer recycleCookie(secCookie)

    var username string
    secCookie.Decode("sessionid", cookie, &username)
    return username
}
