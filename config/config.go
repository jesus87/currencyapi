package config

import (
	"github.com/spf13/viper"
)

const (
	freeCurrencyReloadLatestTimeMs = "FreeCurrencyReloadLatestTimeMs"
	freeCurrencyTimeOutMs          = "FreeCurrencyTimeOutMs"
	freeCurrencyApiKey             = "FreeCurrencyApiKey"
	postgresHost                   = "PostgresHost"
	postgresDatabase               = "PostgresDatabase"
	postgresUser                   = "PostgresUser"
	postgresPass                   = "PostgresPass"
	postgresPort                   = "PostgresPort"
)

type ConfigServer struct {
	FreeCurrencyReloadLatestTimeMs int
	FreeCurrencyTimeOutMs          int
	FreeCurrencyApiKey             string
	PostgresHost                   string
	PostgresDatabase               string
	PostgresUser                   string
	PostgresPass                   string
	PostgresPort                   string
}

func LoadConfig() ConfigServer {

	viper.SetConfigFile("config.json")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			panic(err)
		} else {
			// Config file was found but another error was produced
			panic(err)
		}
	}

	return ConfigServer{
		FreeCurrencyReloadLatestTimeMs: viper.GetInt(freeCurrencyReloadLatestTimeMs),
		FreeCurrencyTimeOutMs:          viper.GetInt(freeCurrencyTimeOutMs),
		FreeCurrencyApiKey:             viper.GetString(freeCurrencyApiKey),
		PostgresHost:                   viper.GetString(postgresHost),
		PostgresDatabase:               viper.GetString(postgresDatabase),
		PostgresUser:                   viper.GetString(postgresUser),
		PostgresPass:                   viper.GetString(postgresPass),
		PostgresPort:                   viper.GetString(postgresPort),
	}
}
