package database

import "testing"

func TestUserExists(t *testing.T) {
	Flushall()
	SaveNewUser("Kostiantyn", "Yevchuk",
		"kostiantyn.yevchuk@nure.ua", "nure1961!", "admin")
	SaveNewUser("Kostia", "Yevchuk",
		"kostyantyn.yewchuk@gmail.com", "nure1962!", "instructor")
	SaveNewUser("Katy", "Perry",
		"katy.perry@gmail.com", "nure1963!", "student")
	SaveNewUser("Lady", "Gaga",
		"lady.gaga@gmail.com", "nure1964!", "instructor")

	var tests = []struct {
		expectedOutput  bool
		enteredUsername string
	}{
		{true, "kostiantyn.yevchuk@nure.ua"},
		{false, "steve.jobs@nure.ua"},
		{true, "kostyantyn.yewchuk@gmail.com"},
		{true, "katy.perry@gmail.com"},
		{true, "lady.gaga@gmail.com"},
		{false, "bill.gates@gmail.com"},
		{false, "lg@gmail.com"},
	}

	for _, test := range tests {
		if expectedOutput := UserExists(test.enteredUsername); expectedOutput != test.expectedOutput {
			t.Fatalf("UserExists(%q) = %q, want %q.", test.enteredUsername,
				expectedOutput, test.expectedOutput)
		}
	}

}

func TestIsPasswordCorrect(t *testing.T) {
	Flushall()
	SaveNewUser("Kostiantyn", "Yevchuk",
		"kostiantyn.yevchuk@nure.ua", "nure1961!", "admin")
	SaveNewUser("Kostia", "Yevchuk",
		"kostyantyn.yewchuk@gmail.com", "nure1962!", "instructor")
	SaveNewUser("Katy", "Perry",
		"katy.perry@gmail.com", "nure1963!", "student")
	SaveNewUser("Lady", "Gaga",
		"lady.gaga@gmail.com", "nure1964!", "instructor")

	var tests = []struct {
		expectedOutput  bool
		enteredUsername string
		enteredPassword string
	}{
		{true, "kostiantyn.yevchuk@nure.ua", "nure1961!"},
		{true, "kostyantyn.yewchuk@gmail.com", "nure1962!"},
		{true, "katy.perry@gmail.com", "nure1963!"},
		{true, "lady.gaga@gmail.com", "nure1964!"},

		{false, "kostiantyn.yevchuk@nure.ua", "nure1960!"},
		{false, "kostyantyn.yewchuk@gmail.com", "nure1961!"},
		{false, "katy.perry@gmail.com", "nure1962!"},
		{false, "lady.gaga@gmail.com", "nure1963!"},
		{false, "lg@gmail.com", "helloWorld"},
	}

	for _, test := range tests {
		if expectedOutput := IsPasswordCorrect(test.enteredUsername, test.enteredPassword); expectedOutput != test.expectedOutput {
			t.Fatalf("IsPasswordCorrect(%q, %q) = %q, want %q.", test.enteredUsername,
				test.enteredPassword, expectedOutput, test.expectedOutput)
		}
	}
}

func TestGetUserPosition(t *testing.T) {
	Flushall()
	SaveNewUser("Kostiantyn", "Yevchuk",
		"kostiantyn.yevchuk@nure.ua", "nure1961!", "admin")
	SaveNewUser("Kostia", "Yevchuk",
		"kostyantyn.yewchuk@gmail.com", "nure1962!", "instructor")
	SaveNewUser("Katy", "Perry",
		"katy.perry@gmail.com", "nure1963!", "student")
	SaveNewUser("Lady", "Gaga",
		"lady.gaga@gmail.com", "nure1964!", "instructor")

	var tests = []struct {
		expectedPosition string
		enteredUsername  string
	}{
		{"admin", "kostiantyn.yevchuk@nure.ua"},
		{"instructor", "kostyantyn.yewchuk@gmail.com"},
		{"student", "katy.perry@gmail.com"},
		{"instructor", "lady.gaga@gmail.com"},
		{"", "steve.jobs@nure.ua"},
		{"", "bill.gates@nure.ua"},
		{"", "michael.jordan@nure.ua"},
	}

	for _, test := range tests {
		if expectedPosition := GetUserPosition(test.enteredUsername); expectedPosition != test.expectedPosition {
			t.Fatalf("GetUserPosition(%q, %q) = %q, want %q.", test.enteredUsername,
				expectedPosition, test.expectedPosition)
		}
	}
}
