package config

import (
    "log"
    "io/ioutil"
    "encoding/json"
)

const jsonFileLocation string = "./config.json"

var Config configType

func init() {
    file, err := ioutil.ReadFile(jsonFileLocation)
    if err != nil {
        log.Fatal(err)
    }

    err = json.Unmarshal(file, &Config)
    if err != nil {
        log.Fatal(err)
    }
    log.Println(Config)
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
