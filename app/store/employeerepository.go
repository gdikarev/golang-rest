package store

import "github.com/gdikarev/golang-rest/app/model"

// UserRepository ...
type EmployeeRepository struct {
	store *Store
}

// Create ...
func (r *EmployeeRepository) Create(e *model.Employee) (*model.Employee, error) {
	if err := r.store.db.QueryRow(
		"INSERT INTO employees (first_name, second_name, last_name) VALUES ($1, $2, $3); SELECT LAST_INSERT_ID()",
		e.FistName,
		e.SecondName,
		e.LastName
	).Scan(&e.ID); err != nil {
		return nil, err
	}

	return e, nil
}
