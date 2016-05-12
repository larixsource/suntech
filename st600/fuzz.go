package st600

func Fuzz(data []byte) int {
	p := ParseBytes(data, ParserOpts{})

	var results []int
	for p.Next() {
		frame := p.Msg()
		if frame == nil {
			panic("nil frame")
		}
		if len(frame.Frame) == 0 {
			panic("empty raw frame")
		}
		if frame.ParsingError != nil {
			results = append(results, 0)
			continue
		}
		switch frame.Type {
		case STTReport:
			if frame.STT == nil {
				panic("nil STT")
			}
		case EMGReport:
			if frame.EMG == nil {
				panic("nil EMG")
			}
		case EVTReport:
			if frame.EVT == nil {
				panic("nil EVT")
			}
		case ALTReport:
			if frame.ALT == nil {
				panic("nil ALT")
			}
		case ALVReport:
			if frame.ALV == nil {
				panic("nil ALV")
			}
		case UEXReport:
			if frame.UEX == nil {
				panic("nil UEX")
			}
		case UnknownMsg:
		default:
			panic("invalid Type")
		}

		// good frame
		results = append(results, 1)
	}

	// count results (zeroes and ones)
	zeroCount := 0
	oneCount := 0
	for _, r := range results {
		switch r {
		case 0:
			zeroCount++
		case 1:
			oneCount++
		default:
			panic("fuzz programming error")
		}
	}

	switch {
	case oneCount == 0:
		return 0
	case zeroCount == 0 || zeroCount == 1: // at most one error permitted
		return 1
	default:
		return 0
	}
}
