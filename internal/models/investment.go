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
	ID             int
	CustomID 			 string
	Name           string
	FloorArea      int64
	Layout         string
	Floor          int64
	FlatNumber     string
	BuildingNumber string
	RoomsNumber    int64
	Balcony        bool
	InvestmentID   int32
}

type ParkingPlace struct {
	ID           int
	Type         string
	FloorArea    int64
	InvestmentID int32
}

type FlatStatus struct {
	Price     int64
	Currency  string
	Status    string
	CreatedAt time.Time
	FlatID    int32
}

type ParkingPlaceStatus struct {
	Price          int64
	Currenct       string
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
	comment string
	importancy int32
	FlatID int32
}
