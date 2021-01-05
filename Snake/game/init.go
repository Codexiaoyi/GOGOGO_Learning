package game

import tl "github.com/JoelOtter/termloop"

//准备游戏
func PrepareGame() {
	//创建蛇笼
	//createCage(50, 50)
	game := tl.NewGame()
	level := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorWhite,
	})

	cage := createCage(50, 25)
	level.AddEntity(cage)
	//level.AddEntity(tl.NewEntity(10, 10, 100, 100))
	game.Screen().SetLevel(level)
	game.Start()
}
