package env

import (
	"os"
	"strconv"
)

func Env(key, fallback string) string {

	val := os.Getenv(key)

	if val == "" {

		return fallback
	}

	return val
}

func EnvInt(key string, fallback int) int {

	val := Env(key, "")

	if val == "" {

		return fallback
	}

	ret, err := strconv.Atoi(val)

	if err != nil {

		return fallback
	}

	return ret
}

func EnvBool(key string) bool {

	val := Env(key, "")

	ret, err := strconv.ParseBool(val)

	if err != nil {

		return false
	}

	return ret
}
