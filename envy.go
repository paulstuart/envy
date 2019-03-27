// Package envy provides helpers for getting environment variables
package envy

import (
	"log"
	"os"
	"strconv"
)

var (
	// Debug will print returned values if set true
	Debug = false

	dead    = false
	fatal   = true
	verbose = true
)

func debug(s string, args ...interface{}) {
	if Debug && verbose {
		log.Printf(s, args...)
	}
}

func flog(s string, args ...interface{}) {
	if fatal {
		log.Fatalf(s, args...)
	}
	if verbose {
		log.Printf(s, args...)
	}
	dead = true
	verbose = false
}

// String reads an environment variable from the OS
func String(key string) string {
	return os.Getenv(key)
}

// StringDefault returns an env value or the default if not specified
func StringDefault(key string, value string) string {
	if s := String(key); s != "" {
		debug("%s=%s\n", key, s)
		return s
	}
	debug("%s=%s #DEFAULT\n", key, value)
	return value
}

// StringMust returns an env value or exits if not specified
func StringMust(key string) string {
	s := String(key)
	if s == "" {
		flog("required variable %q was not set or is empty", key)
	}
	debug("%s=%s\n", key, s)
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
		flog("error converting key:%s value:%q to integer:%v\n", key, value, err)
	}
	debug("%s=%d\n", key, i)
	return i
}

// IntDefault returns the variable as an int
func IntDefault(key string, value int) int {
	s := os.Getenv(key)
	if s == "" {
		debug("%s=%d #DEFAULT\n", key, value)
		return value
	}
	i, err := strconv.Atoi(s)
	if err != nil {
		flog("error converting key:%s value:%q to integer:%v\n", key, s, err)
	}
	debug("%s=%d\n", key, i)
	return i
}

// IntMust returns the variable as an int
func IntMust(key string) int {
	s := os.Getenv(key)
	if s == "" {
		flog("required variable %q was not set or is empty", key)
	}
	i, err := strconv.Atoi(s)
	if err != nil {
		flog("error converting key:%s value:%q to integer:%v\n", key, s, err)
	}
	return i
}

// Bool returns the variable as a boolean
func Bool(key string) bool {
	if s := os.Getenv(key); s != "" {
		b, err := strconv.ParseBool(s)
		if err != nil {
			flog("error converting key:%s value:%q to boolean:%v\n", key, s, err)
		}
		return b
	}
	return false
}

// BoolMust returns the variable as a boolean or exits if not set or invalid
func BoolMust(key string) bool {
	s := os.Getenv(key)
	if s == "" {
		flog("required variable %q was not set or is empty", key)
	}
	b, err := strconv.ParseBool(s)
	if err != nil {
		flog("error converting key:%s value:%q to boolean:%v\n", key, s, err)
	}
	return b
}
