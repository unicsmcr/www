package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
}

// CheckHaveenv checks if all the env variables given as args exist.
func CheckHaveenv(vars ...string) bool {
	missing := make([]string, 0)

	for _, envvar := range vars {
		if os.Getenv(envvar) == "" {
			missing = append(missing, envvar)
		}
	}

	if len(missing) == 0 {
		return true
	}

	for _, envvar := range missing {
		log.Printf("Environment variable %s is not assigned.\n", envvar)
	}

	return false
}
