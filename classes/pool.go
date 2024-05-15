package classes

import "errors"

type Pool struct {
	schemas map[string]*Schema
}

func InitPool() *Pool {
	return &Pool{
		make(map[string]*Schema),
	}
}

func (pool *Pool) GetSchema(schemaName string) (*Schema, error) {
	return_el, ok := (*pool).schemas[schemaName]
	if !ok {
		return nil, errors.New("Элемент не найден!")
	}

	return return_el, nil
}

func (pool *Pool) AddSchema(name string) {
	pool.schemas[name] = InitSchema()
}

func (pool *Pool) PopSchema(name string) {
	delete(pool.schemas, name)
}
