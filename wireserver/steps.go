package main

import (
	"github.com/raphaelmeyer/gourd"
	"github.com/raphaelmeyer/gourd-example/calc"
)

type Context struct {
	calc calc.Calculator
}

func main() {
	cucumber := new(gourd.Cucumber)

	cucumber.Given("Regex").Do(
		func(context *gourd.Context) {
			cucumber.Expect(true)
		})

	cucumber.Start()
}
