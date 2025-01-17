package nasType

// UESecurityCapability 9.11.3.54
// EA0_5G Row, sBit, len = [0, 0], 8 , 1
// EA1_128_5G Row, sBit, len = [0, 0], 7 , 1
// EA2_128_5G Row, sBit, len = [0, 0], 6 , 1
// EA3_128_5G Row, sBit, len = [0, 0], 5 , 1
// EA4_5G Row, sBit, len = [0, 0], 4 , 1
// EA5_5G Row, sBit, len = [0, 0], 3 , 1
// EA6_5G Row, sBit, len = [0, 0], 2 , 1
// EA7_5G Row, sBit, len = [0, 0], 1 , 1
// IA0_5G Row, sBit, len = [1, 1], 8 , 1
// IA1_128_5G Row, sBit, len = [1, 1], 7 , 1
// IA2_128_5G Row, sBit, len = [1, 1], 6 , 1
// IA3_128_5G Row, sBit, len = [1, 1], 5 , 1
// IA4_5G Row, sBit, len = [1, 1], 4 , 1
// IA5_5G Row, sBit, len = [1, 1], 3 , 1
// IA6_5G Row, sBit, len = [1, 1], 2 , 1
// IA7_5G Row, sBit, len = [1, 1], 1 , 1
// EEA0 Row, sBit, len = [2, 2], 8 , 1
// EEA1_128 Row, sBit, len = [2, 2], 7 , 1
// EEA2_128 Row, sBit, len = [2, 2], 6 , 1
// EEA3_128 Row, sBit, len = [2, 2], 5 , 1
// EEA4 Row, sBit, len = [2, 2], 4 , 1
// EEA5 Row, sBit, len = [2, 2], 3 , 1
// EEA6 Row, sBit, len = [2, 2], 2 , 1
// EEA7 Row, sBit, len = [2, 2], 1 , 1
// EIA0 Row, sBit, len = [3, 3], 8 , 1
// EIA1_128 Row, sBit, len = [3, 3], 7 , 1
// EIA2_128 Row, sBit, len = [3, 3], 6 , 1
// EIA3_128 Row, sBit, len = [3, 3], 5 , 1
// EIA4 Row, sBit, len = [3, 3], 4 , 1
// EIA5 Row, sBit, len = [3, 3], 3 , 1
// EIA6 Row, sBit, len = [3, 3], 2 , 1
// EIA7 Row, sBit, len = [3, 3], 1 , 1
// Spare Row, sBit, len = [4, 7], 8 , 32
type UESecurityCapability struct {
	Iei    uint8   `json:"-"`
	Len    uint8   `json:"-"`
	Buffer []uint8 `json:"-"`

    EA0_5G	bool
    EA1_128_5G	bool
    EA2_128_5G	bool
    EA3_128_5G	bool
    EA4_5G	bool
    EA5_5G	bool
    EA6_5G	bool
    EA7_5G	bool

    IA0_5G	bool
    IA1_128_5G	bool
    IA2_128_5G	bool
    IA3_128_5G	bool
    IA4_5G	bool
    IA5_5G	bool
    IA6_5G	bool
    IA7_5G	bool

    EEA0	bool
    EEA1_128	bool
    EEA2_128	bool
    EEA3_128	bool
    EEA4	bool
    EEA5	bool
    EEA6	bool
    EEA7	bool
    
    EIA0	bool
    EIA1_128	bool
    EIA2_128	bool
    EIA3_128	bool
    EIA4	bool
    EIA5	bool
    EIA6	bool
    EIA7	bool
}

