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

	// 2. Criamos uma requisição HTTP "falsa"
	req, _ := http.NewRequest(http.MethodGet, "/api/todos", nil)

	// Criamos um "gravador" de resposta
	w := httptest.NewRecorder()

	// "Anexamos" o userId ao contexto, simulando o que o nosso AuthMiddleware faria
	req.Header.Set("Content-Type", "application/json") // Embora não seja necessário para GET, é boa prática

	// Criamos um contexto Gin falso e passamos o userId
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	c.Set("userId", 1) // ID do utilizador de teste

	// 3. Executamos a função do handler
	// Encontramos a função handler para a rota que estamos testando
	handlerFunc := router.Handler().ServeHTTP
	handlerFunc(w, req) // Executa a requisição contra o router

	// 4. Verificamos os resultados
	assert.Equal(t, http.StatusOK, w.Code, "O código de status deve ser 200 OK")

	// --- NOVA VERIFICAÇÃO DO CORPO ---
	// O nosso mock sempre devolve a mesma lista de "todos" falsos
	expectedBody := `[{"id":"1","title":"Tarefa de Teste 1","completed":false},{"id":"2","title":"Tarefa de Teste 2","completed":true}]`

	// A biblioteca 'assert' tem uma função especial para comparar JSONs,
	// que ignora problemas de espaçamento e ordem de chaves.
	assert.JSONEq(t, expectedBody, w.Body.String(), "O corpo da resposta deve ser o JSON esperado")

	t.Log("Corpo da resposta:", w.Body.String())
}
