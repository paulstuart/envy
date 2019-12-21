// Package envy provides helpers for getting environment variables
package envy

import (
	"log"
	"os"
	"strconv"
)

// String reads an environment variable from the OS
func String(key string) string {
	return os.Getenv(key)
}

// StringDefault returns an env value or the default if not specified
func StringDefault(key string, value string) string {
	if s := String(key); s != "" {
		return s
	}
	return value
}

// StringMust returns an env value or exits if not specified
func StringMust(key string) string {
	s := String(key)
	if s == "" {
		log.Fatalf("required key %q was not set or is empty", key)
	}
	return s
}

// Int returns the variable as an int
func Int(key string) int {
	value := os.Getenv(key)
	if value == "" {
		return 0
	}
	i, err := strconv.Atoi(value)
	if err != nil {
		log.Fatalf("error converting key:%s value:%q to integer:%v\n", key, value, err)
	}
	return i
}

// IntDefault returns the variable as an int
func IntDefault(key string, value int) int {
	s := os.Getenv(key)
	if s == "" {
		return value
	}
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("error converting key:%s value:%q to integer:%v\n", key, s, err)
	}
	return i
}

// IntMust returns the variable as an int
func IntMust(key string) int {
	s := os.Getenv(key)
	if s == "" {
		log.Fatalf("required variable %q was not set or is empty", key)
	}
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("error converting key:%s value:%q to integer:%v\n", key, s, err)
	}
	return i
}

// Bool returns the variable as a boolean
func Bool(key string) bool {
	if s := os.Getenv(key); s != "" {
		b, err := strconv.ParseBool(s)
		if err != nil {
			log.Fatalf("error converting key:%s value:%q to boolean:%v\n", key, s, err)
		}
		return b
	}
	return false
}

// BoolMust returns the variable as a boolean or exits if not set or invalid
func BoolMust(key string) bool {
	s := os.Getenv(key)
	if s == "" {
		log.Fatalf("required variable %q was not set or is empty\n", key)
	}
	b, err := strconv.ParseBool(s)
	if err != nil {
		log.Fatalf("error converting key:%s value:%q to boolean:%v\n", key, s, err)
	}
	return b
}
