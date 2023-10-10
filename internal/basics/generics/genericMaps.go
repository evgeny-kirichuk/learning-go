package generics

import "fmt"

type CustomMap[K comparable, V any] struct {
	data map[K]V
}

func (m *CustomMap[K, V]) Insert(k K, v V) error {
	m.data[k] = v
	return nil
}

func NewCustomMap[K comparable, V any]() *CustomMap[K, V] {
	return &CustomMap[K, V]{
		data: make(map[K]V),
	}
}

func RunGenericMaps() {
	m1 := NewCustomMap[string, int64]()
	m2 := NewCustomMap[string, float64]()

	m1.Insert("first", 34)
	m1.Insert("second", 12)

	m2.Insert("first", 35.98)
	m2.Insert("second", 26.99)

	fmt.Printf("Generic maps: %v and %v\n", m1, m2  )

}