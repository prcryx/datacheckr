package datacheckr

// Validation Rules

func EmailValidation(value any) bool {
	if email, ok := validateString(value); ok {
		return EmailRegexp.MatchString(email)
	}
	return false
}

//url validation

func UrlSchemeValidation(value any) bool {
	if url, ok := validateString(value); ok {
		return URLSubdomainRegexp.MatchString(url)
	}
	return false
}

// min-length validation

func MinStrLenValidation(minLen int) ValidationRule {
	return func(value any) bool {
		if str, ok := validateString(value); ok {
			return minLen < len(str)
		}
		return false
	}
}

// max-length validation

func MaxLenValidation(maxLen int) ValidationRule {
	return func(value any) bool {
		if str, ok := ValidateString(value); ok {
			return maxLen > len(str)
		}
		return false
	}
}