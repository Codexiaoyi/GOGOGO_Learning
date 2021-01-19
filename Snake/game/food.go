package game

import (
	"math/rand"
	"time"

	tl "github.com/JoelOtter/termloop"
)

//创建食物
func createFood() *Food {
	food := new(Food)
	food.Entity = tl.NewEntity(0, 0, 0, 0)
	rand.Seed(time.Now().UnixNano())
	food.movePosition()
	return food
}

//绘制食物
func (food *Food) Draw(screen *tl.Screen) {
	screen.RenderCell(food.Site.X, food.Site.Y, &tl.Cell{
		Fg: tl.ColorRed,
		Ch: '*',
	})
}

//termloop的DynamicPhysical接口实现碰撞检测（需要搭配下面两个函数）
func (food *Food) Collide(collision tl.Physical) {
	switch collision.(type) {
	case *Snake:
		food.movePosition()
	}
}

func (food *Food) Position() (int, int) {
	return food.Site.X, food.Site.Y
}

func (food *Food) Size() (int, int) {
	return 1, 1
}

//移动位置
func (food *Food) movePosition() {
	newX := randInRange(1, cage.width-1)
	newY := randInRange(1, cage.height-1)
	food.Site.X, food.Site.Y = newX, newY
	food.SetPosition(newX, newY)
}

//一定范围内随机数
func randInRange(min, max int) int {
	return rand.Intn(max-min) + min
}
