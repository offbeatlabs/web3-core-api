package config

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"os"
	"strings"
)

func LoadConfig(configFile string, cfg interface{}) (err error) {
	viper.SetConfigFile(configFile)
	viper.SetConfigType("json")
	err = viper.ReadInConfig()
	if err != nil {
		return errors.Wrap(err, "error reading config file")
	}
	for _, k := range viper.AllKeys() {
		value := viper.GetString(k)
		if strings.HasPrefix(value, "${") && strings.HasSuffix(value, "}") {
			envValue, err := getEnvOrPanic(strings.TrimSuffix(strings.TrimPrefix(value, "${"), "}"))
			if err != nil {
				return err
			}
			viper.Set(k, envValue)
		}
	}
	err = viper.Unmarshal(cfg)
	if err != nil {
		return errors.Wrap(err, "unable to decode config to struct")
	}
	return
}

func getEnvOrPanic(env string) (string, error) {
	res := os.Getenv(env)
	if len(res) == 0 {
		return "", errors.New("missing mandatory env variable: " + env)
	}
	return res, nil
}
