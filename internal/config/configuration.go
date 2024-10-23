package config

import (
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// LoadConfig loads configuration file
func InitConfig() (*viper.Viper, error) {

	log.Info("loading config...")

	v := viper.New()

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	viper.SetDefault("app.version", "0.0.1")

	viper.SetDefault("server.host", "localhost")
	viper.SetDefault("server.port", "7001")
	viper.SetDefault("server.log.level", "debug") // valid values are trace, debug, info, warn, error, fatal
	viper.SetDefault("server.timeout.write", "15")
	viper.SetDefault("server.timeout.read", "15")
	viper.SetDefault("server.request.timeout", "30") // in seconds

	viper.SetDefault("psql.dbname", "mydb")
	viper.SetDefault("psql.user", "seeduser")
	viper.SetDefault("psql.pass", "seeduser123")
	viper.SetDefault("psql.host", "localhost")
	viper.SetDefault("psql.port", "5432")

	viper.SetDefault("jwt.accessKey", "sample_key")
	viper.SetDefault("jwt.refreshKey", "sample_key")

	return v, nil
}
