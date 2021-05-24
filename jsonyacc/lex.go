package jsonyacc

import (
	"bytes"
	"unicode"
)

func Parse(input []byte) (map[string]interface{}, error) {

	l := newLex(input)
	_ = yyParse(l)
	return l.result, l.err
}

type lex struct {
	input  *bytes.Buffer
	result map[string]interface{}
	err    error
}

func nexLex(input []byte) *lex {

	return &lex{

		input: bytes.NewBuffer(input),
	}
}

const eof = 0

func (l *lex) Lex(lval *yySymType) int {

	for {
		r, _, err := l.input.ReadRune()
		if err != nil {
			return eof
		}
		switch {
		case unicode.IsSpace(r):
			continue
		case r == '"':
			return l.scanString(lval)
		case unicode.IsDigit(r) || r == '+' || r == '-':
			_ = l.input.UnreadRune()
			return l.scanNum(lval)
		case unicode.IsLetter(r):
			_ = l.input.UnreadRune()
			return l.scanLiteral(lval)
		default:
			return int(r)
		}
	}
}
