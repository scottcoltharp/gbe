package methods

import "fmt"

type rect struct {
	width, height int
}

func newRect(width, height int) *rect {
	return &rect{width: width, height: height}
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r *rect) perim() int {
	return 2*r.width + 2*r.height
}

func Methods() {
	r := newRect(10, 5)
	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())
}
