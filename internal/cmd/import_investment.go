package cmd

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/m-lejwoda/batna-lublin/internal/dto"
	service "github.com/m-lejwoda/batna-lublin/internal/services"
)

func ImportInterbudInvestment(ctx context.Context, NewInvestmentService service.InvestmentService) {
	fmt.Println("Import Interbud Investment")
	_, err := NewInvestmentService.GetInvestmentByName(ctx, "interbud")
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			investment := dto.CreateInvestmentRequest{
				CustomID:                "interbud",
				Name:                    "interbud",
				FinishDate:              time.Date(2027, time.July, 31, 12, 0, 0, 0, time.Local),
				OfficeLocation:          "Lublin, Al. Racławickie 8 lok. 39",
				Phone:                   "691735100",
				UndergroundParkingPlace: true,
				SurfaceParkingPlace:     true,
				Garage:                  false,
			}
			NewInvestmentService.CreateInvestment(ctx, investment)
			log.Println("Created Interbud investment")
		}
	}
}

func ImportMultimelInvestment(ctx context.Context, NewInvestmentService service.InvestmentService) {
	fmt.Println("Import Multimel Investment")
	_, err := NewInvestmentService.GetInvestmentByName(ctx, "multimel")
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			investment := dto.CreateInvestmentRequest{
				CustomID:                "multimel",
				Name:                    "multimel",
				FinishDate:              time.Date(2027, time.February, 28, 12, 0, 0, 0, time.Local),
				OfficeLocation:          "Dominów 130A 20-388 Lublin",
				Phone:                   "502231423",
				UndergroundParkingPlace: true,
				SurfaceParkingPlace:     true,
			}
			NewInvestmentService.CreateInvestment(ctx, investment)
		}
	}
}
