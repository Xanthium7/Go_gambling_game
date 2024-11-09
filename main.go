package main

import (
	"fmt"
	"math/rand"
)

func getName() string {
	name := ""
	fmt.Println("ELDORADO CASINO WELCOMES YOU!! ")
	fmt.Print("GIB YOUR Name: ")
	_, err := fmt.Scanln(&name)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Print("Lets gets started, ", name)
	return name
}

func getBet(balance uint) uint {
	var bet uint

	for true {
		fmt.Printf("Enter your bet, or 0 to quit (balance = $%d): ", balance)
		_, err := fmt.Scan(&bet)
		if err != nil {
			fmt.Println("Error: ", err)
		}
		if bet > balance {
			fmt.Println("Bet cannot be larger than balance")
		} else {
			break
		}
	}

	return bet
}

func GenerateSymbolArray(symbols map[string]uint) []string {
	symbolArray := []string{}
	for symbol, count := range symbols {
		for i := uint(0); i < count; i++ {
			symbolArray = append(symbolArray, symbol)
		}
	}
	return symbolArray
}

func getRandomNumber(min, max int) int {
	return rand.Intn(max-min+1) + min
}

func getSpin(reel []string, rows int, cols int) [][]string {
	result := [][]string{}

	for i := 0; i < rows; i++ {
		result = append(result, []string{})
	}
	for i := 0; i < cols; i++ {
		selected := map[int]bool{}

		for j := 0; j < rows; j++ {
			for true {
				randomIndex := getRandomNumber(0, len(reel)-1)
				_, exisits := selected[randomIndex]
				if !exisits {
					selected[randomIndex] = true
					result[j] = append(result[j], reel[randomIndex])
					break
				}

			}
		}
	}
	return result
}

func PrintWin(spin [][]string) {
	for _, row := range spin {
		for j, symbol := range row {
			fmt.Printf(symbol)
			if j != len(row)-1 {
				fmt.Printf("\t|\t")
			}
		}
		fmt.Println()
	}
}

func main() {
	symbols := map[string]uint{"cherry": 4, "lemon": 7, "orange": 12, "plum": 13, "bell": 16, "bar": 20}

	// multipliers := map[string]uint{
	// 	"cherry": 20,
	// 	"lemon": 15,
	// 	"orange": 10,
	// 	"plum": 5,
	// 	"bell": 2,
	// 	"bar": 1,
	// }

	symbolArr := GenerateSymbolArray(symbols)

	balance := uint(200)

	for balance > 0 {
		bet := getBet(balance)
		if bet == 0 {
			break
		}
		balance -= bet
		spin := getSpin(symbolArr, 3, 3)
		PrintWin(spin)
	}

	fmt.Printf("you left with $%d.\n", balance)
	//fmt.Println("Bet: ", bet)
	//name := getName()

}
