package main

import "flag"

const (
	todoFile = ".todo.json" //. makes it hidden
)

func main() {

	add := flag.Bool("add", false, "add a new todo")
	flag.Parse()

	todos := &todo.Todos{}

}
