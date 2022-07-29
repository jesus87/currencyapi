package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/adnvilla/boletia/repository"
	"github.com/gin-gonic/gin"
)

const (
	layout = "2006-01-02T15:04:05"
)

func getCurrenyHandler(rc *repository.CurrencyRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		currency := c.Params.ByName("currency")

		if currency == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Currency is empty"})
		}

		var param Parameters
		var finit string
		var fend string
		if c.ShouldBind(&param) == nil {
			fmt.Println(param.Finit)
			fmt.Println(param.Fend)
			finit = param.Finit
			fend = param.Fend
		}

		_, err := time.Parse(layout, finit)
		if err != nil && finit != "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "finit is not date"})
			return
		}
		_, err = time.Parse(layout, fend)
		if err != nil && fend != "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "fend is not date"})
			return
		}
		currencies, err := rc.Query(currency, finit, fend)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ups!"})
			return
		}

		if len(currencies) == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "We have no currency records"})
			return
		}

		c.JSON(http.StatusOK, currencies)
	}
}

type Parameters struct {
	Finit string `form:"finit"`
	Fend  string `form:"fend"`
}

func timeOutResponse(c *gin.Context) {
	c.String(http.StatusRequestTimeout, "timeout")

	// registrar intento fallido
}
