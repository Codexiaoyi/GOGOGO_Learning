package game

import tl "github.com/JoelOtter/termloop"

var cage *Cage
var game *tl.Game

//准备游戏
func PrepareGame() {
	//创建蛇笼
	//createCage(50, 50)
	game = tl.NewGame()
	level := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorWhite,
	})

	cage = createCage(50, 25)
	snake := CreateSnake()
	level.AddEntity(cage)
	level.AddEntity(snake)
	//level.AddEntity(tl.NewEntity(10, 10, 100, 100))
	game.Screen().SetLevel(level)
	game.Screen().SetFps(10)
	game.Start()
}

//结束游戏
func EndGame() {
	game.Screen().SetFps(0)
}
