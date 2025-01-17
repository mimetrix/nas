package nasType

// ReplayedS1UESecurityCapabilities 9.11.3.48A
// EEA0 Row, sBit, len = [0, 0], 8 , 1
// EEA1_128 Row, sBit, len = [0, 0], 7 , 1
// EEA2_128 Row, sBit, len = [0, 0], 6 , 1
// EEA3_128 Row, sBit, len = [0, 0], 5 , 1
// EEA4 Row, sBit, len = [0, 0], 4 , 1
// EEA5 Row, sBit, len = [0, 0], 3 , 1
// EEA6 Row, sBit, len = [0, 0], 2 , 1
// EEA7 Row, sBit, len = [0, 0], 1 , 1
// EIA0 Row, sBit, len = [1, 1], 8 , 1
// EIA1_128 Row, sBit, len = [1, 1], 7 , 1
// EIA2_128 Row, sBit, len = [1, 1], 6 , 1
// EIA3_128 Row, sBit, len = [1, 1], 5 , 1
// EIA4 Row, sBit, len = [1, 1], 4 , 1
// EIA5 Row, sBit, len = [1, 1], 3 , 1
// EIA6 Row, sBit, len = [1, 1], 2 , 1
// EIA7 Row, sBit, len = [1, 1], 1 , 1
// UEA0 Row, sBit, len = [2, 2], 8 , 1
// UEA1 Row, sBit, len = [2, 2], 7 , 1
// UEA2 Row, sBit, len = [2, 2], 6 , 1
// UEA3 Row, sBit, len = [2, 2], 5 , 1
// UEA4 Row, sBit, len = [2, 2], 4 , 1
// UEA5 Row, sBit, len = [2, 2], 3 , 1
// UEA6 Row, sBit, len = [2, 2], 2 , 1
// UEA7 Row, sBit, len = [2, 2], 1 , 1
// UIA1 Row, sBit, len = [3, 3], 7 , 1
// UIA2 Row, sBit, len = [3, 3], 6 , 1
// UIA3 Row, sBit, len = [3, 3], 5 , 1
// UIA4 Row, sBit, len = [3, 3], 4 , 1
// UIA5 Row, sBit, len = [3, 3], 3 , 1
// UIA6 Row, sBit, len = [3, 3], 2 , 1
// UIA7 Row, sBit, len = [3, 3], 1 , 1
// GEA1 Row, sBit, len = [4, 4], 7 , 1
// GEA2 Row, sBit, len = [4, 4], 6 , 1
// GEA3 Row, sBit, len = [4, 4], 5 , 1
// GEA4 Row, sBit, len = [4, 4], 4 , 1
// GEA5 Row, sBit, len = [4, 4], 3 , 1
// GEA6 Row, sBit, len = [4, 4], 2 , 1
// GEA7 Row, sBit, len = [4, 4], 1 , 1
type ReplayedS1UESecurityCapabilities struct {
	Iei    uint8   `json:"Iei,omitempty"`
	Len    uint8   `json:"Len,omitempty"`
	Buffer []uint8 `json:"Buffer,omitempty"`
}

func NewReplayedS1UESecurityCapabilities(iei uint8) (replayedS1UESecurityCapabilities *ReplayedS1UESecurityCapabilities) {
	replayedS1UESecurityCapabilities = &ReplayedS1UESecurityCapabilities{}
	replayedS1UESecurityCapabilities.SetIei(iei)
	return replayedS1UESecurityCapabilities
}

// ReplayedS1UESecurityCapabilities 9.11.3.48A
// Iei Row, sBit, len = [], 8, 8
func (a *ReplayedS1UESecurityCapabilities) GetIei() (iei uint8) {
	return a.Iei
}

// ReplayedS1UESecurityCapabilities 9.11.3.48A
// Iei Row, sBit, len = [], 8, 8
func (a *ReplayedS1UESecurityCapabilities) SetIei(iei uint8) {
	a.Iei = iei
}

// ReplayedS1UESecurityCapabilities 9.11.3.48A
// Len Row, sBit, len = [], 8, 8
func (a *ReplayedS1UESecurityCapabilities) GetLen() (len uint8) {
	return a.Len
}

// ReplayedS1UESecurityCapabilities 9.11.3.48A
// Len Row, sBit, len = [], 8, 8
func (a *ReplayedS1UESecurityCapabilities) SetLen(len uint8) {
	a.Len = len
	a.Buffer = make([]uint8, a.Len)
}

