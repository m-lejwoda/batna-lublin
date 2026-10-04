package handler

import (
	"fmt"
	"net/http"

	service "github.com/m-lejwoda/batna-lublin/internal/services"
)

type InvestmentHandler struct {
	service service.InvestmentService
}

func NewInvestmentHandler(service service.InvestmentService) InvestmentHandler {
	return InvestmentHandler{service: service}
}

func RegisterInvestmentRoutes(mux *http.ServeMux, h *InvestmentHandler) {
	mux.HandleFunc("GET /investments", h.listInvestmentHandler)
	mux.HandleFunc("GET /investment/{id}", h.getInvestment)
	mux.HandleFunc("POST /investment", h.createInvestment)
	mux.HandleFunc("GET /flat/{id}", h.getFlat)
}

func (h *InvestmentHandler) listInvestmentHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("investments list")
}

func (h *InvestmentHandler) getInvestment(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get Investment")
}

func (h *InvestmentHandler) createInvestment(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create Investment")

}

func (h *InvestmentHandler) getFlat(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get Flat")
}
