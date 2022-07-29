package server

import (
	"time"

	"github.com/adnvilla/boletia/config"
	"github.com/adnvilla/boletia/repository"
	"github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"
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
