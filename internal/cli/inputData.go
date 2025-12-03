package cli

import (
	"fmt"
	"strconv"
)

type CalculatorData struct {
	Num1      float64
	Num2      float64
	Operation string
	Result    float64
}

func NewCalculatorData() *CalculatorData {
	cld := &CalculatorData{
		Num1:      0.0,
		Num2:      0.0,
		Operation: "",
		Result:    0.0,
	}
	return cld
}
func (cl *CalculatorData) GetFirstNumber() error {
	fmt.Println("Введите первое число")
	str := ""
	fmt.Scan(&str)
	num, err := strconv.ParseFloat(str, 64)
	if err != nil {
		//fmt.Println("Ошибка:", err)
		return err
	}
	cl.Num1 = num
	return nil
}

func (cl *CalculatorData) GetSecondNumber() error {
	fmt.Println("Введите второе число")
	str := ""
	fmt.Scan(&str)
	num, err := strconv.ParseFloat(str, 64)
	if err != nil {
		//fmt.Println("Ошибка:", err)
		return err
	}
	cl.Num2 = num
	return nil
}

func (cl *CalculatorData) GetOperator() {
	fmt.Println("Введите знак \"+,-,/,* \"  операции которую планируете  выполнить")
	Operation := ""
	fmt.Scan(&Operation)

	cl.Operation = Operation
}
func (cl *CalculatorData) DoOperation() error {
	switch cl.Operation {
	case "+":
		cl.Result = cl.Num1 + cl.Num2
		return nil
	case "-":
		cl.Result = cl.Num1 - cl.Num2
		return nil
	case "*":
		cl.Result = cl.Num1 * cl.Num2
		return nil
	case "/":
		if cl.Num2 == 0 {
			return fmt.Errorf("деление на ноль")
		}
		cl.Result = cl.Num1 / cl.Num2
		return nil
	default:
		return fmt.Errorf("неподдерживаемая операция: %s", cl.Operation)
	}
}
func (cl *CalculatorData) PrintResult() {
	fmt.Println("Результат вычисления равен = ", cl.Result)
}
