package ransomNote

import (
	"strings"
)

func canConstruct(ransomNote string, magazine string) bool {
	for _, char := range ransomNote {
		index := strings.Index(magazine, string(char))
		if index == -1 {
			return false
		}
		magazine = magazine[:index] + magazine[index+1:]
	}
	return true
}
