package isogram

import "strings"

const (
	space  = 32
	hyphen = 45
)

func IsIsogram(word string) bool {
	lowercase := strings.ToLower((word))
	for idx, char := range lowercase {
		if char != space && char != hyphen {
			charStr := strings.ToLower(string(char))
			if strings.Contains(lowercase[:idx], charStr) {
				return false
			}
		}
	}
	return true
}
