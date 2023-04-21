package service_test

import (
	"clean-architecture-api/tests/setup"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	code := m.Run()
	setup.TeardownDB()
	os.Exit(code)
}