// ReplayedS1UESecurityCapabilities 9.11.3.48A
// EEA0 Row, sBit, len = [0, 0], 8 , 1
func (a *ReplayedS1UESecurityCapabilities) GetEEA0() (eEA0 uint8) {
	return a.Buffer[0] & GetBitMask(8, 7) >> (7)
}

// ReplayedS1UESecurityCapabilities 9.11.3.48A
// EEA0 Row, sBit, len = [0, 0], 8 , 1
func (a *ReplayedS1UESecurityCapabilities) SetEEA0(eEA0 uint8) {
	a.Buffer[0] = (a.Buffer[0] & 127) + ((eEA0 & 1) << 7)
}

// ReplayedS1UESecurityCapabilities 9.11.3.48A
// EEA1_128 Row, sBit, len = [0, 0], 7 , 1
func (a *ReplayedS1UESecurityCapabilities) GetEEA1_128() (eEA1_128 uint8) {
	return a.Buffer[0] & GetBitMask(7, 6) >> (6)
}

// ReplayedS1UESecurityCapabilities 9.11.3.48A
// EEA1_128 Row, sBit, len = [0, 0], 7 , 1
func (a *ReplayedS1UESecurityCapabilities) SetEEA1_128(eEA1_128 uint8) {
	a.Buffer[0] = (a.Buffer[0] & 191) + ((eEA1_128 & 1) << 6)
}

// ReplayedS1UESecurityCapabilities 9.11.3.48A
// EEA2_128 Row, sBit, len = [0, 0], 6 , 1
func (a *ReplayedS1UESecurityCapabilities) GetEEA2_128() (eEA2_128 uint8) {
	return a.Buffer[0] & GetBitMask(6, 5) >> (5)
}

// ReplayedS1UESecurityCapabilities 9.11.3.48A
// EEA2_128 Row, sBit, len = [0, 0], 6 , 1
func (a *ReplayedS1UESecurityCapabilities) SetEEA2_128(eEA2_128 uint8) {
	a.Buffer[0] = (a.Buffer[0] & 223) + ((eEA2_128 & 1) << 5)
}

// ReplayedS1UESecurityCapabilities 9.11.3.48A
// EEA3_128 Row, sBit, len = [0, 0], 5 , 1
func (a *ReplayedS1UESecurityCapabilities) GetEEA3_128() (eEA3_128 uint8) {
	return a.Buffer[0] & GetBitMask(5, 4) >> (4)
}

// ReplayedS1UESecurityCapabilities 9.11.3.48A
// EEA3_128 Row, sBit, len = [0, 0], 5 , 1
func (a *ReplayedS1UESecurityCapabilities) SetEEA3_128(eEA3_128 uint8) {
	a.Buffer[0] = (a.Buffer[0] & 239) + ((eEA3_128 & 1) << 4)
}

// ReplayedS1UESecurityCapabilities 9.11.3.48A
// EEA4 Row, sBit, len = [0, 0], 4 , 1
func (a *ReplayedS1UESecurityCapabilities) GetEEA4() (eEA4 uint8) {
	return a.Buffer[0] & GetBitMask(4, 3) >> (3)
}

// ReplayedS1UESecurityCapabilities 9.11.3.48A
// EEA4 Row, sBit, len = [0, 0], 4 , 1
func (a *ReplayedS1UESecurityCapabilities) SetEEA4(eEA4 uint8) {
	a.Buffer[0] = (a.Buffer[0] & 247) + ((eEA4 & 1) << 3)
}

// ReplayedS1UESecurityCapabilities 9.11.3.48A
// EEA5 Row, sBit, len = [0, 0], 3 , 1
func (a *ReplayedS1UESecurityCapabilities) GetEEA5() (eEA5 uint8) {
	return a.Buffer[0] & GetBitMask(3, 2) >> (2)
}

// ReplayedS1UESecurityCapabilities 9.11.3.48A
// EEA5 Row, sBit, len = [0, 0], 3 , 1
func (a *ReplayedS1UESecurityCapabilities) SetEEA5(eEA5 uint8) {
	a.Buffer[0] = (a.Buffer[0] & 251) + ((eEA5 & 1) << 2)
}

