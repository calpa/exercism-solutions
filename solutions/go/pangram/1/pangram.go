package pangram

import "strings"

func IsPangram(str string) bool {
	letters := make(map[rune]bool)

    str = strings.ToLower(str)
    for _, c := range(str) {
        if c >= 'a' && c <= 'z' {
            letters[c] = true
        }
    }

    return len(letters) == 26
}
