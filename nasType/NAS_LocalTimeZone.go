package nasType

// LocalTimeZone 9.11.3.52
// TimeZone Row, sBit, len = [0, 0], 8 , 8
/*
From 3gpp 23.040
The Time Zone indicates the difference, expressed in quarters of an hour, between the local time and GMT. In the first
of the two semi-octets, the first bit (bit 3 of the seventh octet of the TP-Service-Centre-Time-Stamp field) represents the
algebraic sign of this difference (0: positive, 1: negative).
*/
type LocalTimeZone struct {
	Iei   uint8 `json:"-"`
	Octet uint8 `json:"Timezone"`
}

func NewLocalTimeZone(iei uint8) (localTimeZone *LocalTimeZone) {
	localTimeZone = &LocalTimeZone{}
	localTimeZone.SetIei(iei)
	return localTimeZone
}

// LocalTimeZone 9.11.3.52
// Iei Row, sBit, len = [], 8, 8
func (a *LocalTimeZone) GetIei() (iei uint8) {
	return a.Iei
}

// LocalTimeZone 9.11.3.52
// Iei Row, sBit, len = [], 8, 8
func (a *LocalTimeZone) SetIei(iei uint8) {
	a.Iei = iei
}

// LocalTimeZone 9.11.3.52
// TimeZone Row, sBit, len = [0, 0], 8 , 8
func (a *LocalTimeZone) GetTimeZone() (timeZone uint8) {
	return a.Octet
}

// LocalTimeZone 9.11.3.52
// TimeZone Row, sBit, len = [0, 0], 8 , 8
func (a *LocalTimeZone) SetTimeZone(timeZone uint8) {
	a.Octet = timeZone
}
