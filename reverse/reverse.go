package reverse

import (
	_ "fmt"


	)  

func Reverse(in []string) (out []string) {
	for i := len(in) - 1; i >= 0; i-- {
		out = append(out, in[i])
	}

	return out

}
