package main

import "fmt"

// interface Shape
// Shape implement Circle
//		Clone()
//		GetInfo()
// Shape implement Rectangle
//		Clone()
//		GetInfo()

// circle1Object = circle{}
// circle2Object = circle1Object.Clone()

// rectangeeObject = rectangle{}
// rectangle2Object = rectangle1Object.Clone()


// Shape is the prototype interface
type Shape interface {
    Clone() Shape
    GetInfo() string
}

// Circle is a concrete prototype
type Circle struct {
    Radius int
}

func (c *Circle) Clone() Shape {
    newCircle := *c
    return &newCircle
}

func (c *Circle) GetInfo() string {
    return fmt.Sprintf("Circle with radius %d", c.Radius)
}

// Rectangle is another concrete prototype
type Rectangle struct {
    Width  int
    Height int
}

func (r *Rectangle) Clone() Shape {
    newRectangle := *r
    return &newRectangle
}

func (r *Rectangle) GetInfo() string {
    return fmt.Sprintf("Rectangle with width %d and height %d", r.Width, r.Height)
}

func main() {
    circle1 := &Circle{Radius: 5}
    circle2 := circle1.Clone()

    rectangle1 := &Rectangle{Width: 4, Height: 6}
    rectangle2 := rectangle1.Clone()

    fmt.Println(circle1.GetInfo()) // Output: Circle with radius 5
    fmt.Println(circle2.GetInfo()) // Output: Circle with radius 5

    fmt.Println(rectangle1.GetInfo()) // Output: Rectangle with width 4 and height 6
    fmt.Println(rectangle2.GetInfo()) // Output: Rectangle with width 4 and height 6
}
