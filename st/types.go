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
	_ // 15
	_
	_
	_
	_
	ST600R
	ST600V
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
	ParkingMode ModeType = iota + 1
	DrivingMode
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
	// 1: Start driving faster than SPEED_LIMIT.
	StartOverSpeedAlt AlertType = iota + 1

	// 2: Ended over-speed condition.
	StopOverSpeedAlt

	// 3: Disconnected GPS antenna.
	DisconnectedGPSAntennaAlt

	// 4: Reconnected GPS antenna after disconnected.
	ReconnectedGPSAntennaAlt

	// 5: The vehicle exited the geo-fenced area that has the following ID.
	ExitedGeoFenceAlt

	// 6: The vehicle entered the geo-fenced area that has the following ID.
	EnteredGeoFenceAlt

	_

	// 8: Shorted GPS antenna (This alert may not support depending on GPS chipset model).
	ShortedGPSAntennaAlt

	// 9: Enter the Deep Sleep Mode.
	EnterDeepSleepModeAlt

	// 10: Exit from the Deep Sleep Mode.
	ExitDeepSleepModeAlt

	_
	_

	// 13: Backup battery error.
	BackupBatteryErrorAlt

	// 14: Vehicle battery goes down to a very low level.
	BatteryLowLevelAlt

	// 15: Shocked.
	ShockedAlt

	// 16: Collision has occurred to the vehicle.
	CollisionAlt

	_

	// 18: Deviated from the predefined route.
	DeviatedFromRouteAlt

	// 19: Entered into the predefined route.
	EnteredIntoRouteAlt

	_
	_

	// 22: Engine :Exceed Speed
	EngineExceedSpeedAlt

	// 23: Engine :Vehicle Speed Violation
	EngineVehicleSpeedAlt

	// 24: Engine :Coolant Temperature Violation
	EngineCoolantTempAlt

	// 25: Engine :Oil Pressure Violation
	EngineOilPressureAlt

	// 26: Engine :RPM Violation
	EngineRPMAlt

	// 27: Engine :Sudden Hard Brake Violation
	EngineHardBrakeAlt

	// 28: Engine :Error Code(DTC)
	EngineErrCodeAlt

	_
	_
	_
	_

	// 33: Ignition ON
	IgnitionOnAlt

	// 34: Ignition OFF
	IgnitionOffAlt

	_
	_
	_
	_
	_

	// 40: Connected to the Main Power source.
	ConnectedToMainPowerAlt

	// 41: Disconnected from the Main Power source.
	DisconnectedFromMainPowerAlt

	_
	_

	// 44: Connected to the Back-up Battery.
	ConnectedToBackupBatteryAlt

	// 45: Disconnected from the Back-up Battery.
	DisconnectedToBackupBatteryAlt

	// 46: Alert of fast acceleration from Driver Pattern Analysis.
	FastAccelerationFromDPAAlt

	// 47: Alert of fast braking from Driver Pattern Analysis.
	FastBrakingFromDPAAlt

	// 48: Alert of sharp turn from Driver Pattern Analysis.
	SharpTurnFromDPAAlt

	// 49: Alert of over speed from Driver Pattern Analysis.
	OverSpeedFromDPAAlt

	// 50: Jamming detected.
	JammingDetectedAlt

	_
	_
	_
	_
	_
	_
	_
	_

	// 59: Inserted I-Button.
	InsertedIButtonAlt

	// 60: Removed I-Button.
	RemovedIButtonAlt

	// 61: The vehicle turns on but doesn’t drive during predefined time.
	DriveLessThanPredefinedTimeAlt

	// 62: Stopped more than predefined time.
	StoppedMoreThanPredefinedTimeAlt

	// 63: Dead center.
	DeadCenterAlt

	// 64: Over RPM.
	OverRPMAlt

	// 65: Completed automatic RPM calibration.
	CompletedAutoRPMCalibrationAlt

	// 66: Completed automatic odometer calibration. (by ignition or by command).
	CompletedAutoOdometerCalibrationAlt

	// 67: Completed automatic odometer calibration as another type in dual gear system.
	CompletedAutoOdometerCalibrationDualGearSystemAlt

	// 68: Alert of Stop limit at Ignition ON.
	StopLimitAtIgnitionONAlt

	// 69: Alert of Moving after Alert 68.
	MovingAfterStopLimitAtIgnitionONAlt

	_
	_
	_

	// 73: Alert of rapid reduction of the Fuel.
	RapidFuelReductionAlt
)
