package st

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/larixsource/suntech/lexer"
)

const (
	Separator  = ';'
	EndOfFrame = '\r'

	tsLayout = "20060102;15:04:05;"
)

var (
	ErrSeparator        = errors.New("invalid separator, a ';' was expected")
	ErrEndOfFrame       = errors.New("invalid end of frame, a CR was expected")
	ErrInvalidDevID     = errors.New("invalid DevID")
	ErrInvalidModel     = errors.New("invalid Model")
	ErrInvalidSwVer     = errors.New("invalid SwVer")
	ErrInvalidDate      = errors.New("invalid Date")
	ErrInvalidTime      = errors.New("invalid Time")
	ErrInvalidCell      = errors.New("invalid Cell")
	ErrInvalidLat       = errors.New("invalid Latitude")
	ErrInvalidLng       = errors.New("invalid Longitude")
	ErrInvalidSpeed     = errors.New("invalid Speed")
	ErrInvalidCourse    = errors.New("invalid Course")
	ErrInvalidSatt      = errors.New("invalid Satt")
	ErrInvalidFix       = errors.New("invalid Fix")
	ErrInvalidDist      = errors.New("invalid Dist")
	ErrInvalidPowerVolt = errors.New("invalid PowerVolt")
	ErrInvalidMode      = errors.New("invalid Mode")
	ErrInvalidMsgNum    = errors.New("invalid MsgNum")
	ErrInvalidHMeter    = errors.New("invalid HMeter")
	ErrInvalidMsgType   = errors.New("invalid MsgType")
	ErrInvalidEmgID     = errors.New("invalid EmgID")
	ErrInvalidEvtID     = errors.New("invalid EvtID")
	ErrInvalidAltID     = errors.New("invalid AltID")
	ErrInvalidHLen      = errors.New("invalid Length")
)

func AsciiDevID(lex *lexer.Lexer) (devID string, token lexer.Token, err error) {
	token, err = lex.NextFixed(10)
	if err != nil {
		return
	}
	if !token.OnlyDigits() {
		err = ErrInvalidDevID
		return
	}
	if !token.EndsWith(Separator) {
		err = ErrSeparator
		return
	}
	devID = string(token.WithoutSuffix())
	return
}

func AsciiDevIDAtEnd(lex *lexer.Lexer) (devID string, token lexer.Token, err error) {
	token, err = lex.NextFixed(10)
	if err != nil {
		return
	}
	if !token.OnlyDigits() {
		err = ErrInvalidDevID
		return
	}
	if !token.EndsWith(EndOfFrame) {
		err = ErrSeparator
		return
	}
	devID = string(token.WithoutSuffix())
	return
}

func AsciiDevIDOrRes(lex *lexer.Lexer) (isDevID bool, devID string, token lexer.Token, err error) {
	token, err = lex.Next(10, Separator)
	if err != nil {
		return
	}

	switch {
	case token.OnlyDigits() && len(token.WithoutSuffix()) == 9:
		isDevID = true
		devID = string(token.WithoutSuffix())
		return
	case bytes.Equal(ResLiteral, token.WithoutSuffix()):
		return
	default:
		err = ErrInvalidDevID
		return
	}
}

func AsciiModel(lex *lexer.Lexer) (model Model, token lexer.Token, err error) {
	token, err = lex.NextFixed(3)
	if err != nil {
		return
	}
	if !token.OnlyDigits() {
		err = ErrInvalidModel
		return
	}
	if !token.EndsWith(Separator) {
		err = ErrSeparator
		return
	}
	var md uint64
	md, err = strconv.ParseUint(string(token.WithoutSuffix()), 10, 8)
	if err != nil {
		return
	}
	model = Model(md)
	return
}

func AsciiSwVer(lex *lexer.Lexer) (swVer uint16, token lexer.Token, err error) {
	token, err = lex.NextFixed(4)
	if err != nil {
		return
	}
	if !token.OnlyDigits() {
		err = ErrInvalidSwVer
		return
	}
	if !token.EndsWith(Separator) {
		err = ErrSeparator
		return
	}
	var swv uint64
	swv, err = strconv.ParseUint(string(token.WithoutSuffix()), 10, 16)
	if err != nil {
		return
	}
	swVer = uint16(swv)
	return
}

