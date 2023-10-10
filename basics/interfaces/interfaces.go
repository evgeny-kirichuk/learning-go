package interfaces

type Putter interface {
	Put(id int, val any) error
}

type Getter interface {
	Get(id int) (any, error)
}

type Storage interface {
	Putter
	Getter
}


type MyStorage struct {
	Storage
}

func (m MyStorage) Get(id int) (any, error) {
	return "value", nil
}

func (m MyStorage) Put(id int, val any) error {
	return nil
}

type Server struct {
	store Storage
}

func UpdateValue(id int, val any, p Putter) error {
	return p.Put(id, val)
}

func GetValue(id int, g Getter) (any, error) {
	return g.Get(id)
}

func RunInterfaces() {
	s := &Server{
		store: &MyStorage{},
	}
	UpdateValue(1, "value", s.store)
	GetValue(1, s.store)
}