func (u *UESecurityCapability) DecodeNASType() error{

    if(u.Len >= 1){
        u.EA0_5G = u.GetEA0_5G()==1
        u.EA1_128_5G = u.GetEA1_128_5G()==1
        u.EA2_128_5G = u.GetEA2_128_5G()==1
        u.EA3_128_5G = u.GetEA3_128_5G()==1
        u.EA4_5G = u.GetEA4_5G()==1
        u.EA5_5G = u.GetEA5_5G()==1
        u.EA6_5G = u.GetEA6_5G()==1
        u.EA7_5G = u.GetEA7_5G()==1
    }
    if(u.Len >= 2){
        u.IA0_5G = u.GetIA0_5G()==1
        u.IA1_128_5G = u.GetIA1_128_5G()==1
        u.IA2_128_5G = u.GetIA2_128_5G()==1
        u.IA3_128_5G = u.GetIA3_128_5G()==1
        u.IA4_5G = u.GetIA4_5G()==1
        u.IA5_5G = u.GetIA5_5G()==1
        u.IA6_5G = u.GetIA6_5G()==1
        u.IA7_5G = u.GetIA7_5G()==1
    }

    if(u.Len >= 3){
        u.EEA0 = u.GetEEA0()==1
        u.EEA1_128 = u.GetEEA1_128()==1
        u.EEA2_128 = u.GetEEA2_128()==1
        u.EEA3_128 = u.GetEEA3_128()==1
        u.EEA4 = u.GetEEA4()==1
        u.EEA5 = u.GetEEA5()==1
        u.EEA6 = u.GetEEA6()==1
        u.EEA7 = u.GetEEA7()==1
    }

    if(u.Len >= 4){
        u.EIA0 = u.GetEIA0()==1
        u.EIA1_128 = u.GetEIA1_128()==1
        u.EIA2_128 = u.GetEIA2_128()==1
        u.EIA3_128 = u.GetEIA3_128()==1
        u.EIA4 = u.GetEIA4()==1
        u.EIA5 = u.GetEIA5()==1
        u.EIA6 = u.GetEIA6()==1
        u.EIA7 = u.GetEIA7()==1
    }
        
    return nil
}

func NewUESecurityCapability(iei uint8) (uESecurityCapability *UESecurityCapability) {
	uESecurityCapability = &UESecurityCapability{}
	uESecurityCapability.SetIei(iei)
	return uESecurityCapability
}

// UESecurityCapability 9.11.3.54
// Iei Row, sBit, len = [], 8, 8
func (a *UESecurityCapability) GetIei() (iei uint8) {
	return a.Iei
}

// UESecurityCapability 9.11.3.54
// Iei Row, sBit, len = [], 8, 8
func (a *UESecurityCapability) SetIei(iei uint8) {
	a.Iei = iei
}

// UESecurityCapability 9.11.3.54
// Len Row, sBit, len = [], 8, 8
func (a *UESecurityCapability) GetLen() (len uint8) {
	return a.Len
}

// UESecurityCapability 9.11.3.54
// Len Row, sBit, len = [], 8, 8
func (a *UESecurityCapability) SetLen(len uint8) {
	a.Len = len
	a.Buffer = make([]uint8, a.Len)
}

// UESecurityCapability 9.11.3.54
// EA0_5G Row, sBit, len = [0, 0], 8 , 1
func (a *UESecurityCapability) GetEA0_5G() (eA0_5G uint8) {
	return a.Buffer[0] & GetBitMask(8, 7) >> (7)
}

// UESecurityCapability 9.11.3.54
// EA0_5G Row, sBit, len = [0, 0], 8 , 1
func (a *UESecurityCapability) SetEA0_5G(eA0_5G uint8) {
	a.Buffer[0] = (a.Buffer[0] & 127) + ((eA0_5G & 1) << 7)
}

// UESecurityCapability 9.11.3.54
// EA1_128_5G Row, sBit, len = [0, 0], 7 , 1
func (a *UESecurityCapability) GetEA1_128_5G() (eA1_128_5G uint8) {
	return a.Buffer[0] & GetBitMask(7, 6) >> (6)
}

// UESecurityCapability 9.11.3.54
// EA1_128_5G Row, sBit, len = [0, 0], 7 , 1
func (a *UESecurityCapability) SetEA1_128_5G(eA1_128_5G uint8) {
	a.Buffer[0] = (a.Buffer[0] & 191) + ((eA1_128_5G & 1) << 6)
}

// UESecurityCapability 9.11.3.54
// EA2_128_5G Row, sBit, len = [0, 0], 6 , 1
func (a *UESecurityCapability) GetEA2_128_5G() (eA2_128_5G uint8) {
	return a.Buffer[0] & GetBitMask(6, 5) >> (5)
}

// UESecurityCapability 9.11.3.54
// EA2_128_5G Row, sBit, len = [0, 0], 6 , 1
func (a *UESecurityCapability) SetEA2_128_5G(eA2_128_5G uint8) {
	a.Buffer[0] = (a.Buffer[0] & 223) + ((eA2_128_5G & 1) << 5)
}

// UESecurityCapability 9.11.3.54
// EA3_128_5G Row, sBit, len = [0, 0], 5 , 1
func (a *UESecurityCapability) GetEA3_128_5G() (eA3_128_5G uint8) {
	return a.Buffer[0] & GetBitMask(5, 4) >> (4)
}

