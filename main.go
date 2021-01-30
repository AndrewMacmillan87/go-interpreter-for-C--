package main

import (
	"fmt"

	"github.com/sedexdev/interpreter/evaluator"
	"github.com/sedexdev/interpreter/lexer"
	"github.com/sedexdev/interpreter/parser"
	"github.com/sedexdev/interpreter/symbol"
	"github.com/sedexdev/interpreter/testcode"
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
