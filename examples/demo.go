package main

import (
	"fmt"

	"github.com/prcryx/datacheckr"
)

func main() {
	validator := datacheckr.NewValidatorInstance()

	email := "prcryx87@email.in"

	validator.AddValidationRules(email, datacheckr.EmailValidation)

	isValid := validator.Validate()

	fmt.Printf("Is ValidEmail: %v", isValid)
}
