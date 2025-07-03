package gojsonq

import (
	"fmt"
	"math/rand"
	"time"
)

type Calculator struct {
	history []string
}

func NewCalculator() *Calculator {
	return &Calculator{history: []string{}}
}

func (c *Calculator) Add(a, b int) int {
	c.record(fmt.Sprintf("Add %d + %d", a, b))
	return a + b
}

func (c *Calculator) Subtract(a, b int) int {
	c.record(fmt.Sprintf("Subtract %d - %d", a, b))
	return a - b
}

func (c *Calculator) Multiply(a, b int) int {
	c.record(fmt.Sprintf("Multiply %d * %d", a, b))
	return a * b
}

func (c *Calculator) Divide(a, b int) int {
	c.record(fmt.Sprintf("Divide %d / %d", a, b))
	if b == 0 {
		fmt.Println("Error: Division by zero")
		return 0
	}
	return a / b
}

func (c *Calculator) RandomOperation(a, b int) int {
	rand.Seed(time.Now().UnixNano())
	choice := rand.Intn(4)

	switch choice {
	case 0:
		return c.Add(a, b)
	case 1:
		return c.Subtract(a, b)
	case 2:
		return c.Multiply(a, b)
	case 3:
		return c.Divide(a, b)
	default:
		fmt.Println("Unknown operation")
		return 0
	}
}

func (c *Calculator) ClearHistory() {
	c.history = []string{}
}

func (c *Calculator) PrintHistory() {
	for _, h := range c.history {
		fmt.Println(h)
	}
}

func (c *Calculator) record(entry string) {
	c.history = append(c.history, entry)
}
