package lexer

import "fmt"

type Item struct {
	typ itemType
	val string
}

type itemType int

// tentative list of newick-appropriate tokens (NHX is to do later)
const (
	ItemError        itemType = iota // error occurred; value is text of error
	ItemLBracket                     // begin new node                
	ItemRBracket                     // end new node
	ItemLSquare                      // begin comment
	ItemRSquare                      // end comment
	ItemColon                        // branch length delimiter
	ItemSemiColon                    // end of tree
	ItemComma                        // sister-node separator
	ItemSingleQuote                  // Single-quoted string
	ItemDoubleQuote                  // Double-quoted string
	ItemEquals                       // NHX key-value assignment
	ItemLabel                        // Node / Leaf name
	ItemBranchLength                 // Branch length (int or float)
	ItemSupportValue                 // Support value (int or float)
	ItemComment                      // Newick format allows comments inside square brackets
	ItemText
	ItemNHXKey
	ItemNHXValue
	ItemEOF
)

const eof = -1

var itemNamer = map[itemType]string{
	ItemError:        "Item <Error>",
	ItemLBracket:     "Item <(>",
	ItemRBracket:     "Item <)>",
	ItemLSquare:      "Item <[>",
	ItemRSquare:      "Item <]>",
	ItemColon:        "Item <:>",
	ItemSemiColon:    "Item <;>",
	ItemComma:        "Item <,>",
	ItemSingleQuote:  "Item <'>",
	ItemDoubleQuote:  "Item <\">",
	ItemEquals:       "Item <=>",
	ItemLabel:        "Item <Label>",
	ItemBranchLength: "Item <BranchLength>",
	ItemSupportValue: "Item <SupportValue>",
	ItemComment:      "Item <Comment>",
	ItemText:         "Item <Text>",
	ItemNHXKey:       "Item <NHX key>",
	ItemNHXValue:     "Item <NHX value>",
	ItemEOF:          "Item <EOF>",
}

var singleCharItems = []itemType{ItemLBracket, ItemRBracket, ItemLSquare, ItemRSquare, ItemColon, ItemSemiColon, ItemComma, ItemSingleQuote, ItemDoubleQuote, ItemEquals}

func (i itemType) String() string {
	s := itemNamer[i]
	if s == "" {
		return fmt.Sprintf("item%d", int(i))
	}
	return s
}

// Add method "String" to item: now Printf knows how to print it
func (i Item) String() string {
	switch {
	case i.typ == ItemEOF:
		return "EOF"
	case i.typ == ItemError:
		return i.val
	case len(i.val) > 10:
		return fmt.Sprintf("%.10q...", i.val)
	}
	return fmt.Sprintf("%q", i.val)
}

func NewItem() Item {
	i := Item{
		typ: ItemText,
		val: "",
	}
	return i
}

func (i Item) IsEOF() bool {
	return i.typ == ItemEOF
}

func (i Item) WhatIs() itemType {
	return i.typ
}
