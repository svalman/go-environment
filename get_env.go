package environment

import (
	"os"
	"strconv"
	"strings"
)

func GetEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return strings.TrimSpace(value)
	}

	if len(defaultVal) == 0 {
		panic("Не установлена переменная окружения " + key + " и не предложено значение по умолчанию")
	}

	return defaultVal
}

// read variable with "name", else with defaultVal
func GetEnvAsInt(name string, defaultVal int) int {

	if valueStr, exists := os.LookupEnv(name); exists {
		if value, err := strconv.Atoi(valueStr); err == nil {
			return value
		}
	}

	return defaultVal
}

// Helper to read an environment variable into a bool or return default value
func GetEnvAsBool(name string, defaultVal bool) bool {
	valStr := GetEnv(name, "0")
	if len(valStr) > 0 {
		if val, err := strconv.ParseBool(valStr); err == nil {
			return val
		}
	}
	return defaultVal
}

// Helper to read an environment variable into a string slice or return default value
func GetEnvAsSlice(name string, defaultVal []string, sep string) []string {
	valStr := GetEnv(name, "")

	if valStr == "" {
		return defaultVal
	}

	val := strings.Split(valStr, sep)

	return val
}
