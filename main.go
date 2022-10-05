package main

import (
	"github.com/lundin/cloudnativemtrl/reverse"
	"fmt"
)

func main() {
	fmt.Println("start")

	in := []string{"Pontus", "Nils", "Kalle"}
	out := reverse.Reverse(in)
	fmt.Println(out)
}
