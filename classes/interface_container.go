package classes

type Container interface {
	Insert(key string, value interface{}) error
	Get(key string) (string, error)
	GetRange(minValue, maxValue string) ([]string, error)
	Update(key, value interface{}) error
	Remove(key string) error
}
