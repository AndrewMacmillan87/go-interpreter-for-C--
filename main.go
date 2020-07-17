package main

import (
	"fmt"

	"github.com/AMacmi01/interpreter/evaluator"
	"github.com/AMacmi01/interpreter/lexer"
	"github.com/AMacmi01/interpreter/parser"
	"github.com/AMacmi01/interpreter/symbol"
	"github.com/AMacmi01/interpreter/testcode"
)

func main() {

	code := testcode.GetProgram()

	lexedCode := lexer.CreateLexer(code)

	program := parser.CreateParser(lexedCode)
	parsedProgram := program.ParseProgram()

	// Check length of program.GetErrors before continuing
	// to see if any syntax errors occurred
	errors := program.GetErrors()
	for _, err := range errors {
		fmt.Println(err)
	}

	if len(errors) == 0 {
		symbolTable := symbol.CreateSymbolTable()
		evaluated := evaluator.Evaluate(parsedProgram, symbolTable)
		fmt.Print(evaluated.GetValue())
	}
}
