package config

import (
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		// Poll runtime.Caller to find config.go's path.
		_, filename, _, ok := runtime.Caller(0)
		if !ok {
			log.Fatal("Error loading .env file ", "No caller information")
		}

		// Go one directory higher.
		rootdir := filepath.Dir(filepath.Dir(filename))

		// Look for ../.env.
		err = godotenv.Load(filepath.Join(rootdir, ".env"))

		if err != nil {
			log.Fatal("Error loading .env file ", err)
		}
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
