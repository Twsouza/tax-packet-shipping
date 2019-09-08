package common

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"time"
)

// Config app
type Config struct {
	DbParams struct {
		Host    string        `json:"host"`
		Name    string        `json:"name"`
		TimeOut time.Duration `json:"timeout"`
	} `json:"database_params"`
	ServerParams struct {
		Addr            string        `json:"address"`
		WriteTimeout    time.Duration `json:"write_timeout"`
		ReadTimeout     time.Duration `json:"read_timeout"`
		IdleTimeout     time.Duration `json:"idle_timeout"`
		GracefulTimeout time.Duration `json:"graceful_timeout"`
	} `json:"server_params"`
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
