package main

import (
	"fmt"

	"github.com/arjunajithtp/travis-coveralls-integration/calculation"
)

func main() {
	a, b := 5, 2
	fmt.Printf("values: %v %v\n", a, b)
	fmt.Println(calculation.Sum(a, b))
	fmt.Println(calculation.Sub(a, b))
}
