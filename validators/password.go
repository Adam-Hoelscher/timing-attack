package validators

import (
	"math/rand"
)

type PasswordRules struct {
	Alphabet string
	MinLen   int
	MaxLen   int
}

func (rules *PasswordRules) RandomPassword() string {

	actualLen := rules.MinLen + rand.Intn(rules.MaxLen-rules.MinLen+1)

	password := make([]byte, actualLen)
	for i := range password {
		password[i] = rules.Alphabet[rand.Intn(len(rules.Alphabet))]
	}

	return string(password)
}
