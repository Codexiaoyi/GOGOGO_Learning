package game

//创建一条新的蛇
func createSnake() *Snake {
	snake := new(Snake)
	//初始为向右
	snake.Direction = right
	//创建初始身体坐标
	snake.BodyPositions = []Coordinates{
		{1, 1},
		{2, 1},
		{3, 1},
	}
	return snake
}
