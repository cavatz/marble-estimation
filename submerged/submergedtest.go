package submergedtest

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func TestSubmerged() {

	for {
		const g float64 = 9.81
		const P float64 = 1000.0
		const METERS_CONVERSION float64 = 0.0254
		const MARBLE_WEIGHT int = 4
		const p int = 1240

		scanner := bufio.NewScanner(os.Stdin)

		fmt.Printf("volume of filled in boat or x to stop: ")
		scanner.Scan()

		Vtext := scanner.Text()
		if Vtext == "x" {
			break
		}

		V_total, err := strconv.ParseFloat(Vtext, 64)
		handleError(err)

		V_total *= math.Pow(METERS_CONVERSION, 3)

		fmt.Printf("volume of hollowed boat: ")
		scanner.Scan()

		v_hollow, err := strconv.ParseFloat(scanner.Text(), 64)
		handleError(err)

		v_hollow *= math.Pow(METERS_CONVERSION, 3)

		mKg := (float64(p) * v_hollow)

		density := mKg / v_hollow

		v_sumberged := (density / float64(p)) * V_total

		grams := (v_sumberged * P) * 1000

		mG := mKg * 1000

		marbles_held := (grams - mG) / float64(MARBLE_WEIGHT)
		fmt.Printf("%v marbles held before sinking. \n", marbles_held)
	}
	//F = (P * V) * g

}

func handleError(err error) {
	if err != nil {
		fmt.Printf("Error parsing volume of the boat.")
		os.Exit(1)
	}
}
