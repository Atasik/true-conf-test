package config

import (
	"os"
	"time"

	"github.com/spf13/viper"
)

const (
	defaultHTTPPort                = "3333"
	defaultHTTPRWTimeout           = 10 * time.Second
	defaultHTTPMaxHeaderMegabytes  = 1
	defaultDatabaseRefreshInterval = 30 * time.Second
)

type (
	Config struct {
		HTTP HTTPConfig
		DB   DBConfig
	}

	HTTPConfig struct {
		Host               string
		Port               string        `mapstructure:"port"`
		ReadTimeout        time.Duration `mapstructure:"readTimeout"`
		WriteTimeout       time.Duration `mapstructure:"writeTimeout"`
		MaxHeaderMegaBytes int           `mapstructure:"maxHeaderMegaBytes"`
	}

	DBConfig struct {
		UsersStore string `mapstructure:"users"`
	}
)

func InitConfig(configPath string) (*Config, error) {
	setDefaults()

	if err := parseConfigFile(configPath); err != nil {
		return nil, err
	}

	var cfg Config
	if err := unmarshal(&cfg); err != nil {
		return nil, err
	}

	setFromEnv(&cfg)

	return &cfg, nil
}

func setFromEnv(cfg *Config) {
	cfg.HTTP.Host = os.Getenv("HTTP_HOST")
}

func unmarshal(cfg *Config) error {
	if err := viper.UnmarshalKey("db", &cfg.DB); err != nil {
		return err
	}

	return viper.UnmarshalKey("http", &cfg.HTTP)
}

func parseConfigFile(folder string) error {
	viper.AddConfigPath(folder)
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return viper.MergeInConfig()
}

func setDefaults() {
	viper.SetDefault("http.port", defaultHTTPPort)
	viper.SetDefault("http.maxHeaderMegaBytes", defaultHTTPMaxHeaderMegabytes)
	viper.SetDefault("http.readTimeout", defaultHTTPRWTimeout)
	viper.SetDefault("http.writeTimeout", defaultHTTPRWTimeout)
}
