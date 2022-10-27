package config

import (
	"github.com/arhamj/offbeat-api/commons/logger"
)

type Config struct {
	LogConfig    logger.Config `mapstructure:"log_config"`
	SqliteConfig SqliteConfig  `mapstructure:"sqlite_config"`
	HelperFlags  HelperFlags   `mapstructure:"helper_flags"`
	FeatureFlags FeatureFlags  `mapstructure:"feature_flags"`
}

type SqliteConfig struct {
	Path string `json:"path" validate:"required"`
}

type HelperFlags struct {
	RunMigrations bool `json:"run_migrations"`
}

type FeatureFlags struct {
}

func NewConfig(configFile string) (Config, error) {
	var cfg Config
	err := loadConfig(configFile, &cfg)
	if err != nil {
		return Config{}, err
	}
	return cfg, nil
}
