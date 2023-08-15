package model

type PortfolioType int

const (
	HighRisk PortfolioType = iota
	Retirement
)

// Portfolio represents the financial portfolio of a customer
type Portfolio struct {
	Name           PortfolioType
	OnetimeDeposit float64
	MonthlyDeposit float64
	TotalDeposit   float64
}

// Customer represents a client with an ID and a list of portfolios
type Customer struct {
	ID         int
	Portfolios []Portfolio
}
