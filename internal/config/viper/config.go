package viper

import (
	"fmt"

	"github.com/crypto-grill/app/internal/config"
	"github.com/spf13/viper"
)

func LoadConfig() (*config.Config, error) {
	cfg := config.Default()

	viper.AutomaticEnv()

	_ = viper.BindEnv("Storage.Endpoint", "STORAGE_DSN")
	_ = viper.BindEnv("SecretKey", "SECRET_KEY")
	_ = viper.BindEnv("Delivery.HTTP.BindPort", "PORT")

	if err := viper.Unmarshal(cfg); err != nil {
		return nil, fmt.Errorf("unable to decode environment into config struct: %w", err)
	}

	if err := cfg.Validate(); err != nil {
		return nil, err
	}
	return cfg, nil
}
