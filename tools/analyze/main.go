package main

import (
	"flag"
	"log"
	"os"

	"github.com/MakeNowJust/memefish/pkg/analyzer"
	"github.com/MakeNowJust/memefish/pkg/parser"
	"github.com/olekukonko/tablewriter"
)

func init() {
	flag.Parse()
}

func main() {
	if flag.NArg() < 1 {
		log.Fatal("usage: ./parse [SQL query]")
	}

	query := flag.Arg(0)

	log.Printf("query: %q", query)

	p := &parser.Parser{
		Lexer: &parser.Lexer{
			File: parser.NewFile("[query]", query),
		},
	}

	log.Printf("start parsing")
	stmt, err := p.ParseQuery()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("finish parsing successfully")

	log.Printf("start analyzing")
	a := &analyzer.Analyzer{
		File: p.File,
	}
	a.AnalyzeQueryStatement(stmt)
	log.Printf("finish analyzing")

	list := a.NameLists[stmt.Query]
	if list == nil {
		log.Fatal("missing name list")
	}

	table := tablewriter.NewWriter(os.Stdout)
	var header []string
	for _, c := range list.Columns {
		if c.Path.ImplicitAliasID > 0 {
			header = append(header, "(unspecified)")
		} else {
			header = append(header, c.Path.Name)
		}
	}
	table.SetHeader(header)

	var types []string
	for _, c := range list.Columns {
		if c.Type == nil {
			types = append(types, "(null)")
		} else {
			types = append(types, c.Type.String())
		}
	}
	table.Append(types)
	table.Render()
}
