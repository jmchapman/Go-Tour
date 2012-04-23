package main

import (
       "tour/wc"
       "strings"
)

func WordCount(s string) (m map[string]int) {
     m = make(map[string]int)
     f := strings.Fields(s)
     for _, v := range f {
     	 m[v]++
	 }
	 return m
}

func main() {
     wc.Test(WordCount)
}