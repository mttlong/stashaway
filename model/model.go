package model

type PortfolioType int

const (
	HighRisk PortfolioType = iota
	Retirement
	USEquities
	GovernmentBonds
	RealEstate
)

// Portfolio represents the financial portfolio of a customer
type Portfolio struct {
	Name         PortfolioType
	Allocation   float64
	TotalDeposit float64
}
