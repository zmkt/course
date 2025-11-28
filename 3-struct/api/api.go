package api

import (
	"3-struct/config"
)

type Config struct {
	Key string
}

func NewConfig() (Config, error) {

	key, err := config.LoadKey()
	if err != nil {
		return Config{}, err
	}
	return Config{
		Key: key,
	}, nil
}
