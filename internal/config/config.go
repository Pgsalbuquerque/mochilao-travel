package config

import (
	"sync"

	"github.com/spf13/viper"
)

var currentConfig *Config

var config *viper.Viper = viper.New()

var configOnce *sync.Once = &sync.Once{}

type Config struct {
	MongoConnectionString string
	GRPCPort              string
	GRPCGatewayPort       string
	DBName                string
}

const (
	MongoConnectionString = "MONGO_CONNECTION_STRING"
	GRPCPort              = "GRPC_PORT"
	GRPCGatewayPort       = "GRPC_GATEWAY_PORT"
	DBName                = "DB_NAME"
)

func Get() *Config {
	configOnce.Do(func() {
		config.BindEnv(MongoConnectionString)
		config.BindEnv(GRPCPort)
		config.BindEnv(GRPCGatewayPort)
		config.BindEnv(DBName)

		currentConfig = &Config{
			MongoConnectionString: config.GetString(MongoConnectionString),
			GRPCPort:              config.GetString(GRPCPort),
			GRPCGatewayPort:       config.GetString(GRPCGatewayPort),
			DBName:                config.GetString(DBName),
		}

	})
	return currentConfig
}
