package collection

// MapStrings takes a slice of strings and applies a function, fn
// to each value. A new, mapped slice is returned.
func MapStrings(strings []string, fn func(string) string) []string {
	mapped := make([]string, len(strings))
	for i, v := range strings {
		mapped[i] = fn(v)
	}
	return mapped
}
