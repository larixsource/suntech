package st600

import (
	"errors"

	"github.com/larixsource/suntech/lexer"
	"github.com/larixsource/suntech/st"
)

var ErrInvalidIO = errors.New("invalid IO")

func knownModel(model st.Model) bool {
	switch model {
	case st.ST600V, st.ST600R:
		return true
	default:
		return false
	}
}

func asciiIO(lex *lexer.Lexer) (ioStatus string, token lexer.Token, err error) {
	token, err = lex.Next(9, st.Separator)
	if err != nil {
		return
	}
	if token.Type != lexer.BitsToken {
		err = ErrInvalidIO
		return
	}
	if !token.EndsWith(st.Separator) {
		err = st.ErrSeparator
		return
	}
	ioStatus = string(token.WithoutSuffix())
	return
}
