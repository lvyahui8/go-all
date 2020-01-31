package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

type MyFloat2 float64

func (f MyFloat2) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex3 struct {
	X, Y float64
}

func (v *Vertex3) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	var a Abser
	f := MyFloat2(-math.Sqrt2)
	v := Vertex3{3, 4}

	a = f  // a MyFloat  实现了 Abser
	a = &v // a *Vertex 实现了 Abser -- 在指针类型上实现了方法， 而不是在原类型上实现方法

	// 下面一行 ， v是一个 Vertex （而不是* Vertex（指针））
	// 所以没有实现Abser
	//a = v

	fmt.Println(a.Abs())
}
