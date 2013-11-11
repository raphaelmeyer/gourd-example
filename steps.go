package main

import (
	"github.com/raphaelmeyer/gourd-example/calc"
	"github.com/raphaelmeyer/gourd"
)

type Context struct {
	calc calc.Calculator
}

func main() {
	cuke := new(gourd.Cucumber)

	cuke.Given("Regex").Do(
		func(context * gourd.Context) {
			cuke.Expect(true)
		})

	cuke.Start();
}

