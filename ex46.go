package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
     x := 0
     y := 1
     var z int
     return func() int {
     	    z = x + y
	      x = y
	      	y = z
		  return z
		  }
}

func main() {
     f := fibonacci()
     for i := 0; i < 10; i++ {
     	 fmt.Println(f())
	 }
}