package cookies

import (
    "net/http"
    "log"
    "time"
)

func GetCookieVal(r *http.Request, name string) (string) {
    /*
     * From a given http.Request gets a cookie value.
     * Returns a string (empty one if there is no such cookie).
     */

    c, err := r.Cookie(name)
    if err != nil {
        if err == http.ErrNoCookie {
            return ""
        } else {
            log.Println(err)
            return ""
        }
    }

    return c.Value
}

func SetSessionId(w http.ResponseWriter, sessionid string, remember bool) {
    /* 
     * Writes a new (rewrites) cookie to set a sessionid.
     * If remember is false, cookie's expire field is not set
     * and cookie will expire when the session ends.
     * Otherwise it exipres in 2 weeks.
     */

    var C http.Cookie
    C.Name = "sessionid"
    C.Value = sessionid
    C.Path = "/"
    if remember {
        C.Expires = time.Now().AddDate(0, 0, 14)
    }
    http.SetCookie(w, &C)

}
