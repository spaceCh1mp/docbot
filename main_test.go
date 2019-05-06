package main

import (
	"os"
	"testing"
)

func TestPath(t *testing.T) {
	var dir, home, got, expected string
	println("TestPath...")
	//Test case 1 {{ if path is absolute}}
	dir = validate("log/015")
	home, _ = os.UserHomeDir()
	got = Path(dir, false)
	expected = home + dir
	if got != expected {
		t.Errorf("\nTest case 1:\ngot: %s\nexpected : %s", got, expected)
	}

	//Test case 2 {{if Path is relative[default] }}
	dir = validate("log/015")
	got = Path(dir, true)
	expected = dir
	if got != expected {
		t.Errorf("\nTest case 2 :\ngot: %s\nexpected : %s", got, expected)
	}

}

func TestValidate(t *testing.T) {
	var dir, got string
	println("TestValidate...")
	expected := "/Documents/log/"
	//Test case 1 {{ no leading and trailing backslashes }}
	dir = "Documents/log"
	got = validate(dir)
	if got != expected {
		t.Errorf("\nTest case 1:\ngot: %s\nexpected : %s", got, expected)
	}
	//Test case 2 {{ no leading backslash }}
	dir = "Documents/log/"
	got = validate(dir)
	if got != expected {
		t.Errorf("\nTest case 2:\ngot: %s\nexpected : %s", got, expected)
	}
	//Test case 3 {{ no trailing backslash }}
	dir = "/Documents/log"
	got = validate(dir)
	if got != expected {
		t.Errorf("\nTest case 3:\ngot: %s\nexpected : %s", got, expected)
	}
}
