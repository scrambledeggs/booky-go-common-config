package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
)

const (
	ENV_VAR_KEY_APP_ENV = "APP_ENV"
)

var (
	DEFAULT_CONFIG_PATH = []string{"/config.yml", "/config/config.yml"}
	_, absPath, _, _    = runtime.Caller(0)
	basepath            = filepath.Dir(absPath)
)

// NewConfig initialize
func NewConfig() {
	configPath, error := locateConfigFile()
	fmt.Println(configPath)
	fmt.Println(error)
}

func locateConfigFile() (configPath string, err error) {
	isFound := false
	fmt.Println("Locating configuration file...")
	viper.SetEnvPrefix(os.Getenv(ENV_VAR_KEY_APP_ENV))
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	for index := 0; index < len(DEFAULT_CONFIG_PATH); index++ {
		configPath = basepath + DEFAULT_CONFIG_PATH[index]
		viper.AddConfigPath(basepath + DEFAULT_CONFIG_PATH[index])
		err := viper.ReadInConfig() // Find and read the config file
		if err != nil {             // Handle errors reading the config file
			continue
		}
	}

	if isFound {
		return configPath, nil
	}

	return "", errors.New("config file not found")

}

func loadConfigFile() {
	if os.Getenv(ENV_VAR_KEY_APP_ENV) == "" {
		os.Setenv(ENV_VAR_KEY_APP_ENV, "development")
	}
}
