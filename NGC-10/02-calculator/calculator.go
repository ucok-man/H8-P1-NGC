package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type operation string

const (
	add operation = "add"
	sub operation = "sub"
	mul operation = "mul"
	div operation = "div"
)

type Calculator struct {
	operation operation
	val_a     float64
	val_b     float64
}

func (c Calculator) Welcome() string {
	var buff bytes.Buffer
	buff.WriteString("pilih operasi aritmatika: ")
	buff.WriteString("add | sub | mul | div")
	return buff.String()
}

func (c *Calculator) promptOperation(reader io.Reader) operation {
	for {
		fmt.Print("operation ('q' for): ")

		input := c.prompt(reader)
		if strings.ToLower(input) == "q" {
			os.Exit(1)
		}

		err := c.validateOP(input)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			continue
		}

		if nil == err {
			return operation(input)
		}
	}
}

func (c *Calculator) promptNumber(reader io.Reader, question string) float64 {
	for {
		fmt.Print(question)
		input_1 := c.prompt(reader)
		num, err := c.validateNum(input_1)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			continue
		}

		if nil == err {
			return num
		}
	}
}

func (c Calculator) prompt(reader io.Reader) string {
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return scanner.Text()
}

func (c Calculator) validateOP(input string) error {
	validop := []string{"add", "sub", "mul", "div"}
	var valid bool
	for _, vop := range validop {
		if vop == strings.ToLower(input) {
			valid = true
		}
	}
	if !valid {
		return errValidation("make sure input is in range [add | sub | mul | div]")
	}
	return nil
}

func (c Calculator) validateNum(input string) (float64, error) {
	num, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, errValidation(fmt.Sprintf("(%v) make sure you enter a number", input))
	}
	return num, nil
}

func (c Calculator) calculate(op operation, num1, num2 float64) string {
	switch op {
	case add:
		return fmt.Sprintf("%.2f\n", num1+num2)
	case sub:
		return fmt.Sprintf("%.2f\n", num1-num2)
	case mul:
		return fmt.Sprintf("%.2f\n", num1*num2)
	case div:
		return fmt.Sprintf("%.2f\n", num1/num2)
	default:
		panic(fmt.Sprintf("unhandled posible case: %v", op))
	}
}
