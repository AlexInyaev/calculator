package appMemory

import (
	"fmt"
	"github.com/AlexInyaev/calculator/internal/cli"
	"strings"
)

var mamoryOperations = make(map[string]*cli.CalculatorData, 10)

func SaveOptration(opiration *cli.CalculatorData) {
	fmt.Println("Сохранить операцию \"Yes/No\"Y/N")

	confirmed := ""
	_, err := fmt.Scan(&confirmed)

	if err != nil {
		fmt.Println(err)
		return
	}

	if strings.ToLower(confirmed) == "y" {
		fmt.Println(" Введите название операции")
		var nameOperation string
		_, err := fmt.Scan(&nameOperation)
		if err != nil {
			fmt.Println(err)
			return
		}
		mamoryOperations[nameOperation] = opiration
	}

}
