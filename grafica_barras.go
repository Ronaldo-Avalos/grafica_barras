package main

import "fmt"

func main() {

	var num = []int{3,0,2,6}
	var symbol = "*"
	numax := findMax(num)
	grafica(numax, num, symbol)

}

func findMax(num []int) (numax int) {
	numax = num[0]
	for _, value := range num {
		if value > numax {
			numax = value
		}
	}
	return numax
}

func grafica(numax int, num []int, symbol string) {
	for i := numax; i >= 0; i-- {
		if i != numax {
			fmt.Print(" |")

		}
		//range devuelve el índice del elemento como un número entero.
		for _, values := range num {
			if values > i {
				fmt.Print(symbol)
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
	fmt.Print("  ")
	for range num {
		fmt.Print("¯")

	}
}
