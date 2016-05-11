package st600

import (
	"errors"
	"strconv"

	"github.com/larixsource/suntech/lexer"
	"github.com/larixsource/suntech/st"
)

var (
	ErrInvalidMCC         = errors.New("invalid MCC")
	ErrInvalidMNC         = errors.New("invalid MNC")
	ErrInvalidLAC         = errors.New("invalid LAC")
	ErrInvalidSignalLevel = errors.New("invalid SignalLevel")
)

type Cell3G struct {
	CellID      string
	MCC         string
	MNC         string
	LAC         string
	SignalLevel float32
}

type CellType int

const (
	Cell2GType CellType = iota
	Cell3GType
)

type Cell struct {
	Type   CellType
	Cell2G string
	Cell3G
}

// asciiCell3G reads a 2G or 3G cell
func asciiCell3G(lex *lexer.Lexer) (Cell, []lexer.Token, error) {
	var tokens []lexer.Token
	token, err := lex.Next(9, st.Separator)
	tokens = append(tokens, token)
	if err != nil {
		return Cell{}, tokens, err
	}
	if !token.IsHex() {
		return Cell{}, tokens, st.ErrInvalidCell
	}

	// an cell of length 9 is interpreted as a 3G cell
	if len(token.Literal) == 9 {
		cell := Cell{
			Type: Cell3GType,
			Cell3G: Cell3G{
				CellID: string(token.WithoutSuffix()),
			},
		}

		mcc, mccToken, mccErr := asciiMCC(lex)
		tokens = append(tokens, mccToken)
		if mccErr != nil {
			return cell, tokens, ErrInvalidMCC
		}
		cell.MCC = mcc

		mnc, mncToken, mncErr := asciiMNC(lex)
		tokens = append(tokens, mncToken)
		if mncErr != nil {
			return cell, tokens, ErrInvalidMNC
		}
		cell.MNC = mnc

		lac, lacToken, lacErr := asciiLAC(lex)
		tokens = append(tokens, lacToken)
		if lacErr != nil {
			return cell, tokens, ErrInvalidLAC
		}
		cell.LAC = lac

		sl, slToken, slErr := asciiSignalLevel(lex)
		tokens = append(tokens, slToken)
		if slErr != nil {
			return cell, tokens, ErrInvalidSignalLevel
		}
		cell.SignalLevel = sl

		return cell, tokens, nil
	}

	// otherwise, it's a 2G cell
	cell := Cell{
		Type:   Cell2GType,
		Cell2G: string(token.WithoutSuffix()),
	}
	return cell, tokens, nil
}

func asciiMCC(lex *lexer.Lexer) (mcc string, token lexer.Token, err error) {
	token, err = lex.Next(4, st.Separator)
	if err != nil {
		return
	}
	if !token.IsHex() {
		err = ErrInvalidMCC
		return
	}
	mcc = string(token.WithoutSuffix())
	return
}

func asciiMNC(lex *lexer.Lexer) (mnc string, token lexer.Token, err error) {
	token, err = lex.Next(4, st.Separator)
	if err != nil {
		return
	}
	if !token.IsHex() {
		err = ErrInvalidMNC
		return
	}
	mnc = string(token.WithoutSuffix())
	return
}

func asciiLAC(lex *lexer.Lexer) (lac string, token lexer.Token, err error) {
	token, err = lex.Next(5, st.Separator)
	if err != nil {
		return
	}
	if !token.IsHex() {
		err = ErrInvalidLAC
		return
	}
	lac = string(token.WithoutSuffix())
	return
}

func asciiSignalLevel(lex *lexer.Lexer) (signalLevel float32, token lexer.Token, err error) {
	token, err = lex.Next(4, st.Separator)
	if err != nil {
		return
	}
	if !token.IsFloat() {
		err = ErrInvalidSignalLevel
		return
	}
	sl64, parseErr := strconv.ParseFloat(string(token.WithoutSuffix()), 32)
	if parseErr != nil {
		err = parseErr
		return
	}
	signalLevel = float32(sl64)
	return
}
