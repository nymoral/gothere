package cookies

import (
    "net/http"
    "log"
)

func GetCookieVal(r *http.Request, name string) (string) {
    c, err := r.Cookie(name)
    if err != nil {
        if err == http.ErrNoCookie {
            return ""
        } else {
            log.Fatal(err)
        }
    }

    return c.Value
}

func SetSessionId(w http.ResponseWriter, sessionid string, remember bool) {
    var C http.Cookie
    C.Name = "sessionid"
    C.Value = sessionid
    C.Path = "/"
    http.SetCookie(w, &C)

}
