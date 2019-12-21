package envy

import (
	"fmt"
	"os"
	"testing"
)

func reset(t *testing.T) {
	t.Helper()
}

func TestStringMustNotExist(t *testing.T) {
	t.Skip("need to figure out capturing os.Exit")
	reset(t)
	const name = "TEST_ENV"
	os.Unsetenv(name)
	fmt.Println("started pre panic in func")
	value := StringMust(name)
	fmt.Println("now post panic in func")
	if value != "" {
		t.Fatalf("unexpected value: %q for key: %q\n", value, name)
	}
}

func TestStringMustExist(t *testing.T) {
	reset(t)
	const name = "TEST_ENV"
	const expect = "test_value"
	os.Setenv(name, expect)
	value := StringMust(name)
	if value != expect {
		t.Fatalf("expected value: %q - got %q for key: %q\n", expect, value, name)
	}
}

func TestIntMustNotExist(t *testing.T) {
	t.Skip("need to figure out capturing os.Exit")
	reset(t)
	const name = "TEST_ENV"
	os.Unsetenv(name)
	value := IntMust(name)
	if value != 0 {
		t.Fatalf("expected value: %d -- got %d for key: %q\n", 0, value, name)
	}
}

func TestIntMustExist(t *testing.T) {
	reset(t)
	const name = "TEST_ENV"
	const expect = 23
	os.Setenv(name, fmt.Sprint(expect))
	value := IntMust(name)
	if value != expect {
		t.Fatalf("expected value: %q - got %q for key: %q\n", expect, value, name)
	}
}

func TestIntInvalid(t *testing.T) {
	t.Skip("need to figure out capturing os.Exit")
	reset(t)
	const name = "TEST_ENV"
	os.Unsetenv(name)
	os.Setenv(name, "badnumbertext")
	value := IntMust(name)
	if value != 0 {
		t.Fatalf("expected value: %d -- got %d for key: %q\n", 0, value, name)
	}
}

func TestBoolMustNotExist(t *testing.T) {
	t.Skip("need to figure out capturing os.Exit")
	reset(t)
	const name = "TEST_ENV"
	os.Unsetenv(name)
	value := BoolMust(name)
	if value {
		t.Fatalf("expected value: %t -- got %t for key: %q\n", false, value, name)
	}
}

func TestBoolMustExist(t *testing.T) {
	reset(t)
	const name = "TEST_ENV"
	const expect = true
	os.Setenv(name, fmt.Sprint(expect))
	value := BoolMust(name)
	if value != expect {
		t.Fatalf("expected value: %t -- got %t for key: %q\n", expect, value, name)
	}
}
