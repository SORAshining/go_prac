package main

import (
	"fmt"
)

const Pi = 3.14

func main() {
	var c []int
	//c = make([]int, 5) //length = 5
	c = make([]int, 0, 5) //length = 0 capasity = 5
	for i := 0; i < 5; i++ {
		c = append(c, i)
		fmt.Println(c)
	}
	fmt.Println(c)
}
