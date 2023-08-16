package customers

import (
	"fmt"
	"log"

	"github.com/mttlong/stashaway/model"
)

// Customer represents a client with an ID and a list of portfolios
type customer struct {
	ID         int
	Portfolios []model.Portfolio
}

type customerManager struct {
	nextID        int
	customerTable map[int]*customer
}

// New creates a new customer and auto-increments the ID.
func (cm *customerManager) new() *customer {
	c := &customer{
		ID: cm.nextID,
	}
	cm.nextID++
	return c
}

func (c *customer) addPortfolio(p *model.Portfolio) {
	c.Portfolios = append(c.Portfolios, *p)
}

var cm *customerManager

func Init() {
	if cm == nil {
		cm = &customerManager{}
		cm.customerTable = make(map[int]*customer)
	} else {
		log.Println("cm is already initialised")
	}
}

func CreateCustomer() (int, error) {
	if cm == nil {
		return -1, fmt.Errorf("call the Init function first")
	}
	customer := cm.new()
	cm.customerTable[customer.ID] = customer
	return customer.ID, nil
}

func AddPortfolio(CustomerID int, PortfolioType model.PortfolioType, Allocation float64) {
	value, ok := cm.customerTable[CustomerID]

	if ok {
		value.addPortfolio(&model.Portfolio{
			Name:       PortfolioType,
			Allocation: Allocation,
		})
	} else {
		log.Println("Cannot find the Customer ID provided.")
	}
}
