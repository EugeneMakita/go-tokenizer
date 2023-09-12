package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	paser := CreateParser()
	parsedValue := paser.Program()
	fmt.Println(parsedValue)
}

type Parser struct {
	Id int
}

func CreateParser() *Parser {
	return &Parser{Id: 1}
}

func (p *Parser) Parse(s string) {
	fmt.Println("hello i am Add more")
}

func (p *Parser) Program() string {
	program := map[string]interface{}{
		"type": "Program",
		"body": p.NumericLiteral(45),
	}

	programJson, err := json.Marshal(program)
	if err != nil {
		fmt.Println(err.Error())
	}

	return string(programJson)
}

func (p *Parser) NumericLiteral(s int32) map[string]interface{} {
	return map[string]interface{}{
		"type":  "NumericLiteral",
		"value": s,
	}
}
