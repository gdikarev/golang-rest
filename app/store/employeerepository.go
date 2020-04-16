package store

import "github.com/gdikarev/golang-rest/app/model"

// UserRepository ...
type EmployeeRepository struct {
	store *Store
}

// Create ...
func (r *EmployeeRepository) Create(e *model.Employee) (*model.Employee, error) {
	// if err := r.store.db.QueryRow(
	// 	"INSERT INTO employees (first_name, second_name, last_name) VALUES (?, ?, ?)",
	// 	e.FistName,
	// 	e.SecondName,
	// 	e.LastName,
	// ).Scan(&e.ID); err != nil {
	// 	return nil, err
	// }

	// return e, nil

	res, err := r.store.db.Exec(
		"INSERT INTO employees (first_name, second_name, last_name) VALUES (?, ?, ?);",
		e.FistName,
		e.SecondName,
		e.LastName,
	)

	if err != nil {
		println("Execution error", err)
	} else {
		id, err := res.LastInsertId()
		if err != nil {
			println("Error", err)
		}
		e.ID = id
		// println("Last id", id)
	}

	return e, nil
}
