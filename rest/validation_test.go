package rest

import "testing"

func TestValidateInfo(t *testing.T) {
	var tests = []struct {
		isValid bool
		info    string
	}{
		{true, "Hello World"},
		{true, "Michael"},
		{false, "_"},
		{false, "0"},
		{false, ""},
		{true, "Kyrie Irving"},
	}

	for _, test := range tests {
		if isValid := ValidateInfo(test.info); isValid != test.isValid {
			t.Fatalf("ValidateInfo(%q) = %q, want %q.", test.info,
				isValid, test.isValid)
		}
	}
}

func TestValidateEmail(t *testing.T) {
	var tests = []struct {
		isValid bool
		email   string
	}{
		{true, "email@domain.com"},
		{true, "firstname.lastname@domain.com"},
		{true, "email@subdomain.domain.com"},
		{true, "firstname+lastname@domain.com"},
		{true, "1234567890@domain.com"},
		{true, "email@domain-one.com"},
		{true, "_______@domain.com"},
		{true, "email@domain.name"},
		{true, "email@domain.co.jp"},
		{true, "firstname-lastname@domain.com"},

		{false, "plainaddress"},
		{false, "#@%^%#$@#$@#.com"},
		{false, "@domain.com"},
		{false, "Joe Smith <email@domain.com>"},
		{false, "email.domain.com"},
		{false, "email@domain@domain.com"},
		{false, "あいうえお@domain.comあいうえお@domain.com"},
		{false, "email@domain.com (Joe Smith)"},
		{false, "email@domain"},
	}

	for _, test := range tests {
		if isValid := ValidateEmail(test.email); isValid != test.isValid {
			t.Fatalf("ValidateEmail(%q) = %q, want %q.", test.email,
				isValid, test.isValid)
		}
	}
}

func TestValidatePassword(t *testing.T) {
	var tests = []struct {
		isValid  bool
		password string
	}{
		{true, "HelloWorld"},
		{true, "Michael"},
		{false, "_"},
		{false, "0"},
		{false, ""},
		{true, "KyrieIrving"},
	}

	for _, test := range tests {
		if isValid := ValidatePassword(test.password); isValid != test.isValid {
			t.Fatalf("ValidatePassword(%q) = %q, want %q.", test.password,
				isValid, test.isValid)
		}
	}
}
