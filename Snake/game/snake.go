package game

import (
	tl "github.com/JoelOtter/termloop"
)

//创建一条新的蛇
func createSnake() *Snake {
	snake := new(Snake)
	snake.Entity = tl.NewEntity(3, 5, 1, 1)
	//初始为向右
	snake.Direction = right
	//创建初始身体坐标
	snake.BodyPositions = []Coordinates{
		{1, 5},
		{2, 5},
		{3, 5},
	}
	snake.BodyLen = len(snake.BodyPositions)
	return snake
}

//绘制蛇
func (snake *Snake) Draw(screen *tl.Screen) {
	if snake == nil {
		return
	}

	//根据不同方向，蛇向前走一步
	newHead := *snake.head()
	switch snake.Direction {
	case right:
		newHead.X++
	case left:
		newHead.X--
	case up:
		newHead.Y--
	case down:
		newHead.Y++
	}

	if snake.isGrowing() {
		//变大变强
		snake.BodyPositions = append(snake.BodyPositions, newHead)
	} else {
		//普通往前走，切片把第一个去掉
		snake.BodyPositions = append(snake.BodyPositions[1:], newHead)
	}
	snake.SetPosition(newHead.X, newHead.Y)

	//走完判断是否炸了
	if snake.isHitCage() || snake.isHitSelf() {
		EndGame()
	}

	if snake.isHitFood() {
		snake.grow()
	}

	//开画！
	for _, c := range snake.BodyPositions {
		screen.RenderCell(c.X, c.Y, &tl.Cell{
			Fg: tl.ColorGreen,
			Ch: '*',
		})
	}
}

//Tick函数来自Termloop用来处理键盘输入
func (snake *Snake) Tick(event tl.Event) {
	if event.Type == tl.EventKey {
		switch event.Key {
		case tl.KeyArrowRight:
			if snake.Direction != left {
				snake.Direction = right
			}
		case tl.KeyArrowLeft:
			if snake.Direction != right {
				snake.Direction = left
			}
		case tl.KeyArrowUp:
			if snake.Direction != down {
				snake.Direction = up
			}
		case tl.KeyArrowDown:
			if snake.Direction != up {
				snake.Direction = down
			}
		}
	}
}

func (snake *Snake) Collide(collision tl.Physical) {
	switch collision.(type) {
	case *Food:
		snake.grow()
	case *Cage:
		EndGame()
	}
}

//获取头部点位
func (snake *Snake) head() *Coordinates {
	return &snake.BodyPositions[len(snake.BodyPositions)-1]
}

//变大变强！
func (snake *Snake) grow() {
	snake.BodyLen++
}

//本次是否有变大变强
func (snake *Snake) isGrowing() bool {
	return snake.BodyLen > len(snake.BodyPositions)
}

//是否吃到食物
func (snake *Snake) isHitFood() bool {
	return snake.head().X == food.Site.X && snake.head().Y == food.Site.Y
}

//是否撞到自己的身体
func (snake *Snake) isHitSelf() bool {
	for i := 0; i < len(snake.BodyPositions)-1; i++ {
		if *snake.head() == snake.BodyPositions[i] {
			return true
		}
	}
	return false
}

//是否撞到笼子
func (snake *Snake) isHitCage() bool {
	return cage.Contains(*snake.head())
}
