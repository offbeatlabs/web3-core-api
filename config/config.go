package config

import (
	commonConfig "github.com/arhamj/go-commons/pkg/config"
	"github.com/offbeatlabs/web3-core-api/pkg/db"
)

type Config struct {
	SqliteConfig db.SqliteConfig `mapstructure:"sqlite_config"`
	HelperFlags  HelperFlags     `mapstructure:"helper_flags"`
	ServerConfig ServerConfig    `mapstructure:"server_config"`
	FeatureFlags FeatureFlags    `mapstructure:"feature_flags"`
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
