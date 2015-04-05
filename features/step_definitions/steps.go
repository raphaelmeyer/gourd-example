package main

import (
	"github.com/raphaelmeyer/gourd"
	"github.com/raphaelmeyer/gourd-example/calc"
)

type calc_context struct {
	calc calc.Calculator
}

func new_scenario() interface{} {
	return &calc_context{}
}

func main() {
	cucumber := gourd.NewCucumber(new_scenario)

	cucumber.Given("^I have entered (\\d+) into the calculator$").Do(
		func(context interface{}, arguments gourd.Arguments) {
			scenario, _ := context.(*calc_context)
			input := arguments.Int(0)

			scenario.calc.Push(input)
		})

	cucumber.When("^I press add$").Do(
		func(context interface{}, arguments gourd.Arguments) {
			scenario, _ := context.(*calc_context)

			scenario.calc.Add()
		})

	cucumber.Then("^the result should be (\\d+) on the screen$").Do(
		func(context interface{}, arguments gourd.Arguments) {
			scenario, _ := context.(*calc_context)
			result := arguments.Int(0)
			if scenario.calc.Result() != result {
				panic("Wrong result")
			}
		})

	cucumber.Run()
}