// UESecurityCapability 9.11.3.54
// EA3_128_5G Row, sBit, len = [0, 0], 5 , 1
func (a *UESecurityCapability) SetEA3_128_5G(eA3_128_5G uint8) {
	a.Buffer[0] = (a.Buffer[0] & 239) + ((eA3_128_5G & 1) << 4)
}

// UESecurityCapability 9.11.3.54
// EA4_5G Row, sBit, len = [0, 0], 4 , 1
func (a *UESecurityCapability) GetEA4_5G() (eA4_5G uint8) {
	return a.Buffer[0] & GetBitMask(4, 3) >> (3)
}

// UESecurityCapability 9.11.3.54
// EA4_5G Row, sBit, len = [0, 0], 4 , 1
func (a *UESecurityCapability) SetEA4_5G(eA4_5G uint8) {
	a.Buffer[0] = (a.Buffer[0] & 247) + ((eA4_5G & 1) << 3)
}

// UESecurityCapability 9.11.3.54
// EA5_5G Row, sBit, len = [0, 0], 3 , 1
func (a *UESecurityCapability) GetEA5_5G() (eA5_5G uint8) {
	return a.Buffer[0] & GetBitMask(3, 2) >> (2)
}

// UESecurityCapability 9.11.3.54
// EA5_5G Row, sBit, len = [0, 0], 3 , 1
func (a *UESecurityCapability) SetEA5_5G(eA5_5G uint8) {
	a.Buffer[0] = (a.Buffer[0] & 251) + ((eA5_5G & 1) << 2)
}

// UESecurityCapability 9.11.3.54
// EA6_5G Row, sBit, len = [0, 0], 2 , 1
func (a *UESecurityCapability) GetEA6_5G() (eA6_5G uint8) {
	return a.Buffer[0] & GetBitMask(2, 1) >> (1)
}

// UESecurityCapability 9.11.3.54
// EA6_5G Row, sBit, len = [0, 0], 2 , 1
func (a *UESecurityCapability) SetEA6_5G(eA6_5G uint8) {
	a.Buffer[0] = (a.Buffer[0] & 253) + ((eA6_5G & 1) << 1)
}

// UESecurityCapability 9.11.3.54
// EA7_5G Row, sBit, len = [0, 0], 1 , 1
func (a *UESecurityCapability) GetEA7_5G() (eA7_5G uint8) {
	return a.Buffer[0] & GetBitMask(1, 0)
}

// UESecurityCapability 9.11.3.54
// EA7_5G Row, sBit, len = [0, 0], 1 , 1
func (a *UESecurityCapability) SetEA7_5G(eA7_5G uint8) {
	a.Buffer[0] = (a.Buffer[0] & 254) + (eA7_5G & 1)
}

// UESecurityCapability 9.11.3.54
// IA0_5G Row, sBit, len = [1, 1], 8 , 1
func (a *UESecurityCapability) GetIA0_5G() (iA0_5G uint8) {
	return a.Buffer[1] & GetBitMask(8, 7) >> (7)
}

// UESecurityCapability 9.11.3.54
// IA0_5G Row, sBit, len = [1, 1], 8 , 1
func (a *UESecurityCapability) SetIA0_5G(iA0_5G uint8) {
	a.Buffer[1] = (a.Buffer[1] & 127) + ((iA0_5G & 1) << 7)
}

// UESecurityCapability 9.11.3.54
// IA1_128_5G Row, sBit, len = [1, 1], 7 , 1
func (a *UESecurityCapability) GetIA1_128_5G() (iA1_128_5G uint8) {
	return a.Buffer[1] & GetBitMask(7, 6) >> (6)
}

// UESecurityCapability 9.11.3.54
// IA1_128_5G Row, sBit, len = [1, 1], 7 , 1
func (a *UESecurityCapability) SetIA1_128_5G(iA1_128_5G uint8) {
	a.Buffer[1] = (a.Buffer[1] & 191) + ((iA1_128_5G & 1) << 6)
}

// UESecurityCapability 9.11.3.54
// IA2_128_5G Row, sBit, len = [1, 1], 6 , 1
func (a *UESecurityCapability) GetIA2_128_5G() (iA2_128_5G uint8) {
	return a.Buffer[1] & GetBitMask(6, 5) >> (5)
}

