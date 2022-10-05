package main

import (
	"github.com/lundin/reverser/reverse"
	"fmt"
)

func main() {
	fmt.Println("start")

	in := []string{"Pontus", "Nils", "Kalle"}
	out := reverse.Reverse(in)
	fmt.Println(out)
}
