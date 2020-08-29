package String

func judgeCircle(moves string) bool {
	x := 0
	y := 0
	for _, direction := range moves {
		switch direction {
		case 'R':
			x++
		case 'L':
			x--
		case 'U':
			y++
		case 'D':
			y--
		}
	}
	return x == 0 && y == 0
}
