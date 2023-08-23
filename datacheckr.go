package datacheckr

type ValidationRule = func(any) bool

type Validator struct {
	ValidationRules map[any][]ValidationRule
}

// we can apply multiple validations to the data
func NewValidatorInstance() *Validator {
	return &Validator{
		ValidationRules: make(map[any][]ValidationRule),
	}
}

// to get validation rules are applied to the data
func (validator *Validator) AddValidationRules(data any, validations ...ValidationRule) {
	validator.ValidationRules[data] = validations
}


func (validator *Validator) Validate() bool{
	// add logic !! 
	return false
}