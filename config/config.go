package config

import (
	commonConfig "github.com/arhamj/go-commons/pkg/config"
)

type Config struct {
	SqliteConfig SqliteConfig `mapstructure:"sqlite_config"`
	HelperFlags  HelperFlags  `mapstructure:"helper_flags"`
	ServerConfig ServerConfig `mapstructure:"server_config"`
	FeatureFlags FeatureFlags `mapstructure:"feature_flags"`
}

type SqliteConfig struct {
	Path string `mapstructure:"path"`
}

type HelperFlags struct {
	RunMigrations bool `mapstructure:"run_migrations"`
}

type FeatureFlags struct {
	EnableTokenSync bool `mapstructure:"enable_token_sync"`
	EnablePriceSync bool `mapstructure:"enable_price_sync"`
}

type ServerConfig struct {
	BaseUrlForSwagger string `mapstructure:"base_url_for_swagger"`
}

func NewConfig(configFile string) (Config, error) {
	var cfg Config
	err := commonConfig.LoadConfig(configFile, &cfg)
	if err != nil {
		return Config{}, err
	}
	return cfg, nil
}
