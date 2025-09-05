package handlers

import (
	"net/http"
	"strconv"
	"time"
	"todo-api/internal/auth"
	"todo-api/internal/models"
	"todo-api/internal/storage"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type Handler struct {
	Store storage.Storer
}

func SetupRoutes(store storage.Storer) *gin.Engine {
	router := gin.Default()
	h := &Handler{Store: store}

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
	var input models.RegisterInput
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Payload inválido"})
		return
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Erro ao gerar o hash da senha"})
		return
	}
	newUserID, err := h.Store.CreateUser(input, string(hash))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Erro ao criar o utilizador"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Utilizador criado com sucesso!", "userId": newUserID})
}

func (h *Handler) LoginUser(c *gin.Context) {
	var input models.LoginInput
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Payload inválido"})
		return
	}
	user, err := h.Store.GetUserByEmail(input.Email)
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
	tokenString, err := token.SignedString(auth.JwtSecretKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Erro ao gerar o token"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func (h *Handler) GetTodos(c *gin.Context) {
	userID, _ := c.Get("userId")
	todosList, err := h.Store.GetTodosByUserID(userID.(int))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Erro ao buscar os 'todos'"})
		return
	}
	c.JSON(http.StatusOK, todosList)
}

func (h *Handler) CreateTodo(c *gin.Context) {
	userID, _ := c.Get("userId")
	var input models.CreateTodoInput
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Payload inválido"})
		return
	}
	newTodo, err := h.Store.CreateTodo(input, userID.(int))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Erro ao inserir o 'todo'"})
		return
	}
	c.JSON(http.StatusCreated, newTodo)
}

func (h *Handler) GetTodoByID(c *gin.Context) {
	userID, _ := c.Get("userId")
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "ID inválido"})
		return
	}
	todo, err := h.Store.GetTodoByID(id, userID.(int))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Todo não encontrado"})
		return
	}
	c.JSON(http.StatusOK, todo)
}

func (h *Handler) UpdateTodo(c *gin.Context) {
	userID, _ := c.Get("userId")
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "ID inválido"})
		return
	}
	var input models.UpdateTodoInput
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "JSON inválido"})
		return
	}
	updatedTodo, err := h.Store.UpdateTodo(id, userID.(int), input)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Todo não encontrado ou erro ao atualizar"})
		return
	}
	c.JSON(http.StatusOK, updatedTodo)
}

func (h *Handler) DeleteTodo(c *gin.Context) {
	userID, _ := c.Get("userId")
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "ID inválido"})
		return
	}
	rowsAffected, err := h.Store.DeleteTodo(id, userID.(int))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Erro ao deletar o 'todo'"})
		return
	}
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "Todo não encontrado"})
		return
	}
	c.Status(http.StatusNoContent)
}
