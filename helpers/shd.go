package helpers

import (
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/charmbracelet/log"
)

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

// Ko helper function for terminal output decoration in red with ❌.
// Not OK
func Ko(line string) string {
	return strings.Join([]string{red, KO, line, end}, " ")
}

// Ok helper function for terminal output decoration in green with ✅.
func Ok(line string) string {
	return strings.Join([]string{green, OK, line, end}, " ")
}

// Ah helper function for terminal output decoration in yellow with 😯
// Debugging and exploartion purpose.
func Ah(line string) string {
	return strings.Join([]string{blue, AH, line, end}, " ")
}

func NotNil(t *testing.T, data any) {
	t.Helper()
	if data == nil {
		t.Errorf(Ko("was not expecting %#v"), nil)
	}
}

func NoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Errorf(Ko("was not expecting error...\n\tGOT: %#v"), err)
	}
}

func Equal[T comparable](t *testing.T, actual, expected T) {
	t.Helper()
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf(Ko("Not Equal => \nEXP: %#v\nGOT: %#v"), expected, actual)
	}
	reflect.DeepEqual(actual, expected)
}

func NotEqual[T comparable](t *testing.T, actual, expected T) {
	t.Helper()
	if reflect.DeepEqual(actual, expected) {
		t.Errorf(Ko("Equal => \nEXP: %#v\nGOT: %#v"), expected, actual)
	}
}

// DurationLog measure the duration of a function
//
// Usage:
//
//	defer helpers.DurationLog(time.Now())
func DurationLog(start time.Time, name string) {
	log.Info("DURATION", "func", name, "duration", time.Since(start))
}
