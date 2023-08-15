package main

import (
	"fmt"

	"github.com/mttlong/stashaway/model"
)

func main() {
	cm := &model.CustomerManager{}

	customer1 := cm.New()
	customer2 := cm.New()

	customer1.Portfolios = append(customer1.Portfolios, model.Portfolio{
		Name:           model.HighRisk,
		OnetimeDeposit: 1000.0,
		MonthlyDeposit: 100.0,
		TotalDeposit:   1200.0,
	})

	fmt.Println(customer1)
	fmt.Println(customer2)
}
