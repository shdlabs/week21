package service

import (
	"strings"
	"testing"
)

func TestTheBeggining(t *testing.T) {
	t.Parallel()

	t.Run("NewUser", func(t *testing.T) {
		t.Parallel()

		user := NewUser("John", "New York", "123456789", 1.75, true)
		if user == (User{}) {
			t.Error(Ko("Expected not empty"))
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

		t.Logf(Ok("%v"), db)

		if len(db) != 3 {
			t.Error(Ko("Expected 3"))
		}

		NotNil(t, db)
	})
}

// helpers
func NotNil(t *testing.T, data any) {
	t.Helper()

	if data == nil {
		t.Error(Ko("Expected not empty"))
	}
}

const (
	end     = "\033[0m"
	red     = "\033[31;3;1m"
	green   = "\033[32m"
	yellow  = "\033[33m;3;1m"
	blue    = "\033[34;3m"
	magenta = "\033[35m"
	cyan    = "\033[36m"
	gray    = "\033[37;1m"
	white   = "\033[97;1m"
)

const (
	OK = "✅"
	KO = "❌"
	OH = "😯"
	AH = "🤨"
)

func Ko(line string) string {
	return strings.Join([]string{red, KO, line, end}, " ")
}

func Ok(line string) string {
	return strings.Join([]string{green, OK, line, end}, " ")
}

func Ah(line string) string {
	return strings.Join([]string{blue, AH, line, end}, " ")
}
