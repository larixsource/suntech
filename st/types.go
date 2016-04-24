// Package st provides common support to Suntech GPS parsers
package st

import "errors"

// Model is the device model type
type Model uint8

const (
	UnknownModel Model = 0
	ST300RI      Model = 1
	ST340        Model = 2
	ST340LC      Model = 3
	ST300H       Model = 4
	ST350        Model = 5
	ST480        Model = 6
	ST300A       Model = 7
	ST300R       Model = 8
	ST300B       Model = 9
	ST300V       Model = 10
	ST300C       Model = 11
	ST300K       Model = 12
	ST300P       Model = 13
	ST300F       Model = 14
)

const (
	// STX is the start of ZIP msg mark byte
	STX = 0x02

	// ETX is the end of ZIP msg mark byte
	ETX = 0x03
)

var (
	ErrZipUnsupported   = errors.New("zip msg unsupported")
	ErrUnsupportedModel = errors.New("unsupported model")
)

type ModeType int

const (
	IdleMode     ModeType = 1
	ActiveMode   ModeType = 2
	DistanceMode ModeType = 4
	AngleMode    ModeType = 5
)
