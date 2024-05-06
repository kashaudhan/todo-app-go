package actions

import (
	"time"
)

func (todos *Todos) Add(task string) {
	todo := Item{
		Task: task,
		Done: false,
		Created_at: time.Now(),
		Completed_at: time.Time{},
	}

	*todos = append(*todos, todo)
}