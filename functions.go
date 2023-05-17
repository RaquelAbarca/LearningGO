package main

import "fmt"

func welcome(name string) {
  fmt.Println("Hello, " + name)
}

func sum(a int, b int){
	fmt.Println(a+b)
}

func div(x, y int) int {
	return x/y
}

func main() {
  welcome("David")
  welcome("James")

  sum(42, 8)

  result := div(42, 8)
  fmt.Println(result)
}

//A useful feature of Go is that functions can return multiple values!

func swap(x, y int) (int, int) {
    return y, x
} 

//A defer statement ensures that the function is called only after the surrounding function returns.

func welcome() {
	fmt.Println("Welcome")
}
  
func main() {
	defer welcome()
	fmt.Println("Hey")
}

/*The code above will first output "Hey" and only after that output the result of the welcome() function.
This happens because the call to welcome() is deferred, meaning it waits until main() finishes execution and only then calls it.
defer is often used for cleanup, for example, to release resources used by the code, such as files, connections, etc.*/

//If you have deferred multiple function calls, they will execute in last-in-first-out order.

func main() {
	fmt.Println("start")
  
	for i := 0; i < 5; i++ {
	  defer fmt.Println(i)
	}
	fmt.Println("end")
  }

