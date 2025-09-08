package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {

	users := map[int]User{
		1: {ID: 1, Name: "Maria Silva"},
		2: {ID: 2, Name: "João Costa"},
	}

	router := gin.Default()

	router.GET("users/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})

			return
		}

		user, exists := users[id]
		if !exists {
			c.JSON(http.StatusNotFound, gin.H{"error": "Utilizador não encontrado"})
			return
		}

		c.JSON(http.StatusOK, user)

	})

	router.Run(":8081")

}
