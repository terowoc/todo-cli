package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CommandFlags struct {
	Add    string
	Del    int
	Edit   string
	Toggle int
	List   bool
}

func NewCommandFlags() *CommandFlags {
	cmdf := CommandFlags{}

	flag.StringVar(&cmdf.Add, "add", "", "Add a new todo specify title")
	flag.IntVar(&cmdf.Del, "del", -1, "Delete a todo by index")
	flag.StringVar(&cmdf.Edit, "edit", "", "Edit a todo by index id:new title")
	flag.IntVar(&cmdf.Toggle, "toggle", -1, "Toggle a todo by index")
	flag.BoolVar(&cmdf.List, "list", false, "List all todos")
	flag.Parse()

	return &cmdf
}

func (cmdf *CommandFlags) Execute(todos *Todos) {
	switch {

	case cmdf.Add != "":
		todos.Add(cmdf.Add)

	case cmdf.Del != -1:
		todos.Delete(cmdf.Del)

	case cmdf.Edit != "":
		parts := strings.SplitN(cmdf.Edit, ":", 2)
		if len(parts) != 2 {
			fmt.Println("invalid format for edit")
			os.Exit(1)
		}

		index, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("invalid index")
			os.Exit(1)
		}

		todos.Edit(index, parts[1])

	case cmdf.Toggle != -1:
		todos.Toggle(cmdf.Toggle)

	case cmdf.List:
		todos.Print()

	default:
		fmt.Println("invalid command")
		os.Exit(1)
	}
}
