package tests

import (
    "goThere/cookies"
    "goThere/utils"
    "goThere/password"
    "goThere/database"
    "fmt"
)

func cookieTest() {
    fmt.Println()
    fmt.Println("<-- COOKIE TEST BEGIN -->")

    cookie := cookies.GenerateCookie()
    testUname := utils.RandomStr(16)
    sessionid := cookies.GenerateSessionId(cookie, testUname)
    fmt.Printf("SessionID: %s...\n", sessionid[:20])
    username := cookies.UsernameFromCookie(cookie, sessionid)
    fmt.Printf("Username from cookie test: ")
    if username == testUname {
        fmt.Printf("PASSED\n")
    } else {
        fmt.Printf("FAILED\n")
    }

    fmt.Println("<-- COOKIE TEST END -->")
    fmt.Println()
}

func randomTest() {
    fmt.Println()
    fmt.Println("<-- RANDOM TEST BEGIN -->")

    t8 := utils.RandomStr(8)
    t16 := utils.RandomStr(16)
    t32 := utils.RandomStr(32)
    t64 := utils.RandomStr(64)
    fmt.Printf("8 char len: %d %s\n", len(t8), t8)
    fmt.Printf("16 char len: %d %s\n", len(t16), t16)
    fmt.Printf("32 char len: %d %s\n", len(t32), t32)
    fmt.Printf("64 char len: %d %s\n", len(t64), t64)

    fmt.Println("<-- RANDOM TEST END -->")
    fmt.Println()

}

func passwordTest() {
    fmt.Println()
    fmt.Println("<-- PASSWORD TEST BEGIN -->")

    p := utils.RandomStr(8)
    hashed := password.NewPassword(p)
    fmt.Printf("%s ->  %s\n", p, hashed)
    fmt.Printf("Same password test: ")
    if password.Authenticate(p, hashed) {
        fmt.Printf("PASSED\n")
    } else {
        fmt.Printf("FAILED\n")
    }

    fmt.Printf("Different password test: ")
    p = utils.RandomStr(8)
    if password.Authenticate(p, hashed) {
        fmt.Printf("FAILED\n")
    } else {
        fmt.Printf("PASSED\n")
    }

    fmt.Println("<-- PASSWORD TEST END -->")
    fmt.Println()

}

func dbTest() {
    fmt.Println()
    fmt.Println("<-- DB TEST BEGIN -->")

    fmt.Printf("Connection test: ")
    db, err := database.DbInit()
    if err != nil {
        fmt.Printf("FAILED\n")
        fmt.Println(err)
    } else {
        fmt.Printf("PASSED\n")

        fmt.Printf("Closing connection to DB.\n")
        database.DbClose(db)
    }

    fmt.Println("<-- DB TEST END -->")
    fmt.Println()

}

func Test() {
    cookieTest()
    randomTest()
    passwordTest()
    dbTest()
}


