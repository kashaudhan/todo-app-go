package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
	"todo_app/actions"
)

var todosPath = "todo.json"

func main() {

	add := flag.Bool("add", false, "add a task")
	list := flag.Bool("list", false, "list all tasks")
	delete := flag.Int("delete", -1, "delete a task")
	complete := flag.Int("complete", -1, "complete a task")

	flag.Parse()

	var todos = actions.Todos{}
	// var 

	err := todos.Load(todosPath);
	if err != nil {
		fmt.Println("Failed to load json file")
		return
	}

	switch {
	case *add:
		todo, err := getInput(flag.Args()...)

		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}

		todos.Add(todo)

		err = todos.Save(todosPath)

		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	case *list:
		todos.List()
	case *delete > 0 && *delete <= len(todos):
		err = todos.Delete(*delete)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
		err = todos.Save(todosPath)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	case *complete > 0 && *complete <= len(todos):
		err = todos.Complete(*complete)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
		err = todos.Save(todosPath)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	default:
		fmt.Fprintln(os.Stdout, "invalid command")
		os.Exit(0)
	}

}

func getInput(args ...string) (string, error) {
	if len(args) > 0 {
		return strings.Join(args, " "), nil
	}

	return "", errors.New("invalid input")
}