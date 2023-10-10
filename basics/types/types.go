package types

import (
	"fmt"
	"strings"
)

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

type TransformFunc func(s string) string

func Uppercase(s string) string {
	return strings.ToUpper(s)
}

func Prefixer(prefix string) TransformFunc {
	return func(s string) string {
		return prefix + s
	}
}

func transformString(s string, fn TransformFunc) string {
	return fn(s)
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

	fmt.Println(entyty, transformString(specialEntyty.Name, Uppercase))
	fmt.Println(transformString(specialEntyty.Name, Prefixer("Mr. ")))
}
