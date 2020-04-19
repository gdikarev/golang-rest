package store

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// Store ...
type Store struct {
	config             *Config
	db                 *sql.DB
	employeeRepository *EmployeeRepository
}

// New ...
func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

// Open ..
func (s *Store) Open() error {
	db, err := sql.Open("mysql", s.config.DatabaseURL)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	s.db = db

	fmt.Println("Successfully connected to MySQL database")

	return nil
}

// Close ..
func (s *Store) Close() {
	s.db.Close()
}

// Employee ...
func (s *Store) Employee() *EmployeeRepository {
	if s.employeeRepository != nil {
		return s.employeeRepository
	}

	s.employeeRepository = &EmployeeRepository{
		store: s,
	}

	return s.employeeRepository
}
