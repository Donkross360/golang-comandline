package main

import (
	"fmt"
	"os"
	"strings"

	"pragprog.com/rggo/interacting/todo/todo"
)

// Hardcoding the file name
const todoFilename = ".todo.json"

func main() {
	// define an items list
	l := &todo.List{}

	// use the get method to read to do items from file
	if err := l.Get(todoFilename); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Decide what to do based on the number of argument provided
	switch {
	// For no extra argument, print the list
	case len(os.Args) == 1:
		// List current todo items
		for _, item := range *l {
			fmt.Println(item.Task)
		}

	// Concatenate all provided argument with a space
	// and add to the list as an item
	default:
		// Concatenate all argument with a space
		item := strings.Join(os.Args[1:], " ")

		// Add the task
		l.Add(item)

		//Save the new list
		if err := l.Save(todoFilename); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
}
