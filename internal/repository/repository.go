package repository

type InvestmentRepository interface {
	GetInvestments()
	GetSingleInvestment(id int32)
}
