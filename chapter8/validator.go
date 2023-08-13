package chapter8

import (
	"errors"
	"net/mail"
	"strings"
)

type Validator struct {
	validator func() error
}

func (v *Validator) Validate() error {
	return v.validator()
}

func (v *Validator) Set(callable func() error) {
	v.validator = callable
}

func ValidateList(validators ...*Validator) error {
	for _, validator := range validators {
		err := validator.Validate()
		if err != nil {
			return err
		}
	}
	return nil
}

func NewCardTokenValidator(cardToken string) *Validator {
	validator := &Validator{}

	validator.Set(func() error {
		if len(cardToken) != 12 {
			return errors.New("invalid card token")
		}
		return nil
	})

	return validator
}

func NewEmailValidator(email string) *Validator {
	validator := &Validator{}

	validator.Set(func() error {
		_, err := mail.ParseAddress(email)
		if err != nil {
			return errors.New("invalid email format")
		}
		return nil
	})

	return validator
}

func NewEmailSuffixValidator(email string) *Validator {
	validator := &Validator{}

	validator.Set(func() error {
		mailSuffix := strings.Split(email, "@")[1]
		switch mailSuffix {
		case "gmail.com",
			"yahoo.com",
			"msn.com":
			return nil
		default:
			return errors.New("invalid email suffix")
		}
	})

	return validator
}
