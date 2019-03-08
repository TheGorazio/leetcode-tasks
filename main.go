package main

import "fmt"

func longestPrefix(strings []string) string {
	var prefix string

	if len(strings) == 0 {
		return prefix
	}

	var flag bool = true

	for i := 0; i < len(strings[0]); i++ {
		letter := strings[0][i]

		for j := 1; j < len(strings); j++ {
			word := strings[j]

			if i >= len(word) {
				flag = false
				break
			}

			flag = word[i] == letter

			if flag == false {
				break
			}
		}

		if flag {
			prefix += string(letter)
		} else {
			break
		}

	}

	return prefix
}

func main() {
	fmt.Println(longestPrefix([]string{"flower", "flow", "flight", "flon"}))

	fmt.Println(longestPrefix([]string{}))
}
