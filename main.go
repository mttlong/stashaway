package main

import (
	"fmt"

	customers "github.com/mttlong/stashaway/manager"
)

func main() {

	customers.Init()
	referenceID, error := customers.CreateCustomer()
	if error != nil {
		fmt.Println("Error: ", error)
	} else {
		fmt.Println("Customer referenceID: ", referenceID)
	}

}
