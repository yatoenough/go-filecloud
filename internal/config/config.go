package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	Env        string `yaml:"env" env:"ENV" json:"env" env-default:"development"`
	HTTPConfig `yaml:"http_config" json:"httpConfig"`
}

type HTTPConfig struct {
	Address        string        `yaml:"address" env:"ADDRESS" json:"address" env-default:"localhost:3000"`
	RequestTimeout time.Duration `yaml:"req_timeout" env:"REQ_TIMEOUT" json:"reqTimeout" env-default:"4s"`
	IdleTimeout    time.Duration `yaml:"idle_timeout" env:"IDLE_TIMEOUT" json:"idleTimeout" env-default:"60s"`
}

func MustLoad() *Config {
	cfgPath := os.Getenv("CONFIG_PATH")
	if cfgPath == "" {
		log.Print(cfgPath)
		panic("CONFIG_PATH is not set")
	}

	if _, err := os.Stat(cfgPath); os.IsNotExist(err) {
		panic(fmt.Sprintf("config file does not exist: %s", cfgPath))
	}

	var cfg Config
	if err := cleanenv.ReadConfig(cfgPath, &cfg); err != nil {
		panic(fmt.Sprintf("cannot read config: %s", err))
	}

	return &cfg
}
