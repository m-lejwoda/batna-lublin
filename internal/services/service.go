package service

import (
	"context"
	"fmt"

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
	investments, err := i.investmentRepo.GetInvestments(ctx)
	if err != nil {
		fmt.Println("List Investment Error")
	}
	return investments, nil
}

func (i InvestmentService) GetInvestment(ctx context.Context, id int32) (dto.InvestmentResponse, error) {

	investment, err := i.investmentRepo.GetInvestment(ctx, id)
	if err != nil {
		fmt.Println("List Investment Error")
	}
	return investment, nil

}

func (i InvestmentService) CreateInvestment(ctx context.Context, investment dto.CreateInvestmentRequest) (dto.InvestmentResponse, error) {
	createdInvestment, err := i.investmentRepo.CreateInvestment(ctx, investment)
	if err != nil {
		fmt.Println("List Investment Error")
	}
	return createdInvestment, nil

}

func (i InvestmentService) GetFlat(ctx context.Context, id int32) (dto.FlatResponse, error) {
	flat, err := i.investmentRepo.GetFlat(ctx, id)
	if err != nil {
		fmt.Println("List Investment Error")
	}
	return flat, nil

}
