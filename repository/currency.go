package repository

import (
	"errors"
	"strings"
	"time"

	"gorm.io/gorm"
)

var ErrFinitDateParse = errors.New("The finit is invald")
var ErrFendDateParse = errors.New("The fend is invald")

const (
	layout = "2006-01-02T15:04:05"
)

type CurrencyRepository struct {
	db *gorm.DB
}

type Currencies struct {
	ID            uint   `gorm:"primarykey"`
	Currency      string `gorm:"index:idx_currency_last_updated_at"`
	Value         string
	CreatedAt     time.Time
	LastUpdatedAt time.Time `gorm:"index:idx_currency_last_updated_at"`
}

func NewCurrency(db *gorm.DB) *CurrencyRepository {
	return &CurrencyRepository{
		db: db,
	}
}

func (r *CurrencyRepository) Save(currencies []Currencies, updatedAt time.Time) error {
	errt := r.db.Transaction(func(tx *gorm.DB) error {

		for _, currency := range currencies {
			if err := tx.
				Where(Currencies{Currency: currency.Currency, LastUpdatedAt: currency.LastUpdatedAt, Value: currency.Value}).
				FirstOrCreate(&currency).Error; err != nil {
				return err
			}
		}
		return nil
	})

	return errt
}

func (r *CurrencyRepository) Query(currency string, start, end string) ([]Currencies, error) {
	curencies := []Currencies{}
	currency = strings.ToUpper(currency)

	tstart, err := time.Parse(layout, start)
	if err != nil && start != "" {
		return nil, ErrFinitDateParse
	}
	tend, err := time.Parse(layout, end)
	if err != nil && end != "" {
		return nil, ErrFendDateParse
	}

	if currency == "FULL" {
		if start != "" && end != "" {
			if err := r.db.Where("created_at > ? AND created_at < ?", tstart, tend).Find(&curencies).Error; err != nil {
				return nil, err
			}
		}
		if start == "" && end != "" {
			if err := r.db.Where("AND created_at < ?", tend).Find(&curencies).Error; err != nil {
				return nil, err
			}
		}
		if start != "" && end == "" {
			if err := r.db.Where("AND created_at > ?", tstart).Find(&curencies).Error; err != nil {
				return nil, err
			}
		}
		if start == "" && end == "" {
			if err := r.db.Find(&curencies).Error; err != nil {
				return nil, err
			}
		}
		return curencies, nil
	}

	if start != "" && end != "" {
		if err := r.db.Where("currency = ? AND created_at > ? AND created_at < ?", currency, tstart, tend).Find(&curencies).Error; err != nil {
			return nil, err
		}
	}
	if start == "" && end != "" {
		if err := r.db.Where("currency = ? AND created_at < ?", currency, tend).Find(&curencies).Error; err != nil {
			return nil, err
		}
	}
	if start != "" && end == "" {
		if err := r.db.Where("currency = ? AND created_at > ?", currency, tstart).Find(&curencies).Error; err != nil {
			return nil, err
		}
	}
	if start == "" && end == "" {
		if err := r.db.Where("currency = ?", currency).Find(&curencies).Error; err != nil {
			return nil, err
		}
	}

	return curencies, nil
}