func AsciiTimestamp(lex *lexer.Lexer) (ts time.Time, tokens []lexer.Token, err error) {
	tokens = make([]lexer.Token, 0, 2)

	// date
	dateToken, dateErr := lex.NextFixed(9)
	tokens = append(tokens, dateToken)
	if dateErr != nil {
		err = dateErr
		return
	}
	if !dateToken.OnlyDigits() {
		err = ErrInvalidDate
		return
	}
	if !dateToken.EndsWith(Separator) {
		err = ErrSeparator
		return
	}

	// time
	timeToken, timeErr := lex.NextFixed(9)
	tokens = append(tokens, timeToken)
	if timeErr != nil {
		err = timeErr
		return
	}
	if timeToken.Type != lexer.DataToken {
		err = ErrInvalidTime
		return
	}
	if !timeToken.EndsWith(Separator) {
		err = ErrSeparator
		return
	}

	// to timestamp
	buf := bytes.NewBuffer(make([]byte, 0, 18))
	buf.Write(dateToken.Literal)
	buf.Write(timeToken.Literal)
	ts, err = time.Parse(tsLayout, buf.String())
	return
}

func AsciiCell(lex *lexer.Lexer) (cell string, token lexer.Token, err error) {
	token, err = lex.Next(7, Separator)
	if err != nil {
		return
	}
	if !token.IsHex() {
		err = ErrInvalidCell
		return
	}
	cell = string(token.WithoutSuffix())
	return
}

func AsciiLat(lex *lexer.Lexer) (lat float32, token lexer.Token, err error) {
	token, err = lex.Next(11, Separator)
	if err != nil {
		return
	}
	if !token.IsFloat() {
		err = ErrInvalidLat
		return
	}
	lat64, parseErr := strconv.ParseFloat(string(token.WithoutSuffix()), 32)
	if parseErr != nil {
		err = parseErr
		return
	}
	lat = float32(lat64)
	return
}

func AsciiLon(lex *lexer.Lexer) (lng float32, token lexer.Token, err error) {
	token, err = lex.Next(12, Separator)
	if err != nil {
		return
	}
	if !token.IsFloat() {
		err = ErrInvalidLng
		return
	}
	lng64, parseErr := strconv.ParseFloat(string(token.WithoutSuffix()), 32)
	if parseErr != nil {
		err = parseErr
		return
	}
	lng = float32(lng64)
	return
}

func AsciiSpeed(lex *lexer.Lexer) (speed float32, token lexer.Token, err error) {
	token, err = lex.Next(8, Separator)
	if err != nil {
		return
	}
	if !token.IsFloat() {
		err = ErrInvalidSpeed
		return
	}
	spd, parseErr := strconv.ParseFloat(string(token.WithoutSuffix()), 32)
	if parseErr != nil {
		err = parseErr
		return
	}
	speed = float32(spd)
	return
}

func AsciiCourse(lex *lexer.Lexer) (speed float32, token lexer.Token, err error) {
	token, err = lex.Next(7, Separator)
	if err != nil {
		return
	}
	if !token.IsFloat() {
		err = ErrInvalidCourse
		return
	}
	crs, parseErr := strconv.ParseFloat(string(token.WithoutSuffix()), 32)
	if parseErr != nil {
		err = parseErr
		return
	}
	speed = float32(crs)
	return
}

func AsciiSatellites(lex *lexer.Lexer) (satellites uint8, token lexer.Token, err error) {
	token, err = lex.Next(3, Separator)
	if err != nil {
		return
	}
	if !token.OnlyDigits() {
		err = ErrInvalidSatt
		return
	}
	sat, parseErr := strconv.ParseUint(string(token.WithoutSuffix()), 10, 8)
	if parseErr != nil {
		err = parseErr
		return
	}
	satellites = uint8(sat)
	return
}

func AsciiFix(lex *lexer.Lexer) (fix bool, token lexer.Token, err error) {
	token, err = lex.Next(3, Separator)
	if err != nil {
		return
	}
	if token.Type != lexer.BitsToken {
		err = ErrInvalidFix
		return
	}
	switch token.Literal[0] {
	case '0':
		fix = false
	case '1':
		fix = true
	default:
		err = ErrInvalidFix
	}
	return
}

func AsciiDistance(lex *lexer.Lexer) (distance uint32, token lexer.Token, err error) {
	token, err = lex.Next(11, Separator)
	if err != nil {
		return
	}
	if !token.OnlyDigits() {
		err = ErrInvalidDist
		return
	}
	dist, parseErr := strconv.ParseUint(string(token.WithoutSuffix()), 10, 32)
	if parseErr != nil {
		err = parseErr
		return
	}
	distance = uint32(dist)
	return
}

func AsciiPowerVolt(lex *lexer.Lexer) (powerVolt float32, token lexer.Token, err error) {
	token, err = lex.Next(11, Separator)
	if err != nil {
		return
	}
	if !token.IsFloat() {
		err = ErrInvalidPowerVolt
		return
	}
	pv, parseErr := strconv.ParseFloat(string(token.WithoutSuffix()), 32)
	if parseErr != nil {
		err = parseErr
		return
	}
	powerVolt = float32(pv)
	return
}

