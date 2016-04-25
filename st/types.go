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

// AlertType is the type of and alert report.
type AlertType int

const (
	// Start driving faster than SPEED_LIMIT.
	StartOverSpeedAlt AlertType = iota + 1

	// Ended over-speed condition.
	StopOverSpeedAlt

	// Disconnected GPS antenna.
	DisconnectedGPSAntennaAlt

	// Reconnected GPS antenna after disconnected.
	ReconnectedGPSAntennaAlt

	// The vehicle exited the geo-fenced area that has the following ID.
	ExitedGeoFenceAlt

	// The vehicle entered the geo-fenced area that has the following ID.
	EnteredGeoFenceAlt

	_

	// Shorted GPS antenna (This alert may not support depending on GPS chipset model).
	ShortedGPSAntennaAlt

	// Enter the Deep Sleep Mode.
	EnterDeepSleepModeAlt

	// Exit from the Deep Sleep Mode.
	ExitDeepSleepModeAlt

	_
	_

	// Backup battery error.
	BackupBatteryErrorAlt

	// Vehicle battery goes down to a very low level.
	BatteryLowLevelAlt

	// Shocked.
	ShockedAlt

	// Collision has occurred to the vehicle.
	CollisionAlt

	_

	// Deviated from the predefined route.
	DeviatedFromRouteAlt

	// Entered into the predefined route.
	EnteredIntoRouteAlt

	_
	_

	// Engine :Exceed Speed
	EngineExceedSpeedAlt

	// Engine :Vehicle Speed Violation
	EngineVehicleSpeedAlt

	// Engine :Coolant Temperature Violation
	EngineCoolantTempAlt

	// Engine :Oil Pressure Violation
	EngineOilPressureAlt

	// Engine :RPM Violation
	EngineRPMAlt

	// Engine :Sudden Hard Brake Violation
	EngineHardBrakeAlt

	// Engine :Error Code(DTC)
	EngineErrCodeAlt

	_
	_
	_
	_

	// Ignition ON
	IgnitionOnAlt

	// Ignition OFF
	IgnitionOffAlt

	_
	_
	_
	_
	_

	// Connected to the Main Power source.
	ConnectedToMainPowerAlt

	// Disconnected from the Main Power source.
	DisconnectedFromMainPowerAlt

	_
	_

	// Connected to the Back-up Battery.
	ConnectedToBackupBatteryAlt

	// Disconnected from the Back-up Battery.
	DisconnectedToBackupBatteryAlt

	// Alert of fast acceleration from Driver Pattern Analysis.
	FastAccelerationFromDPAAlt

	// Alert of fast braking from Driver Pattern Analysis.
	FastBrakingFromDPAAlt

	// Alert of sharp turn from Driver Pattern Analysis.
	SharpTurnFromDPAAlt

	// Alert of over speed from Driver Pattern Analysis.
	OverSpeedFromDPAAlt

	// Jamming detected.
	JammingDetectedAlt

	_
	_
	_
	_
	_
	_
	_
	_

	// Inserted I-Button.
	InsertedIButtonAlt

	// Removed I-Button.
	RemovedIButtonAlt

	// The vehicle turns on but doesn’t drive during predefined time.
	DriveLessThanPredefinedTimeAlt

	// Stopped more than predefined time.
	StoppedMoreThanPredefinedTimeAlt

	// Dead center.
	DeadCenterAlt

	// Over RPM.
	OverRPMAlt

	// Completed automatic RPM calibration.
	CompletedAutoRPMCalibrationAlt

	// Completed automatic odometer calibration. (by ignition or by command).
	CompletedAutoOdometerCalibrationAlt

	// Completed automatic odometer calibration as another type in dual gear system.
	CompletedAutoOdometerCalibrationDualGearSystemAlt

	// Alert of Stop limit at Ignition ON.
	StopLimitAtIgnitionONAlt

	// Alert of Moving after Alert 68.
	MovingAfterStopLimitAtIgnitionONAlt

	_
	_
	_

	// Alert of rapid reduction of the Fuel.
	RapidFuelReductionAlt
)
