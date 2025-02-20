package main

import (
	"fmt"
	"strconv"
)

type Stack struct {
	items []string
}

// Pushes a new element on the Stacl
func (stack *Stack) Push(item string) {
	stack.items = append(stack.items, item)
}

// Pops the first element from the Stack
func (s *Stack) Pop() (string, error) {
	if len(s.items) == 0 {
		return "", fmt.Errorf("Stack ist leer")
	}
	lastIndex := len(s.items) - 1
	item := s.items[lastIndex]
	s.items = s.items[:lastIndex]
	return item, nil
}

func getTopTwoValuesAsFloat(stack *Stack) (string, string, float64, float64) {
	stack.Pop()
	first, _ := stack.Pop()
	second, _ := stack.Pop()

	firstFloat, _ := strconv.ParseFloat(first, 64)
	secondFloat, _ := strconv.ParseFloat(second, 64)

	return first, second, firstFloat, secondFloat
}

func getTopValueAsFloat(stack *Stack) float64{
	stack.Pop()
	value, _ := stack.Pop()
	valueFloat, _ := strconv.ParseFloat(value, 64)
	return valueFloat
}