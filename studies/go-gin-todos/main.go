package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5" // Corrigido para a v5 que instalamos
	_ "github.com/jackc/pgx/v5/stdlib"
	"golang.org/x/crypto/bcrypt"
)

type Todo struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type CreateTodoInput struct {
	Title string `json:"title" binding:"required"`
}

type UpdateTodoInput struct {
	Title     *string `json:"title"`
	Completed *bool   `json:"completed"`
}

type RegisterInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

var jwtSecretKey = []byte("UMA_CHAVE_SECRETA_MUITO_LONGA_E_SEGURA")

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(401, gin.H{"message": "Header de autorização em falta"})
			return
		}
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			c.AbortWithStatusJSON(401, gin.H{"message": "Header de autorização mal formatado"})
			return
		}
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("método de assinatura inesperado: %v", token.Header["alg"])
			}
			return jwtSecretKey, nil
		})
		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(401, gin.H{"message": "Token inválido"})
			return
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			userID := int(claims["sub"].(float64))
			c.Set("userId", userID)
			c.Next()
		} else {
			c.AbortWithStatusJSON(401, gin.H{"message": "Claims do token inválidas"})
			return
		}
	}
}

func main() {
	connStr := "user=user-todos password=pass-todos dbname=db-todos host=localhost port=5432 sslmode=disable"
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Fatal("Não foi possível conectar à base de dados:", err)
	}
	defer db.Close()
	if err = db.Ping(); err != nil {
		log.Fatal("Não foi possível 'pingar' a base de dados:", err)
	}
	fmt.Println("Conexão com a base de dados PostgreSQL estabelecida com sucesso!")

	// CORREÇÃO 1: Faltava uma vírgula antes de 'user_id'
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS todos (
		id SERIAL PRIMARY KEY,
		title TEXT NOT NULL,
		completed BOOLEAN NOT NULL DEFAULT FALSE,
		user_id INTEGER NOT NULL,
		CONSTRAINT fk_user
			FOREIGN KEY(user_id) 
			REFERENCES users(id)
			ON DELETE CASCADE
	);`
	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatal("Não foi possível criar a tabela 'todos':", err)
	}
	fmt.Println("Tabela 'todos' verificada/criada com sucesso!")

	// CORREÇÃO 2: A tabela 'users' estava com os nomes das colunas errados e SQL quebrado
	createUserTableSQL := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		email TEXT NOT NULL UNIQUE,
		password_hash TEXT NOT NULL
	);`
	_, err = db.Exec(createUserTableSQL)
	if err != nil {
		log.Fatal("Não foi possível criar a tabela 'users':", err)
	}
	fmt.Println("Tabela 'users' verificada/criada com sucesso!")

	router := gin.Default()

	// --- ROTAS PÚBLICAS (não precisam de token) ---
	router.POST("/register", func(c *gin.Context) {
		var input RegisterInput
		if err := c.BindJSON(&input); err != nil {
			c.JSON(400, gin.H{"message": "Payload inválido"})
			return
		}
		hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(500, gin.H{"message": "Erro ao gerar o hash da senha"})
			return
		}
		sqlStatement := `INSERT INTO users (email, password_hash) VALUES ($1, $2) RETURNING id`
		var newUserID int
		err = db.QueryRow(sqlStatement, input.Email, string(hash)).Scan(&newUserID)
		if err != nil {
			c.JSON(500, gin.H{"message": "Erro ao criar o utilizador"})
			return
		}
		c.JSON(201, gin.H{"message": "Utilizador criado com sucesso!", "userId": newUserID})
	})

	router.POST("/login", func(c *gin.Context) {
		var input LoginInput
		var user struct {
			ID           int
			Email        string
			PasswordHash string
		}
		if err := c.BindJSON(&input); err != nil {
			c.JSON(400, gin.H{"message": "Payload inválido"})
			return
		}
		sqlStatement := `SELECT id, email, password_hash FROM users WHERE email = $1`
		err := db.QueryRow(sqlStatement, input.Email).Scan(&user.ID, &user.Email, &user.PasswordHash)
		if err != nil {
			c.JSON(401, gin.H{"message": "Credenciais inválidas"})
			return
		}
		err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(input.Password))
		if err != nil {
			c.JSON(401, gin.H{"message": "Credenciais inválidas"})
			return
		}
		claims := jwt.MapClaims{
			"sub": user.ID,
			"exp": time.Now().Add(time.Hour * 8).Unix(),
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString(jwtSecretKey)
		if err != nil {
			c.JSON(500, gin.H{"message": "Erro ao gerar o token"})
			return
		}
		c.JSON(200, gin.H{"token": tokenString})
	})

	api := router.Group("/api")
	api.Use(AuthMiddleware())
	{
		api.GET("/todos", func(c *gin.Context) {
			rows, err := db.Query("SELECT id, title, completed FROM todos ORDER BY id ASC")
			if err != nil {
				c.JSON(500, gin.H{"message": "Erro ao buscar os 'todos'"})
				return
			}
			defer rows.Close()

			var todosList []Todo

			for rows.Next() {
				var todo Todo
				if err := rows.Scan(&todo.ID, &todo.Title, &todo.Completed); err != nil {
					c.JSON(500, gin.H{"message": "Erro ao escanear o 'todo'"})
					return
				}
				todosList = append(todosList, todo)
			}

			c.JSON(200, todosList)
		})

		api.GET("/todos/:id", func(c *gin.Context) {
			id := c.Param("id")

			sqlStatement := `SELECT id, title, completed FROM todos WHERE id = $1;`
			var todo Todo

			err := db.QueryRow(sqlStatement, id).Scan(&todo.ID, &todo.Title, &todo.Completed)
			if err != nil {
				if err == sql.ErrNoRows {
					c.JSON(404, gin.H{"message": "Todo não encontrado"})
					return
				}
				c.JSON(500, gin.H{"message": "Erro ao buscar o 'todo'"})
				return
			}

			c.JSON(200, todo)

		})

		api.POST("/todos", func(c *gin.Context) {
			var payload CreateTodoInput

			if err := c.BindJSON(&payload); err != nil {
				c.JSON(400, gin.H{"message": "Payload inválido. O título é obrigatório."})
				return
			}

			sqlStatement := `INSERT INTO todos (title) VALUES ($1) RETURNING id`
			var newID int

			err := db.QueryRow(sqlStatement, payload.Title).Scan(&newID)

			if err != nil {
				c.JSON(500, gin.H{"message": "Erro ao inserir o 'todo' na base de dados"})
				return
			}
			newTodo := Todo{
				ID:        strconv.Itoa(newID),
				Title:     payload.Title,
				Completed: false,
			}

			c.JSON(201, newTodo)
		})

		api.PATCH("/todos/:id", func(c *gin.Context) {
			id := c.Param("id")
			var input UpdateTodoInput

			if err := c.BindJSON(&input); err != nil {
				c.JSON(400, gin.H{"message": "JSON inválido"})
				return
			}

			sqlSelect := `SELECT id, title, completed FROM todos WHERE id = $1;`
			var currentTodo Todo
			err := db.QueryRow(sqlSelect, id).Scan(&currentTodo.ID, &currentTodo.Title, &currentTodo.Completed)
			if err != nil {
				c.JSON(404, gin.H{"message": "Todo não encontrado"})
				return
			}

			if input.Title != nil {
				currentTodo.Title = *input.Title
			}
			if input.Completed != nil {
				currentTodo.Completed = *input.Completed
			}

			sqlUpdate := `UPDATE todos SET title = $1, completed = $2 WHERE id = $3;`
			_, err = db.Exec(sqlUpdate, currentTodo.Title, currentTodo.Completed, id)
			if err != nil {
				c.JSON(500, gin.H{"message": "Erro ao atualizar o 'todo'"})
				return
			}

			c.JSON(200, currentTodo)

		})

		api.DELETE("/todos/:id", func(c *gin.Context) {
			id := c.Param("id")

			sqlStatement := `DELETE FROM todos WHERE id = $1;`

			res, err := db.Exec(sqlStatement, id)
			if err != nil {
				c.JSON(500, gin.H{"message": "Erro ao deletar o 'todo'"})
				return
			}

			rowsAffected, err := res.RowsAffected()
			if err != nil {
				c.JSON(500, gin.H{"message": "Erro ao verificar as linhas afetadas"})
				return
			}

			if rowsAffected == 0 {
				c.JSON(404, gin.H{"message": "Todo não encontrado"})
				return
			}

			c.Status(204)
		})
	}

	router.Run(":8080")
}
