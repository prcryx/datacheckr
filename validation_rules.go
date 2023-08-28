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

func MaxStrLenValidation(maxLen int) ValidationRule {
	return func(value any) bool {
		if str, ok := validateString(value); ok {
			return maxLen > len(str)
		}
		return false
	}
}

// max val validation

func MaxValValidation[Num Number](maxVal Num) ValidationRule {
	return func(value any) bool {
		if num, err := validateNum[Num](value); err == nil {
			return maxVal > num
		}
		return false
	}
}

// min val validation

func MinValValidation[Num Number](minVal Num) ValidationRule {
	return func(value any) bool {
		if num, err := validateNum[Num](value); err == nil {
			return minVal < num
		}
		return false
	}
}
