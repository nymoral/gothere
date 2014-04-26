package password

import (
    "log"
    "code.google.com/p/go.crypto/bcrypt"
)

func hashBC(plain string) string {
    hashed, err := bcrypt.GenerateFromPassword([]byte(plain), 10)
    var hashedStr string
    if err != nil {
        log.Println(err)
        return ""
    } else {
        hashedStr = string(hashed)
        return hashedStr
    }
}

func NewPassword(plain string) (string) {
    return hashBC(plain)
}

func Authenticate(plain, hashed string) bool{
    // Returns true if passwords match,
    // false otherwise.
    // Checks if a given password matches
    // hashed one retrieved from a DB,
    // Uses bcrypts own comparison.

    if hashed == "" {
        // No user in the db.
        return false
    }
    hashBytes :=  []byte(hashed)
    plainBytes := []byte(plain)

    err := bcrypt.CompareHashAndPassword(hashBytes, plainBytes)
    if err != nil {
        log.Println(err)
        return false
    }
    return true
}
