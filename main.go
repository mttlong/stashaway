package main

import (
	"fmt"

	"github.com/mttlong/stashaway/model"
)

func main() {
	cm := &model.CustomerManager{}
	customerTable := make(map[int]*model.Customer)

	customer1 := cm.New()
	customer1.AddPortfolio(&model.Portfolio{
		Name:           model.HighRisk,
		OnetimeDeposit: 1000.0,
		MonthlyDeposit: 100.0,
	})
	fmt.Println("Customer ID: ", customer1.ID)
	customerTable[customer1.ID] = customer1

	customer2 := cm.New()
	customer2.AddPortfolio(&model.Portfolio{
		Name:           model.Retirement,
		OnetimeDeposit: 1500.0,
		MonthlyDeposit: 50.0,
	})
	customerTable[customer2.ID] = customer2
	fmt.Println("Customer ID: ", customer2.ID)

	for id, c := range customerTable {
		fmt.Println("ID: ", id)
		fmt.Println("Customer: ", *c)
	}
}
