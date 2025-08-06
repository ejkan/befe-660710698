package main

import (
	"errors"
	"fmt"
)

func divine(a,b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("you couldn't divine by zero")
	}
	return a/b,  nil
}

func main(){
	result, err := divine (10,0)
	if err != nil{
		fmt.Println("Error!", err)
	}
	fmt.Println("Result is : ", result)
}