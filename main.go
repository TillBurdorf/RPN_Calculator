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

	//gameloop
	for {
		var input string
		_, err := fmt.Scan(&input)
		if err != nil {
			fmt.Println("Eingabefehler")
		}
		if input == "exit" {
			return
		}

		stack.Push(input)

		switch input {
		case "+":
			firstFloat, secondFloat := getTopTwoValuesAsFloat(stack)

			result := secondFloat + firstFloat

			fmt.Println("Es wird addiert: ", secondFloat, "+", firstFloat, "=", result)
			stack.Push(strconv.FormatFloat(result, 'f', 2, 64))

		case "-":
			firstFloat, secondFloat := getTopTwoValuesAsFloat(stack)

			result := secondFloat - firstFloat

			fmt.Println("Es wird subtrahiert: ", secondFloat, "-", firstFloat, "=", result)
			stack.Push(strconv.FormatFloat(result, 'f', 2, 64))

		case "*":
			firstFloat, secondFloat := getTopTwoValuesAsFloat(stack)

			result := firstFloat * secondFloat

			fmt.Println("Es wird multipliziert: ", secondFloat, "*", firstFloat, "=", result)
			stack.Push(strconv.FormatFloat(result, 'f', 2, 64))

		case "/":

			firstFloat, secondFloat := getTopTwoValuesAsFloat(stack)

			result := secondFloat / firstFloat

			fmt.Println("Es wird dividiert: ", secondFloat, "/", firstFloat, "=", result)
			stack.Push(strconv.FormatFloat(result, 'f', 2, 64))

		case "^":
			firstFloat, secondFloat := getTopTwoValuesAsFloat(stack)

			result := math.Pow(secondFloat, firstFloat)

			fmt.Println("Es wird potenziert: ", secondFloat, "^", firstFloat, "=", result)
			stack.Push(strconv.FormatFloat(result, 'f', 2, 64))

		case "sqrt":
			value := getTopValueAsFloat(stack)
			result := math.Sqrt(value)
			fmt.Println("Es wird die Quadratwurzel von", value, "gezogen: ", result)
			stack.Push(strconv.FormatFloat(result, 'f', 2, 64))

		case "log":
			value := getTopValueAsFloat(stack)
			result := math.Log10(value)
			fmt.Println("Es wird der Logarithmus von", value, "zur Basis 10 berechnet: ", result)
			stack.Push(strconv.FormatFloat(result, 'f', 2, 64))

		case "!":

		case "abs":
			value := getTopValueAsFloat(stack)
			result := math.Abs(value)
			fmt.Println("Betrag von", value, "=", result)
			stack.Push(strconv.FormatFloat(result, 'f', 2, 64))

		case "++":
		case "**":
		}
	}
}
