package config

import (
	"encoding/json"
	"os"
)

//configuration file of service
type Config struct {
	Port string `json:"port"`
}

func LoadConfiguration(fileName string) (Config, error)  {
	var config Config
	configFile, err := os.Open(fileName)
	defer configFile.Close()
	if err != nil{
		return config, err
	}
	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)
	return config, err

}