// ReplayedS1UESecurityCapabilities 9.11.3.48A
// EEA6 Row, sBit, len = [0, 0], 2 , 1
func (a *ReplayedS1UESecurityCapabilities) GetEEA6() (eEA6 uint8) {
	return a.Buffer[0] & GetBitMask(2, 1) >> (1)
}

// ReplayedS1UESecurityCapabilities 9.11.3.48A
// EEA6 Row, sBit, len = [0, 0], 2 , 1
func (a *ReplayedS1UESecurityCapabilities) SetEEA6(eEA6 uint8) {
	a.Buffer[0] = (a.Buffer[0] & 253) + ((eEA6 & 1) << 1)
}

// ReplayedS1UESecurityCapabilities 9.11.3.48A
// EEA7 Row, sBit, len = [0, 0], 1 , 1
func (a *ReplayedS1UESecurityCapabilities) GetEEA7() (eEA7 uint8) {
	return a.Buffer[0] & GetBitMask(1, 0)
}

// ReplayedS1UESecurityCapabilities 9.11.3.48A
// EEA7 Row, sBit, len = [0, 0], 1 , 1
func (a *ReplayedS1UESecurityCapabilities) SetEEA7(eEA7 uint8) {
	a.Buffer[0] = (a.Buffer[0] & 254) + (eEA7 & 1)
}

// ReplayedS1UESecurityCapabilities 9.11.3.48A
// EIA0 Row, sBit, len = [1, 1], 8 , 1
func (a *ReplayedS1UESecurityCapabilities) GetEIA0() (eIA0 uint8) {
	return a.Buffer[1] & GetBitMask(8, 7) >> (7)
}

// ReplayedS1UESecurityCapabilities 9.11.3.48A
// EIA0 Row, sBit, len = [1, 1], 8 , 1
func (a *ReplayedS1UESecurityCapabilities) SetEIA0(eIA0 uint8) {
	a.Buffer[1] = (a.Buffer[1] & 127) + ((eIA0 & 1) << 7)
}

// ReplayedS1UESecurityCapabilities 9.11.3.48A
// EIA1_128 Row, sBit, len = [1, 1], 7 , 1
func (a *ReplayedS1UESecurityCapabilities) GetEIA1_128() (eIA1_128 uint8) {
	return a.Buffer[1] & GetBitMask(7, 6) >> (6)
}

// ReplayedS1UESecurityCapabilities 9.11.3.48A
// EIA1_128 Row, sBit, len = [1, 1], 7 , 1
func (a *ReplayedS1UESecurityCapabilities) SetEIA1_128(eIA1_128 uint8) {
	a.Buffer[1] = (a.Buffer[1] & 191) + ((eIA1_128 & 1) << 6)
}

// ReplayedS1UESecurityCapabilities 9.11.3.48A
// EIA2_128 Row, sBit, len = [1, 1], 6 , 1
func (a *ReplayedS1UESecurityCapabilities) GetEIA2_128() (eIA2_128 uint8) {
	return a.Buffer[1] & GetBitMask(6, 5) >> (5)
}

// ReplayedS1UESecurityCapabilities 9.11.3.48A
// EIA2_128 Row, sBit, len = [1, 1], 6 , 1
func (a *ReplayedS1UESecurityCapabilities) SetEIA2_128(eIA2_128 uint8) {
	a.Buffer[1] = (a.Buffer[1] & 223) + ((eIA2_128 & 1) << 5)
}

// ReplayedS1UESecurityCapabilities 9.11.3.48A
// EIA3_128 Row, sBit, len = [1, 1], 5 , 1
func (a *ReplayedS1UESecurityCapabilities) GetEIA3_128() (eIA3_128 uint8) {
	return a.Buffer[1] & GetBitMask(5, 4) >> (4)
}

// ReplayedS1UESecurityCapabilities 9.11.3.48A
// EIA3_128 Row, sBit, len = [1, 1], 5 , 1
func (a *ReplayedS1UESecurityCapabilities) SetEIA3_128(eIA3_128 uint8) {
	a.Buffer[1] = (a.Buffer[1] & 239) + ((eIA3_128 & 1) << 4)
}

// ReplayedS1UESecurityCapabilities 9.11.3.48A
// EIA4 Row, sBit, len = [1, 1], 4 , 1
func (a *ReplayedS1UESecurityCapabilities) GetEIA4() (eIA4 uint8) {
	return a.Buffer[1] & GetBitMask(4, 3) >> (3)
}

