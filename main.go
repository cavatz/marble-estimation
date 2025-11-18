package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {

	for {
		const g float64 = 9.81
		const P float64 = 1000.0
		const METERS_CONVERSION float64 = 0.0254
		const MARBLE_WEIGHT int = 8
		const p int = 1240

		scanner := bufio.NewScanner(os.Stdin)

		fmt.Printf("volume of filled in boat or x to stop: ")
		scanner.Scan()

		Vtext := scanner.Text()
		if Vtext == "x" {
			break
		}

		V, err := strconv.ParseFloat(Vtext, 64)
		handleError(err)

		V *= math.Pow(METERS_CONVERSION, 3)

		grams := (V * P) * 1000

		fmt.Printf("volume of hollowed boat: ")
		scanner.Scan()

		v, err := strconv.ParseFloat(scanner.Text(), 64)
		handleError(err)

		v *= math.Pow(METERS_CONVERSION, 3)

		m := (float64(p) * v) * 1000 //weight of b oat print it

		marbles_held := math.Round((grams - m) / float64(MARBLE_WEIGHT))
		weight := roundFloat(m, 2)

		fmt.Printf("%v marbles held before sinking. \n", marbles_held)
		fmt.Printf("Boat weighs %v grams. \n", weight)
		fmt.Printf("Total mass: %v grams. \n", math.Round(grams))
	}
	//F = (P * V) * g

}

func handleError(err error) {
	if err != nil {
		fmt.Printf("Error parsing volume of the boat.")
		os.Exit(1)
	}
}

func roundFloat(val float64, precision uint) float64 {
	r := math.Pow(10, float64(precision))
	return math.Round(val*r) / r
}
