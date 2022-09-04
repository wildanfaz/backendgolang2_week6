package main

import (
	"fmt"

	"github.com/wildanfaz/backendgolang2_week6/task1"
	"github.com/wildanfaz/backendgolang2_week6/task2"
	"github.com/wildanfaz/backendgolang2_week6/task3"
)

func main() {
	fmt.Println("\nTask 1")
	fmt.Printf("%g0\n", task1.Rounding(4.37))
	fmt.Printf("%g0\n", task1.Rounding(4.32))
	fmt.Printf("%g0\n", task1.Rounding(4.324))

	fmt.Println("\nTask 2")
	num := task2.Number{N: 40}
	fmt.Println("Prime =", num.Prime())
	fmt.Println("Odd =", num.Odd())
	fmt.Println("Even =", num.Even())
	fmt.Println("Fibonacci =", num.Fibonacci())

	fmt.Println("\nTask3")
	var count task3.AllCalculate = &task3.Cube{Side: 10}
	fmt.Println("Count 2D")
	fmt.Println("Area =", count.Area(), "cm2")
	fmt.Println("Perimeter =", count.Perimeter(), "cm")
	fmt.Println("Count 3D")
	fmt.Println("Volume =", count.Volume(), "cm3")
}
