package api

import (
	"encoding/json"
	"net/http"
)

func (s *Server) getListTasksHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	tasks := s.task.GetListTasks()

	jsonResponse, err := json.Marshal(tasks)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func (s *Server) getTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	var requestPayload struct {
		ID uint `json:"id"`
	}

	err := json.NewDecoder(r.Body).Decode(&requestPayload)
	if err != nil {
		http.Error(w, "Corpo da requisição inválido ou mal formatado", http.StatusBadRequest)
		return
	}

	task := s.task.GetTask(requestPayload.ID)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)

}

func (s *Server) addTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	var requestPayload struct {
		Description string `json:"description"`
	}

	err := json.NewDecoder(r.Body).Decode(&requestPayload)
	if err != nil {
		http.Error(w, "Corpo da requisição inválido ou mal formatado", http.StatusBadRequest)
		return
	}

	if requestPayload.Description == "" {
		http.Error(w, "O campo 'description' não pode ser vazio", http.StatusBadRequest)
		return
	}

	newTask := s.task.AddTask(requestPayload.Description)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTask)
}

func (s *Server) deleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	var requestPayload struct {
		ID uint `json:"id"`
	}

	err := json.NewDecoder(r.Body).Decode(&requestPayload)
	if err != nil {
		http.Error(w, "Corpo da requisição inválido ou mal formatado", http.StatusBadRequest)
		return
	}

	s.task.DeleteTask(requestPayload.ID)
	w.WriteHeader(http.StatusNoContent)

}

func (s *Server) updateTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	var requestPayload struct {
		ID          uint   `json:"id"`
		Description string `json:"description"`
		Completed   bool   `json:"completed"`
	}

	err := json.NewDecoder(r.Body).Decode(&requestPayload)
	if err != nil {
		http.Error(w, "Corpo da requisição inválido ou mal formatado", http.StatusBadRequest)
		return
	}

	if requestPayload.Description == "" {
		http.Error(w, "O campo 'description' não pode ser vazio", http.StatusBadRequest)
		return
	}

	taskUpdated := s.task.UpdateTask(requestPayload.ID, requestPayload.Description, requestPayload.Completed)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(taskUpdated)
}
