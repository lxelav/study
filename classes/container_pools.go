package classes

import "errors"

type AllPools struct {
	pools map[string]*Pool
}

func InitAllPools() *AllPools {
	return &AllPools{
		make(map[string]*Pool),
	}
}

func (pools *AllPools) AddPool(name string) {
	(*pools).pools[name] = InitPool()
}

func (pools *AllPools) PopPool(name string) {
	delete((*pools).pools, name)
}

func (pools *AllPools) GetPool(name string) (*Pool, error) {
	return_el, ok := (*pools).pools[name]
	if !ok {
		return nil, errors.New("Элемент не найден!")
	}

	return return_el, nil
}