// ReplayedS1UESecurityCapabilities 9.11.3.48A
// EIA4 Row, sBit, len = [1, 1], 4 , 1
func (a *ReplayedS1UESecurityCapabilities) SetEIA4(eIA4 uint8) {
	a.Buffer[1] = (a.Buffer[1] & 247) + ((eIA4 & 1) << 3)
}

// ReplayedS1UESecurityCapabilities 9.11.3.48A
// EIA5 Row, sBit, len = [1, 1], 3 , 1
func (a *ReplayedS1UESecurityCapabilities) GetEIA5() (eIA5 uint8) {
	return a.Buffer[1] & GetBitMask(3, 2) >> (2)
}

// ReplayedS1UESecurityCapabilities 9.11.3.48A
// EIA5 Row, sBit, len = [1, 1], 3 , 1
func (a *ReplayedS1UESecurityCapabilities) SetEIA5(eIA5 uint8) {
	a.Buffer[1] = (a.Buffer[1] & 251) + ((eIA5 & 1) << 2)
}

// ReplayedS1UESecurityCapabilities 9.11.3.48A
// EIA6 Row, sBit, len = [1, 1], 2 , 1
func (a *ReplayedS1UESecurityCapabilities) GetEIA6() (eIA6 uint8) {
	return a.Buffer[1] & GetBitMask(2, 1) >> (1)
}

// ReplayedS1UESecurityCapabilities 9.11.3.48A
// EIA6 Row, sBit, len = [1, 1], 2 , 1
func (a *ReplayedS1UESecurityCapabilities) SetEIA6(eIA6 uint8) {
	a.Buffer[1] = (a.Buffer[1] & 253) + ((eIA6 & 1) << 1)
}

// ReplayedS1UESecurityCapabilities 9.11.3.48A
// EIA7 Row, sBit, len = [1, 1], 1 , 1
func (a *ReplayedS1UESecurityCapabilities) GetEIA7() (eIA7 uint8) {
	return a.Buffer[1] & GetBitMask(1, 0)
}

// ReplayedS1UESecurityCapabilities 9.11.3.48A
// EIA7 Row, sBit, len = [1, 1], 1 , 1
func (a *ReplayedS1UESecurityCapabilities) SetEIA7(eIA7 uint8) {
	a.Buffer[1] = (a.Buffer[1] & 254) + (eIA7 & 1)
}

// ReplayedS1UESecurityCapabilities 9.11.3.48A
// UEA0 Row, sBit, len = [2, 2], 8 , 1
func (a *ReplayedS1UESecurityCapabilities) GetUEA0() (uEA0 uint8) {
	return a.Buffer[2] & GetBitMask(8, 7) >> (7)
}

// ReplayedS1UESecurityCapabilities 9.11.3.48A
// UEA0 Row, sBit, len = [2, 2], 8 , 1
func (a *ReplayedS1UESecurityCapabilities) SetUEA0(uEA0 uint8) {
	a.Buffer[2] = (a.Buffer[2] & 127) + ((uEA0 & 1) << 7)
}

// ReplayedS1UESecurityCapabilities 9.11.3.48A
// UEA1 Row, sBit, len = [2, 2], 7 , 1
func (a *ReplayedS1UESecurityCapabilities) GetUEA1() (uEA1 uint8) {
	return a.Buffer[2] & GetBitMask(7, 6) >> (6)
}

// ReplayedS1UESecurityCapabilities 9.11.3.48A
// UEA1 Row, sBit, len = [2, 2], 7 , 1
func (a *ReplayedS1UESecurityCapabilities) SetUEA1(uEA1 uint8) {
	a.Buffer[2] = (a.Buffer[2] & 191) + ((uEA1 & 1) << 6)
}

// ReplayedS1UESecurityCapabilities 9.11.3.48A
// UEA2 Row, sBit, len = [2, 2], 6 , 1
func (a *ReplayedS1UESecurityCapabilities) GetUEA2() (uEA2 uint8) {
	return a.Buffer[2] & GetBitMask(6, 5) >> (5)
}

