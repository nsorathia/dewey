package common

import (
	"os"
)

//GetEnvironmentVar fetches an environment varibale or returns the fallback value
func GetEnvironmentVar(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
