package main

import (
	"math"
	"strings"

	"github.com/fogleman/gg"
)

type point struct {
	x, y float64
}

func main() {
	dc := gg.NewContext(1000, 1000)
	dc.SetLineWidth(5)
	dc.Translate(1000/2, 1000)
	dc.SetRGB(55, 55, 55)
	// p1 := point{}
	// len := -100.0
	// p2 := point{x: p1.x, y: p1.y + len}
	// dc.DrawLine(p1.x, p1.y, p2.x, p2.y)

	// tp1, tp2 := p1, p2
	// p1, p2 = translate(p1, p2, 45)
	// dc.DrawLine(p1.x, p1.y, p2.x, p2.y)

	// p1, p2 = tp1, tp2
	// p1, p2 = translate(p1, p2, 0)
	// dc.DrawLine(p1.x, p1.y, p2.x, p2.y)

	str := "FF+[+F-F-F]-[-F+F+F]"
	str = strings.Replace(str, "F", "FF+[+F-F-F]-[-F+F+F]", -1)
	str = strings.Replace(str, "F", "FF+[+F-F-F]-[-F+F+F]", -1)
	str = strings.Replace(str, "F", "FF+[+F-F-F]-[-F+F+F]", -1)

	interpret(str, dc)
	dc.Stroke()

	dc.SavePNG("out.png")
}

func interpret(val string, dc *gg.Context) {
	angle := 0.0
	p1 := point{}
	p2 := point{x: p1.x, y: p1.y - 100}
	stack := Stack{}

	for _, v := range val {
		switch v {
		case 70: // F
			dc.DrawLine(p1.x, p1.y, p2.x, p2.y)
			p1, p2 = translate(p1, p2, angle)
		case 43: // +
			angle += 45
		case 45: // -
			angle -= 45
		case 91: // [
			stack.Push(p1, p2)
		case 93: // ]
			p1, p2 = stack.Pop()
		}
	}
}

func translate(p1, p2 point, a float64) (point, point) {
	s := math.Sin(a)
	c := math.Cos(a)

	p2.x -= p1.x
	p2.y -= p1.y

	p1 = p2

	X := p2.x*c - p2.y*s
	Y := p2.x*s + p2.y*c
	// fmt.Println(p1, p2)
	return p1, point{x: X + p1.x, y: Y + p1.y}
}

type Stack struct {
	p1 []point
	p2 []point
}

func (s *Stack) Push(tp1, tp2 point) {
	s.p1 = append(s.p1, tp1)
	s.p2 = append(s.p2, tp2)
}

func (s *Stack) Pop() (point, point) {
	tp1 := s.p1[len(s.p1)-1]
	tp2 := s.p2[len(s.p2)-1]
	if len(s.p1) > 0 {
		s.p1 = s.p1[:len(s.p1)-1]
	}

	if len(s.p2) > 0 {
		s.p2 = s.p2[:len(s.p2)-1]
	}

	return tp1, tp2
}
