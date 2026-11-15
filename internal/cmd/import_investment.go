package cmd

import (
	"context"
	"fmt"

	service "github.com/m-lejwoda/batna-lublin/internal/services"
)

func ImportInterbudInvestment(ctx context.Context, NewInvestmentService service.InvestmentService) {
	fmt.Println("Import Interbud Investment")
	NewInvestmentService.GetInvestmentByName(ctx, "interbud")
}

func ImportMultimelInvestment() {
	fmt.Println("Import Multimel Investment")
}
