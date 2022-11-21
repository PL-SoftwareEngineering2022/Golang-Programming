// Each source file begins with a package declaration that indicate what package the file is part of.
// Every go program ,ust be included in a package.
// eg. package main means that the file belongs to the main package, and this is an executable package that generates
//a standalong executable file that can be run. The name of an executable package is predefined and is always called main.

// the package must also have a function called main, defined inside.

// the package declaration is followed by any import statements and then a sequence of package level declarations of types, variables, 
//constants and functions in any order.

// to run a go program you need to be in the directory that has the main.go file. command to run the program: go run main.go
// if you run into problems running the code, due to a module error you can run; go mod init then go mod tidy then run go run main.go


package main

import "fmt" // the name of the package should be written in quotes; fmt comes from format and is the name of a standard library package included in go
//sometimes it is pronounced as "fampt" and is used to format strings and print out messages to the standard output.
const secondsInHour = 3600

func main(){
	fmt.Println("Hello Go World!") // in the main function, we declare another function declared in the fmt package
	distance := 60.8 //distance in kilometers
	fmt.Printf("The distance in miles is %f \n", distance * 0.621) //%f is called a verb and tells the print function how to format and print the output
			// Hello Go World!
			// The distance in miles is 37.756800 
}

//Coding - Go Application Structure
/////////////////////////////////
// The Typical Structure of a Go Application
// Go Playground: https://play.golang.org/p/vY_IeYBb1GN
/////////////////////////////////
 
// a package clause starts every source file
// main is a special name declaring an executable rather than a library (package)
package main
 
// import declaration declares packages used in this file
import "fmt"
 
// package scoped variables and constants
var x int = 100
 
const y = 0
 
// a function declaration. main is a special function name
// it is the entry point for the executable program
// Go uses brace brackets to delimit code blocks
func main() {
 
    // Local Scoped Variables and Constants Declarations, calling functions etc
    var a int = 7
    var b float64 = 3.5
    const c int = 10
 
    // Println() function prints out a line to stdout
    // It belongs to package fmt
    fmt.Println("Hello Go world!") // => Hello Go world!
    fmt.Println(a, b, c)           // => 7 3.5 10
 
}

// Compiling(go build) and Running Go Applications (go run)
// go is a tool for managing Go source code
// 1. go run : it compiles and runs the application. it does not produce an executab;e
	// go run gile.go compiles and immediately runs go programs
// 2. go build: it just compiles the application. It produces an executable
	// go build file.go compiles a bunch of go source code files. It complies packages and dependencies
	// if you run go build, it will compile files in the current directoy and will produce an executable file with the name of the current working directory.
	// go build -o app will produce an executable filled called app (app.exe)

// to see a lot of output run go with the -x option.
// eg. go run -x main.go

// to give the executable file another name, you use the -o option. 
// go build -o myapp.exe

// compiling for Windows: GOOS=windows GOARCH=amd64 go build -o winapp.exe
// compiling for Linux: GOOS=linux GOARCH=amd64 go build -o linuxapp
// compiling for Mac: GOOS=mac GOARCH=amd64 go build -o macapp

// these options are GOOS for the operating system and GOARCH for the architecture or the CPU type
//3. go install
// both go install and go build will compile teh package in the current directory
// if the package is main, go build will place the resulting executable in the current directory 
	//and go install will move teh executable to GOPATH/bin
// when running go install you use paths relative to GOPATH/src

// Formatting Go SourceCode (gofmt)
// go strongly suggests certain styles
// gofmt which comes from golang formatter will format a program's source code in an idionamtic way that is easy to read and understand
// this too, gofmt, is built into the language runtime and it formats Go code according to a set of stable, well understoof language rules.
// We can run it manually at the command line or we can configure the IDE to run gofmt each time a file is saved.
// eg. gofmt -w main.go
// gofmt -w -l application_structure ==> output will show what file was modified in the specified directory 
 // the -w option writes the results back to the source file. it actually overwrites the file
 // with the -l option it will modify all the go files in the specified directory
 // while in the directory holding all the go files that needs to be modified, you can run "go fmt" without the -w or -l option



