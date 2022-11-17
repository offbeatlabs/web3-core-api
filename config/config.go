package config

import (
	commonConfig "github.com/arhamj/offbeat-api/commons/config"
	"github.com/arhamj/offbeat-api/commons/logger"
)

type Config struct {
	LogConfig    logger.Config `mapstructure:"log_config"`
	SqliteConfig SqliteConfig  `mapstructure:"sqlite_config"`
	HelperFlags  HelperFlags   `mapstructure:"helper_flags"`
	FeatureFlags FeatureFlags  `mapstructure:"feature_flags"`
}

type SqliteConfig struct {
	Path string `mapstructure:"path"`
}

type HelperFlags struct {
	RunMigrations bool `mapstructure:"run_migrations"`
}

type FeatureFlags struct {
	EnableTokenSync    bool `mapstructure:"enable_token_sync"`
	EnablePriceSync    bool `mapstructure:"enable_price_sync"`
	EnableHttpErrDebug bool `mapstructure:"enable_http_err_debug"`
}

func NewConfig(configFile string) (Config, error) {
	var cfg Config
	err := commonConfig.LoadConfig(configFile, &cfg)
	if err != nil {
		return Config{}, err
	}
	return cfg, nil
}
