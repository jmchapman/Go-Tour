package main

import (
       "fmt"
)

func Sqrt(x float64) (z float64) {
     z = 10
     for i := 0; i < 10; i++ {
     	 z = z - (z * z - x) / (2 * z)
	 }
	 return
}

func main() {
     fmt.Println(Sqrt(2))
}