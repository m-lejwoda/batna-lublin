package models

import "time"

type Investment struct {
	ID                      int32
	CustomID 								string
	Name                    string
	FinishDate              time.Time
	OfficeLocation          string
	Phone                   string
	UndergroundParkingPlace bool
	SurfaceParkingPlace     bool
	Garage                  bool
}

type Flat struct {
	ID             int32
	CustomID 			 string
	Name           string
	FloorArea      int32
	Layout         string
	Floor          int32
	FlatNumber     string
	BuildingNumber string
	RoomsNumber    int32
	Balcony        bool
	InvestmentID   int32
}

type ParkingPlace struct {
	ID           int
	Type         string
	FloorArea    int32
	InvestmentID int32
}

type FlatStatus struct {
	Price     int32
	Currency  string
	Status    string
	CreatedAt time.Time
	FlatID    int32
}

type ParkingPlaceStatus struct {
	Price          int64
	Currency       string
	Status         string
	CreatedAt      time.Time
	ParkingPlaceID int32
}

type InvestmentComment struct {
	Comment string
	Importancy int32
	InvestmentID int32 
}

type FlatComment struct {
	Comment string
	Importancy int32
	FlatID int32
}
