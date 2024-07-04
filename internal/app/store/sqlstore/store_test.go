package sqlstore_test

import (
	"os"
	"testing"
)

// Конфигурация тестировочной БД
var (
	databaseURL string
)

func TestMain(m *testing.M) {
	databaseURL = os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = "host=localhost user=user password=1 dbname=dotatest sslmode=disable"
	}

	os.Exit(m.Run())
}
