package common

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

// Config app
type Config struct {
	MongoHost    string `json:"mongo_host"`
	DatabaseName string `json:"database_name"`
}

// GetConfig from file
func GetConfig() (*Config, error) {
	config := &Config{}

	raw, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Println("err on readfile - GetConfig", os.Args[1], raw, err)
		return config, err
	}

	err = json.Unmarshal(raw, &config)
	if err != nil {
		// it'll return below
		log.Println("err on Unmarshal - GetConfig")
	}
	return config, err
}
