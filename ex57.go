package main

import (
       "fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
     s := "cannot Sqrt negative number: %v"
     return fmt.Sprintf(s,float64(e))
}

func Sqrt(f float64) (z float64, err error) {
     z = 10
     if f < 0 {
     	return 0, ErrNegativeSqrt(f)
     }
     for i := 0; i < 10; i++ {
     	  z = z - (z * z - f) / (2 * z)
	   }
	    return z , nil
}

func main() {
     fmt.Println(Sqrt(2))
     fmt.Println(Sqrt(-2))
}