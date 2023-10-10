package types

import "fmt"

type BasicData struct {
	Name string
	Age  int
}

type Entyty struct {
	BasicData

	Country string
}

type SpecialEntyty struct {
	Entyty
	SpecialField bool
}

func BuildEntyty() {
	entyty := Entyty{
		BasicData: BasicData{
			Name: "Evgeny",
			Age:  31,
		},
		Country: "Poland",
	}

	specialEntyty := SpecialEntyty{
		Entyty: Entyty{
			BasicData: BasicData{
				Name: "Evgeny",
				Age:  31,
			},
			Country: "Poland",
		},
		SpecialField: true,
	}

	fmt.Println(entyty, specialEntyty.Name)
}
