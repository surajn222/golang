package main

import ("fmt")

func main(){
	fmt.Printf("HW")
}

func getShape(shapeType string){
	switch shapeType {
	case "circle":
		return &circle
	case "square":
		return &square
	}
}

type Shape interface{
	Draw() string
}

type circle struct{}

func (c *circle) Draw() {
	return "Circle"
}

type square struct{}

func (s *square) Draw() {
	return "square"
}

// Interface Shape, method Draw
// Implement Shape -> Circle
// Implement Shape -> Square
// getShape -> switch case -> Circle, Square
