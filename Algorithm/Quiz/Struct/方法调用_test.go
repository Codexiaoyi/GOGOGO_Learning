package struct1

import (
	"fmt"
	"testing"
)

type Combine struct {
}

type CombineChild struct {
	Combine
}

func (c *Combine) Print() {
	fmt.Println("Combine print...")
}

func (c *CombineChild) Print() {
	fmt.Println("CombineChild print...")
}

func TestPrint(t *testing.T) {
	c := &CombineChild{}
	c.Print()
}
