package repository

import "context"

type InvestmentRepository interface {
	GetInvestments(ctx context.Context) ([]*Investment, error)
	GetSingleInvestment(ctx context.Context, id int32) (*Investment, error)
}
