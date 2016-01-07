package main

import (
	"github.com/nsf/termbox-go"
	"math/rand"
	"time"
)

var start = true

const (
	fx    = 5
	fy    = 5
	fsize = 20
)

/*
Point consists of a position (x, y) that is a body of snake
*/
type Point struct {
	X int
	Y int
}

/*
Snake indicates a body of snake includes the positions and direction
*/
type Snake struct {
	Elm []Point
	Dir int // 0: up 1: down 2: right 3: left
}

func detectCollision(snake *Snake) bool {
	head := snake.Elm[0]
	if head.X == fx || head.X == fx+fsize {
		return true
	}
	if head.Y == fy || head.Y == fy+fsize {
		return true
	}
	for _, pos := range snake.Elm[1:] {
		if head.X == pos.X && head.Y == pos.Y {
			return true
		}
	}
	return false
}

func drawMap() {
	bg := termbox.ColorRed

	for i := 0; i <= fsize; i++ {
		termbox.SetCell(fx, fy+i, ' ', termbox.ColorDefault, bg)
		termbox.SetCell(fx+fsize, fy+i, ' ', termbox.ColorDefault, bg)

		if i == 0 || i == fsize {
			for j := 1; j <= fsize; j++ {
				termbox.SetCell(fx+j, fy+i, ' ', termbox.ColorDefault, bg)
			}
		}
	}
}

func initFeed() []Point {
	feed := make([]Point, 0)
	for i := 0; i < 5; i++ {
		feed = append(feed, createFeed())
	}
	return feed
}

func createFeed() Point {
	for {
		x, y := rand.Int()%fsize+fx, rand.Int()%fsize+fy
		if x > fx && x < fsize+fx && y > fy && y < fsize+fy {
			return Point{x, y}
		}
	}
}

func drawFeed(feed []Point) {
	for _, f := range feed {
		termbox.SetCell(f.X, f.Y, '+', termbox.ColorDefault, termbox.ColorDefault)
	}
}

func (snake *Snake) eatFeed(feed []Point) []Point {
	head := snake.Elm[0]
	for i, f := range feed {
		if head.X == f.X && head.Y == f.Y {
			tmp := []Point{}
			for j, elm := range feed {
				if i != j {
					tmp = append(tmp, elm)
				}
			}
			snake.grow()
			tmp = append(tmp, createFeed())
			return tmp
		}
	}
	return feed
}

func draw(s *Snake) {
	feed := initFeed()
	for {
		termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
		drawMap()
		drawFeed(feed)
		for _, e := range s.Elm {
			termbox.SetCell(e.X, e.Y, '@', termbox.ColorDefault, termbox.ColorDefault)
		}
		termbox.Flush()
		if detectCollision(s) {
			start = true
			return
		}
		feed = s.eatFeed(feed)
		time.Sleep(300 * time.Millisecond)
		s.updatePos()
	}
}

func (s *Snake) grow() {
	s.Elm = append(s.Elm, s.Elm[len(s.Elm)-1])
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
			case termbox.KeySpace:
				if start {
					start = false
					snake = initSnake()
					go draw(&snake)
				}
			}
		}
	}
}
