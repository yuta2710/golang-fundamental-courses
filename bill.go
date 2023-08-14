package main

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

func newBill(name string) bill {
	newBill := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0.0,
	}

	return newBill
}
