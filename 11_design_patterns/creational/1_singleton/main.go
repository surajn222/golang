package main

import ("fmt")

func main(){
	println("HW")
	
	singletonObject := getSingletonObj()
	singletonObject.addOne()
	singletonObject.addOne()
	fmt.Printf("%+v", singletonObject)

	singletonObject2 := getSingletonObj()
	singletonObject2.addOne()
	singletonObject2.addOne()
	fmt.Printf("%+v", singletonObject2)

}


var singletonVar *SingletonClass

func getSingletonObj() *SingletonClass{
	if singletonVar == nil {
		singletonVar = newSingletonClass()
	}
	return singletonVar
}


type SingletonClass struct {
	counter int
}

func (s *SingletonClass) addOne() {
	s.counter++
}

func newSingletonClass() *SingletonClass {
	return &SingletonClass{counter: 0}
}

// struct singleton, var counter
// method addOne, increment counter
// Constuctor singleton struct

// var singletonVar *SingletonClass
// func getSingletonObject