// UESecurityCapability 9.11.3.54
// IA2_128_5G Row, sBit, len = [1, 1], 6 , 1
func (a *UESecurityCapability) SetIA2_128_5G(iA2_128_5G uint8) {
	a.Buffer[1] = (a.Buffer[1] & 223) + ((iA2_128_5G & 1) << 5)
}

// UESecurityCapability 9.11.3.54
// IA3_128_5G Row, sBit, len = [1, 1], 5 , 1
func (a *UESecurityCapability) GetIA3_128_5G() (iA3_128_5G uint8) {
	return a.Buffer[1] & GetBitMask(5, 4) >> (4)
}

// UESecurityCapability 9.11.3.54
// IA3_128_5G Row, sBit, len = [1, 1], 5 , 1
func (a *UESecurityCapability) SetIA3_128_5G(iA3_128_5G uint8) {
	a.Buffer[1] = (a.Buffer[1] & 239) + ((iA3_128_5G & 1) << 4)
}

// UESecurityCapability 9.11.3.54
// IA4_5G Row, sBit, len = [1, 1], 4 , 1
func (a *UESecurityCapability) GetIA4_5G() (iA4_5G uint8) {
	return a.Buffer[1] & GetBitMask(4, 3) >> (3)
}

// UESecurityCapability 9.11.3.54
// IA4_5G Row, sBit, len = [1, 1], 4 , 1
func (a *UESecurityCapability) SetIA4_5G(iA4_5G uint8) {
	a.Buffer[1] = (a.Buffer[1] & 247) + ((iA4_5G & 1) << 3)
}

// UESecurityCapability 9.11.3.54
// IA5_5G Row, sBit, len = [1, 1], 3 , 1
func (a *UESecurityCapability) GetIA5_5G() (iA5_5G uint8) {
	return a.Buffer[1] & GetBitMask(3, 2) >> (2)
}

// UESecurityCapability 9.11.3.54
// IA5_5G Row, sBit, len = [1, 1], 3 , 1
func (a *UESecurityCapability) SetIA5_5G(iA5_5G uint8) {
	a.Buffer[1] = (a.Buffer[1] & 251) + ((iA5_5G & 1) << 2)
}

// UESecurityCapability 9.11.3.54
// IA6_5G Row, sBit, len = [1, 1], 2 , 1
func (a *UESecurityCapability) GetIA6_5G() (iA6_5G uint8) {
	return a.Buffer[1] & GetBitMask(2, 1) >> (1)
}

// UESecurityCapability 9.11.3.54
// IA6_5G Row, sBit, len = [1, 1], 2 , 1
func (a *UESecurityCapability) SetIA6_5G(iA6_5G uint8) {
	a.Buffer[1] = (a.Buffer[1] & 253) + ((iA6_5G & 1) << 1)
}

// UESecurityCapability 9.11.3.54
// IA7_5G Row, sBit, len = [1, 1], 1 , 1
func (a *UESecurityCapability) GetIA7_5G() (iA7_5G uint8) {
	return a.Buffer[1] & GetBitMask(1, 0)
}

// UESecurityCapability 9.11.3.54
// IA7_5G Row, sBit, len = [1, 1], 1 , 1
func (a *UESecurityCapability) SetIA7_5G(iA7_5G uint8) {
	a.Buffer[1] = (a.Buffer[1] & 254) + (iA7_5G & 1)
}

// UESecurityCapability 9.11.3.54
// EEA0 Row, sBit, len = [2, 2], 8 , 1
func (a *UESecurityCapability) GetEEA0() (eEA0 uint8) {
	return a.Buffer[2] & GetBitMask(8, 7) >> (7)
}

// UESecurityCapability 9.11.3.54
// EEA0 Row, sBit, len = [2, 2], 8 , 1
func (a *UESecurityCapability) SetEEA0(eEA0 uint8) {
	a.Buffer[2] = (a.Buffer[2] & 127) + ((eEA0 & 1) << 7)
}

// UESecurityCapability 9.11.3.54
// EEA1_128 Row, sBit, len = [2, 2], 7 , 1
func (a *UESecurityCapability) GetEEA1_128() (eEA1_128 uint8) {
	return a.Buffer[2] & GetBitMask(7, 6) >> (6)
}

// UESecurityCapability 9.11.3.54
// EEA1_128 Row, sBit, len = [2, 2], 7 , 1
func (a *UESecurityCapability) SetEEA1_128(eEA1_128 uint8) {
	a.Buffer[2] = (a.Buffer[2] & 191) + ((eEA1_128 & 1) << 6)
}

