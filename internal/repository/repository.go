package repository

import (
	"context"

	"github.com/m-lejwoda/batna-lublin/internal/dto"
)

type InvestmentRepository interface {
	GetInvestments(ctx context.Context) ([]dto.InvestmentResponse, error)
	GetInvestment(ctx context.Context, id int32) (dto.InvestmentResponse, error)
	GetInvestmentByName(ctx context.Context, name string) (dto.InvestmentResponse, error)
	CreateInvestment(ctx context.Context, investment dto.CreateInvestmentRequest) (dto.InvestmentResponse, error)
	GetFlats(ctx context.Context, id int32) ([]dto.FlatResponse, error)
	GetFlat(ctx context.Context, flatID int32) (dto.FlatResponse, error)
	GetAllFlatStatuses(ctx context.Context, investmentId int32) (dto.FlatStatusResponse, error)
	CreateFlat(ctx context.Context, createFlat dto.CreateFlatRequest) (dto.FlatResponse, error)
	GetFlatsFromInvestment(ctx context.Context, investment string) ([]dto.FlatResponse, error)
}
