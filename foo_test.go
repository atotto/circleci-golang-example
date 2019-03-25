package foo

import (
	"testing"
	"time"

	_ "golang.org/x/crypto/bcrypt"
)

func TestTooHeavyTest(t *testing.T) {
	time.Sleep(10*time.Second)
}
