package misc

import (
	"log"
	"testing"
)

func TestInt64Env(t *testing.T) {
	env := Int64Env("something", 1)
	log.Print(env)
}