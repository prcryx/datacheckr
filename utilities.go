package datacheckr

func ValidateString(value any) (string, bool) {
	if value == nil || value == "" {
		return "", false
	}
	val, ok := value.(string)
	if !ok {
		return "", false
	}
	return val, true
}
