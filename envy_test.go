package envy

import (
	"fmt"
	"os"
	"testing"
)

const testEnvName = "TEST_ENV"

func reset(t *testing.T) {
	t.Helper()
}

func TestStringMustNotExist(t *testing.T) {
	t.Skip("need to figure out capturing os.Exit")
	reset(t)
	os.Unsetenv(testEnvName)
	fmt.Println("started pre panic in func")
	value := StringMust(testEnvName)
	fmt.Println("now post panic in func")
	if value != "" {
		t.Fatalf("unexpected value: %q for key: %q\n", value, testEnvName)
	}
}

func TestStringMustExist(t *testing.T) {
	reset(t)
	const expect = "test_value"
	os.Setenv(testEnvName, expect)
	value := StringMust(testEnvName)
	if value != expect {
		t.Fatalf("expected value: %q - got %q for key: %q\n", expect, value, testEnvName)
	}
}

func TestIntMustNotExist(t *testing.T) {
	t.Skip("need to figure out capturing os.Exit")
	reset(t)
	os.Unsetenv(testEnvName)
	value := IntMust(testEnvName)
	if value != 0 {
		t.Fatalf("expected value: %d -- got %d for key: %q\n", 0, value, testEnvName)
	}
}

func TestIntMustExist(t *testing.T) {
	reset(t)
	const expect = 23
	os.Setenv(testEnvName, fmt.Sprint(expect))
	value := IntMust(testEnvName)
	if value != expect {
		t.Fatalf("expected value: %q - got %q for key: %q\n", expect, value, testEnvName)
	}
}

func TestIntInvalid(t *testing.T) {
	t.Skip("need to figure out capturing os.Exit")
	reset(t)
	os.Unsetenv(testEnvName)
	os.Setenv(testEnvName, "badnumbertext")
	value := IntMust(testEnvName)
	if value != 0 {
		t.Fatalf("expected value: %d -- got %d for key: %q\n", 0, value, testEnvName)
	}
}

func TestBoolMustNotExist(t *testing.T) {
	t.Skip("need to figure out capturing os.Exit")
	reset(t)
	os.Unsetenv(testEnvName)
	value := BoolMust(testEnvName)
	if value {
		t.Fatalf("expected value: %t -- got %t for key: %q\n", false, value, testEnvName)
	}
}

func TestBoolMustExist(t *testing.T) {
	reset(t)
	const expect = true
	os.Setenv(testEnvName, fmt.Sprint(expect))
	value := BoolMust(testEnvName)
	if value != expect {
		t.Fatalf("expected value: %t -- got %t for key: %q\n", expect, value, testEnvName)
	}
}
