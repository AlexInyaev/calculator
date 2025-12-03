package main

import (
	"fmt"
	"github.com/AlexInyaev/calculator/internal/appMemory"
	"github.com/AlexInyaev/calculator/internal/cli"
	"strings"
)

func main() {
	flowOperations()
	ContinueOperation()
}

func flowOperations() {
	CalcOne := cli.NewCalculatorData()
	CalcOne.GetFirstNumber()
	CalcOne.GetSecondNumber()
	CalcOne.GetOperator()
	CalcOne.DoOperation()
	CalcOne.PrintResult()
	appMemory.SaveOptration(CalcOne)
}
func ContinueOperation() {
	for {
		fmt.Println("Выполнить еще одну операцию \"Yes/No\"Y/N")

		continueOperation := ""
		fmt.Scan(&continueOperation)

		if strings.ToLower(continueOperation) == "y" {
			flowOperations()

		} else {
			break
		}

	}

}
