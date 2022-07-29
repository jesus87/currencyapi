package repository

import (
	"time"

	"gorm.io/gorm"
)

type HistoryApiRepository struct {
	db *gorm.DB
}

type HistoryApi struct {
	ID             uint `gorm:"primarykey"`
	CreatedAt      time.Time
	TimeResponseMs int
	Error          string
}

func NewHistoryApi(db *gorm.DB) *HistoryApiRepository {
	return &HistoryApiRepository{
		db: db,
	}
}

func (r *HistoryApiRepository) Save(historyApi HistoryApi) error {

	errt := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&historyApi).Error; err != nil {
			return err
		}
		return nil
	})

	return errt
}
