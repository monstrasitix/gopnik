package env

import (
	"os"

	"github.com/joho/godotenv"
)

func Configure() {
	godotenv.Load()
}

func env(name string) func () string {
	return func() string {
		return os.Getenv(name)
	}
}

