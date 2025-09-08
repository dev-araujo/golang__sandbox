package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Tarefa struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	UserID int    `json:"-"`
}

type TarefaCompleta struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	UserName string `json:"userName"`
}

type User struct {
	Name string `json:"name"`
}

func main() {
	tarefas := map[int]Tarefa{
		101: {ID: 101, Title: "Aprender sobre Microserviços", UserID: 1},
		102: {ID: 102, Title: "Construir um exemplo prático", UserID: 2},
	}

	router := gin.Default()

	router.GET("/tarefas/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "ID de tarefa inválido"})
			return
		}

		tarefa, exists := tarefas[id]
		if !exists {
			c.JSON(http.StatusNotFound, gin.H{"message": "Tarefa não encontrada"})
			return
		}

		userURL := fmt.Sprintf("http://localhost:8081/users/%d", tarefa.UserID)

		resp, err := http.Get(userURL)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Erro ao comunicar com o serviço de utilizadores"})
			return
		}
		defer resp.Body.Close()

		var user User
		if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Erro ao processar a resposta do serviço de utilizadores"})
			return
		}

		tarefaCompleta := TarefaCompleta{
			ID:       tarefa.ID,
			Title:    tarefa.Title,
			UserName: user.Name,
		}

		c.JSON(http.StatusOK, tarefaCompleta)
	})

	router.Run(":8082")
}
