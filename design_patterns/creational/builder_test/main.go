package main

import ("fmt")
func main(){
	fmt.Printf("HW")
	cb := newCarbuilder()

	cb = cb.SetModel()
	cb = cb.SetEngine()
	cb = cb.SetBrand()
	c := cb.Build()
	fmt.Printf("C %+v", c)
}

type Car struct {
	Engine string
	Model string
	Brand string
}

type carBuilder struct {
	car Car
}

func newCarbuilder() *carBuilder{
	return &carBuilder{}
}

func (b *carBuilder) SetModel() *carBuilder{
	fmt.Println("Setting Model")
	b.car.Model = "SetModel"
	return b
}

func  (b *carBuilder) SetEngine() *carBuilder{
	fmt.Println("Setting Engine")
	b.car.Engine = "SetBrand"
	return b
}

func  (b *carBuilder) SetBrand() *carBuilder{
	fmt.Println("Setting Brand")
	b.car.Brand = "SetBrand"
	return b
}

func (b *carBuilder) Build() Car{
	fmt.Println("Car built")
	return b.car
}
