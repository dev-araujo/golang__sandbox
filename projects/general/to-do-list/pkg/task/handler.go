package task

import (
	"encoding/json"
	"net/http"
)

func Routes(s Service) *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("/todo/list", getListTasksHandler(s))
	router.HandleFunc("/todo/add", addTaskHandler(s))
	router.HandleFunc("/todo/delete", deleteTaskHandler(s))
	router.HandleFunc("/todo/update", updateTaskHandler(s))
	router.HandleFunc("/todo/get", getTaskHandler(s))
	return router
}

func getListTasksHandler(s Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
			return
		}

		tasks := s.GetListTasks()

		jsonResponse, err := json.Marshal(tasks)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonResponse)
	}
}

func getTaskHandler(s Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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

		task, err := s.GetTask(requestPayload.ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(task)
	}
}

func addTaskHandler(s Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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

		newTask := s.AddTask(requestPayload.Description)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newTask)
	}
}

func deleteTaskHandler(s Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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

		if err := s.DeleteTask(requestPayload.ID); err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

func updateTaskHandler(s Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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

		taskUpdated, err := s.UpdateTask(requestPayload.ID, requestPayload.Description, requestPayload.Completed)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(taskUpdated)
	}
}
