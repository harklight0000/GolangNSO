package utils

import "strings"

type String string

func (s String) Contains(str string) bool {
	return strings.Contains(string(s), str)
}

func (s String) Replace(old string, newStr string) string {
	return strings.Replace(string(s), old, newStr, -1)
}

func (s String) Split(sep string) []string {
	return strings.Split(string(s), sep)
}

func (s String) ToLower() string {
	return strings.ToLower(string(s))
}
