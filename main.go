package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"` // "pending" or "completed"
}

// In-memory storage (simulating a database)
var tasks = []Task{}
var taskId = 1

func getTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	json.NewEncoder(w).Encode(tasks)
}
func createTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	var task Task

	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "Invalid read request body", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content/Type", "application/json")
	task.ID = taskId
	taskId++
	tasks = append(tasks, task)
	json.NewEncoder(w).Encode(task)
}

func main() {
	http.HandleFunc("/gettasks", getTaskHandler)
	http.HandleFunc("/createtask", createTaskHandler)
	pNumber := ":8092"
	fmt.Printf("Server is running on the port: %s\n", pNumber)
	http.ListenAndServe(pNumber, nil)
}
