package main

import (
	"fmt"
)





func sum(srcs ...int) int {
	//可変長引数はようそがスライスで返される
	num := 0
	for _, src := range srcs {
		num += src
	}
	return num
}

