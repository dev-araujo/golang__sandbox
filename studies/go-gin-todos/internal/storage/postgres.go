package storage

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"todo-api/internal/models"

	_ "github.com/jackc/pgx/v5/stdlib"
)

// Storer define o contrato para todas as operações de base de dados.
type Storer interface {
	GetUserByEmail(email string) (*models.User, error)
	CreateUser(input models.RegisterInput, hash string) (int, error)
	GetTodosByUserID(userID int) ([]models.Todo, error)
	CreateTodo(input models.CreateTodoInput, userID int) (*models.Todo, error)
	GetTodoByID(todoID int, userID int) (*models.Todo, error)
	UpdateTodo(todoID int, userID int, input models.UpdateTodoInput) (*models.Todo, error)
	DeleteTodo(todoID int, userID int) (int64, error)
}

type PostgresStore struct {
	DB *sql.DB
}

func NewPostgresStore() (Storer, error) {
	connStr := os.Getenv("DB_SOURCE")
	if connStr == "" {
		log.Fatal("A variável de ambiente DB_SOURCE não está definida.")
	}

	db, err := sql.Open("pgx", connStr)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	fmt.Println("Conexão com a base de dados PostgreSQL estabelecida com sucesso!")
	migrateDB(db)
	return &PostgresStore{DB: db}, nil
}

// --- MÉTODOS QUE CUMPREM O CONTRATO ---

func (s *PostgresStore) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	sqlStatement := `SELECT id, email, password_hash FROM users WHERE email = $1`
	err := s.DB.QueryRow(sqlStatement, email).Scan(&user.ID, &user.Email, &user.PasswordHash)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *PostgresStore) CreateUser(input models.RegisterInput, hash string) (int, error) {
	sqlStatement := `INSERT INTO users (email, password_hash) VALUES ($1, $2) RETURNING id`
	var newUserID int
	err := s.DB.QueryRow(sqlStatement, input.Email, hash).Scan(&newUserID)
	if err != nil {
		return 0, err
	}
	return newUserID, nil
}

func (s *PostgresStore) GetTodosByUserID(userID int) ([]models.Todo, error) {
	rows, err := s.DB.Query("SELECT id, title, completed FROM todos WHERE user_id = $1 ORDER BY id ASC", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todosList []models.Todo
	for rows.Next() {
		var todo models.Todo
		var id int
		if err := rows.Scan(&id, &todo.Title, &todo.Completed); err != nil {
			return nil, err
		}
		todo.ID = strconv.Itoa(id)
		todosList = append(todosList, todo)
	}
	return todosList, nil
}

func (s *PostgresStore) CreateTodo(input models.CreateTodoInput, userID int) (*models.Todo, error) {
	sqlStatement := `INSERT INTO todos (title, user_id) VALUES ($1, $2) RETURNING id;`
	var newID int
	err := s.DB.QueryRow(sqlStatement, input.Title, userID).Scan(&newID)
	if err != nil {
		return nil, err
	}
	newTodo := &models.Todo{
		ID:        strconv.Itoa(newID),
		Title:     input.Title,
		Completed: false,
	}
	return newTodo, nil
}

func (s *PostgresStore) GetTodoByID(todoID int, userID int) (*models.Todo, error) {
	sqlStatement := `SELECT id, title, completed FROM todos WHERE id = $1 AND user_id = $2;`
	var todo models.Todo
	var id int
	err := s.DB.QueryRow(sqlStatement, todoID, userID).Scan(&id, &todo.Title, &todo.Completed)
	if err != nil {
		return nil, err
	}
	todo.ID = strconv.Itoa(id)
	return &todo, nil
}

func (s *PostgresStore) UpdateTodo(todoID int, userID int, input models.UpdateTodoInput) (*models.Todo, error) {
	currentTodo, err := s.GetTodoByID(todoID, userID)
	if err != nil {
		return nil, err
	}
	if input.Title != nil {
		currentTodo.Title = *input.Title
	}
	if input.Completed != nil {
		currentTodo.Completed = *input.Completed
	}
	sqlUpdate := `UPDATE todos SET title = $1, completed = $2 WHERE id = $3 AND user_id = $4;`
	_, err = s.DB.Exec(sqlUpdate, currentTodo.Title, currentTodo.Completed, todoID, userID)
	if err != nil {
		return nil, err
	}
	return currentTodo, nil
}

func (s *PostgresStore) DeleteTodo(todoID int, userID int) (int64, error) {
	sqlStatement := `DELETE FROM todos WHERE id = $1 AND user_id = $2;`
	res, err := s.DB.Exec(sqlStatement, todoID, userID)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func migrateDB(db *sql.DB) {
	createUserTableSQL := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		email TEXT NOT NULL UNIQUE,
		password_hash TEXT NOT NULL
	);`
	_, err := db.Exec(createUserTableSQL)
	if err != nil {
		log.Fatal("Não foi possível criar a tabela 'users':", err)
	}
	fmt.Println("Tabela 'users' verificada/criada com sucesso!")

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
}
