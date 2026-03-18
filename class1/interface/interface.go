package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Perimeter () float64
	Area () float64
}

type Circle struct {
	radius float64
}

type Rectangle struct {
	length, width float64
}

type Square struct {
	length float64
}

func (c Circle) Area() float64 {
	return  math.Pi * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return  2 * math.Pi * c.radius
}

func (r Rectangle) Area() float64 {
	return  r.length * r.width
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.length + r.width)
}

func (s Square) Area() float64{
	return s.length * s.length
}

func (s Square) Perimeter() float64 {
	return  4 * s.length
}

func PrintArea(s interface{}){
	value, ok := s.(Shape)

	if ok {
		fmt.Println("Area: ",value.Area())
	}else{
		fmt.Println("Not a Shape interface")
	}
}

func calculateArea(shape interface{}) {
	switch s := shape.(type) {
	case Circle:
		fmt.Printf("Circle area: %.2f\n",s.Area())
	case Rectangle:
		fmt.Printf("Rectangle area: %.2f\n",s.Area())
	case Square:
		fmt.Printf("Square area: %.2f\n",s.Area())
	default:
		fmt.Println("Unknown shape")
	} 
}

///Implementaction 
func main () {
	var s Shape

	fmt.Println("Circle: ")
	s = Circle{radius: 5}
	fmt.Println("C Area: ",s.Area())
	fmt.Println("C Perimeter: ",s.Perimeter())

	fmt.Println("Rectangle: ")
	s = Rectangle{length: 4, width: 3}
	fmt.Println("R Area: ",s.Area())
	fmt.Println("R Perimeter: ",s.Perimeter())

	fmt.Println("Square: ")
	s = Square{length: 5}
	fmt.Println("S Area: ",s.Area())
	fmt.Println("S Perimeter: ",s.Perimeter())

	fmt.Println("Type Assertion")
	sq := Square{length: 5}
	PrintArea(sq)

	noS := 5
	PrintArea(noS)

	fmt.Println("switch Type ")
	c := Circle{radius: 5}
	r := Rectangle{length: 4,width: 3}
	
	calculateArea(c)
	calculateArea(r)
	calculateArea(sq)
	calculateArea(noS)
}