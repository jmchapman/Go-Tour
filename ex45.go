package main

import "tour/pic"

func Pic(dx, dy int) (p [][]uint8) {
     p = make([][]uint8,dy)
     for y := 0; y < dy; y++ {
     	 p[y] = make([]uint8,dx)
	      for x := 0; x < dx; x++ {
	      	    p[y][x] = uint8(x ^ y)
		    	    }
			    }
			    return
}

func main() {
     pic.Show(Pic)
}