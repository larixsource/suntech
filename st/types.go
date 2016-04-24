// Package st provides common support to Suntech GPS parsers
package st

import "errors"

// Model is the device model type
type Model uint8

const (
	UnknownModel Model = iota
	ST300
	ST340
	ST340LC
	ST300H
	ST350
	ST480
	ST300A
	ST300R
	ST300B
	ST300V
	ST300C
	ST300K
	ST300P
	ST300F
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

	ResLiteral = []byte{'R', 'e', 's'}
)

type ModeType int

const (
	IdleMode ModeType = iota + 1
	ActiveMode
	_
	DistanceMode
	AngleMode
)

// EmergencyType is the type of an emergency report
type EmergencyType int

const (
	// PanicButton is issued when the panic button is pressed.
	PanicButtonEmg EmergencyType = iota + 1

	ParkingLockEmg

	// RemovingMainPower is issued when the main power is removed (available only in  with backup battery).
	RemovingMainPowerEmg

	_

	AntiTheftEmg

	AntiTheftDoorEmg

	MotionEmg

	AntiTheftShockEmg
)

// EventType is the type of an event report.
type EventType int

const (
	Input1GroundEvt EventType = iota + 1

	Input1OpenEvt

	Input2GroundEvt

	Input2OpenEvt

	Input3GroundEvt

	Input3OpenEvt
)
