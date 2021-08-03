package main

import (
	"fmt"
	"reflect"
)
/*
func add(x float64, y float64)float64{
	return x+y
}*/

// This can be refactored as

func add(x, y float64)float64{
	return x+y
}

func concatenate(str1,str2 string) (string,string){
	return str1, str2
}

func main(){
	/*
	var num1 float64= 5.6
	var num2 float64= 10.4
	OR
	var num1, num2 float64 = 5.6, 10.4
	*/

	// After refactoring this piece of code
	// This can be used only within a function for declaration with implicit type.

	num1, num2:= 5.6, 10.4
	fmt.Println(add(num1, num2))

	//Function with multiple return values

	str1, str2:= concatenate("Hello", "World!")
	fmt.Println(str1, str2)

	//Type-casting
	var a = 16
	var b float64 = float64(a)
	fmt.Println(b)
	//Type-inference
	var x float64
	y := x //Now y is also float64
	fmt.Println("y is a", reflect.TypeOf(y),"variable")
}


