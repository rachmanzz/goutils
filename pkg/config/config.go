package config

import (
	"fmt"
	"os"
	"strings"
)

func SetConfig(key, val string) error {
	key = strings.TrimSpace(key)
	if key == "" {
		return fmt.Errorf("config key cannot be empty")
	}
	return os.Setenv(key, val)
}

func GetConfig(key string, fallback ...string) string {
	val := os.Getenv(key)
	if val != "" {
		return val
	}
	if len(fallback) > 0 {
		return fallback[0]
	}
	return ""
}

func GetConfigBool(key string, fallback ...bool) bool {
	val := strings.ToLower(strings.TrimSpace(os.Getenv(key)))

	switch val {
	case "true", "1":
		return true
	case "false", "0":
		return false
	case "":
		if len(fallback) > 0 {
			return fallback[0]
		}
		return false // default fallback
	default:
		if len(fallback) > 0 {
			return fallback[0]
		}
		return false // default fallback
	}
}

func GetConfigNum(key string, fallback ...int) int {
	val := os.Getenv(key)
	if val == "" {
		if len(fallback) > 0 {
			return fallback[0]
		}
		return -1 // default fallback
	}

	var num int
	_, err := fmt.Sscanf(val, "%d", &num)
	if err != nil {
		if len(fallback) > 0 {
			return fallback[0]
		}
		return -1 // default fallback
	}

	return num
}
