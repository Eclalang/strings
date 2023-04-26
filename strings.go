package strings

import (
	"strconv"
	"strings"
)

// Contains returns true if the string contains the substring
func Contains(str string, substr string) bool {
	return strings.Contains(str, substr)
}

// ContainsAny returns true if the string contains any characters in the substring
func ContainsAny(str string, substr string) bool {
	return strings.ContainsAny(str, substr)
}

// Count returns the number of times the substring appears in the string
func Count(str string, substr string) int {
	return strings.Count(str, substr)
}

// Cut returns the strings before and after the substring, and a boolean indicating if the substring was found
func Cut(str string, substr string) [3]string {
	before, after, found := strings.Cut(str, substr)
	return [3]string{before, after, strconv.FormatBool(found)}
}

// HasPrefix returns true if the string starts with the prefix
func HasPrefix(str string, substr string) bool {
	return strings.HasPrefix(str, substr)
}

// HasSuffix returns true if the string ends with the suffix
func HasSuffix(str string, substr string) bool {
	return strings.HasSuffix(str, substr)
}

// IndexOf returns the index of the first instance of the substring in the string, or -1 if not found
func IndexOf(str string, substr string) int {
	return strings.Index(str, substr)
}

// Join concatenates the strings in the slice with the separator
func Join(elems []string, sep string) string {
	return strings.Join(elems, sep)
}

// Replace returns a copy of the string with old replaced by new and the number of replacements
func Replace(str string, old string, new string, nb int) string {
	return strings.Replace(str, old, new, nb)
}

// ReplaceAll returns a copy of the string with old replaced by new
func ReplaceAll(str string, old string, new string) string {
	return strings.ReplaceAll(str, old, new)
}

// Split returns a slice of strings split by the separator
func Split(str string, sep string) []string {
	return strings.Split(str, sep)
}

// SplitAfter returns a slice of strings split after the separator, with the separator included
func SplitAfter(str string, sep string) []string {
	return strings.SplitAfter(str, sep)
}

// SplitAfterN returns a slice of strings split after the separator, with the separator included, with a maximum of n splits
func SplitAfterN(str string, sep string, n int) []string {
	return strings.SplitAfterN(str, sep, n)
}

// SplitN returns a slice of strings split by the separator, with a maximum of n splits
func SplitN(str string, sep string, n int) []string {
	return strings.SplitN(str, sep, n)
}

// ToLower returns a copy of the string with all characters converted to lowercase
func ToLower(str string) string {
	return strings.ToLower(str)
}

// ToUpper returns a copy of the string with all characters converted to uppercase
func ToUpper(str string) string {
	return strings.ToUpper(str)
}

// Trim returns a copy of the string with all leading and trailing characters in cutset removed
func Trim(str string, cutset string) string {
	return strings.Trim(str, cutset)
}
