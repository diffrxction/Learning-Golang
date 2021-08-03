package main

import "fmt"

func main(){
	/*
	var x = 16
	var a = &x
	*/
	x := 16
	a := &x
	//Accessing address and value at address through pointers
	fmt.Println("Address of x is", a)
	fmt.Println("Value stored at the address is", *a)
	//Changes made in the memory at the address stored in pointer also reflects on the variable.
	*a = 40
	fmt.Println("Value is now = ",x)
	//Pointers during arithmetic operations
	*a = *a**a
	fmt.Println("New value is ",*a)
	fmt.Println("x also changes to",x)
	//Address can be accessed by
	fmt.Println(&x)
	fmt.Println(a)

}
