package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"todo-api/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type MockStore struct {
}

func (m *MockStore) GetTodosByUserID(userID int) ([]models.Todo, error) {
	fakeTodos := []models.Todo{
		{ID: "1", Title: "Tarefa de Teste 1", Completed: false},
		{ID: "2", Title: "Tarefa de Teste 2", Completed: true},
	}
	return fakeTodos, nil
}

func (m *MockStore) GetUserByEmail(email string) (*models.User, error)               { return nil, nil }
func (m *MockStore) CreateUser(input models.RegisterInput, hash string) (int, error) { return 0, nil }

func (m *MockStore) CreateTodo(input models.CreateTodoInput, userID int) (*models.Todo, error) {
	return nil, nil
}
func (m *MockStore) GetTodoByID(todoID int, userID int) (*models.Todo, error) { return nil, nil }
func (m *MockStore) UpdateTodo(todoID int, userID int, input models.UpdateTodoInput) (*models.Todo, error) {
	return nil, nil
}
func (m *MockStore) DeleteTodo(todoID int, userID int) (int64, error) { return 0, nil }

func TestGetTodos(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockStore := &MockStore{}

	router := SetupRoutes(mockStore)

	req, _ := http.NewRequest(http.MethodGet, "/api/todos", nil)

	w := httptest.NewRecorder()

	req.Header.Set("Content-Type", "application/json")

	c, _ := gin.CreateTestContext(w)
	c.Request = req
	c.Set("userId", 1)

	handlerFunc := router.Handler().ServeHTTP
	handlerFunc(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "O código de status deve ser 200 OK")

	expectedBody := `[{"id":"1","title":"Tarefa de Teste 1","completed":false},{"id":"2","title":"Tarefa de Teste 2","completed":true}]`

	assert.JSONEq(t, expectedBody, w.Body.String(), "O corpo da resposta deve ser o JSON esperado")

	t.Log("Corpo da resposta:", w.Body.String())
}
