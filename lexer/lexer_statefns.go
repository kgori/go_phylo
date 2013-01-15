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

func LexNHX(l Lexer) stateFn {
	b, itemType := l.MatchSingleChar()
	if b {
		switch itemType {
		case ItemColon:
			l.Ignore()
			return LexKey
		case ItemEquals:
			// l.Emit(ItemEquals)
			l.Ignore()
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
		// l.Emit(itemType)
		switch itemType {
		case ItemLBracket:
			l.Emit(itemType)
			if l.Peek() == ':' || l.Peek() == ',' {
				return LexLabel
			}
			return LexInsideTree
		case ItemRBracket:
			l.Emit(itemType)
			b, _ := l.Accept("0123456789.-+")
			if b {
				l.Backup()
				return LexSupportValue
			} else {
				return LexInsideTree
			}
		case ItemLSquare:
			l.Ignore()
			return LexComment
		case ItemRSquare:
			l.Ignore()
			return LexInsideTree
		case ItemColon:
			l.Ignore()
			return LexBranchLength
		case ItemSemiColon:
			l.Emit(itemType)
			return LexOutsideTree
		case ItemComma:
			l.Ignore()
			return LexInsideTree
		case ItemSingleQuote:
			l.Ignore()
			return LexSingleQuotedString
		case ItemDoubleQuote:
			l.Ignore()
			return LexDoubleQuotedString
		case ItemEquals:
			l.Ignore()
			return LexInsideTree
		default:
			return l.Errorf("unrecognized character in tree: %#U", itemType)
		}
	}
	if isAlphaNumeric(l.Peek()) {
		return LexLabel
	}
	// if isNumeric(l.Peek()) {
	// 	return LexNumber
	// }
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

func LexBranchLength(l Lexer) stateFn {
	if l.MatchNumber() {
		l.Emit(ItemBranchLength)
	}
	return LexInsideTree
}

func LexSupportValue(l Lexer) stateFn {
	if l.MatchNumber() {
		l.Emit(ItemSupportValue)
	}
	return LexInsideTree
}

func LexLabel(l Lexer) stateFn {
	if l.MatchLabel() {
		l.Emit(ItemLabel)
	}
	return LexInsideTree
}

func LexSingleQuotedString(l Lexer) stateFn {
Loop:
	for {
		switch l.Next() {
		case eof:
			return l.Errorf("unterminated comment")
		case '\'':
			break Loop
		}
	}
	l.Backup()
	l.Emit(ItemLabel)
	l.Next()
	l.Ignore()
	return LexInsideTree
}

func LexDoubleQuotedString(l Lexer) stateFn {
Loop:
	for {
		switch l.Next() {
		case eof:
			return l.Errorf("unterminated comment")
		case '"':
			break Loop
		}
	}
	l.Backup()
	l.Emit(ItemLabel)
	l.Next()
	l.Ignore()
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
