package nasType

// T3512Value 9.11.2.5
// Unit Row, sBit, len = [0, 0], 8 , 3
// TimerValue Row, sBit, len = [0, 0], 5 , 5
type T3512Value struct {
	Iei   uint8 `json:"-"`
	Len   uint8 `json:"-"`
	Octet uint8 `json:"-"`
    IncrementedEvery string
    TimerValue uint8
}

var IncrementUnits = map[uint8]string{
    0:"10 minutes",
    1:"1 hour",
    2:"10 hours",
    3:"2 seconds",
    4:"30 seconds",
    5:"1 minute",
    6:"320 hours",
    7:"timer is deactivated",

}

func (t *T3512Value) DecodeNASType() error {
    t.IncrementedEvery = IncrementUnits[t.GetUnit()]
    t.TimerValue = t.GetTimerValue()
    return nil
}

func NewT3512Value(iei uint8) (t3512Value *T3512Value) {
	t3512Value = &T3512Value{}
	t3512Value.SetIei(iei)
	return t3512Value
}

// T3512Value 9.11.2.5
// Iei Row, sBit, len = [], 8, 8
func (a *T3512Value) GetIei() (iei uint8) {
	return a.Iei
}

// T3512Value 9.11.2.5
// Iei Row, sBit, len = [], 8, 8
func (a *T3512Value) SetIei(iei uint8) {
	a.Iei = iei
}

// T3512Value 9.11.2.5
// Len Row, sBit, len = [], 8, 8
func (a *T3512Value) GetLen() (len uint8) {
	return a.Len
}

// T3512Value 9.11.2.5
// Len Row, sBit, len = [], 8, 8
func (a *T3512Value) SetLen(len uint8) {
	a.Len = len
}

// T3512Value 9.11.2.5
// Unit Row, sBit, len = [0, 0], 8 , 3
func (a *T3512Value) GetUnit() (unit uint8) {
	return a.Octet & GetBitMask(8, 5) >> (5)
}

// T3512Value 9.11.2.5
// Unit Row, sBit, len = [0, 0], 8 , 3
func (a *T3512Value) SetUnit(unit uint8) {
	a.Octet = (a.Octet & 31) + ((unit & 7) << 5)
}

// T3512Value 9.11.2.5
// TimerValue Row, sBit, len = [0, 0], 5 , 5
func (a *T3512Value) GetTimerValue() (timerValue uint8) {
	return a.Octet & GetBitMask(5, 0)
}

// T3512Value 9.11.2.5
// TimerValue Row, sBit, len = [0, 0], 5 , 5
func (a *T3512Value) SetTimerValue(timerValue uint8) {
	a.Octet = (a.Octet & 224) + (timerValue & 31)
}
