package handlers

import (
	"database/sql"
	"net/http"
	"strconv"
	"time"
	"todo-api/internal/auth"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type Handler struct {
	DB *sql.DB
}

var jwtSecretKey = []byte("UMA_CHAVE_SECRETA_MUITO_LONGA_E_SEGURA")

func SetupRoutes(db *sql.DB) *gin.Engine {
	router := gin.Default()
	h := &Handler{DB: db}

	router.POST("/register", h.RegisterUser)
	router.POST("/login", h.LoginUser)

	api := router.Group("/api")
	api.Use(auth.AuthMiddleware())
	{
		api.GET("/todos", h.GetTodos)
		api.POST("/todos", h.CreateTodo)
		api.GET("/todos/:id", h.GetTodoByID)
		api.PATCH("/todos/:id", h.UpdateTodo)
		api.DELETE("/todos/:id", h.DeleteTodo)
	}

	return router
}

func (h *Handler) RegisterUser(c *gin.Context) {
	var input RegisterInput
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Payload inválido"})
		return
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Erro ao gerar o hash da senha"})
		return
	}
	sqlStatement := `INSERT INTO users (email, password_hash) VALUES ($1, $2) RETURNING id`
	var newUserID int
	err = h.DB.QueryRow(sqlStatement, input.Email, string(hash)).Scan(&newUserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Erro ao criar o utilizador"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Utilizador criado com sucesso!", "userId": newUserID})
}

func (h *Handler) LoginUser(c *gin.Context) {
	var input LoginInput
	var user struct {
		ID           int
		Email        string
		PasswordHash string
	}
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Payload inválido"})
		return
	}
	sqlStatement := `SELECT id, email, password_hash FROM users WHERE email = $1`
	err := h.DB.QueryRow(sqlStatement, input.Email).Scan(&user.ID, &user.Email, &user.PasswordHash)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Credenciais inválidas"})
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(input.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Credenciais inválidas"})
		return
	}
	claims := jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 8).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecretKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Erro ao gerar o token"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func (h *Handler) GetTodos(c *gin.Context) {
	userID, _ := c.Get("userId")
	rows, err := h.DB.Query("SELECT id, title, completed FROM todos WHERE user_id = $1 ORDER BY id ASC", userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Erro ao buscar os 'todos'"})
		return
	}
	defer rows.Close()
	var todosList []Todo
	for rows.Next() {
		var todo Todo
		if err := rows.Scan(&todo.ID, &todo.Title, &todo.Completed); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Erro ao escanear o 'todo'"})
			return
		}
		todosList = append(todosList, todo)
	}
	c.JSON(http.StatusOK, todosList)
}

func (h *Handler) CreateTodo(c *gin.Context) {
	userID, _ := c.Get("userId")
	var input CreateTodoInput
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Payload inválido"})
		return
	}
	sqlStatement := `INSERT INTO todos (title, user_id) VALUES ($1, $2) RETURNING id;`
	var newID int
	err := h.DB.QueryRow(sqlStatement, input.Title, userID).Scan(&newID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Erro ao inserir o 'todo'"})
		return
	}
	newTodo := Todo{
		ID:        strconv.Itoa(newID),
		Title:     input.Title,
		Completed: false,
	}
	c.JSON(http.StatusCreated, newTodo)
}

func (h *Handler) GetTodoByID(c *gin.Context) {
	userID, _ := c.Get("userId")
	id := c.Param("id")
	sqlStatement := `SELECT id, title, completed FROM todos WHERE id = $1 AND user_id = $2;`
	var todo Todo
	err := h.DB.QueryRow(sqlStatement, id, userID).Scan(&todo.ID, &todo.Title, &todo.Completed)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"message": "Todo não encontrado"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Erro ao buscar o 'todo'"})
		return
	}
	c.JSON(http.StatusOK, todo)
}

func (h *Handler) UpdateTodo(c *gin.Context) {
	userID, _ := c.Get("userId")
	id := c.Param("id")
	var input UpdateTodoInput
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "JSON inválido"})
		return
	}
	sqlSelect := `SELECT id, title, completed FROM todos WHERE id = $1 AND user_id = $2;`
	var currentTodo Todo
	err := h.DB.QueryRow(sqlSelect, id, userID).Scan(&currentTodo.ID, &currentTodo.Title, &currentTodo.Completed)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Todo não encontrado"})
		return
	}
	if input.Title != nil {
		currentTodo.Title = *input.Title
	}
	if input.Completed != nil {
		currentTodo.Completed = *input.Completed
	}
	sqlUpdate := `UPDATE todos SET title = $1, completed = $2 WHERE id = $3 AND user_id = $4;`
	_, err = h.DB.Exec(sqlUpdate, currentTodo.Title, currentTodo.Completed, id, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Erro ao atualizar o 'todo'"})
		return
	}
	c.JSON(http.StatusOK, currentTodo)
}

func (h *Handler) DeleteTodo(c *gin.Context) {
	userID, _ := c.Get("userId")
	id := c.Param("id")
	sqlStatement := `DELETE FROM todos WHERE id = $1 AND user_id = $2;`
	res, err := h.DB.Exec(sqlStatement, id, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Erro ao deletar o 'todo'"})
		return
	}
	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "Todo não encontrado"})
		return
	}
	c.Status(http.StatusNoContent)
}
