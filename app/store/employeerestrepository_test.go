package store_test

import (
	"testing"

	"github.com/gdikarev/golang-rest/app/model"
	"github.com/gdikarev/golang-rest/app/store"
	"github.com/stretchr/testify/assert"
)

func TestEmployeeRepository_Create(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("employees")

	e, err := s.Employee().Create(&model.Employee{
		FistName:   "Warren",
		SecondName: "John",
		LastName:   "Buffet",
	})

	assert.NoError(t, err)
	assert.NotNil(t, e)
}

func TestEmployeeRepository_FingByUuid(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("employees")

	uuid := "793e1192-442e-422b-a9f1-5ea6a26f16b0"
	_, err := s.Employee().FindByUuid(uuid)

	assert.Error(t, err)

	// res, err := s.Employee().Create(&model.Employee{
	// 	FistName:   "Warren",
	// 	SecondName: "John",
	// 	LastName:   "Buffet",
	// })

	// e, err := s.Employee().FindByUuid(res.UUID)
	// assert.NoError(t, err)
	// assert.NotNil(t, e)
}
