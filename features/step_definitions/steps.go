package main

import (
	"github.com/raphaelmeyer/gourd"
	"github.com/raphaelmeyer/gourd-example/calc"
)

type scenario struct {
	calc calc.Calculator
}

func new_scenario() interface{} {
	return &scenario{}
}

func main() {
	cucumber := gourd.NewCucumber(new_scenario)

	cucumber.Given("^I have entered (\\d+) into the calculator$").Pass()

	cucumber.When("^I press add$").Pass()

	cucumber.Then("^the result should be (\\d+) on the screen$").Pass()

	cucumber.Run()
}
