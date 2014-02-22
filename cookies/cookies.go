package cookies

import (
    "github.com/gorilla/securecookie"
    "log"
    "gothere/config"
)

func generateCookie() (*securecookie.SecureCookie) {
    /*
    * Generates SecureCookie type object and returns a pointer to it.
    * It is used to Encode/Decode plain data to/from a cookie.
    */

    S := securecookie.New([]byte(config.Secret1), []byte(config.Secret1))
    return S
}

var secCookie *securecookie.SecureCookie = generateCookie()

func GenerateSessionId(username string) (string) {
    /*
    * Given a SecureCookie object and a plain-text string encodes
    * it to a new cookie and returns SessionID to be used in
    * HTTP response.
    */

    encoded_cookie, err := secCookie.Encode("sessionid", username)
    // This should always work independant of input.
    if err != nil {
        log.Fatal(err)
    }
    return string(encoded_cookie)
}

func UsernameFromCookie(cookie string) (string) {
    /*
    * Given a SecureCookie object and a SessionID from request
    * decodes it and returns plain-text data from it (username).
    */

    var username string
    secCookie.Decode("sessionid", cookie, &username)
    return username
}
