package iteration

import "strings"

// Repeat takes a character and returns it repeated X times
func Repeat(character string, times int) string {
	var repeated strings.Builder
	for i := 0; i < times; i++ {
		repeated.WriteString(character)
	}
	return repeated.String()
}
