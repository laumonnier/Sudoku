package sudoku

func ReturnSquare(Slice []string, SquareId int) []string {
	var values []string
	SquareLine := 0
	for SquareId >= 3 {
		SquareId = SquareId - 3
		SquareLine++
	}
	for Line := SquareLine * 3; Line <= SquareLine*3+2; Line++ {
		for Column := SquareId * 3; Column <= SquareId*3+2; Column++ {
			values = append(values, string(Slice[Line][Column]))
		}
	}
	return values
}

func ReturnRemainingLine(s string) []string {
	var result []string
	for i := '1'; i <= '9'; i++ {
		variable := false
		for _, char := range s {
			if char == i {
				variable = true
			}
		}
		if variable == false {
			result = append(result, string(i))
		}
	}
	return result
}

func ReturnRemainingColumn(Slice []string, ColumnId int) []string {
	var Remaining []string
	SliceOfValues := []rune{'1', '2', '3', '4', '5', '6', '7', '8', '9'}
	for Number := 1; Number <= 9; Number++ {
		IsNumberIn := false
		for _, str := range Slice {
			if rune(str[ColumnId]) == SliceOfValues[Number-1] {
				IsNumberIn = true
			}
		}
		if !IsNumberIn {
			Remaining = append(Remaining, string(SliceOfValues[Number-1]))
		}
	}
	return Remaining
}

func ReturnRemainingSquare(Slice []string, SquareId int) []string {
	Slice = ReturnSquare(Slice, SquareId)
	var Remaining []string
	for Number := '1'; Number <= '9'; Number++ {
		IsNumberIn := false
		for _, str := range Slice {
			if str == string(Number) {
				IsNumberIn = true
			}
		}
		if IsNumberIn == false {
			Remaining = append(Remaining, string(Number))
		}
	}
	return Remaining
}

func GetPreciseNumber(x int, y int, Grille []string) []string {
	var TestTroisFonctions []string
	number := x / 3
	number = number + y/3*3
	Line := ReturnRemainingLine(Grille[y])
	Column := ReturnRemainingColumn(Grille, x)
	Square := ReturnRemainingSquare(Grille, number)
	for _, GetElementFromLine := range Line {
		for _, GetElementFromColumn := range Column {
			for _, GetElementFromSquare := range Square {
				if GetElementFromSquare == GetElementFromLine && GetElementFromLine == GetElementFromColumn {
					TestTroisFonctions = append(TestTroisFonctions, GetElementFromLine)
				}
			}
		}
	}
	return TestTroisFonctions
}
