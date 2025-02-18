package nasType

// BackoffTimerValue 9.11.2.5
// UnitTimerValue Row, sBit, len = [0, 0], 8 , 3
// TimerValue Row, sBit, len = [0, 0], 5 , 5
type BackoffTimerValue struct {
	Iei   uint8 `json:"Iei,omitempty"`
	Len   uint8 `json:"Len,omitempty"`
	Octet uint8 `json:"Octet,omitempty"`
}

func NewBackoffTimerValue(iei uint8) (backoffTimerValue *BackoffTimerValue) {
	backoffTimerValue = &BackoffTimerValue{}
	backoffTimerValue.SetIei(iei)
	return backoffTimerValue
}

// BackoffTimerValue 9.11.2.5
// Iei Row, sBit, len = [], 8, 8
func (a *BackoffTimerValue) GetIei() (iei uint8) {
	return a.Iei
}

// BackoffTimerValue 9.11.2.5
// Iei Row, sBit, len = [], 8, 8
func (a *BackoffTimerValue) SetIei(iei uint8) {
	a.Iei = iei
}

// BackoffTimerValue 9.11.2.5
// Len Row, sBit, len = [], 8, 8
func (a *BackoffTimerValue) GetLen() (len uint8) {
	return a.Len
}

// BackoffTimerValue 9.11.2.5
// Len Row, sBit, len = [], 8, 8
func (a *BackoffTimerValue) SetLen(len uint8) {
	a.Len = len
}

// BackoffTimerValue 9.11.2.5
// UnitTimerValue Row, sBit, len = [0, 0], 8 , 3
func (a *BackoffTimerValue) GetUnitTimerValue() (unitTimerValue uint8) {
	return a.Octet & GetBitMask(8, 5) >> (5)
}

// BackoffTimerValue 9.11.2.5
// UnitTimerValue Row, sBit, len = [0, 0], 8 , 3
func (a *BackoffTimerValue) SetUnitTimerValue(unitTimerValue uint8) {
	a.Octet = (a.Octet & 31) + ((unitTimerValue & 7) << 5)
}

// BackoffTimerValue 9.11.2.5
// TimerValue Row, sBit, len = [0, 0], 5 , 5
func (a *BackoffTimerValue) GetTimerValue() (timerValue uint8) {
	return a.Octet & GetBitMask(5, 0)
}

// BackoffTimerValue 9.11.2.5
// TimerValue Row, sBit, len = [0, 0], 5 , 5
func (a *BackoffTimerValue) SetTimerValue(timerValue uint8) {
	a.Octet = (a.Octet & 224) + (timerValue & 31)
}
