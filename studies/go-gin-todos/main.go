package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type CreateTodo struct {
	Title string `json:"title" binding:"required"`
}
type UpdateTodo struct {
	Title     *string `json:"title" binding:"required"`
	Completed *bool   `json:"completed" binding:"required"`
}

var todos = []Todo{
	{ID: "1", Title: "Aprender Go", Completed: false},
	{ID: "2", Title: "Criar API com Gin", Completed: false},
	{ID: "3", Title: "Tomar um café", Completed: false},
}

func main() {
	router := gin.Default()

	router.GET("/todos", func(c *gin.Context) {
		c.JSON(200, todos)
	})

	router.GET("/todos/:id", func(c *gin.Context) {
		id := c.Param("id")

		for _, todo := range todos {
			if todo.ID == id {
				c.JSON(200, todo)
				return
			}
		}

		c.JSON(404, gin.H{"message": "Todo not find"})

	})

	router.POST("/todos", func(c *gin.Context) {
		var payload CreateTodo

		if err := c.BindJSON(&payload); err != nil {
			c.JSON(400, gin.H{"message": "Payload inválido. O título é obrigatório."})
			return
		}

		newTodo := Todo{
			ID:        strconv.Itoa(len(todos) + 1),
			Title:     payload.Title,
			Completed: false,
		}

		todos = append(todos, newTodo)

		c.JSON(201, newTodo)
	})

	router.PATCH("/todos/:id", func(c *gin.Context) {
		id := c.Param("id")

		var payload UpdateTodo

		if err := c.BindJSON(&payload); err != nil {
			c.JSON(400, gin.H{"message": "Payload inválido. O título e o status são obrigatórios."})
			return
		}

		for i, todo := range todos {
			if todo.ID == id {

				if payload.Title != nil {
					todos[i].Title = *payload.Title
				}

				if payload.Completed != nil {
					todos[i].Completed = *payload.Completed
				}

				c.JSON(200, todos[i])
				return

			}
		}
		c.JSON(404, gin.H{"message": "Todo não encontrado"})

	})

	router.DELETE("/todos/:id", func(c *gin.Context) {
		id := c.Param("id")

		for i, todo := range todos {
			if todo.ID == id {
				todos = append(todos[:i], todos[i+1:]...)
				c.Status(204)
				return
			}
		}

		c.JSON(404, gin.H{"message": "Todo não encontrado"})
	})

	router.Run(":8080")
}
