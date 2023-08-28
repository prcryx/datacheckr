package datacheckr

import (
	"errors"
	"math"
	"regexp"

	"golang.org/x/exp/constraints"
)

// Global variables

var (
	EmailRegexp        *regexp.Regexp = regexp.MustCompile(EmailRegexStr)
	URLSubdomainRegexp *regexp.Regexp = regexp.MustCompile(URLSubdomain)
)

// Generic types

type Number interface {
	constraints.Integer | constraints.Float
}

// utility functions

func validateString(value any) (string, bool) {
	if value == nil {
		return "", false
	}
	val, ok := value.(string)
	if !ok {
		return "", false
	}
	return val, true
}

func validateNum[Num Number](value any) (Num, error) {
	if value == nil {
		return 0, errors.New(NotAValidIntegerErrMsg)
	}

	if err := isINForNaN(value); err != nil {
		return 0, err
	}

	if val, ok := value.(Num); !ok {
		return 0, errors.New(NotAValidIntegerErrMsg)
	} else {
		return val, nil
	}
}

func isINForNaN(value any) error {
	if v, ok := value.(float64); ok {
		if math.IsInf(v, 0) || math.IsNaN(v) {
			return errors.New(NotAValidIntegerErrMsg)
		}
	}
	return nil
}
