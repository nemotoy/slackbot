package config

import (
)

type Config struct {
  SlackBot   SlackBotConfig
}

type SlackBotConfig struct {
  Token string  `toml:"token"`
}
