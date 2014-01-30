package utils

import (
    "crypto/rand"
)

func RandomStr(L  int) string {
    /*
     * Generates a random string of lenght L.
     */
    const source = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
    var bytes = make([]byte, L)
    rand.Read(bytes)
    for i, b := range bytes {
        bytes[i] = source[b % byte(len(source))]
    }
    return string(bytes)
}

