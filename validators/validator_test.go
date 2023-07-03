package validators

import "testing"

func TestSimpleValidatorCheck(t *testing.T) {

	testCases := []struct {
		name  string
		guess string
		valid bool
	}{
		{"match", "abc", true},
		{"wrong", "FAKE", false},
	}

	rules := PasswordRules{Alphabet: "abc", MinLen: 3, MaxLen: 3}
	password := "abc"
	validator, _ := NewSimpleValidtor(&rules, &password)

	for _, tc := range testCases {

		t.Run(tc.name, func(t *testing.T) {
			if got := validator.Check(tc.guess); got != tc.valid {
				t.Errorf(
					"validator.Check(\"%vc\") should be %v; got %v",
					tc.guess,
					tc.valid,
					got)
			}
		})

	}
}

func TestNewSimpleValidatorDefault(t *testing.T) {
	_, err := NewSimpleValidtor(nil, nil)
	if err != nil {
		t.Errorf("NewSimpleValidator generated an error using default values.")
	}
}

func TestNewSimpleValidatorErrors(t *testing.T) {

	password3 := "aaa"
	testCases := []struct {
		name     string
		rules    *PasswordRules
		password *string
	}{
		{"too long",
			&PasswordRules{Alphabet: "abc", MinLen: 3, MaxLen: 2},
			&password3},
		{"too short",
			&PasswordRules{Alphabet: "abc", MinLen: 4, MaxLen: 3},
			&password3},
		{"bad byyes",
			&PasswordRules{Alphabet: "xyz", MinLen: 3, MaxLen: 3},
			&password3},
	}

	for _, tc := range testCases {

		t.Run(tc.name, func(t *testing.T) {
			_, err := NewSimpleValidtor(tc.rules, tc.password)
			if err == nil {
				t.Errorf(
					"NewSimpleValidator(%#v, %#v) did not generate an error",
					*tc.rules,
					*tc.password)
			}
		})

	}
}
