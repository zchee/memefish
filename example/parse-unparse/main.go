package main

import (
	"fmt"
	"log"

	"github.com/MakeNowJust/memefish/pkg/parser"
	"github.com/MakeNowJust/memefish/pkg/token"
	"github.com/k0kubun/pp"
)

func main() {
	// Create a new Parser instance.
	file := &token.File{
		Buffer: `
CREATE TABLE Tests (
  Id STRING(20) NOT NULL,
  Generated STRING(100) AS (CONCAT(Id, "XXX")) STORED OPTIONS (
    allow_commit_timestamp = true
  ),
) PRIMARY KEY(Id)`,
	}
	p := &parser.Parser{
		Lexer: &parser.Lexer{File: file},
	}

	// Do parsing!
	stmt, err := p.ParseDDL()
	if err != nil {
		log.Fatal(err)
	}

	// Show AST.
	log.Print("AST")
	_, _ = pp.Println(stmt)

	// Unparse AST to SQL source string.
	log.Print("Unparse")
	fmt.Println(stmt.SQL())
}
