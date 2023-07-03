package validators

import "testing"

func TestNewRandomPasswordValid(t *testing.T) {

	testCases := []struct {
		name  string
		rules PasswordRules
	}{
		{"empty", PasswordRules{"abc", 0, 0}},
		{"8 bytes", PasswordRules{"abc", 8, 8}},
		{"variable", PasswordRules{"abc", 0, 9}},
	}

	for _, tc := range testCases {

		t.Run(tc.name, func(t *testing.T) {

			valid := map[byte]bool{}
			for _, b := range []byte(tc.rules.Alphabet) {
				valid[b] = true
			}

			pwd := tc.rules.RandomPassword()

			if len(pwd) < tc.rules.MinLen {
				t.Errorf(
					"NewRandomPassword(%+v): %v has fewer than %v bytes",
					tc.rules,
					pwd,
					tc.rules.MinLen)
			}

			if len(pwd) > tc.rules.MaxLen {
				t.Errorf(
					"NewRandomPassword(%+v): %v has more than %v bytes",
					tc.rules,
					pwd,
					tc.rules.MaxLen)
			}

			for _, b := range []byte(pwd) {
				if !valid[b] {
					t.Errorf(
						"NewRandomPassword(%+v): %v has invalid byte %v",
						tc.rules,
						pwd,
						string(b))

				}
			}

		})
	}
}
