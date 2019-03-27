package envy

import (
	"flag"
	"fmt"
	"os"
	"testing"
)

func init() {
	fatal = false
	flag.Parse()
	verbose = testing.Verbose()
}

func reset(t *testing.T) {
	t.Helper()
	verbose = testing.Verbose()
	dead = false
}

func TestStringMustNotExist(t *testing.T) {
	reset(t)
	const name = "TEST_ENV"
	os.Unsetenv(name)
	value := StringMust(name)
	if value != "" {
		t.Fatalf("unexpected value: %q for key: %q\n", value, name)
	}
	if !dead {
		t.Fatalf("expected to be dead\n")
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
	if dead {
		t.Fatalf("unexpected to be dead\n")
	}
}

func TestIntMustNotExist(t *testing.T) {
	reset(t)
	const name = "TEST_ENV"
	os.Unsetenv(name)
	value := IntMust(name)
	if value != 0 {
		t.Fatalf("expected value: %d -- got %d for key: %q\n", 0, value, name)
	}
	if !dead {
		t.Fatalf("expected to be dead\n")
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
	if dead {
		t.Fatalf("unexpected to be dead\n")
	}
}

func TestIntInvalid(t *testing.T) {
	reset(t)
	const name = "TEST_ENV"
	os.Unsetenv(name)
	os.Setenv(name, "badnumbertext")
	value := IntMust(name)
	if value != 0 {
		t.Fatalf("expected value: %d -- got %d for key: %q\n", 0, value, name)
	}
	if !dead {
		t.Fatalf("expected to be dead\n")
	}
}

func TestBoolMustNotExist(t *testing.T) {
	reset(t)
	const name = "TEST_ENV"
	os.Unsetenv(name)
	value := BoolMust(name)
	if value {
		t.Fatalf("expected value: %t -- got %t for key: %q\n", false, value, name)
	}
	if !dead {
		t.Fatalf("expected to be dead\n")
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
	if dead {
		t.Fatalf("unexpected to be dead\n")
	}
}
