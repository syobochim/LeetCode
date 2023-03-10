package roman;

import (
    "fmt"
    "strings"
)

func romanToInt(s string) int {
    result := 0;

    if strings.Contains(s, "IV") {
        s = strings.ReplaceAll(s, "IV", "")
        result += 4
    }
    if  strings.Contains(s, "IX") {
        s = strings.ReplaceAll(s, "IX", "")
        result += 9
    }
    if strings.Contains(s, "XL") {
        s = strings.ReplaceAll(s, "XL", "")
        result += 40
    }
    if strings.Contains(s, "XC") {
        s = strings.ReplaceAll(s, "XC", "")
        result += 90
    }
    if strings.Contains(s, "CD") {
        s = strings.ReplaceAll(s, "CD", "")
        result += 400
    }
    if strings.Contains(s, "CM"){
        s = strings.ReplaceAll(s, "CM", "")
        result += 900
    }
    fmt.Printf("string : %s, result: %d \n", s, result)
    for _, char := range s {
        switch char {
        case 'I': result += 1
        case 'V': result += 5
        case 'X': result += 10
        case 'L': result += 50
        case 'C': result += 100
        case 'D': result += 500
        case 'M': result += 1000
        }
    }
    return result
}