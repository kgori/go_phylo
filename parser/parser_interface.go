package parser

import (
	"github.com/kgori/go_phylo/lexer"
	"github.com/kgori/go_phylo/node"
)

type parser struct {
	lex   lexer.Lexer
	stack node.Stack
	state parseFn
	next  lexer.Item
	peek  lexer.Item
}

type parseFn func(Parser) parseFn

type Parser interface {
	Parse() bool
	Next() lexer.Item
	Peek() lexer.Item
	advance()
}

func New(input string) Parser {
	l := lexer.New(input)
	s := node.NewStack()
	// q := node.NewQueue(1)
	p := &parser{
		lex:   l,
		stack: s,
		// state: parse,
	}

	p.next = l.NextItem()
	p.peek = l.NextItem()

	return p
}
