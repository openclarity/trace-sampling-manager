package config

import (
	"github.com/spf13/viper"
)

const (
	RestServerPort     = "REST_SERVER_PORT"
	GRPCServerPort     = "GRPC_SERVER_PORT"
	HealthCheckAddress = "HEALTH_CHECK_ADDRESS"
)

type Config struct {
	RestServerPort     int
	GRPCServerPort     int
	HealthCheckAddress string
}

func LoadConfig() *Config {
	config := &Config{}

	config.RestServerPort = viper.GetInt(RestServerPort)
	config.GRPCServerPort = viper.GetInt(GRPCServerPort)
	config.HealthCheckAddress = viper.GetString(HealthCheckAddress)

	return config
}
