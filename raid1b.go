package student

import "github.com/01-edu/z01"

func Top(x int) {
	for i := 1; i <= x; i++ { //1 massiv
		if i == 1 {
			z01.PrintRune('/')
		} else if i == x {
			z01.PrintRune(92)
		} else {
			z01.PrintRune('*')
		} 
		}
	z01.PrintRune(10)
}

func Middle(x, y int) {
	for i := 2; i <= y-1; i++ {
		for j := 1; j <= x; j++ {
			if j == 1 || j == x {
				z01.PrintRune('*')
			} else {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune(10)
	}
}

func Bottom(x int) {
	for i := 1; i <= x; i++ { //1 massiv
		if i == 1 {
			z01.PrintRune(92)
		} else if i == x {
			z01.PrintRune('/')
		} else {
			z01.PrintRune('*')
		} 
	}
	z01.PrintRune(10)
}
func Raid1a(x, y int) {
	if x >= 1 && y >= 1 {
		Top(x)
		Middle(x,y)
		if y != 1 {
			Bottom(x)
		}
	}
}
