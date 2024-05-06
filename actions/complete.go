package actions

import (
	"errors"
	"time"
)

func (todos *Todos) Complete(index int) error {
	list := *todos

	if index < 0 || index > len(list) {
		return errors.New("index out of range")
	}


	list[index - 1].Done = true
	list[index - 1].Completed_at = time.Now()

	return nil
}