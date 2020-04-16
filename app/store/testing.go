package store

import (
	"testing"
)

func TestStore(t *testing.T, databaseURL string) (*Store, func(...string)) {
	t.Helper()

	config := NewConfig()
	config.DatabaseURL = databaseURL
	s := New(config)
	if err := s.Open(); err != nil {
		t.Fatal(err)
	}

	return s, func(tables ...string) {
		if len(tables) > 0 {
			// Refactor for loop on tables to truncate
			// if _, err := s.db.Exec(fmt.Sprintf("TRUNCATE TABLE %s", strings.Join(tables, ", "))); err != nil {
			// 	t.Fatal(err)
			// }
		}

		s.Close()
	}
}
