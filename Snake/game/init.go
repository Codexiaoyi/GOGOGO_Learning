package game

import tl "github.com/JoelOtter/termloop"

var cage *Cage
var food *Food
var game *tl.Game

//准备游戏
func PrepareGame() {
	//创建蛇笼
	game = tl.NewGame()
	level := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorWhite,
	})

	cage = createCage(50, 25)
	food = createFood()
	snake := createSnake()

	level.AddEntity(cage)
	level.AddEntity(food)
	level.AddEntity(snake)
	game.Screen().SetLevel(level)
	game.Screen().SetFps(5)
	game.Start()
}

//结束游戏
func EndGame() {
	game.Screen().SetFps(0)
}
