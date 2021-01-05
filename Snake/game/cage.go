package game

import (
	tl "github.com/JoelOtter/termloop"
)

//创建一个新的蛇笼
func createCage(width int, height int) *Cage {
	cage := new(Cage)
	cage.Entity = tl.NewEntity(0, 0, 0, 0)
	cage.width, cage.height = width-1, height-1
	cage.Coords = make(map[Coordinates]int)

	//绘制顶部和底部
	for x := 0; x < cage.width; x++ {
		cage.Coords[Coordinates{x, 0}] = 1
		cage.Coords[Coordinates{x, cage.height}] = 1
	}

	//绘制左边和右边
	for y := 0; y < cage.height; y++ {
		cage.Coords[Coordinates{0, y}] = 1
		cage.Coords[Coordinates{cage.width, y}] = 1
	}

	return cage
}

//
func (c *Cage) Contains(coordinates Coordinates) bool {
	_, exists := c.Coords[coordinates]
	return exists
}

// 绘制笼子
func (cage *Cage) Draw(screen *tl.Screen) {
	if cage == nil {
		return
	}

	for c := range cage.Coords {
		screen.RenderCell(c.X, c.Y, &tl.Cell{
			Bg: tl.ColorGreen,
		})
	}
}
