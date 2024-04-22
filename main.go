package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	width   = 50
	height  = 50
	cls     = "\033c\x0c"
	deadsq  = "⬛"
	alivesq = "🟩"
)

type Map [][]bool

func createMap() Map {
	m := make(Map, height)
	for i := range m {
		m[i] = make([]bool, width)
	}
	return m
}

func (m Map) Random() {
	for _, i := range m {
		for j := range i {
			if rand.Intn(4) == 1 {
				i[j] = true
			}
		}
	}
}

func (m Map) Display() {
	fmt.Print("\033[H") // Move cursor to the top-left corner of the screen
output := ""
	for _, i := range m {
		for _, j := range i {
			switch {
			case j:
				output += alivesq
			default:
				output += deadsq
			}
		}
		output += "\n"
	}
	fmt.Print(output)
}

func (m Map) Alive(x, y int) bool {
	y = (height + y) % height
	x = (width + x) % width
	return m[y][x]
}

func (m Map) Neibourghs(x, y int) int {
	var count int
	for i := y - 1; i <= y+1; i++ {
		for j := x - 1; j <= x+1; j++ {
			if i == y && j == x {
				continue
			}
			if m.Alive(j, i) {
				count++
			}
		}
	}
	return count
}

func (m Map) Rulescheck(x, y int) bool {
	cell := m.Neibourghs(x, y)
	alive := m.Alive(x, y)
	if cell < 4 && cell > 1 && alive {
		return true
	} else if cell == 3 && !alive {
		return true
	} else {
		return false
	}
}

func Generate(a, b Map) {
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			b[i][j] = a.Rulescheck(j, i)
		}
	}
}

func main() {
	fmt.Println(cls)
	aWorld := createMap()
	bWorld := createMap()
	aWorld.Random()
	for {
		aWorld.Display()
		Generate(aWorld, bWorld)
		aWorld, bWorld = bWorld, aWorld
		time.Sleep(300 * time.Millisecond)
		fmt.Println(cls)
	}
}
