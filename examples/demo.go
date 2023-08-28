package main

import (
	"fmt"

	"github.com/prcryx/datacheckr"
)

func main() {
	validator := datacheckr.NewValidatorInstance()

	email1 := "prcryx87@email.in" //valid email
	email2 := "prcryx87email.in"  // not valid email

	validator.AddValidationRules(datacheckr.EmailValidation, datacheckr.MinStrLenValidation(5))

	IsEmail1 := validator.Validate(email1)
	fmt.Printf("%v is valid : %v\n", email1, IsEmail1)

	IsEmail2 := validator.Validate(email2)
	fmt.Printf("%v is valid : %v\n", email2, IsEmail2)
}
