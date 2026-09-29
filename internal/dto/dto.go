package dto

import "time"

type InvestmentResponse struct {
	ID                      int32     `json:"id"`
	CustomID                string    `json:"custom_id"`
	Name                    string    `json:"name"`
	FinishDate              time.Time `json:"finish_date"`
	OfficeLocation          string    `json:"office_location"`
	Phone                   string    `json:"phone"`
	UndergroundParkingPlace bool      `json:"underground_parking_place"`
	SurfaceParkingPlace     bool      `json:"surface_parking_place"`
	Garage                  bool      `json:"garage"`
}

type CreateInvestmentRequest struct {
	CustomID                string    `json:"custom_id"`
	Name                    string    `json:"name"`
	FinishDate              time.Time `json:"finish_date"`
	OfficeLocation          string    `json:"office_location"`
	Phone                   string    `json:"phone"`
	UndergroundParkingPlace bool      `json:"underground_parking_place"`
	SurfaceParkingPlace     bool      `json:"surface_parking_place"`
	Garage                  bool      `json:"garage"`
}

type FlatResponse struct {
	ID             int32  `json:"id"`
	CustomID       string `json:"custom_id"`
	Name           string `json:"name"`
	FloorArea      int32  `json:"floor_area"`
	Layout         string `json:"layout"`
	Floor          int32  `json:"floor"`
	FlatNumber     string `json:"flat_number"`
	BuildingNumber string `json:"building_number"`
	RoomsNumber    int32  `json:"rooms_number"`
	Balcony        bool   `json:"balcony"`
	InvestmentID   int32  `json:"investment_id"`
}

type ParkingPlaceResponse struct {
	ID           int32  `json:"id"`
	Type         string `json:"type"`
	FloorArea    int32  `json:"floor_area"`
	InvestmentID int32  `json:"investment_id"`
}

type FlatStatusResponse struct {
	Price     int32     `json:"price"`
	Currency  string    `json:"currency"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	FlatID    int32     `json:"flat_id"`
}

type ParkingPlaceStatusResponse struct {
	Price          int64     `json:"price"`
	Currency       string    `json:"currency"`
	Status         string    `json:"status"`
	CreatedAt      time.Time `json:"created_at"`
	ParkingPlaceId int32     `json:"parking_place_id"`
}

type InvestmentCommentResponse struct {
	Comment      string `json:"comment"`
	Importancy   int32  `json:"importancy"`
	InvestmentID int32  `json:"investment_id"`
}

type FlatCommentResponse struct {
	Comment    string `json:"comment"`
	Importancy int32  `json:"importancy"`
	FlatID     int32  `json:"flat_id"`
}
