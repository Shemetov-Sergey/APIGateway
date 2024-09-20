package config

import "github.com/spf13/viper"

type Config struct {
	Port          string `mapstructure:"PORT"`
	AuthSvcUrl    string `mapstructure:"AUTH_SVC_URL"`
	GoNewsSvcUrl  string `mapstructure:"GONEWS_SVC_URL"`
	CommentSvcUrl string `mapstructure:"COMMENT_SVC_URL"`
	Censored      string `mapstructure:"CENSORED"`
}

func LoadConfig() (Config, error) {
	var c Config
	viper.AddConfigPath("./pkg/config/envs")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()

	if err != nil {
		return Config{}, err
	}

	err = viper.Unmarshal(&c)

	return c, nil
}
