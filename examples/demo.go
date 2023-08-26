package main

import (
	"fmt"

	"github.com/prcryx/datacheckr"
)

func main() {
	validator := datacheckr.NewValidatorInstance()

	email := "prcryx87@email.in"

	validator.AddValidationRules(datacheckr.EmailValidation, datacheckr.MinLenValidation(3))

	isValid := validator.Validate(email)

	fmt.Printf("Is ValidEmail: %v", isValid)
}
