package parser

import "github.com/kgori/go_phylo/lexer"

func (p *parser) Next() lexer.Item {
	defer p.advance()
	return p.next
}

func (p *parser) Peek() lexer.Item {
	return p.peek
}

func (p *parser) Parse() bool {
	return false
}

func (p *parser) advance() {
	if p.peek.IsEOF() {
		p.next = p.peek
	} else {
		p.next, p.peek = p.peek, p.lex.NextItem()
	}
}