func AsciiMode(lex *lexer.Lexer) (mode ModeType, token lexer.Token, err error) {
	token, err = lex.NextFixed(2)
	if err != nil {
		return
	}
	if !token.OnlyDigits() {
		err = ErrInvalidMode
		return
	}
	if !token.EndsWith(Separator) {
		err = ErrSeparator
		return
	}
	switch token.Literal[0] {
	case '1':
		mode = IdleMode
	case '2':
		mode = ActiveMode
	case '4':
		mode = DistanceMode
	case '5':
		mode = AngleMode
	default:
		err = fmt.Errorf("invalid mode value: %v", token.Literal[0])
	}
	return
}

func AsciiMsgNum(lex *lexer.Lexer) (msgNum uint16, token lexer.Token, err error) {
	token, err = lex.NextFixed(5)
	if err != nil {
		return
	}
	if !token.IsHex() {
		err = ErrInvalidMsgNum
		return
	}
	if !token.EndsWith(Separator) {
		err = ErrSeparator
		return
	}
	mnum, parseErr := strconv.ParseUint(string(token.WithoutSuffix()), 10, 16)
	if parseErr != nil {
		err = parseErr
	}
	msgNum = uint16(mnum)
	return
}

func AsciiDrivingHourMeter(lex *lexer.Lexer) (hmeter uint32, token lexer.Token, err error) {
	token, err = lex.Next(8, Separator)
	if err != nil {
		return
	}
	if !token.OnlyDigits() {
		err = ErrInvalidHMeter
		return
	}
	hm, parseErr := strconv.ParseUint(string(token.WithoutSuffix()), 10, 32)
	if parseErr != nil {
		err = parseErr
		return
	}
	hmeter = uint32(hm)
	return
}

func AsciiBackupVolt(lex *lexer.Lexer) (backupVolt float32, token lexer.Token, err error) {
	token, err = lex.Next(11, Separator)
	if err != nil {
		return
	}
	if !token.IsFloat() {
		err = ErrInvalidHMeter
		return
	}
	bv, parseErr := strconv.ParseFloat(string(token.WithoutSuffix()), 32)
	if parseErr != nil {
		err = parseErr
		return
	}
	backupVolt = float32(bv)

	return
}

func AsciiMsgType(lex *lexer.Lexer, last bool) (realTime bool, token lexer.Token, err error) {
	token, err = lex.NextFixed(2)
	if err != nil {
		return
	}
	if token.Type != lexer.BitsToken {
		err = ErrInvalidMsgType
		return
	}
	switch {
	case last && !token.EndsWith(EndOfFrame):
		err = ErrEndOfFrame
		return
	case !last && !token.EndsWith(Separator):
		err = ErrSeparator
		return
	}
	switch token.Literal[0] {
	case '0':
		realTime = false
	case '1':
		realTime = true
	default:
		err = fmt.Errorf("invalid MsgType value: %v", token.Literal[0])
	}
	return
}

func AsciiUnknownTail(lex *lexer.Lexer, max int) (token lexer.Token, err error) {
	token, err = lex.Next(max, EndOfFrame)
	return
}

func AsciiEmgID(lex *lexer.Lexer) (emgType EmergencyType, token lexer.Token, err error) {
	token, err = lex.NextFixed(2)
	if err != nil {
		return
	}
	if !token.OnlyDigits() {
		err = ErrInvalidEmgID
		return
	}
	if !token.EndsWith(Separator) {
		err = ErrSeparator
		return
	}
	switch token.Literal[0] {
	case '1':
		emgType = PanicButtonEmg
	case '2':
		emgType = ParkingLockEmg
	case '3':
		emgType = RemovingMainPowerEmg
	case '5':
		emgType = AntiTheftEmg
	case '6':
		emgType = AntiTheftDoorEmg
	case '7':
		emgType = MotionEmg
	case '8':
		emgType = AntiTheftShockEmg
	default:
		err = fmt.Errorf("unknown EmgID value: %v", token.Literal[0])
	}
	return
}

func AsciiEvtID(lex *lexer.Lexer) (evtType EventType, token lexer.Token, err error) {
	token, err = lex.NextFixed(2)
	if err != nil {
		return
	}
	if !token.OnlyDigits() {
		err = ErrInvalidEvtID
		return
	}
	if !token.EndsWith(Separator) {
		err = ErrSeparator
		return
	}
	switch token.Literal[0] {
	case '1':
		evtType = Input1GroundEvt
	case '2':
		evtType = Input1OpenEvt
	case '3':
		evtType = Input2GroundEvt
	case '4':
		evtType = Input2OpenEvt
	case '5':
		evtType = Input3GroundEvt
	case '6':
		evtType = Input3OpenEvt
	default:
		err = fmt.Errorf("unknown EvtID value: %v", token.Literal[0])
	}
	return
}

