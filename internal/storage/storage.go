package storage

import (
	"context"
	"fmt"

	"github.com/m-lejwoda/batna-lublin/internal/db"
	"github.com/m-lejwoda/batna-lublin/internal/dto"
)

type InvestmentStorage struct {
	db *db.DB
}

func NewInvestmentStorage(db *db.DB) InvestmentStorage {
	return InvestmentStorage{db: db}
}

func (i InvestmentStorage) GetInvestments(ctx context.Context) ([]dto.InvestmentResponse, error) {
	rows, _ := i.db.Pool.Query(ctx, "SELECT * FROM investment")
}

func (i InvestmentStorage) GetInvestment(ctx context.Context, id int32) (dto.InvestmentResponse, error) {
	rows, err := i.db.Pool.Query(ctx, "SELECT id, custom_id, name, finish_date, office_location, phone, underground_parking_place, surface_parking_place, garage FROM investment WHERE id=$1", id)
	if err != nil {
		fmt.Println("Error Occured")
	}
	defer rows.Close()
	var investments []dto.InvestmentResponse
	for rows.Next() {
		var inv dto.InvestmentResponse
		rows.Scan(&inv.ID, &inv.CustomID, &inv.Name, &inv.FinishDate, &inv.OfficeLocation, &inv.Phone, &inv.UndergroundParkingPlace, &inv.SurfaceParkingPlace, &inv.Garage)
		investments = append(investments, inv)
	}
	return investments, nil
}

func (i InvestmentStorage) CreateInvestment(ctx context.Context, investment dto.CreateInvestmentRequest) (dto.InvestmentResponse, error) {
	i.db.Pool.Query()
}

func (i InvestmentStorage) GetFlat(ctx context.Context, id int32) (dto.FlatResponse, error) {
	i.db.Pool.Query()
}
