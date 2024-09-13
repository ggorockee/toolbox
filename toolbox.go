package toolbox

import (
	"log"
	"os"
)

func Getenv(key string, defaultValue ...string) string {
	// get env value
	envValue := os.Getenv(key)

	var dValue string
	// env not set
	if envValue == "" {
		switch len(defaultValue) {
		case 0:
			dValue = ""
			return dValue
		case 1:
			dValue = defaultValue[0]
			return dValue
		default:
			log.Println("defaultValue arguments 1")
			os.Exit(1)
		}
	}

	return envValue
}
