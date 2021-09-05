package store

type StoreType int

const (
	_ StoreType = iota
	StoreTypeMySQL
	StoreTypeMySQLWithCache
)

var (
	defaultSimpleFactory = &SimpleFactory{}
)

type SimpleFactory struct{}

func (factory *SimpleFactory) Create(t StoreType) (storer Storer) {

	switch t {
	case StoreTypeMySQL:
		storer = &StorerMySQL{}

	case StoreTypeMySQLWithCache:
		storer = &StorerMySQLWithCache{
			StorerMySQL: &StorerMySQL{},
		}
	}

	return
}
