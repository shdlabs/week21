package service

import (
	"testing"

	h "github.com/shdlabs/week21/helpers"
)

func TestTheBeggining(t *testing.T) {
	t.Parallel()

	t.Run("NewUser", func(t *testing.T) {
		t.Parallel()

		user := NewUser("John", "New York", "123456789", 1.75, true)
		if user == (User{}) {
			t.Error(h.Ko("Expected not empty"))
		}
	})

	t.Run("NewDBMock", func(t *testing.T) {
		t.Parallel()

		db := NewDBMock()
		db.NewUsers(
			NewUser("John", "NY", "123456789", 1.75, true),
			NewUser("Steve", "LA", "123456789", 1.75, true),
			NewUser("Bill", "LA", "123456789", 1.75, true),
		)

		t.Logf(h.Ok("%v"), db)

		if len(db) != 3 {
			t.Error(h.Ko("Expected 3"))
		}

		h.NotNil(t, db)
	})
}
