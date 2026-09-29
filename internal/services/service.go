package service

import (
	"fmt"

	"github.com/m-lejwoda/batna-lublin/internal/repository"
)

type InvestmentService struct {
	investmentRepo repository.InvestmentRepository
}

func NewInvestmentService(investmentRepo repository.InvestmentRepository) InvestmentService {
	return InvestmentService{investmentRepo: investmentRepo}
}

func (i InvestmentService) GetInvestment() {
	i.investmentRepo.GetInvestments()
}

func (i InvestmentService) GetFlat() {
	i.investmentRepo.GetInvestment()
	fmt.Println("test")
}
