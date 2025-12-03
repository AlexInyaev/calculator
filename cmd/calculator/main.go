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
	err := CalcOne.DoOperation()
	if err != nil {
		fmt.Println(err)
	}
	CalcOne.PrintResult()
	appMemory.SaveOptration(CalcOne)
}
func ContinueOperation() {
	for {
		fmt.Println("Выполнить еще одну операцию \"Yes/No\"Y/N")

		continueOperation := ""
		_, err := fmt.Scan(&continueOperation)
		if err != nil {
			fmt.Println(err)
			return
		}

		if strings.ToLower(continueOperation) == "y" {
			flowOperations()

		} else {
			break
		}

	}

}
