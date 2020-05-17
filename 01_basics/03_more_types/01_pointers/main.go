package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set ii throught the pointer
	fmt.Println(i)  // see the new value of i

	p = &j
	*p = *p / 37   // devide j throuth th pointer
	fmt.Println(j) // see the new value of j
}