// UESecurityCapability 9.11.3.54
// EEA2_128 Row, sBit, len = [2, 2], 6 , 1
func (a *UESecurityCapability) GetEEA2_128() (eEA2_128 uint8) {
	return a.Buffer[2] & GetBitMask(6, 5) >> (5)
}

// UESecurityCapability 9.11.3.54
// EEA2_128 Row, sBit, len = [2, 2], 6 , 1
func (a *UESecurityCapability) SetEEA2_128(eEA2_128 uint8) {
	a.Buffer[2] = (a.Buffer[2] & 223) + ((eEA2_128 & 1) << 5)
}

// UESecurityCapability 9.11.3.54
// EEA3_128 Row, sBit, len = [2, 2], 5 , 1
func (a *UESecurityCapability) GetEEA3_128() (eEA3_128 uint8) {
	return a.Buffer[2] & GetBitMask(5, 4) >> (4)
}

// UESecurityCapability 9.11.3.54
// EEA3_128 Row, sBit, len = [2, 2], 5 , 1
func (a *UESecurityCapability) SetEEA3_128(eEA3_128 uint8) {
	a.Buffer[2] = (a.Buffer[2] & 239) + ((eEA3_128 & 1) << 4)
}

// UESecurityCapability 9.11.3.54
// EEA4 Row, sBit, len = [2, 2], 4 , 1
func (a *UESecurityCapability) GetEEA4() (eEA4 uint8) {
	return a.Buffer[2] & GetBitMask(4, 3) >> (3)
}

// UESecurityCapability 9.11.3.54
// EEA4 Row, sBit, len = [2, 2], 4 , 1
func (a *UESecurityCapability) SetEEA4(eEA4 uint8) {
	a.Buffer[2] = (a.Buffer[2] & 247) + ((eEA4 & 1) << 3)
}

// UESecurityCapability 9.11.3.54
// EEA5 Row, sBit, len = [2, 2], 3 , 1
func (a *UESecurityCapability) GetEEA5() (eEA5 uint8) {
	return a.Buffer[2] & GetBitMask(3, 2) >> (2)
}

// UESecurityCapability 9.11.3.54
// EEA5 Row, sBit, len = [2, 2], 3 , 1
func (a *UESecurityCapability) SetEEA5(eEA5 uint8) {
	a.Buffer[2] = (a.Buffer[2] & 251) + ((eEA5 & 1) << 2)
}

// UESecurityCapability 9.11.3.54
// EEA6 Row, sBit, len = [2, 2], 2 , 1
func (a *UESecurityCapability) GetEEA6() (eEA6 uint8) {
	return a.Buffer[2] & GetBitMask(2, 1) >> (1)
}

// UESecurityCapability 9.11.3.54
// EEA6 Row, sBit, len = [2, 2], 2 , 1
func (a *UESecurityCapability) SetEEA6(eEA6 uint8) {
	a.Buffer[2] = (a.Buffer[2] & 253) + ((eEA6 & 1) << 1)
}

// UESecurityCapability 9.11.3.54
// EEA7 Row, sBit, len = [2, 2], 1 , 1
func (a *UESecurityCapability) GetEEA7() (eEA7 uint8) {
	return a.Buffer[2] & GetBitMask(1, 0)
}

// UESecurityCapability 9.11.3.54
// EEA7 Row, sBit, len = [2, 2], 1 , 1
func (a *UESecurityCapability) SetEEA7(eEA7 uint8) {
	a.Buffer[2] = (a.Buffer[2] & 254) + (eEA7 & 1)
}

// UESecurityCapability 9.11.3.54
// EIA0 Row, sBit, len = [3, 3], 8 , 1
func (a *UESecurityCapability) GetEIA0() (eIA0 uint8) {
	return a.Buffer[3] & GetBitMask(8, 7) >> (7)
}

// UESecurityCapability 9.11.3.54
// EIA0 Row, sBit, len = [3, 3], 8 , 1
func (a *UESecurityCapability) SetEIA0(eIA0 uint8) {
	a.Buffer[3] = (a.Buffer[3] & 127) + ((eIA0 & 1) << 7)
}

// UESecurityCapability 9.11.3.54
// EIA1_128 Row, sBit, len = [3, 3], 7 , 1
func (a *UESecurityCapability) GetEIA1_128() (eIA1_128 uint8) {
	return a.Buffer[3] & GetBitMask(7, 6) >> (6)
}

