// Declare a main package
package main
// Import fmt package which contains functions for formatting text, including printing to the console.
// This is one of the "standard package"
import "fmt"
import "rsc.io/quote"
// This is go main function to print the message to the console
func main(){
	fmt.Println(quote.Go())
}
// type: go run hello.go in the terminal to run the hello.go file