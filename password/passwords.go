package password

import (
    "goThere/utils"
    "crypto/sha256"
    "encoding/hex"
    "strings"
)

func hashPassword(plain, salt string) string {
    /* 
     * Main password hashing function.
     * It takes plain-text password and a salt to use
     * and returns hashed password.
     * Uses SHA256.
     */
    hash := sha256.New()
    hash.Write([]byte(salt+plain))
    md := hash.Sum(nil)
    return hex.EncodeToString(md)
}

func NewPassword(plain string) (string) {
    /* 
     * Generates a random salt and hashes given password
     * to be stored in DB.
     * Uses hashPassword() func from this package.
     * Formated as "SALT|HASH"
     */
    salt := utils.RandomStr(16)
    return salt+"|"+hashPassword(plain, salt)
}

func Authenticate(plain, hashed string) bool{
    /* 
     * Returns true if passwords match,
     * false otherwise.
     * Checks if a given password matches
     * hashed one retrieved from a DB,
     * Uses hashPassword() from this package.
     */
    split := strings.Split(hashed, "|")
    return split[1] == hashPassword(plain, split[0])
}
