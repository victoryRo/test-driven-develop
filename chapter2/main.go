package main

import (
	"flag"
	"github.com/victoryRo/test-driven-develop/calculator"
	"github.com/victoryRo/test-driven-develop/input"
	"log"
)

func main() {
	expr := flag.String("expression", "", "mathematical expression to parse")
	flag.Parse()

	engine := calculator.NewEngine()
	validator := input.NewValidator(engine.GetNumOperands(), engine.GetValidOperators())
	parser := input.NewParser(engine, validator)
	result, err := parser.ProcessExpression(*expr)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(*result)
}
