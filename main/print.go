package sudoku

import "fmt"

func SpaceBetweenChar(s string) string {
	result := ""
	for i, char := range s {
		result += string(char)
		if i < len(s) {
			result += " "
		}
	}
	return result
}

func Show(args []string) {
	for i := 0; i < len(args); i++ {
		fmt.Println(SpaceBetweenChar(args[i]))
	}
}
