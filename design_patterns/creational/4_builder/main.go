package main

import ("fmt")

// func main() {
// 	builder := NewCarBuilder()
// 	car := builder.SetBrand("Tesla").
// 		SetModel("Model S").
// 		SetEngine("Electric").
// 		SetWheels(4).
// 		SetSeats(5).
// 		Build()

// 	fmt.Println(car) // Output: Car [Brand=Tesla, Model=Model S, Engine=Electric, Wheels=4, Seats=5]
// }



// builder := NewCarBuilder()
// builder = builder.Setbrand("B")
// builder = builder.SetModel("Model S")
// car 	   = builder.Build() // returns Car

// Based on the above
// object carbuilder
//	 	Setbrand return carbuilder,
//	    SetModel return carbuilder,
//	    Build return Car,
// 
// object Car, Brand, Model



// CarBuilder defines the interface for building parts of a Car
type CarBuilder interface {
	SetBrand(brand string) CarBuilder
	SetModel(model string) CarBuilder
	SetEngine(engine string) CarBuilder
	SetWheels(wheels int) CarBuilder
	SetSeats(seats int) CarBuilder
	Build() Car
}

// ConcreteCarBuilder is a concrete builder that implements CarBuilder
type ConcreteCarBuilder struct {
	car Car
}

func NewCarBuilder() *ConcreteCarBuilder {
	return &ConcreteCarBuilder{}
}

func (b *ConcreteCarBuilder) SetBrand(brand string) CarBuilder {
	b.car.Brand = brand
	return b
}

func (b *ConcreteCarBuilder) SetModel(model string) CarBuilder {
	b.car.Model = model
	return b
}

func (b *ConcreteCarBuilder) SetEngine(engine string) CarBuilder {
	b.car.Engine = engine
	return b
}

func (b *ConcreteCarBuilder) SetWheels(wheels int) CarBuilder {
	b.car.Wheels = wheels
	return b
}

func (b *ConcreteCarBuilder) SetSeats(seats int) CarBuilder {
	b.car.Seats = seats
	return b
}

func (b *ConcreteCarBuilder) Build() Car {
	return b.car
}

// Car represents the product being built
type Car struct {
	Brand   string
	Model   string
	Engine  string
	Wheels  int
	Seats   int
}

func (c Car) String() string {
	return fmt.Sprintf("Car [Brand=%s, Model=%s, Engine=%s, Wheels=%d, Seats=%d]", c.Brand, c.Model, c.Engine, c.Wheels, c.Seats)
}



