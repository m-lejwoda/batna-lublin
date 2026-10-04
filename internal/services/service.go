package service

import (
	"context"

	"github.com/m-lejwoda/batna-lublin/internal/dto"
	"github.com/m-lejwoda/batna-lublin/internal/repository"
)

type InvestmentService struct {
	investmentRepo repository.InvestmentRepository
}

func NewInvestmentService(investmentRepo repository.InvestmentRepository) InvestmentService {
	return InvestmentService{investmentRepo: investmentRepo}
}

func (i InvestmentService) GetInvestments(ctx context.Context) ([]dto.InvestmentResponse, error) {
	i.investmentRepo.GetInvestments(ctx)
}

func (i InvestmentService) GetInvestment(ctx context.Context, id int32) (dto.InvestmentResponse, error) {
	i.investmentRepo.GetInvestment(id)
}

func (i InvestmentService) CreateInvestment(ctx context.Context, investment dto.CreateInvestmentRequest) (dto.InvestmentResponse, error) {
	i.investmentRepo.CreateInvestment(ctx, investment)
}

func (i InvestmentService) GetFlat(ctx context.Context, id int32) (dto.FlatResponse, error) {
	i.investmentRepo.GetFlat(ctx, id)
}
