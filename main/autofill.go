package sudoku

func AutoFill(Grille []string) bool {
	Changes := false
	for i1 := range Grille {
		str := ""
		for i2 := range Grille[i1] {
			if Grille[i1][i2] == '.' {
				Values := GetPreciseNumber(i2, i1, Grille)
				if len(Values) == 1 {
					Changes = true
					str = str + Values[0]
				} else {
					str = str + string(Grille[i1][i2])
				}
			} else {
				str = str + string(Grille[i1][i2])
			}
		}
		Grille[i1] = str
	}
	return Changes
}

func BruteForceFill(Grille []string) []string {
	var NGrille []string
	for _, copy := range Grille {
		NGrille = append(NGrille, copy)
	}
	AutoFill(NGrille)
	for Line := range Grille {
		str := []byte(NGrille[Line])
		for Column := range Grille[Line] {
			if Grille[Line][Column] == '.' {
				Values := GetPreciseNumber(Column, Line, NGrille)
				for _, Value := range Values {
					str[Column] = Value[0] /*Value étant en string, il nous faut récupérer sa valeur en rune*/
					NGrille[Line] = string(str)
					if CheckSudoku(NGrille) {
						return NGrille
					}
					NextValue := BruteForceFill(NGrille)
					if NextValue != nil {
						return NextValue
					}
				}
			}
		}
	}
	return nil
}
