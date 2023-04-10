package config

import (
	"github.com/spf13/viper"
	"log"
	"strconv"
)

const (
	CONFIG_DIR  = "configs"
	CONFIG_FILE = "main"
)

type Config struct {
	Host        string
	Port        int
	Username_DB string
	DBName      string
	SSLMode     string
	Password    string
}

func unmarshal(cfg *Config) error {
	if err := viper.Unmarshal(&cfg); err != nil {
		return err
	}
	return nil
}
func Init() (*Config, error) {
	if err := setUpViper(); err != nil {
		return nil, err
	}
	//viper.AddConfigPath(CONFIG_DIR)
	//viper.SetConfigName(CONFIG_FILE)
	var cfg Config
	if err := unmarshal(&cfg); err != nil {
		return nil, err
	}

	if err := fromEnv(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
func fromEnv(cfg *Config) error {
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("нету значений")
	}

	cfg.Host = viper.GetString("Host")
	cfg.Port, err = strconv.Atoi(viper.GetString("Port"))
	cfg.Username_DB = viper.GetString("Username_DB")

	cfg.DBName = viper.GetString("DBName")
	cfg.SSLMode = viper.GetString("SSLMode")
	cfg.Password = viper.GetString("Password")

	return nil
}

func setUpViper() error {
	viper.AddConfigPath("configs")

	return viper.ReadInConfig()
}
