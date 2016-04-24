package lexer

import "fmt"

type dfaState int

const (
	emptyState dfaState = iota
	bitsState
	digitsState
	hexState
	signState
	intState
	dotState
	floatState
	dataState
)

func (dst dfaState) next(c byte) dfaState {
	switch dst {
	case emptyState:
		switch {
		case c == '0' || c == '1':
			return bitsState
		case c >= '2' && c <= '9':
			return digitsState
		case c == 'a' || c == 'b' || c == 'c' || c == 'd' || c == 'e' || c == 'f' ||
			c == 'A' || c == 'B' || c == 'C' || c == 'D' || c == 'E' || c == 'F':
			return hexState
		case c == '-' || c == '+':
			return signState
		default:
			return dataState
		}
	case bitsState:
		switch {
		case c == '0' || c == '1':
			return bitsState
		case c >= '2' && c <= '9':
			return digitsState
		case c == 'a' || c == 'b' || c == 'c' || c == 'd' || c == 'e' || c == 'f' ||
			c == 'A' || c == 'B' || c == 'C' || c == 'D' || c == 'E' || c == 'F':
			return hexState
		case c == '.':
			return dotState
		default:
			return dataState
		}
	case digitsState:
		switch {
		case c >= '0' && c <= '9':
			return digitsState
		case c == 'a' || c == 'b' || c == 'c' || c == 'd' || c == 'e' || c == 'f' ||
			c == 'A' || c == 'B' || c == 'C' || c == 'D' || c == 'E' || c == 'F':
			return hexState
		case c == '.':
			return dotState
		default:
			return dataState
		}
	case hexState:
		if (c >= '0' && c <= '9') || (c == 'a' || c == 'b' || c == 'c' || c == 'd' || c == 'e' || c == 'f' ||
			c == 'A' || c == 'B' || c == 'C' || c == 'D' || c == 'E' || c == 'F') {
			return hexState
		}
		return dataState
	case signState:
		if c >= '0' && c <= '9' {
			return intState
		}
		return dataState
	case intState:
		switch {
		case c >= '0' && c <= '9':
			return intState
		case c == '.':
			return dotState
		default:
			return dataState
		}
	case dotState:
		if c >= '0' && c <= '9' {
			return floatState
		}
		return dataState
	case floatState:
		if c >= '0' && c <= '9' {
			return floatState
		}
		return dataState
	case dataState:
		return dataState
	default:
		panic(fmt.Errorf("unknown dfa state: %v", dst))
	}
}
