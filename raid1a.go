package student

import (
	"github.com/01-edu/z01"
)

func Raid1a(x, y int) {
	for i := 1; i <= y; i++ {
		for j := 1; j <= x; j++ {

			if j == 1 && i == 1 ||
				j == 1 && i == y ||
				j == x && i == 1 ||
				j == x && i == y {
				z01.PrintRune('o')
			} else if j > 1 && j < x && i == 1 || i == y {
				z01.PrintRune('-')
			} else if j == 1 || j == x && i < y && i > 1 {
				z01.PrintRune('|')
			} else {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}
