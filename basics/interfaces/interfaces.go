package interfaces

import "fmt"

type NumberStorer interface {
	GetAll() ([]int, error)
	Put(int) error
}

type PotgresNumberStore struct {
	// ...
}

type MongoDBNumberStore struct {
	// ...
}

func (m MongoDBNumberStore) GetAll() ([]int, error) {
	return []int{1,2,3}, nil
}

func (m MongoDBNumberStore) Put(number int) error {
	return nil
}

func (p PotgresNumberStore) GetAll() ([]int, error) {
	return []int{1,2,3}, nil
}

func (p PotgresNumberStore) Put(number int) error {
	return nil
}

type ApiServer struct {
	numberStore NumberStorer
}

func RunInterfaces() {
	apiServer := ApiServer{
		numberStore: MongoDBNumberStore{},
	}
	result, err := apiServer.numberStore.GetAll()
	fmt.Println(result, err)
}