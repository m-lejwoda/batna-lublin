package models

import "time"

type Investment struct {
	ID                      string
	Name                    string
	FinishDate              time.Time
	OfficeLocation          string
	Phone                   string
	UndergroundParkingPlace bool
	SurfaceParkingPlace     bool
	Garage                  bool
}

type Flat struct {
	ID             string
	Name           string
	FloorArea      int64
	Layout         string
	Floor          int64
	FlatNumber     string
	BuildingNumber string
	RoomsNumber    int64
	Balcony        bool
	InvestmentID   string
}

type ParkingPlace struct {
	ID           string
	Type         string
	FloorArea    int64
	InvestmentID string
}

type FlatStatus struct {
	Price     int64
	Currency  string
	Status    string
	FlatID    string
	CreatedAt time.Time
}

type ParkingPlaceStatus struct {
	Price          int64
	Currenct       string
	Status         string
	CreatedAt      time.Time
	ParkingPlaceID string
}
