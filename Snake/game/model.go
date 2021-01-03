package game

type direction int

//方向枚举
const (
	up direction = iota
	down
	left
	right
)

//边界
type Area struct {
	Left   int
	Right  int
	Top    int
	Bottom int
}

//坐标
type Coordinates struct {
	X int
	Y int
}

//食物
type Food struct {
	Position Coordinates
	Show     byte
}

//蛇
type Snake struct {
	Direction     direction
	BodyPositions []Coordinates
}

//蛇笼
type Cage struct {
	Area Area
	Show rune
}
