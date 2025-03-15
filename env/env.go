package env

import (
	"os"
	"strconv"
)

func GetEnvOrDefault(key, defaultValue string) string {
	if value, ok := GetEnv(key); ok {
		return value
	}

	return defaultValue
}

func GetIntEnvOrDefault(key string, defaultValue int) (int, error) {
	if value, ok := GetEnv(key); ok {
		return strconv.Atoi(value)
	}

	return defaultValue, nil
}

func GetBoolEnvOrDefault(key string, defaultValue bool) (bool, error) {
	if value, ok := GetEnv(key); ok {
		return strconv.ParseBool(value)
	}

	return defaultValue, nil
}

func GetEnv(key string) (string, bool) {
	value := os.Getenv(key)
	ok := value != ""
	return value, ok
}
