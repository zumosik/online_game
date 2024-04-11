package config_game

import (
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
)

// Config for game using yaml and cleanenv package.
type Config struct {
	// Server data
	Host string `yaml:"host" env:"HOST" env-default:"localhost"`
	Port int    `yaml:"port" env:"PORT" env-default:"8080"`

	// User data
	Username string `yaml:"username" env:"USERNAME" env-required:"true"`
	Pin      uint32 `yaml:"pin" env:"PIN" env-required:"true"`
}

// ReadConfig reads the configuration from file, if cant find file reads from env.
// Priority: file > env.
// Priority for file name is: flag > env > default.
func ReadConfig() (*Config, error) {
	s := fetchCfgPath()
	if s == "" {
		return ReadConfigFromEnv()
	}
	return ReadConfigFromYaml(s)
}

func ReadConfigFromYaml(cfgPath string) (*Config, error) {
	var cfg Config
	err := cleanenv.ReadConfig(cfgPath, &cfg)
	if err != nil {
		return &Config{}, err
	}

	return &cfg, nil
}

func ReadConfigFromEnv() (*Config, error) {
	var cfg Config
	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		return &Config{}, err
	}
	return &cfg, err
}

func fetchCfgPath() string {
	var res string

	flag.StringVar(&res, "config", "", "path to config file")
	flag.Parse()

	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}

	return res
}
