package datacheckr

import (
	"regexp"
)

// Global variables

var (
	EmailRegexp        *regexp.Regexp = regexp.MustCompile(EmailRegexStr)
	URLSubdomainRegexp *regexp.Regexp = regexp.MustCompile(URLSubdomain)
)

// Validation Rules

func EmailValidation(value any) bool {
	if email, ok := ValidateString(value); ok {
		if !EmailRegexp.MatchString(email) {
			return false
		} else {
			return true
		}
	}
	return false
}

// min-length validation

func MinLenValidation(minLen int) ValidationRule {
	return func(value any) bool {
		if str, ok :=ValidateString(value); ok{
			return minLen < len(str)
		}
		return false
	}
}
