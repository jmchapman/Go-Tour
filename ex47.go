package main

import "fmt"

func Cbrt(x complex128) (z complex128) {
     z = 10
     for i := 0; i < 10; i++ {
     	 z = z - (z*z*z - x) / (3*z*z)
	 }
	 return z
}

func main() {
     fmt.Println(Cbrt(2))
}