package config

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Config struct {
	Delivery struct {
		HTTP struct {
			BindPort       uint          `validate:"gte=1,lte=65535"` // between 1 and 65535
			RequestTimeout time.Duration `validate:"gt=0"`
		}
	}
	Log struct {
		Level string `validate:"required,oneof=debug info warn error"` // allowed values: debug, info, warn, error
	}
	Storage struct {
		Endpoint string `validate:"required,url"`
	}
	Secret struct {
		Key string `validate:"required"`
	}
}

func Default() *Config {
	cfg := new(Config)

	cfg.Delivery.HTTP.BindPort = 80
	cfg.Delivery.HTTP.RequestTimeout = 10 * time.Second

	cfg.Log.Level = "info"

	return cfg
}

func (c *Config) Validate() error {
	validate := validator.New()
	return validate.Struct(c)
}
