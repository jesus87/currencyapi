package server

import (
	"time"

	"github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"
	"github.com/jesus87/currencyapi/config"
	"github.com/jesus87/currencyapi/repository"
)

func SetupRute(cfg config.ConfigServer, currency *repository.CurrencyRepository) *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	timeoutMs := time.Duration(cfg.FreeCurrencyTimeOutMs) * time.Millisecond

	// Get currencies value
	r.GET("/currencies/:currency",
		timeout.New(
			timeout.WithTimeout(timeoutMs),
			timeout.WithHandler(getCurrenyHandler(currency)),
			timeout.WithResponse(timeOutResponse),
		),
	)
	return r
}
