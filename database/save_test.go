package database

import "testing"

func TestSaveNewUser(t *testing.T) {
	Flushall()

	var expectedOutput bool
	expectedOutput = SaveNewUser("Kostiantyn", "Yevchuk",
		"kostiantyn.yevchuk@nure.ua", "nure1961!", "student")

	if expectedOutput != true {
		t.Error("Expected true, got ", expectedOutput)
	}
}

func TestSaveAdmin(t *testing.T) {
	Flushall()
	var expectedOutput bool
	expectedOutput = SaveAdmin()
	if expectedOutput != true {
		t.Error("Expected true, got ", expectedOutput)
	}
}
