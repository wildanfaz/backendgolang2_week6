package main

import (
	"fmt"

	"github.com/wildanfaz/backendgolang2_week6/taskbasic1"
	"github.com/wildanfaz/backendgolang2_week6/taskbasic2"
	"github.com/wildanfaz/backendgolang2_week6/taskbasic3"
)

func main() {
	fmt.Println("\nTask 1")
	fmt.Printf("%g0\n", taskbasic1.Rounding(4.37))
	fmt.Printf("%g0\n", taskbasic1.Rounding(4.32))
	fmt.Printf("%g0\n", taskbasic1.Rounding(4.324))

	fmt.Println("\nTask 2")
	num := taskbasic2.Number{N: 40}
	fmt.Println("Prime =", num.Prime())
	fmt.Println("Odd =", num.Odd())
	fmt.Println("Even =", num.Even())
	fmt.Println("Fibonacci =", num.Fibonacci())

	fmt.Println("\nTask3")
	var cal taskbasic3.AllCalculate = &taskbasic3.Cube{Side: 10}
	fmt.Println("Count 2D")
	fmt.Println("Area =", cal.Area(), "cm2")
	fmt.Println("Perimeter =", cal.Perimeter(), "cm")
	fmt.Println("Count 3D")
	fmt.Println("Volume =", cal.Volume(), "cm3")
}
