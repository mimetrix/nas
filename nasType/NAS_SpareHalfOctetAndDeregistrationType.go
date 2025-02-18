package nasType

// SpareHalfOctetAndDeregistrationType 9.11.3.20 9.5
// SwitchOff Row, sBit, len = [0, 0], 4 , 1
// ReRegistrationRequired Row, sBit, len = [0, 0], 3 , 1
// AccessType Row, sBit, len = [0, 0], 2 , 2
type SpareHalfOctetAndDeregistrationType struct {
	Octet uint8 `json:"Octet,omitempty"`
}

func NewSpareHalfOctetAndDeregistrationType() (spareHalfOctetAndDeregistrationType *SpareHalfOctetAndDeregistrationType) {
	spareHalfOctetAndDeregistrationType = &SpareHalfOctetAndDeregistrationType{}
	return spareHalfOctetAndDeregistrationType
}

// SpareHalfOctetAndDeregistrationType 9.11.3.20 9.5
// SwitchOff Row, sBit, len = [0, 0], 4 , 1
func (a *SpareHalfOctetAndDeregistrationType) GetSwitchOff() (switchOff uint8) {
	return a.Octet & GetBitMask(4, 3) >> (3)
}

// SpareHalfOctetAndDeregistrationType 9.11.3.20 9.5
// SwitchOff Row, sBit, len = [0, 0], 4 , 1
func (a *SpareHalfOctetAndDeregistrationType) SetSwitchOff(switchOff uint8) {
	a.Octet = (a.Octet & 247) + ((switchOff & 1) << 3)
}

// SpareHalfOctetAndDeregistrationType 9.11.3.20 9.5
// ReRegistrationRequired Row, sBit, len = [0, 0], 3 , 1
func (a *SpareHalfOctetAndDeregistrationType) GetReRegistrationRequired() (reRegistrationRequired uint8) {
	return a.Octet & GetBitMask(3, 2) >> (2)
}

// SpareHalfOctetAndDeregistrationType 9.11.3.20 9.5
// ReRegistrationRequired Row, sBit, len = [0, 0], 3 , 1
func (a *SpareHalfOctetAndDeregistrationType) SetReRegistrationRequired(reRegistrationRequired uint8) {
	a.Octet = (a.Octet & 251) + ((reRegistrationRequired & 1) << 2)
}

// SpareHalfOctetAndDeregistrationType 9.11.3.20 9.5
// AccessType Row, sBit, len = [0, 0], 2 , 2
func (a *SpareHalfOctetAndDeregistrationType) GetAccessType() (accessType uint8) {
	return a.Octet & GetBitMask(2, 0)
}

// SpareHalfOctetAndDeregistrationType 9.11.3.20 9.5
// AccessType Row, sBit, len = [0, 0], 2 , 2
func (a *SpareHalfOctetAndDeregistrationType) SetAccessType(accessType uint8) {
	a.Octet = (a.Octet & 252) + (accessType & 3)
}
