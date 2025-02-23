package main

import (
	"fmt"
	//"github.com/emirpasic/gods/stacks/linkedliststack"
	"math"
	"strconv"
)

func main() {
	fmt.Println("Eingabe: ")

	// init Stack for the calculator
	stack := Stack{}
	infix := ""

	//gameloop
	for {
		var input string
		_, err := fmt.Scan(&input)
		if err != nil {
			fmt.Println("Eingabefehler")
		}
		if input == "exit" {
			fmt.Println("Infix Notation: ", infix)
			return
		}

		stack.Push(input)

		switch input {

		case "+":
			first, second, firstFloat, secondFloat := getTopTwoValuesAsFloat(&stack)

			result := secondFloat + firstFloat

			fmt.Println("Es wird addiert: ", secondFloat, "+", firstFloat, "=", result)
			stack.Push(strconv.FormatFloat(result, 'f', 2, 64))

			// update infix notation
			infix = updateInfix(infix, first, second, "+")

		case "-":
			first, second, firstFloat, secondFloat := getTopTwoValuesAsFloat(&stack)

			result := secondFloat - firstFloat

			fmt.Println("Es wird subtrahiert: ", secondFloat, "-", firstFloat, "=", result)
			stack.Push(strconv.FormatFloat(result, 'f', 2, 64))

			// update infix notation
			infix = updateInfix(infix, first, second, "-")

		case "*":
			first, second, firstFloat, secondFloat := getTopTwoValuesAsFloat(&stack)

			result := firstFloat * secondFloat

			fmt.Println("Es wird multipliziert: ", secondFloat, "*", firstFloat, "=", result)
			stack.Push(strconv.FormatFloat(result, 'f', 2, 64))

			// update infix notation
			infix = updateInfix(infix, first, second, "*")

		case "/":
			first, second, firstFloat, secondFloat := getTopTwoValuesAsFloat(&stack)

			result := secondFloat / firstFloat

			fmt.Println("Es wird dividiert: ", secondFloat, "/", firstFloat, "=", result)
			stack.Push(strconv.FormatFloat(result, 'f', 2, 64))

			// update infix notation
			infix = updateInfix(infix, first, second, "/")

		case "^":
			_, _, firstFloat, secondFloat := getTopTwoValuesAsFloat(&stack)

			result := math.Pow(secondFloat, firstFloat)

			fmt.Println("Es wird potenziert: ", secondFloat, "^", firstFloat, "=", result)
			stack.Push(strconv.FormatFloat(result, 'f', 2, 64))

		case "sqrt":
			value := getTopValueAsFloat(&stack)
			result := math.Sqrt(value)
			fmt.Println("Es wird die Quadratwurzel von", value, "gezogen: ", result)
			stack.Push(strconv.FormatFloat(result, 'f', 2, 64))

		case "log":
			value := getTopValueAsFloat(&stack)
			result := math.Log10(value)
			fmt.Println("Es wird der Logarithmus von", value, "zur Basis 10 berechnet: ", result)
			stack.Push(strconv.FormatFloat(result, 'f', 2, 64))

		case "!":

		case "abs":
			value := getTopValueAsFloat(&stack)
			result := math.Abs(value)
			fmt.Println("Betrag von", value, "=", result)
			stack.Push(strconv.FormatFloat(result, 'f', 2, 64))

			infix = "abs( " + infix + " )"

		case "**":
			result := multiplyAll(&stack)
			fmt.Println("Multiplikation aller Werte auf dem Stack = ", result)
			stack.Push(strconv.FormatFloat(result, 'f', 2, 64))

		case "++":
			result := addAll(&stack)
			fmt.Println("Addition aller Werte auf dem Stack = ", result)
			stack.Push(strconv.FormatFloat(result, 'f', 2, 64))
		}
	}
}

// updates the infix notation
func updateInfix(infix string, first string, second string, operator string) string {
	if infix != "" {
		infix = infix + " " + operator + " " + first + " )"
	}
	if infix == "" {
		infix = "( " + second + " " + operator + " " + first + " )"
	}
	return infix
}

func addAll(stack *Stack) float64 {
	var result float64 = 0
	//remove the operator from the stack
	stack.Pop()
	for len(stack.items) > 0 {
		first, _ := stack.Pop()
		firstFloat, _ := strconv.ParseFloat(first, 64)
		result = result + firstFloat
	}
	return result
}

func multiplyAll(stack *Stack) float64 {
	var result float64 = 1
	//remove the operator from the stack
	stack.Pop()
	for len(stack.items) > 0 {
		first, _ := stack.Pop()
		firstFloat, _ := strconv.ParseFloat(first, 64)
		result = result * firstFloat
	}
	return result
}

// updates the latex notation
func updateLatex(latex string, first string, second string, operator string) string {

	return latex
}
