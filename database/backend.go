package database

import (
    "log"
    "fmt"
    "io/ioutil"
    "database/sql"
    _ "github.com/lib/pq"
    "gothere/config"
)

var DbChannel = make(chan *sql.DB, config.MaxConnections)
// All availible db connections will be stored here.

func init() {
    for i := 0; i < config.MaxConnections; i++ {
        // Creating and pushing db connections
        // to main channel.
        DbChannel <- DbInit()
    }

    db := GetConnection()
    defer RecycleConnection(db)

    err := db.Ping()
    // Testing db connectivity.
    if err != nil {
        log.Fatal(err)
    } else {
        log.Printf("Starting %d db connections.\n", config.MaxConnections)
    }

    log.Println("Loading queries.")
}

func GetConnection() (*sql.DB) {
    // Takes a connection from a channel or
    // waits for an availible one.
    return <- DbChannel
}

func RecycleConnection(con *sql.DB) {
    // After a function that user a connection
    // exits, used connection is returned to the channel.
    DbChannel <- con
}

func DbInit() (*sql.DB) {
    // Opens a connection to a postgresql databalse
    // and returns a pointer to sql.DB object.

    uname := " user=" + config.DbUser
    dname := " dbname=" + config.DbName

    var pass string

    if config.DbPass != "" {
        pass = " password=" + config.DbPass
    } else {
        pass = ""
    }

    openStatement := "sslmode=disable" + dname + uname + pass
    db, err := sql.Open("postgres", openStatement)
    if err != nil {
        log.Fatal(err)
    }

    return db
}

func getQuery(name string) (string) {
    // Reads a SQL querie string from a file and returns it.
    dir := config.SqlQueriesDir
    filename := fmt.Sprintf("%s%s.sql", dir, name)
    buffer, err := ioutil.ReadFile(filename)
    if err == nil {
        return string(buffer)
    } else {
        log.Fatal(err)
    }
    return ""
}

var (
    // All the SQL queries are loaded into string variables.
    qCreateGame     = getQuery("CreateGame")
    qOpenGames      = getQuery("OpenGames")
    qToFinish       = getQuery("ToFinish")
    qCloseGame      = getQuery("CloseGame")
    qFinishGame     = getQuery("FinishGame")
    qCheckGuess     = getQuery("CheckGuess")
    qInsertGuess    = getQuery("InsertGuess")
    qUpdateGuess    = getQuery("UpdateGuess")
    qUsersGuesses   = getQuery("UsersGuesses")
    qGetGames       = getQuery("GetGames")
    qGetUsers       = getQuery("GetUsers")
    qCreateUser     = getQuery("CreateUser")
    qGetPassword    = getQuery("GetPassword")
    qGetPkAdmin     = getQuery("GetPkAdmin")
    qGetTable       = getQuery("GetTable")
)
