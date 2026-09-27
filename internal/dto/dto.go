package dto

type InvestmentResponse struct {
	ID int32
	CustomID string
	Name string
	FinishDate time.Time
	OfficeLocation string
	Phone string
	UndergroundParkingPlace bool
	SurfaceParkingPlace bool
	Garage bool
}

type CreateInvestment struct {
	CustomID string
	Name string
	FinishDate time.Time
	OfficeLocation string
	Phone string
	UndergroundParkingPlace bool
	SurfaceParkingPlace bool
	Garage bool
}

type FlatResponse struct {
	ID int32
	CustomID string
	Name string
	FloorArea int32
	Layout string
	Floor int32
	FlatNumber string
	BuildingNumber string
	RoomsNumber int32
	Balcony bool
	InvestmentID int32
}

type ParkingPlaceResponse struct {
	ID int32
	Type string
	FloorArea int32
	InvestmentID int32
}

type FlatStatusResponse struct {
	Price int32
	Currency string
	Status string
	CreatedAt time.Time
	FlatID int32
}

type ParkingPlaceStatusResponse struct {
	Price int64
	Currency string
	Status string
	CreatedAt time.Time
	ParkingPlaceId int32
}

type InvestmentCommentResponse struct {
	Comment string
	Importancy int32
	InvestmentID int32
}

type FlatCommentResponse struct {
	Comment string
	Importancy int32
	FlatID int32
}
