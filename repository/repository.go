package repository

import (
	"errors"
	"time"
)

var (
	errUpdateFailed = errors.New("update failed")
	errDeleteFailed = errors.New("delete failed")
)

type Repository interface {
	Migrate() error
	AllHoldings() ([]Holdings, error)
	GetHoldingByID(id int64) (*Holdings, error)
	InsertHolding(h Holdings) (*Holdings, error)
	UpdateHolding(id int64, updated Holdings) error
	DeleteHolding(id int64) error
}

type Holdings struct {
	ID            int64     `json:"id"`
	Amount        int       `json:"amount"`
	PurchaseDate  time.Time `json:"purchase_date"`
	PurchasePrice int       `json:"purchase_price"`
}
