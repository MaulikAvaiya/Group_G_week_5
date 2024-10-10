package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"` // "pending" or "completed"
}

// In-memory storage
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

// Update the task by himani
func updateTaskhandler(w http.ResponseWriter, r *http.Request) {
	var updatedTask Task
	if err := json.NewDecoder(r.Body).Decode(&updatedTask); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	//	updatedTask.ID = id
	if updateExistTask(&updatedTask) {
		json.NewEncoder(w).Encode(updatedTask)
	} else {
		http.Error(w, "Task not found", http.StatusNotFound)
	}
}
func updateExistTask(updatedTask *Task) bool {
	for i, task := range tasks {

		if task.ID == updatedTask.ID {
			// tasks[i] = *updatedTask
			tasks[i].Title = updatedTask.Title
			tasks[i].Description = updatedTask.Description
			tasks[i].Status = updatedTask.Status
			return true
		}
	}
	return false
}

// Handler to delete a task by ID
func deleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	idStr := strings.TrimPrefix(r.URL.Path, "/delete/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	if deleteExistingUser(id) {
		w.WriteHeader(http.StatusNoContent)

	} else {
		http.Error(w, "User not found", http.StatusNotFound)
	}
}
func deleteExistingUser(id int) bool {
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return true
		}
	}
	return false
}
func main() {
	http.HandleFunc("/gettasks", getTaskHandler)
	http.HandleFunc("/createtask", createTaskHandler)
	http.HandleFunc("/delete/", deleteTaskHandler) // New route for DELETE operation
	pNumber := ":8092"
	// http.HandleFunc("/updatetask", updateExistTask)
	http.HandleFunc("/tasks/", updateTaskhandler)
	fmt.Printf("Server is running on the port: %s\n", pNumber)
	http.ListenAndServe(pNumber, nil)
}
