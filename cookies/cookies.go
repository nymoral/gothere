package cookies

import (
    "log"
    "github.com/gorilla/securecookie"
    "gothere/config"
)

func generateCookie() (*securecookie.SecureCookie) {
    /*
    * Generates SecureCookie type object and returns a pointer to it.
    * It is used to Encode/Decode plain data to/from a cookie.
    */
    //hashKey := byte[](config.Secret1)
    //blockKey := byte[](config.Secret2)
    // Both values ought to be loaded from private file.
    // These values are only for development purposes.

    //S := securecookie.New(hashKey, blockKey)
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
    cookie_name := "sessionid"
    //S := generateCookie()
    encoded_cookie, err := secCookie.Encode(cookie_name, username)
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
    cookie_name := "sessionid"
    //S := generateCookie()
    var username string
    secCookie.Decode(cookie_name, cookie, &username)
    return username
}
