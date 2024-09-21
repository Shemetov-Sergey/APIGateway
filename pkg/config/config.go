package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	GateWayAddr    string `mapstructure:"API_GATEWAY_ADDR"`
	Port           string `mapstructure:"PORT"`
	AuthSvcUrl     string `mapstructure:"AUTH_SVC_URL"`
	GoNewsSvcUrl   string `mapstructure:"GONEWS_SVC_URL"`
	CommentSvcUrl  string `mapstructure:"COMMENT_SVC_URL"`
	CensoredSvcUrl string `mapstructure:"CENSOR_SVC_URL"`
}

func LoadConfig() (Config, error) {
	var c Config
	viper.AddConfigPath("./pkg/config/envs")

	viper.SetConfigName("prod")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()

	if err != nil {
		viper.SetConfigName("dev")
		viper.SetConfigType("env")
		err = viper.Unmarshal(&c)
		viper.AutomaticEnv()

		err = viper.ReadInConfig()

		if err != nil {
			return Config{}, err
		}
		return c, nil
	}

	err = viper.Unmarshal(&c)

	return c, nil
}
