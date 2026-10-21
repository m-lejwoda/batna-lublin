package repository

import (
	"context"

	"github.com/m-lejwoda/batna-lublin/internal/dto"
)

type InvestmentRepository interface {
	GetInvestments(ctx context.Context) ([]dto.InvestmentResponse, error)
	GetInvestment(ctx context.Context, id int32) (dto.InvestmentResponse, error)
	CreateInvestment(ctx context.Context, investment dto.CreateInvestmentRequest) (dto.InvestmentResponse, error)
	GetFlats(ctx context.Context, id int32) ([]dto.FlatResponse, error)
	GetFlat(ctx context.Context, flatID int32) (dto.FlatResponse, error)
	CreateFlat(ctx context.Context, createFlat dto.CreateFlatRequest) (dto.FlatResponse, error)
}
