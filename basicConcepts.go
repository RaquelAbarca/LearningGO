package main

import "fmt"
import "math"

func main() {
	/*
	This is an example program
	that only generates a simple
	text output using the "fmt" package
	*/

	var i int = 8
    fmt.Println(i) 

	var i, j = 8, 42  //If you assign a value to the variable, you can omit the type declaration, as Go will automatically take the type of the value assigned
  	fmt.Println(i) 
  	fmt.Println(j)

	//Go supports short variable declaration using :=.
	i:= 8
	x, y := 10, 5

	// generate the output
    fmt.Println("Hello, World!")

	/*
	float32 - a single-precision floating point value.
	float64 - a double-precision floating point value.
	string - a text value.
	bool - Boolean true or false.
	*/

	var x int = 42
  	var y float32 = 1.37
  	var name string = "James"
  	var online bool = true

  	fmt.Println(name)
  	fmt.Println(x)
  	fmt.Println(y)
  	fmt.Println(online)

	 const pi = 3.14 

  	fmt.Println(pi)

	//Recall the modulo operator -- it can be used to get the remainder after division.

	var input string
	fmt.Scanln(&input)
	fmt.Println(input)

	var num int
	fmt.Scanln(&num)
	fmt.Println(num*2) 

	//if
	x := 42

	if x > 18 {
    	fmt.Println("Allowed")
  	} else {
    	fmt.Println("Not allowed")
  	}

	//if/else
	if x := 42; x > 18 {
    	fmt.Println("Allowed")
  	} else {
    	fmt.Println("Not allowed")
  	}

	//switch
	num := 3
  	switch num {
    	case 1:
    	  	fmt.Println("One")
    	case 2:
      		fmt.Println("Two")
    	case 3:
    	  	fmt.Println("Three")
    	default:
    		 fmt.Println(num)
  	} 

	switch {
    case x>0 && x<10:
      fmt.Println("something")
    case x > 10:
      fmt.Println("something else")

	//Loop
	for i:=0; i<5; i++ {
    	fmt.Println(i)
  	}
		//The init and post statements can be omitted.
		//For example:
		sum := 1
  		res := 0
  		for sum <= 1000 {
	    res += sum
    	sum++
  }
  fmt.Println(res) 

}
