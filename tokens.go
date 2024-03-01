package partitools

import (
	"fmt"

	"github.com/alecthomas/participle/v2/lexer"
	// . "github.com/stevegt/goadapt"
)

// TokenDebugger is a helper struct with methods to debug lexer tokens.
type TokenDebugger struct {
	def            *lexer.StatefulDefinition
	inverseSymbols map[lexer.TokenType]string
}

// NewTokenDebugger creates a new TokenDebugger for the given lexer.
func NewTokenDebugger(def *lexer.StatefulDefinition) (td *TokenDebugger) {
	td = &TokenDebugger{def: def}

	// invert the token symbols for easy lookup
	symbols := td.def.Symbols()
	td.inverseSymbols = make(map[lexer.TokenType]string)
	for name, tokenType := range symbols {
		td.inverseSymbols[tokenType] = name
	}
	return
}

// ShowTokens is a helper function to show the tokens of a lexer.
func (td *TokenDebugger) ShowTokens(txt string) {
	def := td.def
	// Lexer tokenization for debugging
	lex, err := def.LexString("", txt)
	if err != nil {
		fmt.Printf("Error creating lexer: %v\n", err)
		return
	}
	for {
		token, err := lex.Next()
		if err != nil {
			fmt.Printf("Error tokenizing: %v\n", err)
			break
		}
		if token.EOF() {
			break
		}
		// Pprint(token)
		// fmt.Printf("Token: Type[%T] Value[%v]\n", token.Type, token.Value)
		// fmt.Printf("%s\n", token.GoString())
		fmt.Printf("%s\n", td.tokenString(&token))
	}
}

// tokenString returns a string representation of the given token.
func (d *TokenDebugger) tokenString(t *lexer.Token) string {
	sym := d.inverseSymbols[t.Type]
	if t.Pos == (lexer.Position{}) {
		return fmt.Sprintf("Token{%s, %q}", sym, t.Value)
	}
	return fmt.Sprintf("Token@%s{%s, %q}", t.Pos.String(), sym, t.Value)
}
