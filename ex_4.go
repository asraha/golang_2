package main

import "fmt"

type amb int

var x amb

var y int

func main() {

	fmt.Println(x)

	fmt.Printf("%T%", x)

	x = 42

	fmt.Println(x)

	y = int(x)

	fmt.Println(y)

}
