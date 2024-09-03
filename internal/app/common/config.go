package common

import (
	"github.com/RajabovIlyas/golang-crud/internal/app/constants"
	"github.com/spf13/viper"
	"time"
)

type Config struct {
	UrlDB                 string        `mapstructure:"DB_URL"`
	Port                  string        `mapstructure:"PORT"`
	TokenSecret           string        `mapstructure:"TOKEN_SECRET"`
	TokenExpiresIn        time.Duration `mapstructure:"TOKEN_EXPIRED_IN"`
	TokenRefreshExpiresIn time.Duration `mapstructure:"TOKEN_REFRESH_EXPIRED_IN"`
	TokenMaxAge           int           `mapstructure:"TOKEN_MAXAGE"`
	BaseUrl               string        `mapstructure:"BASE_URL"`
}

func GetConfig(path string) (config Config, err error) {

	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName("app")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)

	if config.Port == "" {
		config.Port = constants.PORT
		return
	}
	config.Port = ":" + config.Port
	return
}
