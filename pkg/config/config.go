package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Debug bool

	Database struct {
		Driver string
		Source string
	}

	Server struct {
		Address string
	}

	Context struct {
		Timeout int
	}
}

func LoadConfig(configFile string) (*Config, error) {
	viper.SetConfigFile(configFile)
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	config := &Config{}

	err = viper.Unmarshal(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
