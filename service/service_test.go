package service

import (
	"testing"

	h "github.com/shdlabs/week21/helpers"
)

func TestTheBeginning(t *testing.T) {
	t.Parallel()
	t.Run("NewUser", func(t *testing.T) {
		t.Parallel()

		user := NewUser("John", "NY", "123456789", 1.75, true)

		h.Equal(t, user.ID, int32(0))
		h.Equal(t, user.Fname, "John")
		h.Equal(t, user.City, "NY")
	})

	t.Run("NewDBMock", func(t *testing.T) {
		db := NewDBMock()
		db.NewUsers(
			NewUser("John", "NY", "123456789", 1.85, true),
			NewUser("Anne", "LA", "123456789", 1.75, false),
			NewUser("Bill", "CF", "123456789", 1.90, true),
		)

		h.Equal(t, len(db), 3)
		h.NotNil(t, db)

		t.Run("AddUser", func(t *testing.T) {
			h.NoError(t, db.AddUser(42, "John", "NY", "123456789", 1.75, true))

			if len(db) != 4 {
				t.Errorf(h.Ko("Expected %d"), 4)
			}

			err := db.AddUser(42, "John", "NY", "123456789", 1.75, true)
			if err == nil || err.Error() != "user ID exists" {
				t.Error(h.Ko("Expected error"))
			}
		})

		t.Run("FindUser", func(t *testing.T) {
			user := db.FindUser(42)
			h.NotEqual(t, user, User{})

			user = db.FindUser(52)
			h.Equal(t, user, (User{}))
		})

		t.Logf("%+v themale test \n\t", db)
	})
}
