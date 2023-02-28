package anchor

import (
	"fmt"
)

type Multi struct {
	operations []struct {
		name      string
		operation func(map[string]interface{}) (interface{}, error)
	}
	names map[string]struct{}
}

func NewMulti() *Multi {
	return &Multi{
		operations: []struct {
			name      string
			operation func(map[string]interface{}) (interface{}, error)
		}{},
		names: map[string]struct{}{},
	}
}

func (multi *Multi) RunMulti(name string, op func(map[string]interface{}) (interface{}, error)) error {
	if _, ok := multi.names[name]; ok {
		return fmt.Errorf("%s is already a member of the MultiTask", name)
	}

	multi.names[name] = struct{}{}
	multi.operations = append(multi.operations, struct {
		name      string
		operation func(map[string]interface{}) (interface{}, error)
	}{
		name:      name,
		operation: op,
	})

	return nil
}

func (mt *Multi) Run() (map[string]interface{}, error) {
	results := make(map[string]interface{})

	for i := 0; i < len(mt.operations); i++ {
		name, operation := mt.operations[i].name, mt.operations[i].operation
		result, err := operation(results)
		if err != nil {
			return nil, fmt.Errorf("%s: %s", name, err)
		}
		results[name] = result
	}

	return results, nil
}
