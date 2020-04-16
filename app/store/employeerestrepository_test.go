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
