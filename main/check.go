package sudoku

func UndesirableChar(tab []string) bool {
	for _, s := range tab {
		for _, char := range s {
			if (char >= '1' && char <= '9') == false {
				return false
			}
		}
	}
	return true
}

func UndesirableCharBegin(tab []string) bool {
	for _, s := range tab {
		for _, char := range s {
			if (char >= '1' && char <= '9' || char == '.') == false {
				return false
			}
		}
	}
	return true
}

func DoubleDigitLineCheck(nb string) bool {
	for i := 0; i <= len(nb)-1; i++ {
		for j := i + 1; j < len(nb); j++ {
			if nb[i] == nb[j] {
				return false
			}
		}
	}
	return true
}

func CheckAllLines(tab []string) bool {
	for i := 0; i < 9; i++ {
		if DoubleDigitLineCheck(tab[i]) == false {
			return false
		}
	}
	return true
}

func DoubleDigitColumnCheck(grille []string, nb int) bool {
	for i := 0; i < len(grille)-1; i++ {
		for j := i + 1; j < len(grille); j++ {
			if grille[i][nb] == grille[j][nb] {
				return false
			}
		}
	}
	return true
}

func CheckAllColumns(Grille []string) bool {
	for Column := 0; Column <= 8; Column++ {
		if DoubleDigitColumnCheck(Grille, Column) == false {
			return false
		}
	}
	return true
}

func DoubleDigitSquareCheck(Slice []string, SquareId int) bool {
	values := ReturnSquare(Slice, SquareId)
	for Index1, Element1 := range values {
		for Index2, Element2 := range values {
			if Index1 != Index2 && Element1 == Element2 {
				return false
			}
		}
	}
	return true
}

func CheckAllSquares(Slice []string) bool {
	for SquareId := 0; SquareId <= 8; SquareId++ {
		if DoubleDigitSquareCheck(Slice, SquareId) == false {
			return false
		}
	}
	return true
}

func SizeCheck(tab []string) bool {
	if len(tab) != 9 {
		return false
	}
	for _, el := range tab {
		if len(el) != 9 {
			return false
		}
	}
	return true
}

func CheckSudoku(Grille []string) bool {
	if UndesirableChar(Grille) {
		if CheckAllSquares(Grille) && CheckAllLines(Grille) && CheckAllColumns(Grille) {
			return true
		}
	}
	return false
}
