package helpers

import (
	"log"
	"strings"
	"testing"
	"time"
)

func NotNil(t *testing.T, data any) {
	t.Helper()
	t.Run("NotNil", func(t *testing.T) {
		if data == nil {
			t.Error(Ko("was not expecting `nil`"))
		}
	})
}

func NoError(t *testing.T, err error) {
	t.Helper()
	t.Run("NoError", func(t *testing.T) {
		if err != nil {
			t.Errorf(Ko("was not expecting error...\n\tGOT: `%#v`"), err)
		}
	})
}

func LogDuration(start time.Time) {
	log.Printf(Ah("Done: %v"), time.Since(start))
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
