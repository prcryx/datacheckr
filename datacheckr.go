package datacheckr

type ValidationRule = func(any) bool

// we can apply multiple validations to the data
type Validator struct {
	ValidationRules []ValidationRule
}

func NewValidatorInstance() *Validator {
	return &Validator{
		ValidationRules: nil,
	}
}

// to get validation rules are applied to the data
func (validator *Validator) AddValidationRules(validations ...ValidationRule) {
	validator.ValidationRules = validations
}

func (validator *Validator) Validate(data any) bool {
	return false
}
