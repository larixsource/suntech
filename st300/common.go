package st300

import (
	"errors"

	"github.com/larixsource/suntech/lexer"
	"github.com/larixsource/suntech/st"
)

var ErrInvalidIO = errors.New("invalid IO")

func knownModel(model st.Model) bool {
	switch model {
	case st.ST300, st.ST340, st.ST340LC, st.ST300H, st.ST350, st.ST480, st.ST300A,
		st.ST300R, st.ST300B, st.ST300V, st.ST300C, st.ST300K, st.ST300P, st.ST300F:
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
