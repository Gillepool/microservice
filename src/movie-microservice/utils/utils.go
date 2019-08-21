package utils

import (
	"encoding/json"
	"os"
)

//Configuration information
type Configuration struct {
	Port                string `json:"port"`
	EnableGinConsoleLog bool   `json:"enableGinConsoleLog"`
	EnableGinFileLog    bool   `json:"enableGinFileLog"`

	LogFilename   string `json:"logFilename"`
	LogMaxSize    int    `json:"logMaxSize"`
	LogMaxBackups int    `json:"logMaxBackups"`
	LogMaxAge     int    `json:"logMaxAge"`

	MgAddrs      string `json:"mgAddrs"`
	MgDbName     string `json:"mgDbName"`
	MgDbUsername string `json:"mgDbUsername"`
	MgDbPassword string `json:"mgDbPassword"`

	AuthAddr          string `json:"authAddr"`
	JwtSecretPassword string `json:"jwtSecretPassword"`
	Issuer            string `json:"issuer"`
}

// Global Config
var (
	Config *Configuration
)

//LoadConfig the config for config file
func LoadConfig() error {
	file, err := os.Open("config/config.json")
	if err != nil {
		return err
	}

	Config = new(Configuration)
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&Config)
	if err != nil {
		return err
	}

	return nil
}