// ReplayedS1UESecurityCapabilities 9.11.3.48A
// UEA2 Row, sBit, len = [2, 2], 6 , 1
func (a *ReplayedS1UESecurityCapabilities) SetUEA2(uEA2 uint8) {
	a.Buffer[2] = (a.Buffer[2] & 223) + ((uEA2 & 1) << 5)
}

// ReplayedS1UESecurityCapabilities 9.11.3.48A
// UEA3 Row, sBit, len = [2, 2], 5 , 1
func (a *ReplayedS1UESecurityCapabilities) GetUEA3() (uEA3 uint8) {
	return a.Buffer[2] & GetBitMask(5, 4) >> (4)
}

// ReplayedS1UESecurityCapabilities 9.11.3.48A
// UEA3 Row, sBit, len = [2, 2], 5 , 1
func (a *ReplayedS1UESecurityCapabilities) SetUEA3(uEA3 uint8) {
	a.Buffer[2] = (a.Buffer[2] & 239) + ((uEA3 & 1) << 4)
}

// ReplayedS1UESecurityCapabilities 9.11.3.48A
// UEA4 Row, sBit, len = [2, 2], 4 , 1
func (a *ReplayedS1UESecurityCapabilities) GetUEA4() (uEA4 uint8) {
	return a.Buffer[2] & GetBitMask(4, 3) >> (3)
}

// ReplayedS1UESecurityCapabilities 9.11.3.48A
// UEA4 Row, sBit, len = [2, 2], 4 , 1
func (a *ReplayedS1UESecurityCapabilities) SetUEA4(uEA4 uint8) {
	a.Buffer[2] = (a.Buffer[2] & 247) + ((uEA4 & 1) << 3)
}

// ReplayedS1UESecurityCapabilities 9.11.3.48A
// UEA5 Row, sBit, len = [2, 2], 3 , 1
func (a *ReplayedS1UESecurityCapabilities) GetUEA5() (uEA5 uint8) {
	return a.Buffer[2] & GetBitMask(3, 2) >> (2)
}

// ReplayedS1UESecurityCapabilities 9.11.3.48A
// UEA5 Row, sBit, len = [2, 2], 3 , 1
func (a *ReplayedS1UESecurityCapabilities) SetUEA5(uEA5 uint8) {
	a.Buffer[2] = (a.Buffer[2] & 251) + ((uEA5 & 1) << 2)
}

// ReplayedS1UESecurityCapabilities 9.11.3.48A
// UEA6 Row, sBit, len = [2, 2], 2 , 1
func (a *ReplayedS1UESecurityCapabilities) GetUEA6() (uEA6 uint8) {
	return a.Buffer[2] & GetBitMask(2, 1) >> (1)
}

// ReplayedS1UESecurityCapabilities 9.11.3.48A
// UEA6 Row, sBit, len = [2, 2], 2 , 1
func (a *ReplayedS1UESecurityCapabilities) SetUEA6(uEA6 uint8) {
	a.Buffer[2] = (a.Buffer[2] & 253) + ((uEA6 & 1) << 1)
}

// ReplayedS1UESecurityCapabilities 9.11.3.48A
// UEA7 Row, sBit, len = [2, 2], 1 , 1
func (a *ReplayedS1UESecurityCapabilities) GetUEA7() (uEA7 uint8) {
	return a.Buffer[2] & GetBitMask(1, 0)
}

// ReplayedS1UESecurityCapabilities 9.11.3.48A
// UEA7 Row, sBit, len = [2, 2], 1 , 1
func (a *ReplayedS1UESecurityCapabilities) SetUEA7(uEA7 uint8) {
	a.Buffer[2] = (a.Buffer[2] & 254) + (uEA7 & 1)
}

// ReplayedS1UESecurityCapabilities 9.11.3.48A
// UIA1 Row, sBit, len = [3, 3], 7 , 1
func (a *ReplayedS1UESecurityCapabilities) GetUIA1() (uIA1 uint8) {
	return a.Buffer[3] & GetBitMask(7, 6) >> (6)
}

// ReplayedS1UESecurityCapabilities 9.11.3.48A
// UIA1 Row, sBit, len = [3, 3], 7 , 1
func (a *ReplayedS1UESecurityCapabilities) SetUIA1(uIA1 uint8) {
	a.Buffer[3] = (a.Buffer[3] & 191) + ((uIA1 & 1) << 6)
}

