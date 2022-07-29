package main

import (
	"fmt"

	"github.com/jesus87/currencyapi/config"
	"github.com/jesus87/currencyapi/repository"
	"github.com/jesus87/currencyapi/server"
	"github.com/jesus87/currencyapi/worker"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dsn = "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable"

func main() {
	cfg := config.LoadConfig("config.json")

	dnsf := fmt.Sprintf(dsn, cfg.PostgresHost, cfg.PostgresUser, cfg.PostgresPass, cfg.PostgresDatabase, cfg.PostgresPort)
	db, err := gorm.Open(postgres.Open(dnsf), &gorm.Config{SkipDefaultTransaction: true})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&repository.Currencies{})
	db.AutoMigrate(&repository.HistoryApi{})

	rc := repository.NewCurrency(db)
	hapi := repository.NewHistoryApi(db)

	fmt.Println(cfg.FreeCurrencyReloadLatestTimeMs)
	worker.Start(cfg, rc, hapi)

	r := server.SetupRute(cfg, rc)
	// Listen and Server in 0.0.0.0:8080
	r.Run(":9090")
}
