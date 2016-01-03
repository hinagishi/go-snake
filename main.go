package main

import (
	"github.com/nsf/termbox-go"
	"time"
)

type Point struct {
	X int
	Y int
}

type Snake struct {
	Elm []Point
	Dir int // 0: up 1: down 2: right 3: left
}

func drawMap() {
	x, y := 5, 5
	length := 40
	bg := termbox.ColorRed

	for i := 0; i < length; i++ {
		termbox.SetCell(x, y+i, ' ', termbox.ColorDefault, bg)
		termbox.SetCell(x+length, y+i, ' ', termbox.ColorDefault, bg)

		if i == 0 || i == length-1 {
			for j := 0; j < length; j++ {
				termbox.SetCell(x+j, y+i, ' ', termbox.ColorDefault, bg)
			}
		}
	}
}

func draw(s *Snake) {
	for {
		termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
		drawMap()
		for _, e := range s.Elm {
			termbox.SetCell(e.X, e.Y, '@', termbox.ColorDefault, termbox.ColorDefault)
		}
		termbox.Flush()
		s.updatePos()
		time.Sleep(300 * time.Millisecond)
	}
}

func (s *Snake) updatePos() {
	x := s.Elm[0].X
	y := s.Elm[0].Y
	switch s.Dir {
	case 0:
		y--
	case 1:
		y++
	case 2:
		x++
	case 3:
		x--
	}
	s.Elm = append([]Point{Point{x, y}}, s.Elm[0:len(s.Elm)-1]...)
	return
}

func initSnake() Snake {
	var s Snake
	s.Elm = []Point{Point{20, 20}, Point{20, 19}, Point{20, 18}}
	s.Dir = 1
	return s
}

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	snake := initSnake()
	go draw(&snake)
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				return
			case termbox.KeyArrowUp:
				snake.Dir = 0
			case termbox.KeyArrowDown:
				snake.Dir = 1
			case termbox.KeyArrowRight:
				snake.Dir = 2
			case termbox.KeyArrowLeft:
				snake.Dir = 3
			}
		}
	}
}
