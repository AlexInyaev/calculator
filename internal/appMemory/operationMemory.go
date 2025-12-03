package appMemory

import (
	"fmt"
	"github.com/AlexInyaev/calculator/internal/cli"
	"strings"
)

var mamoryOperations = make(map[string]*cli.CalculatorData, 10)

func S100aveOptration(opiration *cli.CalculatorData) {
	fmt.Println("Сохранить операцию \"Yes/No\"Y/N")

	confirmed := ""
	fmt.Scan(&confirmed)

	if strings.ToLower(confirmed) == "y" {
		fmt.Println(" Введите название операции")
		var nameOperation string
		fmt.Scan(&nameOperation)
		mamoryOperations[nameOperation] = opiration
	}

}
