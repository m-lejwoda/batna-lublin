package handler

import service "github.com/m-lejwoda/batna-lublin/internal/services"

type InvestmentHandler struct {
	service service.InvestmentService
}

func NewInvestmentHandler(service service.InvestmentService) InvestmentHandler {
	return InvestmentHandler{service: service}
}
