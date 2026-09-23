package storage

import "github.com/m-lejwoda/batna-lublin/internal/db"

type InvestmentStorage struct {
	db *db.DB
}

func NewInvestmentStorage(db *db.DB) InvestmentStorage {
	return InvestmentStorage{db: db}
}
