package main

import "fmt"

// START OMIT
func main() {
	b := make([]int, 1024)
	fmt.Println(b)
	b = append(b, 99)
	fmt.Println("len:", len(b), "cap:", cap(b))
}

// END OMIT
