package game

import (
	"fmt"
)

//创建一个新的蛇笼
func createCage(width int, height int) *Cage {
	cage := new(Cage)
	cage.Show = '#'
	cage.Area.Left = 0
	cage.Area.Right = width
	cage.Area.Top = 0
	cage.Area.Bottom = height
	drawCage(cage)
	return cage
}

//绘制蛇笼
func drawCage(cage *Cage) {
	for i := 0; i < cage.Area.Bottom; i++ {
		for j := 0; j < cage.Area.Right; j++ {
			if i == 0 || i == cage.Area.Bottom-1 || j == 0 || j == cage.Area.Right-1 {
				fmt.Printf("%c", cage.Show)
			} else {
				fmt.Printf("%c", ' ')
			}
		}
		//绘制完一行，换行
		fmt.Printf("\n")
	}
}
