package store

import (
	"github.com/gdikarev/golang-rest/app/model"
	"github.com/google/uuid"
)

// EmployeeRepository ...
type EmployeeRepository struct {
	store *Store
}

// Create ...
func (r *EmployeeRepository) Create(e *model.Employee) (*model.Employee, error) {
	if _, err := r.store.db.Exec(
		"INSERT INTO employees (first_name, second_name, last_name, uuid) VALUES (?, ?, ?, ?);",
		e.FistName,
		e.SecondName,
		e.LastName,
		uuid.New(),
	); err != nil {
		println("Execution error", err)
	}

	return e, nil
}

// FindByUuid ...
func (r *EmployeeRepository) FindByUuid(uuid string) (*model.Employee, error) {
	e := &model.Employee{}
	if err := r.store.db.QueryRow("SELECT first_name, last_name, uuid FROM employees WHERE uuid = ?",
		uuid,
	).Scan(
		&e.FistName,
		&e.LastName,
		&e.UUID,
	); err != nil {
		return nil, err
	}

	return e, nil
}
