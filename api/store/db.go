package store

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

type MySQLRepository struct {
	db *sql.DB
}

func NewSQLRepository(cfg mysql.Config) *MySQLRepository {
	db, err := sql.Open("mysql", cfg.FormatDSN())

	if err != nil {
		log.Fatal((err))
	}

	err = db.Ping()
	if err != nil {
		log.Fatal((err))
	}

	log.Println("Connected to MySQL!")

	return &MySQLRepository{db: db}
}

func (s *MySQLRepository) Init() (*sql.DB, error) {
	// init tables

	if err := s.createProjectsTable(); err != nil {
		return nil, err
	}

	if err := s.createUsersTable(); err != nil {
		return nil, err
	}

	if err := s.createTasksTable(); err != nil {
		return nil, err
	}

	return s.db, nil
}

func (s *MySQLRepository) createProjectsTable() error {
	log.Println("Creating user projects")
	_, err := s.db.Exec(`
		CREATE TABLE IF NOT EXISTS projects (
			id 		INT UNSIGNED NOT NULL AUTO_INCREMENT,
			name 	VARCHAR(255) NOT NULL,
			createdAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

			PRIMARY KEY (id)
		);
	`)

	return err
}

func (s *MySQLRepository) createTasksTable() error {
	log.Println("Creating user tasks")
	_, err := s.db.Exec(`
		CREATE TABLE IF NOT EXISTS tasks (
			id 			INT UNSIGNED NOT NULL AUTO_INCREMENT,
			name 		VARCHAR(255) NOT NULL,
			status 		ENUM('TODO', 'IN_PROGRESS', 'IN_TESTING', 'DONE') NOT NULL DEFAULT 'TODO',
			projectId 	INT UNSIGNED NOT NULL,
			assignedTo	INT UNSIGNED NOT NULL,
			createdAt 	TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

			PRIMARY KEY (id),
			FOREIGN KEY (assignedTo) REFERENCES users(id),
			FOREIGN KEY (projectId) REFERENCES projects(id)
		);
	`)

	return err
}

func (s *MySQLRepository) createUsersTable() error {
	log.Println("Creating user table")
	_, err := s.db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id 			INT UNSIGNED NOT NULL AUTO_INCREMENT,
			email 		VARCHAR(255) NOT NULL,
			firstName	VARCHAR(255) NOT NULL,
			lastName	VARCHAR(255) NOT NULL,
			password	VARCHAR(255) NOT NULL,
			createdAt 	TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

			PRIMARY KEY (id),
			UNIQUE KEY (email)
		);
	`)

	return err
}
