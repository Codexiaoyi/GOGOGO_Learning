package game

import (
	tl "github.com/JoelOtter/termloop"
)

type direction int

//方向枚举
const (
	up direction = iota
	down
	left
	right
)

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
	*tl.Entity
	Direction     direction
	BodyPositions []Coordinates
	BodyLen       int
}

//蛇笼
type Cage struct {
	*tl.Entity
	width, height int
	Coords        map[Coordinates]int
}
