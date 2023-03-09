package ransomNote

import (
	"strings"
)

func canConstruct(ransomNote string, magazine string) bool {
	for _, char := range ransomNote {
		if strings.Index(magazine, string(char)) == -1 {
			return false
		}
		magazine = strings.Replace(magazine, string(char), "", 1)
	}
	return true
}
