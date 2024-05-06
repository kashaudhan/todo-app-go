package actions

import "fmt"

func (todos *Todos) List() {
	fmt.Println("index, created at, task, done, completed at")
		for index, val := range *todos {
			fmt.Println(index+1, val.Created_at, val.Task, val.Done, val.Completed_at)
		}
}