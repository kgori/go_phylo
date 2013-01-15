package lexer

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

// Function defs

func (l *lexer) Next() (r rune) {
	if l.pos >= len(l.input) {
		l.width = 0
		return eof
	}
	r, l.width = utf8.DecodeRuneInString(l.input[l.pos:])
	l.pos += l.width
	return r
}

func (l *lexer) Pos() int {
	return l.pos
}

func (l *lexer) Peek() rune {
	r := l.Next()
	l.Backup()
	return r
}

func (l *lexer) Backup() {
	l.pos -= l.width
}

func (l *lexer) Emit(t itemType) {
	fmt.Printf("Emitted %s, %s\n", t, l.input[l.start:l.pos])
	l.items <- Item{t, l.input[l.start:l.pos]}
	l.start = l.pos
}

func (l *lexer) NextItem() Item {
	for {
		select {
		case Item := <-l.items:
			return Item
		default:
			l.state = l.state(l)
		}
	}
	panic("not reached")
}

func (l *lexer) Ignore() {
	l.start = l.pos
}

func (l *lexer) Accept(valid string) (bool, int) {
	i := strings.IndexRune(valid, l.Next())
	if i >= 0 {
		return true, i
	}
	l.Backup()
	return false, -1
}

func (l *lexer) AcceptRun(valid string) {
	for strings.IndexRune(valid, l.Next()) >= 0 {
	}
	l.Backup()
}
func (l *lexer) Line() int {
	return 1 + strings.Count(l.input[:l.pos], "\n")
}

func (l *lexer) NewLine() {
	l.line += 1
	l.column = 0
}

func (l *lexer) Errorf(format string, args ...interface{}) stateFn {
	l.items <- Item{ItemError, fmt.Sprintf(format, args...)}
	return nil
}

func (l *lexer) Column() int { return 0 }

func (l *lexer) MatchSingleChar() (bool, itemType) {
	chars := "()[]:;,'\"="
	_, i := l.Accept(chars)
	if i >= 0 {
		t := singleCharItems[i]
		return true, t
	}
	return false, ItemError
}

func (l *lexer) MatchNumber() bool {
	l.Accept("+-")
	digits := "0123456789"
	l.AcceptRun(digits)
	b, _ := l.Accept(".")
	if b {
		l.AcceptRun(digits)
	}
	b, _ = l.Accept("eE")
	if b {
		l.Accept("+-")
		l.AcceptRun(digits)
	}
	return true
}

func (l *lexer) MatchLabel() bool {
	for {
		r := l.Next()
		if !(strings.ContainsRune("'_-.?!0123456789", r) || unicode.IsLetter(r)) {
			l.Backup()
			break
		}
	}
	return true
}

func (l *lexer) MatchNHX() bool {
	if strings.HasPrefix(l.input[l.start:], "&&NHX") {
		l.pos += len("&&NHX")
		l.Ignore()
		return true
	}
	return false
}

func (l *lexer) UnemittedText() bool {
	return l.pos > l.start
}