func AsciiAltID(lex *lexer.Lexer) (altType AlertType, token lexer.Token, err error) {
	token, err = lex.Next(3, Separator)
	if err != nil {
		return
	}
	if !token.OnlyDigits() {
		err = ErrInvalidAltID
		return
	}
	if !token.EndsWith(Separator) {
		err = ErrSeparator
		return
	}
	switch string(token.WithoutSuffix()) {
	case "1":
		altType = StartOverSpeedAlt
	case "2":
		altType = StopOverSpeedAlt
	case "3":
		altType = DisconnectedGPSAntennaAlt
	case "4":
		altType = ReconnectedGPSAntennaAlt
	case "5":
		altType = ExitedGeoFenceAlt
	case "6":
		altType = EnteredGeoFenceAlt
	case "8":
		altType = ShortedGPSAntennaAlt
	case "9":
		altType = EnterDeepSleepModeAlt
	case "10":
		altType = ExitDeepSleepModeAlt
	case "13":
		altType = BackupBatteryErrorAlt
	case "14":
		altType = BatteryLowLevelAlt
	case "15":
		altType = ShockedAlt
	case "16":
		altType = CollisionAlt
	case "18":
		altType = DeviatedFromRouteAlt
	case "19":
		altType = EnteredIntoRouteAlt
	case "22":
		altType = EngineExceedSpeedAlt
	case "23":
		altType = EngineVehicleSpeedAlt
	case "24":
		altType = EngineCoolantTempAlt
	case "25":
		altType = EngineOilPressureAlt
	case "26":
		altType = EngineRPMAlt
	case "27":
		altType = EngineHardBrakeAlt
	case "28":
		altType = EngineErrCodeAlt
	case "33":
		altType = IgnitionOnAlt
	case "34":
		altType = IgnitionOffAlt
	case "40":
		altType = ConnectedToMainPowerAlt
	case "41":
		altType = DisconnectedFromMainPowerAlt
	case "44":
		altType = ConnectedToBackupBatteryAlt
	case "45":
		altType = DisconnectedToBackupBatteryAlt
	case "46":
		altType = FastAccelerationFromDPAAlt
	case "47":
		altType = FastBrakingFromDPAAlt
	case "48":
		altType = SharpTurnFromDPAAlt
	case "49":
		altType = OverSpeedFromDPAAlt
	case "50":
		altType = JammingDetectedAlt
	case "59":
		altType = InsertedIButtonAlt
	case "60":
		altType = RemovedIButtonAlt
	case "61":
		altType = DriveLessThanPredefinedTimeAlt
	case "62":
		altType = StoppedMoreThanPredefinedTimeAlt
	case "63":
		altType = DeadCenterAlt
	case "64":
		altType = OverRPMAlt
	case "65":
		altType = CompletedAutoRPMCalibrationAlt
	case "66":
		altType = CompletedAutoOdometerCalibrationAlt
	case "67":
		altType = CompletedAutoOdometerCalibrationDualGearSystemAlt
	case "68":
		altType = StopLimitAtIgnitionONAlt
	case "69":
		altType = MovingAfterStopLimitAtIgnitionONAlt
	case "73":
		altType = RapidFuelReductionAlt
	default:
		err = fmt.Errorf("unknown AltID value: %v", token.Literal[0])
	}
	return
}

func AsciiADC(lex *lexer.Lexer, last bool) (adc float32, token lexer.Token, err error) {
	var c byte = Separator
	if last {
		c = EndOfFrame
	}
	token, err = lex.Next(6, c)
	if err != nil {
		return
	}
	if !token.IsFloat() {
		err = ErrInvalidHMeter
		return
	}
	bv, parseErr := strconv.ParseFloat(string(token.WithoutSuffix()), 32)
	if parseErr != nil {
		err = parseErr
		return
	}
	adc = float32(bv)

	return
}

func AsciiLen(lex *lexer.Lexer) (length uint16, token lexer.Token, err error) {
	token, err = lex.Next(6, Separator)
	if err != nil {
		return
	}
	if !token.OnlyDigits() {
		err = ErrInvalidHMeter
		return
	}
	l, parseErr := strconv.ParseUint(string(token.WithoutSuffix()), 10, 16)
	if parseErr != nil {
		err = parseErr
		return
	}
	length = uint16(l)

	return
}

func AsciiChecksum(lex *lexer.Lexer) (chk uint8, token lexer.Token, err error) {
	token, err = lex.Next(6, Separator)
	if err != nil {
		return
	}
	if !token.OnlyDigits() {
		err = ErrInvalidHMeter
		return
	}
	crc, parseErr := strconv.ParseUint(string(token.WithoutSuffix()), 16, 16)
	if parseErr != nil {
		err = parseErr
		return
	}
	chk = uint8(crc)

	return
}
