package config

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"os"
	"strings"
)

func loadConfig(configFile string, cfg interface{}) (err error) {
	viper.SetConfigFile(configFile)
	viper.SetConfigType("json")
	err = viper.ReadInConfig()
	if err != nil {
		return errors.Wrap(err, "error reading config file")
	}
	for _, k := range viper.AllKeys() {
		value := viper.GetString(k)
		if strings.HasPrefix(value, "${") && strings.HasSuffix(value, "}") {
			viper.Set(k, getEnvOrPanic(strings.TrimSuffix(strings.TrimPrefix(value, "${"), "}")))
		}
	}
	err = viper.Unmarshal(cfg)
	if err != nil {
		return errors.Wrap(err, "unable to decode config to struct")
	}
	return
}

func getEnvOrPanic(env string) string {
	res := os.Getenv(env)
	if len(res) == 0 {
		panic("Missing mandatory env variable: " + env)
	}
	return res
}
