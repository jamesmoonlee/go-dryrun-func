package main

import "fmt"
import "math"
import "errors"

func main(){
	var err error
	var result float64

	var x float64 = 9
	result, err = sqrtNum(x)
	if (err != nil) {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

}


func sqrtNum(x float64) (float64, error){
	if (x < 0) {
		return 0, errors.New("Invalid value for sqrt operate")
	} else {
		return math.Sqrt(x), nil
	}

}
