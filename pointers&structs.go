package main

import "fmt"

func main() {
  x := 42
  p := &x
  fmt.Println(*p) //42
  fmt.Println(p)  //0xc00001a050
  fmt.Println(&p) //location

  *p = 8
  fmt.Println(*p) //The * operator can also be used to change the value of the memory address the pointer holds:
  fmt.Println(x)
} 
//We were able to change the value of x using the pointer p. The * operator is called the dereferencing operator.


//We can pass pointers as function parameters.

func change(val int) {
  val = 8
}
func change_ptr(ptr *int) {
  *ptr = 8
}

func main() {
  x := 42

  change(x)
  fmt.Println(x)

  change_ptr(&x)
  fmt.Println(x)
}

/*The change() function takes an integer argument and changes its value.
The change_ptr() function does the same using a pointer.

When you run the code, you will see that the change() function did not change the value of our x variable, because the argument is just a copy of its value, while the change_ptr() did change the actual value of x, because it used its memory address as the argument.
Note that we need to pass the memory address using the & operator to functions that take a pointer as their argument.*/


type Contact struct {
  name string
  age int
}

func main() {
  x := Contact{"James", 42}

  x.age = 8
  fmt.Println(x.age)
  fmt.Println(x.name)
}

//Similar to simple pointers, we can also make pointers to structs using the & operator:

type Contact struct {
  name string
  age int
}

func main() {
  x := Contact{"James", 42}
  p := &x

  fmt.Println(p.age)
}

//We can also use pointers when creating a new struct:


type Contact struct {
  name string
  age int
}

func main() {
  p := &Contact{"John", 15}

  fmt.Println(p.name)
}

//Methods are simply functions with a special receiver argument.
// The receiver appears between the func keyword and the method name. The receiver is the Contact struct.

type Contact struct {
  name string
  age int
}

func (x Contact) welcome() {
  fmt.Println(x.name)
  fmt.Println(x.age)
}

func main() {
  x := Contact{"James", 42}
  x.welcome()
}

//In case we need to change the data of the struct in a method, we can use pointers as method receivers.

type Contact struct {
  name string
  age int
}

func (ptr *Contact) increase(val int) {
  ptr.age += val
}

func main() {
  x := Contact{"James", 42}
  x.increase(8)
  fmt.Prin