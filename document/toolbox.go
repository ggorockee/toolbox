package document

import (
	"log"
	"os"
	"syscall"
)

// Getenv reads the environment variable
// the first argument is "key" and the second variable is the "default" if there is no key value.
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

// Setenv sets the value of the environment variable named by the key.
// It returns an error, if any.
func Setenv(key, value string) error {
	err := syscall.Setenv(key, value)
	if err != nil {
		return err
	}
	return nil
}
