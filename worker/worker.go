package worker

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/adnvilla/boletia/config"
	"github.com/adnvilla/boletia/repository"
)

func Start(cfg config.ConfigServer, rc *repository.CurrencyRepository, rh *repository.HistoryApiRepository) {
	tickMs := time.Duration(cfg.FreeCurrencyReloadLatestTimeMs) * time.Millisecond
	go func() {
		for _ = range time.Tick(tickMs) {
			getCurrencies(cfg, rc, rh)
		}
	}()
}

func getCurrencies(cfg config.ConfigServer, rc *repository.CurrencyRepository, rh *repository.HistoryApiRepository) {
	historyApi := repository.HistoryApi{
		CreatedAt: time.Now(),
	}

	defer func() {
		rh.Save(historyApi)
	}()

	start := time.Now()
	timeoutMs := time.Duration(cfg.FreeCurrencyTimeOutMs) * time.Millisecond
	ctx, cncl := context.WithTimeout(context.Background(), timeoutMs)
	defer cncl()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("https://api.currencyapi.com/v3/latest?apikey=%s", cfg.FreeCurrencyApiKey), nil)
	if err != nil {
		log.Println(err)
		historyApi.Error = err.Error()
		return
	}

	result, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		historyApi.Error = err.Error()
		return
	}
	defer result.Body.Close()

	elapsed := time.Since(start).Milliseconds()
	fmt.Println(elapsed)
	historyApi.TimeResponseMs = int(elapsed)

	body, err := ioutil.ReadAll(result.Body)
	if err != nil {
		log.Println(err)
		historyApi.Error = err.Error()
		return
	}

	var currencyLatest CurrencyLatest
	if err := json.Unmarshal(body, &currencyLatest); err != nil {
		log.Println(err)
		historyApi.Error = err.Error()
		return
	}

	currencies, lastUpdatedAt := ApiModelToRepoModel(currencyLatest)
	go rc.Save(currencies, lastUpdatedAt)
}
