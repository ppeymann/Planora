package env

import (
	"os"

	"github.com/subosito/gotenv"
)

func GetEnv(key, def string) string {
	err := gotenv.Load(".env")
	if err != nil {

		return def
	}

	val := os.Getenv(key)
	if len(val) == 0 {
		return def
	}

	return val
}
