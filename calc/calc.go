package calc

type Calculator struct {
	stack []int
}

func (c *Calculator) Push(number int) {
	c.stack = append(c.stack, number)
}

func (c *Calculator) Add() {
	last := len(c.stack) - 1
	c.stack = append(c.stack[:last-1], c.stack[last-1]+c.stack[last])
}

func (c *Calculator) Result() int {
	return c.stack[len(c.stack)-1]
}
