package lexer

type lexer struct {
	// name   string  // the name of the input; used only for error reports.
	input  string  // the string being scanned.
	state  stateFn // the next lexing function to enter.
	pos    int     // current position in the input.
	start  int     // start position of this item.
	width  int     // width of last rune read from input.
	line   int
	column int
	items  chan Item // channel of scanned items.
}

type stateFn func(Lexer) stateFn

type Lexer interface {
	Next() rune
	Peek() rune
	Backup()
	Emit(itemType)
	NextItem() Item
	Ignore()
	Accept(string) (bool, int)
	AcceptRun(string)
	Line() int
	NewLine()
	Column() int
	Errorf(string, ...interface{}) stateFn
	MatchNumber() bool
	MatchNHX() bool
	MatchSingleChar() (bool, itemType)
	MatchLabel() bool
	UnemittedText() bool
	Pos() int
}

func New(input string) Lexer {
	l := &lexer{
		input: input,
		state: LexOutsideTree,
		items: make(chan Item, 2), // buffer size 1 works, 2 to make sure
		line:  1,
	}
	return l
}
