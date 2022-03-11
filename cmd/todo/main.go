package main

import "main/todolist"

func main() {
	server()
}

func server() {
	todolist.TodoServer()
}
