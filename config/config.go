package config

import (
    "os"
    "log"
    "io/ioutil"
    "encoding/json"
)

var Config configType

func init() {
    args := os.Args
    var jsonFile string
    if len(args) < 2 {
        log.Println("Configuration file not provided.")
        jsonFile = "./config.json"
    } else {
        jsonFile = args[1]
    }
    file, err := ioutil.ReadFile(jsonFile)
    if err != nil {
        log.Println("Failed to read config file.")
        log.Fatal(err)
    }

    err = json.Unmarshal(file, &Config)
    if err != nil {
        log.Println("Failed to parse config file.")
        log.Fatal(err)
    }
}

type configType struct {
    Port           string
    Static         string
    ServeStatic    bool
    DbUser         string
    DbName         string
    DbPass         string
    Secret1        string
    Secret2        string
    TemplateDir    string
    SqlQueriesDir  string
    HashCycles     int
    MaxConnections int
    NrOfCookies    int
}
