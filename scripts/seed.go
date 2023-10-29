package main

import "learning_go/bookingApi/types"

func main() {
	horel := types.Hotel{
		Name:    "Hilton",
		Address: "1234 Main St",
	}

	room := types.Room{
		Type:      types.Single,
		BasePrice: 100,
		Price:     100,
	}
	println("Seeding...")
}
