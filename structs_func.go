package main

import "fmt"
import "math"
import "errors"




func main(){
	var c = Circle{r:5}
	var areaofshape = getArea(c)
	fmt.Printf("Area of Shape is %f", areaofshape)
}




type Shape interface {
	area() float64
}

type Circle struct {
	r float64
}

func (cir Circle) area() float64{
	return math.Pi * cir.r * cir.r
}

func getArea(s Shape) float64{
	return s.area()
}
