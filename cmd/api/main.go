package main

import (
	"encoding/json"
	"net/http"
)

type todo struct {
	ID			string	`json:"id"`
	Item		string	`json:"item"`
	Completed	bool	`json:"completed"`
}

var todos = []todo{
	{ID: "1", Item: "Walk Dog", Completed: false},
	{ID: "2", Item: "Clean Dishes", Completed: false},
	{ID: "3", Item: "Take Out Trash", Completed: false},
}

func getAllTodos(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(todos)
}

func main() {
    http.HandleFunc("/todos", getAllTodos) 
    http.ListenAndServe("localhost:8080", nil)     
}