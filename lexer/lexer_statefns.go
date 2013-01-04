package lexer

import (
	"unicode"
)

func isAlphaNumeric(r rune) bool {
	return r == '_' || unicode.IsLetter(r) || unicode.IsDigit(r)
}

func isAlpha(r rune) bool {
	return r == '_' || unicode.IsLetter(r)
}

func isNumeric(r rune) bool {
	return r == '_' || unicode.IsDigit(r)
}

// isSpace reports whether r is a space character.
func isSpace(r rune) bool {
	switch r {
	case ' ', '\t', '\n', '\r':
		return true
	}
	return false
}

// func LexInsideTree() stateFn         { return LexInsideTree }
// func LexLabel() stateFn        { return nil }
// func LexQuotedString() stateFn { return nil }
// func LexNumber() stateFn       { return nil }
// func LexSingleQuote() stateFn  { return nil }
// func LexDoubleQuote() stateFn  { return nil }
// func LexComment() stateFn      { return nil }

func LexNHX(l Lexer) stateFn {
	b, itemType := l.MatchSingleChar()
	if b {
		switch itemType {
		case ItemColon:
			l.Ignore()
			return LexKey
		case ItemEquals:
			l.Emit(ItemEquals)
			return LexValue
		case ItemRSquare:
			return LexInsideTree
		default:
			return l.Errorf("unrecognized character in NHX block: %#U", itemType)
		}
	}
	return nil
}

func LexKey(l Lexer) stateFn {
Loop:
	for {
		switch l.Next() {
		case ':':
			l.Ignore()
		case '=':
			l.Backup()
			break Loop
		case ']':
			return l.Errorf("No value found for NHX key")
		case eof:
			return l.Errorf("unterminated NHX block found while searching for key")
		case '[':
			return l.Errorf("found '[' in NHX block")
		}
	}
	l.Emit(ItemNHXKey)
	return LexNHX
}

func LexValue(l Lexer) stateFn {
	for {
		switch l.Next() {
		case ':':
			l.Backup()
			l.Emit(ItemNHXValue)
			return LexKey
		case ']':
			l.Backup()
			l.Emit(ItemNHXValue)
			return LexInsideTree
		case eof:
			return l.Errorf("unterminated NHX block found while searching for value")
		}
	}
	return nil
}

func LexInsideTree(l Lexer) stateFn {
	if isSpace(l.Next()) {
		l.Ignore()
	} else {
		l.Backup()
	}
	b, itemType := l.MatchSingleChar()
	if b {
		l.Emit(itemType)
		switch itemType {
		case ItemLBracket:
			return LexInsideTree
		case ItemRBracket:
			b, _ := l.Accept("0123456789.-+")
			if b {
				l.Backup()
				return LexNumber
			} else {
				return LexInsideTree
			}
		case ItemLSquare:
			return LexComment
		case ItemRSquare:
			return LexInsideTree
		case ItemColon:
			return LexNumber
		case ItemSemiColon:
			return LexOutsideTree
		case ItemComma:
			return LexInsideTree
		case ItemSingleQuote:
			return LexQuotedString
		case ItemDoubleQuote:
			return LexQuotedString
		case ItemEquals:
			return LexInsideTree
		default:
			return l.Errorf("unrecognized character in tree: %#U", itemType)
		}
	}
	if isAlpha(l.Peek()) {
		return LexLabel
	}
	if isNumeric(l.Peek()) {
		return LexNumber
	}
	return LexInsideTree
}

func LexOutsideTree(l Lexer) stateFn {
	for {
		r := l.Next()
		if r == '(' {
			l.Backup()
			if l.UnemittedText() {
				l.Emit(ItemText)
			}
			return LexInsideTree
		}
		if r == eof {
			break
		}
	}
	// Correctly reached EOF.
	if l.UnemittedText() {
		l.Emit(ItemText)
	}
	l.Emit(ItemEOF)
	return nil
}

func LexNumber(l Lexer) stateFn {
	if l.MatchNumber() {
		l.Emit(ItemNumber)
	}
	return LexInsideTree
}

func LexLabel(l Lexer) stateFn {
	if l.MatchLabel() {
		l.Emit(ItemLabel)
	}
	return LexInsideTree
}

func LexQuotedString(l Lexer) stateFn {
Loop:
	for {
		switch l.Next() {
		case eof:
			return l.Errorf("unterminated comment")
		case '\'':
			break Loop
		case '"':
			break Loop
		}
	}
	l.Emit(ItemLabel)
	return LexInsideTree
}

func LexComment(l Lexer) stateFn {
	if l.MatchNHX() {
		return LexNHX
	}
Loop:
	for {
		switch l.Next() {
		case eof:
			return l.Errorf("unterminated comment")
		case ']':
			break Loop
		}
	}
	l.Backup()
	l.Emit(ItemComment)
	return LexInsideTree
}
