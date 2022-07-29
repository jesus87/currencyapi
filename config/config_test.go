package config

import (
	"errors"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestGetConfig(t *testing.T) {
	viper.Reset() // is for testing
	cfg := LoadConfig("./data_test/config.json")

	assert.Equal(t, 20, cfg.FreeCurrencyTimeOutMs)
	assert.Equal(t, "qwerty", cfg.FreeCurrencyApiKey)
	assert.Equal(t, "127.0.0.1", cfg.PostgresHost)
	assert.Equal(t, "dbname", cfg.PostgresDatabase)
	assert.Equal(t, "pgyser", cfg.PostgresUser)
	assert.Equal(t, "postgres", cfg.PostgresPass)
	assert.Equal(t, "5436", cfg.PostgresPort)
}

func TestGetConfigPanic(t *testing.T) {
	viper.Reset() // is for testing
	assert.Panics(t, func() {
		cfg := LoadConfig("")

		assert.Error(t, errors.New("Expected panic!"), cfg)
	})

}

func TestGetBadConfigPanic(t *testing.T) {
	viper.Reset() // is for testing
	assert.Panics(t, func() {
		cfg := LoadConfig("./data_test/bad_config.json")

		assert.Error(t, errors.New("Expected panic!"), cfg)
	})

}
