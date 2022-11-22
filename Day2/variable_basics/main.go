// Variables in Go
// A variable is a name for amemory location where a value of a specific type is stored
// in Go, a variable belongs and it is created at runtime
//A declared variable must be used or we get an error!
// _ is the blank identifier and mutes the compile-time error returned by unused variables

// Declaring Variables:
//1. Using the var keyword
// var x int = 7
// var s1 string
// s1 = "Learning Go!"

// 2. Using the short declaration operator (:=)
// age := 30

// in go we use semicolons if we want to write multiple statements on the same line
// statements are instructions that tell go to execute something
// we use keyword identifier type to declare something
package main

import "fmt"
func main()  {
	//int age = 10; //how to declare a variable in c++
	var age int = 30 // var is the keyword;age is the name of the variable; int is the type ; optionally equals and the initial value which is 30
	// if int is ommited; var age = 30, Go will infer from the right side of the equation the type of the variable
	fmt.Println("Age:", age)
		//	Age: 30

	var name = "Dan"
	fmt.Println("Your name is:", name)
		// Your name is: Dan

	var new_name = "Phyllis"
	//fmt.Println("My name is:", new_name)
	_ = new_name
	//in Go code, every declared variable must be used or it will show an error during compilation. 
	//to silence this error (only use with caution),
    //you can use the blank _ = new_name


	// short decaration operator; works only in block scope
	//used when declaring new variables or at least one new variable, cannot be used to declared already used variables
	A := "Learning Golang"
	fmt.Println(A)
		// Learning Golang


	// Multiple Declarations
	car, cost := "Audi", 50000
	fmt.Println(car, cost)
	  // Audi 50000

	//car, cost := "BMW", 2018 ;this will produce an error because there are no new variables to the left side of the go equation, 
	// so short operator canot be used. There has to be at least one new variable to use the short operator, eg.
	car, year := "BMW", 2018
	_ = year // the year variable is not being used in this case so we use the blank _

	var opened = false
	opened, file := true, "a.txt"
	_, _ = opened, file

   // Declaring multiple variables with the var keyword:
   var (
		salary float64
		firstName string
		gender bool
   )
   fmt.Println(salary, firstName, gender)
   		// 0 false

  var a, b, c int
  fmt.Println(a, b, c)
   		// 0 0 0

  var i, j int
  i, j = 5, 8
  _, _ = i, j

  // swap 2 variables using multiple assignments
  // this is also a tuple assignment that allows several variables to be assigned at once.

  var d, f int
  d, f = 6, 8
  f, d = d, f
  fmt.Println(d, f)
   	// 8 6

  sum := 5 + 2.3
  fmt.Println(sum)
   	// 7.3


//Types and Zero Values
 var g = 4
 var h = 5.2
 //g = h
 //fmt.Println(g, h)
   	//cannot use h (variable of type float64) as type int in assignment 
	// the type of a variable cannot be changed once it has been declared
    // to use the value of h, you have to fist convert they type of variable from float to int
 g = int(h)
 fmt.Println(g, h)
   	// 5 5.2

// literal is the value itself 
}


