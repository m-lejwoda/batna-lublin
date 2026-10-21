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
	rows, err := i.db.Pool.Query(ctx, "SELECT id, custom_id, name, finish_date, office_location, phone, underground_parking_place, surface_parking_place, garage FROM investment")
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

func (i InvestmentStorage) GetInvestment(ctx context.Context, id int32) (dto.InvestmentResponse, error) {
	row := i.db.Pool.QueryRow(ctx, "SELECT id, custom_id, name, finish_date, office_location, phone, underground_parking_place, surface_parking_place, garage FROM investment WHERE id=$1", id)
	var inv dto.InvestmentResponse
	row.Scan(&inv.ID, &inv.CustomID, &inv.Name, &inv.FinishDate, &inv.OfficeLocation, &inv.Phone, &inv.UndergroundParkingPlace, &inv.SurfaceParkingPlace, &inv.Garage)
	return inv, nil
}

func (i InvestmentStorage) CreateInvestment(ctx context.Context, investment dto.CreateInvestmentRequest) (dto.InvestmentResponse, error) {
	row := i.db.Pool.QueryRow(ctx, `INSERT INTO (custom_id, name, finish_date, office_location, phone, underground_parking_place, surface_parking_place, garage)
	Values($1, $2, $3, $4, $5, $6, $7, $8)
	Returning id, custom_id, name, finish_date, office_location, phone, underground_parking_place, surface_parking_place, garage
	`, investment.CustomID, investment.Name, investment.FinishDate, investment.OfficeLocation, investment.Phone, investment.UndergroundParkingPlace, investment.SurfaceParkingPlace, investment.Garage)
	var inv dto.InvestmentResponse
	row.Scan(&inv.ID, &inv.CustomID, &inv.Name, &inv.FinishDate, &inv.OfficeLocation, &inv.Phone, &inv.UndergroundParkingPlace, &inv.SurfaceParkingPlace, &inv.Garage)
	return inv, nil
}

func (i InvestmentStorage) GetFlats(ctx context.Context, id int32) ([]dto.FlatResponse, error) {
	rows, err := i.db.Pool.Query(ctx, "SELECT id, custom_id, name, floor_area, layout, floor, flat_number, building_number, rooms_number, balcony, investment_id From flat WHERE investment_id=$1", id)
	if err != nil {
		fmt.Println(err)
	}
	var flats []dto.FlatResponse
	defer rows.Close()
	for rows.Next() {
		var flat dto.FlatResponse
		rows.Scan(&flat.ID, &flat.CustomID, &flat.Name, &flat.FloorArea, &flat.Layout, &flat.Floor, &flat.FlatNumber, &flat.BuildingNumber, &flat.RoomsNumber, &flat.Balcony, &flat.InvestmentID)
		flats = append(flats, flat)
	}
	return flats, nil
}

func (i InvestmentStorage) GetFlat(ctx context.Context, flatID int32) (dto.FlatResponse, error) {
	row := i.db.Pool.QueryRow(ctx, "SELECT id, custom_id, name, floor_area, layout, floor, flat_number, building_number, rooms_number, balcony, investment_id from flat WHERE id=$1", flatID)
	var flat dto.FlatResponse
	row.Scan(&flat.ID, &flat.CustomID, &flat.Name, &flat.FloorArea, &flat.Layout, &flat.Floor, &flat.BuildingNumber, &flat.RoomsNumber, &flat.Balcony, &flat.InvestmentID)
	return flat, nil
}

func (i InvestmentStorage) CreateFlat(ctx context.Context, createFlat dto.CreateFlatRequest) (dto.FlatResponse, error) {
	row := i.db.Pool.QueryRow(ctx, `INSERT INTO (custom_id, name, floor_area, layout, floor, flat_number, building_number, rooms_number, balcony, investment_id)
	Values($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	Returning id, custom_id, name, floor_area, layout, floor, flat_number, building_number, rooms_number, balcony, investment_id`,
		createFlat.CustomID, createFlat.Name, createFlat.FloorArea, createFlat.Layout, createFlat.Floor, createFlat.FlatNumber, createFlat.BuildingNumber, createFlat.RoomsNumber, createFlat.RoomsNumber, createFlat.Balcony, createFlat.InvestmentID)
	var flat dto.FlatResponse
	row.Scan(&flat.ID, &flat.CustomID, &flat.Name, &flat.FloorArea, &flat.Layout, &flat.Floor, &flat.FlatNumber, &flat.BuildingNumber, &flat.RoomsNumber, &flat.Balcony, &flat.InvestmentID)
	return flat, nil
}
