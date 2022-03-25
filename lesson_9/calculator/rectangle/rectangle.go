package rectangle

import "fmt"

func init() {
	fmt.Println("Hello! I'm function in rectangle package")
}

type Rectangle struct {
	A, B  int
	color string
}

func New(a, b int, color string) *Rectangle {
	return &Rectangle{a, b, color}
}

func (r *Rectangle) Perimeter() int {
	return 2 * (r.A + r.B)
}
