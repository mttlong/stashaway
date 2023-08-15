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

type CustomerManager struct {
	nextID int
}

// New creates a new customer and auto-increments the ID.
func (cm *CustomerManager) New() *Customer {
	c := &Customer{
		ID: cm.nextID,
	}
	cm.nextID++
	return c
}
