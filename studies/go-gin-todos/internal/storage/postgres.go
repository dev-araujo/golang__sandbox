package storage

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	// Não se esqueça do _ para o driver
	_ "github.com/jackc/pgx/v5/stdlib"
)

// A nossa função para iniciar a conexão e criar as tabelas
func InitDB() *sql.DB {
	connStr := os.Getenv("DB_SOURCE")
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Fatal("Não foi possível conectar à base de dados:", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal("Não foi possível 'pingar' a base de dados:", err)
	}
	fmt.Println("Conexão com a base de dados PostgreSQL estabelecida com sucesso!")

	// Executa as migrações (criação das tabelas)
	migrateDB(db)

	return db
}

// Função privada para criar as tabelas
func migrateDB(db *sql.DB) {
	// (Lembre-se de colocar a criação da tabela 'users' PRIMEIRO)
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
