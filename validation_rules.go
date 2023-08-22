package datacheckr

import ("regexp")

var EmailRegexp *regexp.Regexp = regexp.MustCompile(`^[a-zA-Z0-9._%+]+@[a-zA-Z0-9]+\.[a-zA-Z]{2,}$`)


