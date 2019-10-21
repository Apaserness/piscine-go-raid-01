package student

import "github.com/01-edu/z01"

func topD(x int) {
	for i := 1; i <= x; i++ { //1 massiv
		if i == 1 {
			z01.PrintRune('A')
		} else if i == x {
			z01.PrintRune('C')
		} else {
			z01.PrintRune('B')
		}
	}
	z01.PrintRune(10)
}

func middleD(x, y int) {
	for i := 2; i <= y-1; i++ {
		for j := 1; j <= x; j++ {
			if j == 1 || j == x {
				z01.PrintRune('B')
			} else {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune(10)
	}
}

func bottomD(x int) {
	for i := 1; i <= x; i++ { //1 massiv
		if i == 1 {
			z01.PrintRune('A')
		} else if i == x {
			z01.PrintRune('C')
		} else {
			z01.PrintRune('B')
		}
	}
	z01.PrintRune(10)
}
func AltRaid1d(x, y int) {
	if x >= 1 && y >= 1 {
		topD(x)
		middleD(x, y)
		if y != 1 {
			bottomD(x)
		}
	}
}