// ReplayedS1UESecurityCapabilities 9.11.3.48A
// UIA2 Row, sBit, len = [3, 3], 6 , 1
func (a *ReplayedS1UESecurityCapabilities) GetUIA2() (uIA2 uint8) {
	return a.Buffer[3] & GetBitMask(6, 5) >> (5)
}

// ReplayedS1UESecurityCapabilities 9.11.3.48A
// UIA2 Row, sBit, len = [3, 3], 6 , 1
func (a *ReplayedS1UESecurityCapabilities) SetUIA2(uIA2 uint8) {
	a.Buffer[3] = (a.Buffer[3] & 223) + ((uIA2 & 1) << 5)
}

// ReplayedS1UESecurityCapabilities 9.11.3.48A
// UIA3 Row, sBit, len = [3, 3], 5 , 1
func (a *ReplayedS1UESecurityCapabilities) GetUIA3() (uIA3 uint8) {
	return a.Buffer[3] & GetBitMask(5, 4) >> (4)
}

// ReplayedS1UESecurityCapabilities 9.11.3.48A
// UIA3 Row, sBit, len = [3, 3], 5 , 1
func (a *ReplayedS1UESecurityCapabilities) SetUIA3(uIA3 uint8) {
	a.Buffer[3] = (a.Buffer[3] & 239) + ((uIA3 & 1) << 4)
}

// ReplayedS1UESecurityCapabilities 9.11.3.48A
// UIA4 Row, sBit, len = [3, 3], 4 , 1
func (a *ReplayedS1UESecurityCapabilities) GetUIA4() (uIA4 uint8) {
	return a.Buffer[3] & GetBitMask(4, 3) >> (3)
}

// ReplayedS1UESecurityCapabilities 9.11.3.48A
// UIA4 Row, sBit, len = [3, 3], 4 , 1
func (a *ReplayedS1UESecurityCapabilities) SetUIA4(uIA4 uint8) {
	a.Buffer[3] = (a.Buffer[3] & 247) + ((uIA4 & 1) << 3)
}

// ReplayedS1UESecurityCapabilities 9.11.3.48A
// UIA5 Row, sBit, len = [3, 3], 3 , 1
func (a *ReplayedS1UESecurityCapabilities) GetUIA5() (uIA5 uint8) {
	return a.Buffer[3] & GetBitMask(3, 2) >> (2)
}

// ReplayedS1UESecurityCapabilities 9.11.3.48A
// UIA5 Row, sBit, len = [3, 3], 3 , 1
func (a *ReplayedS1UESecurityCapabilities) SetUIA5(uIA5 uint8) {
	a.Buffer[3] = (a.Buffer[3] & 251) + ((uIA5 & 1) << 2)
}

// ReplayedS1UESecurityCapabilities 9.11.3.48A
// UIA6 Row, sBit, len = [3, 3], 2 , 1
func (a *ReplayedS1UESecurityCapabilities) GetUIA6() (uIA6 uint8) {
	return a.Buffer[3] & GetBitMask(2, 1) >> (1)
}

// ReplayedS1UESecurityCapabilities 9.11.3.48A
// UIA6 Row, sBit, len = [3, 3], 2 , 1
func (a *ReplayedS1UESecurityCapabilities) SetUIA6(uIA6 uint8) {
	a.Buffer[3] = (a.Buffer[3] & 253) + ((uIA6 & 1) << 1)
}

// ReplayedS1UESecurityCapabilities 9.11.3.48A
// UIA7 Row, sBit, len = [3, 3], 1 , 1
func (a *ReplayedS1UESecurityCapabilities) GetUIA7() (uIA7 uint8) {
	return a.Buffer[3] & GetBitMask(1, 0)
}

// ReplayedS1UESecurityCapabilities 9.11.3.48A
// UIA7 Row, sBit, len = [3, 3], 1 , 1
func (a *ReplayedS1UESecurityCapabilities) SetUIA7(uIA7 uint8) {
	a.Buffer[3] = (a.Buffer[3] & 254) + (uIA7 & 1)
}

// ReplayedS1UESecurityCapabilities 9.11.3.48A
// GEA1 Row, sBit, len = [4, 4], 7 , 1
func (a *ReplayedS1UESecurityCapabilities) GetGEA1() (gEA1 uint8) {
	return a.Buffer[4] & GetBitMask(7, 6) >> (6)
}

