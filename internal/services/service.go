package service

import "github.com/m-lejwoda/batna-lublin/internal/repository"

type InvestmentService struct {
	investmentRepo repository.InvestmentRepository
}

func NewInvestmentService(investmentRepo repository.InvestmentRepository) InvestmentService {
	return InvestmentService{investmentRepo: investmentRepo}
}

func (i InvestmentService) GetFlat() {
	i.investmentRepo.GetInvestment()
}
