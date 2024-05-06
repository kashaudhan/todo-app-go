package actions

import "errors"

func (todos *Todos) Delete(index int) error {

	t := *todos

	if index < 0 || index > len(t) {
		return errors.New("index out of range")
	}

	t = append(t[:index-1], t[index:]...)

	*todos = t

	return nil
}