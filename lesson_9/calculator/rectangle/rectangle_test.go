package rectangle

import (
	"log"
	"testing"
)

func TestPerimeter(t *testing.T) {
	r := New(1, 2, "green")

	if res := r.Perimeter(); res != 6 {
		log.Fatalf("Perimeter with a=%v and b=%v. Expected = 6. Result = %v", r.A, r.B, res)
	}
}
