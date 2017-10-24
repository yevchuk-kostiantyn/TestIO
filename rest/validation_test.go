package rest

import "testing"

func TestValidateInfo(t *testing.T) {
	var expectedOutput bool
	var info string

	info = "Kostiantyn"
	expectedOutput = ValidateInfo(info)
	if expectedOutput != true {
		t.Error("Expected true, got ", expectedOutput)
	}
}

func TestValidateEmail(t *testing.T) {
	var expectedOutput bool
	var email string

	email = "kostiantyn.yevchuk@nure.ua"
	expectedOutput = ValidateEmail(email)
	if expectedOutput != true {
		t.Error("Expected true, got ", expectedOutput)
	}
}

func TestValidatePassword(t *testing.T) {
	var expectedOutput bool
	var password string

	password = "nure1961!"
	expectedOutput = ValidatePassword(password)

	if expectedOutput != true {
		t.Error("Expected true, got ", expectedOutput)
	}
}