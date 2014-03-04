package password

import (
    "fmt"
    "crypto/sha256"
    "encoding/hex"
    "gothere/utils"
    "gothere/config"
)

func hashPassword(plain, salt string, cycles int) string {
    // Main password hashing function.
    // It takes plain-text password a salt and nr of cycles
    // and returns hashed password.
    // Uses SHA256.

    hashed := plain
    for i := 0; i < cycles; i++ {
        hash := sha256.New()
        hash.Write([]byte(salt+hashed))
        md := hash.Sum(nil)
        hashed = hex.EncodeToString(md)
    }

    hash := sha256.New()
    hash.Write([]byte(salt+hashed))
    md := hash.Sum(nil)
    return hex.EncodeToString(md)
}

func NewPassword(plain string) (string) {
    //  Generates a random salt and hashes given password
    // to be stored in DB.
    // Uses hashPassword() func from this package.
    // Formated as "CYCLE SALT HASH"

    cycles := config.HashCycles

    salt := utils.RandomStr(16)
    return fmt.Sprintf("%d %s %s", cycles, salt, hashPassword(plain, salt, cycles))
}

func Authenticate(plain, hashed string) bool{
    // Returns true if passwords match,
    // false otherwise.
    // Checks if a given password matches
    // hashed one retrieved from a DB,
    // Uses hashPassword() from this package.

    if hashed == "" {
        // No user in the db.
        return false
    }
    var cycles int
    var s, h string
    nr, err := fmt.Sscanf(hashed, "%d %s %s", &cycles, &s, &h)
    // Extracts nr of cycles, salt and a hash from db data.
    if nr != 3 || err != nil {
        // Bad formating of a password.
        // Should never happen.
        return false
    }
    return h == hashPassword(plain, s, cycles)
}
