package student

import "github.com/01-edu/z01"

func topC(x int) {
	for i := 1; i <= x; i++ {
		if i == 1 || i == x {
			z01.PrintRune('A')
		} else {
			z01.PrintRune('B')
		}
	}
	z01.PrintRune(10)
}

func middleC(x, y int) {
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

func bottomC(x int) {
	for i := 1; i <= x; i++ {
		if i == 1 || i == x {
			z01.PrintRune('C')
		} else {
			z01.PrintRune('B')
		}
	}
	z01.PrintRune(10)
}
func Raid1c(x, y int) {
	if x >= 1 && y >= 1 {
		topC(x)
		middleC(x, y)
		if y != 1 {
			bottomC(x)
		}
	}
}