// UESecurityCapability 9.11.3.54
// EIA1_128 Row, sBit, len = [3, 3], 7 , 1
func (a *UESecurityCapability) SetEIA1_128(eIA1_128 uint8) {
	a.Buffer[3] = (a.Buffer[3] & 191) + ((eIA1_128 & 1) << 6)
}

// UESecurityCapability 9.11.3.54
// EIA2_128 Row, sBit, len = [3, 3], 6 , 1
func (a *UESecurityCapability) GetEIA2_128() (eIA2_128 uint8) {
	return a.Buffer[3] & GetBitMask(6, 5) >> (5)
}

// UESecurityCapability 9.11.3.54
// EIA2_128 Row, sBit, len = [3, 3], 6 , 1
func (a *UESecurityCapability) SetEIA2_128(eIA2_128 uint8) {
	a.Buffer[3] = (a.Buffer[3] & 223) + ((eIA2_128 & 1) << 5)
}

// UESecurityCapability 9.11.3.54
// EIA3_128 Row, sBit, len = [3, 3], 5 , 1
func (a *UESecurityCapability) GetEIA3_128() (eIA3_128 uint8) {
	return a.Buffer[3] & GetBitMask(5, 4) >> (4)
}

// UESecurityCapability 9.11.3.54
// EIA3_128 Row, sBit, len = [3, 3], 5 , 1
func (a *UESecurityCapability) SetEIA3_128(eIA3_128 uint8) {
	a.Buffer[3] = (a.Buffer[3] & 239) + ((eIA3_128 & 1) << 4)
}

// UESecurityCapability 9.11.3.54
// EIA4 Row, sBit, len = [3, 3], 4 , 1
func (a *UESecurityCapability) GetEIA4() (eIA4 uint8) {
	return a.Buffer[3] & GetBitMask(4, 3) >> (3)
}

// UESecurityCapability 9.11.3.54
// EIA4 Row, sBit, len = [3, 3], 4 , 1
func (a *UESecurityCapability) SetEIA4(eIA4 uint8) {
	a.Buffer[3] = (a.Buffer[3] & 247) + ((eIA4 & 1) << 3)
}

// UESecurityCapability 9.11.3.54
// EIA5 Row, sBit, len = [3, 3], 3 , 1
func (a *UESecurityCapability) GetEIA5() (eIA5 uint8) {
	return a.Buffer[3] & GetBitMask(3, 2) >> (2)
}

// UESecurityCapability 9.11.3.54
// EIA5 Row, sBit, len = [3, 3], 3 , 1
func (a *UESecurityCapability) SetEIA5(eIA5 uint8) {
	a.Buffer[3] = (a.Buffer[3] & 251) + ((eIA5 & 1) << 2)
}

// UESecurityCapability 9.11.3.54
// EIA6 Row, sBit, len = [3, 3], 2 , 1
func (a *UESecurityCapability) GetEIA6() (eIA6 uint8) {
	return a.Buffer[3] & GetBitMask(2, 1) >> (1)
}

// UESecurityCapability 9.11.3.54
// EIA6 Row, sBit, len = [3, 3], 2 , 1
func (a *UESecurityCapability) SetEIA6(eIA6 uint8) {
	a.Buffer[3] = (a.Buffer[3] & 253) + ((eIA6 & 1) << 1)
}

// UESecurityCapability 9.11.3.54
// EIA7 Row, sBit, len = [3, 3], 1 , 1
func (a *UESecurityCapability) GetEIA7() (eIA7 uint8) {
	return a.Buffer[3] & GetBitMask(1, 0)
}

// UESecurityCapability 9.11.3.54
// EIA7 Row, sBit, len = [3, 3], 1 , 1
func (a *UESecurityCapability) SetEIA7(eIA7 uint8) {
	a.Buffer[3] = (a.Buffer[3] & 254) + (eIA7 & 1)
}

// UESecurityCapability 9.11.3.54
// Spare Row, sBit, len = [4, 7], 8 , 32
func (a *UESecurityCapability) GetSpare() (spare [4]uint8) {
	copy(spare[:], a.Buffer[4:8])
	return spare
}

// UESecurityCapability 9.11.3.54
// Spare Row, sBit, len = [4, 7], 8 , 32
func (a *UESecurityCapability) SetSpare(spare [4]uint8) {
	copy(a.Buffer[4:8], spare[:])
}
