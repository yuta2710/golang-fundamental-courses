package main

import "fmt"

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

func newBill(name string) bill {
	newBill := bill{
		name:  name,
		items: map[string]float64{"pie": 0.05, "cake": 999.99},
		tip:   0.0,
	}

	return newBill
}

// Format the bill
func (b bill) format() string {
	fs := "Bill breakdown \n"
	var total float64 = 0

	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ...$%v \n", k+":", v) // -25v: create 25 empty space
		total += v
	}

	fs += fmt.Sprintf("%-25v ...$%0.2f", "total: ", total)

	return fs
}
