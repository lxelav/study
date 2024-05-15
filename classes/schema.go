package classes

import "errors"

type Schema struct {
	containers map[string]Container
}

func InitSchema() *Schema {
	return &Schema{
		make(map[string]Container),
	}
}

func (schema *Schema) GetContainer(name string) (Container, error) {
	return_el, ok := (*schema).containers[name]
	if !ok {
		return nil, errors.New("Элемент не найден!")
	}

	return return_el, nil
}

func (schema *Schema) AddContainer(name, containerType string) error {
	//Какой контейнер Вы хотите испольщовать и дальше я буду добавлять контейнер
	//TODO
	return nil
}

func (schema *Schema) PopContainer(name string) {
	delete(schema.containers, name)
}