// ReplayedS1UESecurityCapabilities 9.11.3.48A
// GEA1 Row, sBit, len = [4, 4], 7 , 1
func (a *ReplayedS1UESecurityCapabilities) SetGEA1(gEA1 uint8) {
	a.Buffer[4] = (a.Buffer[4] & 191) + ((gEA1 & 1) << 6)
}

// ReplayedS1UESecurityCapabilities 9.11.3.48A
// GEA2 Row, sBit, len = [4, 4], 6 , 1
func (a *ReplayedS1UESecurityCapabilities) GetGEA2() (gEA2 uint8) {
	return a.Buffer[4] & GetBitMask(6, 5) >> (5)
}

// ReplayedS1UESecurityCapabilities 9.11.3.48A
// GEA2 Row, sBit, len = [4, 4], 6 , 1
func (a *ReplayedS1UESecurityCapabilities) SetGEA2(gEA2 uint8) {
	a.Buffer[4] = (a.Buffer[4] & 223) + ((gEA2 & 1) << 5)
}

// ReplayedS1UESecurityCapabilities 9.11.3.48A
// GEA3 Row, sBit, len = [4, 4], 5 , 1
func (a *ReplayedS1UESecurityCapabilities) GetGEA3() (gEA3 uint8) {
	return a.Buffer[4] & GetBitMask(5, 4) >> (4)
}

// ReplayedS1UESecurityCapabilities 9.11.3.48A
// GEA3 Row, sBit, len = [4, 4], 5 , 1
func (a *ReplayedS1UESecurityCapabilities) SetGEA3(gEA3 uint8) {
	a.Buffer[4] = (a.Buffer[4] & 239) + ((gEA3 & 1) << 4)
}

// ReplayedS1UESecurityCapabilities 9.11.3.48A
// GEA4 Row, sBit, len = [4, 4], 4 , 1
func (a *ReplayedS1UESecurityCapabilities) GetGEA4() (gEA4 uint8) {
	return a.Buffer[4] & GetBitMask(4, 3) >> (3)
}

// ReplayedS1UESecurityCapabilities 9.11.3.48A
// GEA4 Row, sBit, len = [4, 4], 4 , 1
func (a *ReplayedS1UESecurityCapabilities) SetGEA4(gEA4 uint8) {
	a.Buffer[4] = (a.Buffer[4] & 247) + ((gEA4 & 1) << 3)
}

// ReplayedS1UESecurityCapabilities 9.11.3.48A
// GEA5 Row, sBit, len = [4, 4], 3 , 1
func (a *ReplayedS1UESecurityCapabilities) GetGEA5() (gEA5 uint8) {
	return a.Buffer[4] & GetBitMask(3, 2) >> (2)
}

// ReplayedS1UESecurityCapabilities 9.11.3.48A
// GEA5 Row, sBit, len = [4, 4], 3 , 1
func (a *ReplayedS1UESecurityCapabilities) SetGEA5(gEA5 uint8) {
	a.Buffer[4] = (a.Buffer[4] & 251) + ((gEA5 & 1) << 2)
}

// ReplayedS1UESecurityCapabilities 9.11.3.48A
// GEA6 Row, sBit, len = [4, 4], 2 , 1
func (a *ReplayedS1UESecurityCapabilities) GetGEA6() (gEA6 uint8) {
	return a.Buffer[4] & GetBitMask(2, 1) >> (1)
}

// ReplayedS1UESecurityCapabilities 9.11.3.48A
// GEA6 Row, sBit, len = [4, 4], 2 , 1
func (a *ReplayedS1UESecurityCapabilities) SetGEA6(gEA6 uint8) {
	a.Buffer[4] = (a.Buffer[4] & 253) + ((gEA6 & 1) << 1)
}

// ReplayedS1UESecurityCapabilities 9.11.3.48A
// GEA7 Row, sBit, len = [4, 4], 1 , 1
func (a *ReplayedS1UESecurityCapabilities) GetGEA7() (gEA7 uint8) {
	return a.Buffer[4] & GetBitMask(1, 0)
}

// ReplayedS1UESecurityCapabilities 9.11.3.48A
// GEA7 Row, sBit, len = [4, 4], 1 , 1
func (a *ReplayedS1UESecurityCapabilities) SetGEA7(gEA7 uint8) {
	a.Buffer[4] = (a.Buffer[4] & 254) + (gEA7 & 1)
}
