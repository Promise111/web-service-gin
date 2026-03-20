package NestedStruct

import (
	"fmt"
)

type Address struct {
	HouseNumber string `json:"houseNumber"`
	Street      string `json:"street"`
	City        string `json:"city"`
}

type Person struct {
	Name    string  `json:"name"`
	Address Address `json:"address"`
	Age     int     `json:"age"`
}

func updatePerson(p *Person) {
	p.Address.City = "Abuja"
}

func NestedStruct() {
	promise := Person{
		Name: "Promise Ihunna",
		Address: Address{
			HouseNumber: "17",
			Street:      "Ilubirin",
			City:        "Lagos",
		},
		Age: 29,
	}

	fmt.Println("Before update: ", promise)
	updatePerson(&promise)
	fmt.Println("After update: ", promise)
}
