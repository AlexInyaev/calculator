package cli

import (
	"fmt"
	"strconv"
)

type CalculatorInput struct {
	Num1      float64
	Num2      float64
	Operation string
	Result    float64
}

func (cl *CalculatorInput) GetFirstNumber() error {
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
