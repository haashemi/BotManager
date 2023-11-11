package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	API            API            `yaml:"api"`
	Bot            Bot            `yaml:"bot"`
	TelegramBotAPI TelegramBotAPI `yaml:"telegram_bot_api"`
}

type API struct {
	Addr string `yaml:"addr"`
}

type Bot struct {
	Token  string  `yaml:"token"`
	Admins []int64 `yaml:"admins"`
}

type TelegramBotAPI struct {
	Dir  string `yaml:"dir"`
	Host string `yaml:"host"`
}

func ParseConfig() (*Config, error) {
	file, err := os.Open("config.yaml")
	if err != nil {
		return nil, err
	}

	config := &Config{}

	if err = yaml.NewDecoder(file).Decode(config); err != nil {
		return nil, err
	}

	return config, nil
}
