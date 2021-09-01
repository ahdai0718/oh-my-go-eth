package store

type StoreType int

const (
	StoreTypeMySQL StoreType = iota
)

type SimpleFactory struct{}

func (factory *SimpleFactory) Create(t StoreType) (storer Storer) {

	switch t {
	case StoreTypeMySQL:
		storer = &StorerMySQL{}
	}

	return
}
