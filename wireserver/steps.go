package main

import (
	"github.com/raphaelmeyer/gourd-example/calc"
	"github.com/raphaelmeyer/gourd"
)

type Context struct {
	calc calc.Calculator
}

func main() {
	cucumber := new(gourd.Cucumber)

	cucumber.Given("Regex").Do(
		func(context * gourd.Context) {
			cucumber.Expect(true)
		})

	cucumber.Start();
}

