package main

// Declare a main function, this is the entrypoint into our go module
// That will be run. In our hello world, we won't need this
func main() {}


// This exports an add function.
// It takes in two 32-bit integer values
// And returns a 32-bit integer value.
// To make this function callable from JavaScript, 
// we need to add the: "go:export add" comment above the function
//go:export add
func add(x int, y int) int {
    return x + y;
}

