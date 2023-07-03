package validators

import "errors"

const DefaultAlphabet = " !\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~"

type Validator struct {
	Rules     *PasswordRules
	password  string
	checkFunc func(guess, password string) bool
}

func (v Validator) Check(guess string) bool {
	return v.checkFunc(guess, v.password)
}

func NewSimpleValidtor(rules *PasswordRules, password *string) (*Validator, error) {

	v := new(Validator)

	if rules == nil {
		v.Rules = &PasswordRules{
			Alphabet: DefaultAlphabet,
			MinLen:   8,
			MaxLen:   16,
		}
	} else {
		v.Rules = rules
	}

	if password == nil {
		v.password = v.Rules.RandomPassword()
	} else {
		v.password = *password
	}

	v.checkFunc = func(guess, password string) bool {
		return guess == password
	}

	// Validate that the password and the length rules don't conflict.
	if len(v.password) < v.Rules.MinLen {
		return nil, errors.New("password is too short")
	}
	if len(v.password) > v.Rules.MaxLen {
		return nil, errors.New("password is too long")
	}

	// Validate that the password byes are in the alphabet.
	valid := map[byte]bool{}
	for _, b := range []byte(v.Rules.Alphabet) {
		valid[b] = true
	}
	for _, b := range []byte(v.password) {
		if !valid[b] {
			return nil, errors.New("password contains invalid bytes")
		}
	}

	return v, nil
}
