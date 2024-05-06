package actions

import (
	"encoding/json"
	"os"
)

func (todos *Todos) Save(todosPath string) error {


	t, _ := json.Marshal(&todos)
	err := os.WriteFile(todosPath, t, 0644)

	return err
}