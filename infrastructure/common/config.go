package common

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
)

type AppConfiguration struct {
	AppConfig []Config
}
type Config struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func NewConfiguration() AppConfiguration {
	filePath, _ := filepath.Abs("../config.json")
	data, err := ioutil.ReadFile(filePath)
	Check(err)

	// Decode content of file from Json to object
	appConfig := AppConfiguration{}
	err = json.Unmarshal(data, &appConfig)
	Check(err)
	fmt.Printf("%s", appConfig)

	return appConfig
}

// TODO: Add cache for app config.
func (config *AppConfiguration) ValueOf(key string) string {
	rs := ""
	for _, config := range config.AppConfig {
		if key == config.Key {
			rs = config.Value
		}
	}

	return rs
}
