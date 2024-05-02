package compile

import (
	"encoding/json"
	"fmt"
	"os"

	. "github.com/Ygg-Drasill/Sleipnir/compiler/analysis/lexer"
	ost "github.com/Ygg-Drasill/Sleipnir/compiler/ast"
	"github.com/Ygg-Drasill/Sleipnir/compiler/gocc/parser"
	"github.com/Ygg-Drasill/Sleipnir/compiler/synthesis"
)

func Compile(filePath string) error {
	lexer := NewLexerFromString(filePath)
	tokens := lexer.FindTokens()
	scanner := NewScanner(tokens)
	p := parser.NewParser()
	p.Context = ost.NewParseContext()
	var ast interface{}
	var err error
	if ast, err = p.Parse(scanner); err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println(ast)
	bytes, _ := json.MarshalIndent(ast, "", "\t")
	file, _ := os.OpenFile("AST_out.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	file.Write(bytes)
	file.Close()
	if prog, ok := ast.(ost.Program); ok {
		codeGen := synthesis.GenWrapper(&prog)
		fmt.Printf("%s", codeGen.String())
		files, err := os.OpenFile("codeGen.wat", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
		if err != nil {
			return err
		}
		files.Write(codeGen.Bytes())
		files.Close()
	}
	return nil
}
