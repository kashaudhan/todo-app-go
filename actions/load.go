package actions

import (
	"encoding/json"
	"errors"
	"os"
)

func (todos *Todos) Load(todosPath string) error {
	byteValues, err := os.ReadFile(todosPath)

	if err != nil {
		return errors.New("failed to open file")
	}

	json.Unmarshal(byteValues, todos)

	return nil
}