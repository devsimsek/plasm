package utils_test

import (
	"testing"

	"github.com/devsimsek/plasm/utils"
)

func TestContains(t *testing.T) {
	slice := []string{"a", "b", "c", "d", "e"}

	if !utils.Contains(slice, "a") {
		t.Errorf("Expected true, got false")
	}

	if utils.Contains(slice, "f") {
		t.Errorf("Expected false, got true")
	}
}

func TestStrToInt(t *testing.T) {
	i := utils.StrToInt("123")

	if i != 123 {
		t.Errorf("Expected 123, got %d", i)
	}
}

func TestIntToStr(t *testing.T) {
	s := utils.IntToStr(123)

	if s != "123" {
		t.Errorf("Expected '123', got '%s'", s)
	}
}

func TestStrToFloat(t *testing.T) {
	f := utils.StrToFloat("123.45")

	if f != 123.45 {
		t.Errorf("Expected 123.45, got %f", f)
	}
}

func TestFloatToStr(t *testing.T) {
	s := utils.FloatToStr(123.45)

	if s != "123.45" {
		t.Errorf("Expected '123.45', got '%s'", s)
	}
}

func TestStrToBool(t *testing.T) {
	b := utils.StrToBool("true")

	if !b {
		t.Errorf("Expected true, got false")
	}
}

func TestBoolToStr(t *testing.T) {
	s := utils.BoolToStr(true)

	if s != "true" {
		t.Errorf("Expected 'true', got '%s'", s)
	}
}

func TestStrToUint(t *testing.T) {
	i := utils.StrToUint("123")

	if i != 123 {
		t.Errorf("Expected 123, got %d", i)
	}
}

func TestUintToStr(t *testing.T) {
	s := utils.UintToStr(123)

	if s != "123" {
		t.Errorf("Expected '123', got '%s'", s)
	}
}

func TestStrToInt8(t *testing.T) {
	i := utils.StrToInt8("123")

	if i != 123 {
		t.Errorf("Expected 123, got %d", i)
	}
}

func TestInt8ToStr(t *testing.T) {
	s := utils.Int8ToStr(123)

	if s != "123" {
		t.Errorf("Expected '123', got '%s'", s)
	}
}
