package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

var Config configType

func init() {
	args := os.Args
	var jsonFile string
	if len(args) < 2 {
		jsonFile = "./config/config.json"
	} else {
		jsonFile = args[1]
	}
	log.Printf("Config file: %s\n", jsonFile)
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
	Port             string
	Static           string
	ServeStatic      bool
	DynamicTemplates bool
	DbUser           string
	DbName           string
	DbPass           string
	DbIp             string
	Secret1          string
	Secret2          string
	TemplateDir      string
	MaxConnections   int
	NrOfCookies      int

	MailUsername string
	MailPassword string
	MailHost     string
	MailPort     string

	Register     bool
	ShowClosedNr int
	ShowOpenNr   int
}
