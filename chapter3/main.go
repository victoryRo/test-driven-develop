package main

import (
	"flag"
	"github.com/victoryRo/test-driven-develop/calculator"
	"github.com/victoryRo/test-driven-develop/input"
	"log"
)

func main() {
	expr := flag.String("expression", "7 + 14", "mathematical expression to parser")
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
