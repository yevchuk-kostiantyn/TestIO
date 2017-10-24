package database

import "testing"

func TestUserExists(t *testing.T) {
	Flushall()
	SaveNewUser("Kostiantyn", "Yevchuk",
		"kostiantyn.yevchuk@nure.ua", "nure1961!", "student")

	var entered_username string
	entered_username = "kostiantyn.yevchuk@nure.ua"

	var expectedOutput bool
	expectedOutput = UserExists(entered_username)

	if expectedOutput != true {
		t.Error("Expected true, got ", expectedOutput)
	}
}

func TestIsPasswordCorrect(t *testing.T) {
	Flushall()
	SaveNewUser("Kostiantyn", "Yevchuk",
		"kostiantyn.yevchuk@nure.ua", "nure1961!", "student")

	var entered_username string
	entered_username = "kostiantyn.yevchuk@nure.ua"
	var entered_password string
	entered_password = "nure1961!"

	var expectedOutput bool

	expectedOutput = IsPasswordCorrect(entered_username, entered_password)

	if expectedOutput != true {
		t.Error("Expected true, got ", expectedOutput)
	}
}

func TestGetUserPosition(t *testing.T) {
	Flushall()
	SaveNewUser("Kostiantyn", "Yevchuk",
		"kostiantyn.yevchuk@nure.ua", "nure1961!", "student")

	var entered_username string
	entered_username = "kostiantyn.yevchuk@nure.ua"
	var expectedOutput string
	expectedOutput = GetUserPosition(entered_username)

	if expectedOutput != "student" {
		t.Error("Expected student, got ", expectedOutput)
	}
}