package database

import (
    "log"
    "io/ioutil"
    "database/sql"
    _ "github.com/lib/pq"
    "gothere/config"
)

var dbConnection *sql.DB
// A connection to the db.

func init() {
    dbConnection = dbInit()
    dbConnection.SetMaxOpenConns(config.Config.MaxConnections)
    dbConnection.SetMaxIdleConns(config.Config.MaxConnections)
    // Establish the connection.
    log.Printf("Starting db connections. Max open/idle connections: %d\n", config.Config.MaxConnections)
}

func GetConnection() (*sql.DB) {
    // Passes a connection to a handler.
    return dbConnection
}

func dbInit() (*sql.DB) {
    // Opens a connection to a postgresql databalse
    // and returns a pointer to sql.DB object.

    uname := " user=" + config.Config.DbUser
    dname := " dbname=" + config.Config.DbName

    var pass string

    if config.Config.DbPass != "" {
        pass = " password=" + config.Config.DbPass
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
    dir := config.Config.SqlQueriesDir
    filename := dir +  name + ".sql"
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
    qCreateGame         = getQuery("CreateGame")
    qOpenGames          = getQuery("OpenGames")
    qToFinish           = getQuery("ToFinish")
    qCloseGame          = getQuery("CloseGame")
    qFinishGame         = getQuery("FinishGame")
    qCheckGuess         = getQuery("CheckGuess")
    qInsertGuess        = getQuery("InsertGuess")
    qUpdateGuess        = getQuery("UpdateGuess")
    qUsersGuesses       = getQuery("UsersGuesses")
    qGetGames           = getQuery("GetGames")
    qGetUsers           = getQuery("GetUsers")
    qCreateUser         = getQuery("CreateUser")
    qGetPassword        = getQuery("GetPassword")
    qGetPkAdmin         = getQuery("GetPkAdmin")
    qGetTable           = getQuery("GetTable")

    qGetResult          = getQuery("GetResult")
    qGetPoints          = getQuery("GetPoints")
    qUpdatePoints       = getQuery("UpdatePoints")
    qUpdateGuessPoints  = getQuery("UpdateGuessPoints")
    qInsertGuessPoints  = getQuery("InsertGuessPoints")
)
