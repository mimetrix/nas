package nasType

// Capability5GSM 9.11.4.1
// MH6PDU Row, sBit, len = [0, 0], 2 , 1
// RqoS Row, sBit, len = [0, 0], 1 , 1
// Spare Row, sBit, len = [1, 12], 8 , 96
type Capability5GSM struct {
	Iei   uint8     `json:"-"`
	Len   uint8     `json:"-"`
	Octet [13]uint8 `json:"-"`
    TPMIC bool
    ATSSSST string
    EthernetPDNTypeS1 bool
    MultiHomedIPv6 bool
    ReflectiveQoS bool
    APMQF bool
}

func NewCapability5GSM(iei uint8) (capability5GSM *Capability5GSM) {
	capability5GSM = &Capability5GSM{}
	capability5GSM.SetIei(iei)
	return capability5GSM
}

var SteeringModes = map[uint8]string{
    0: "ATSSS not supported",
    1: "ATSSS Low-Layer functionality with any steering mode supported",
    2: "MPTCP functionality with any steering mode and ATSSS-LL functionality" +
        " with only active-standby steering mode supported",
    3: "MPTCP functionality with any steering mode and ATSSS-LL functionality" +
        "with any steering mode supported",
}

func (c *Capability5GSM ) DecodeNASType() error{
    

    c.TPMIC = ( (c.Octet[2] & 0x80) >> 8) == 1
    c.ATSSSST = SteeringModes[ ((c.Octet[2] & 0x78 ) >> 3) ]
    c.EthernetPDNTypeS1 = ((c.Octet[2] & 0x04) >> 2 )== 1
    c.MultiHomedIPv6 = c.GetMH6PDU() == 1
    c.ReflectiveQoS  = c.GetRqoS() == 1
    c.APMQF = (c.Octet[3] & 0x01) == 1
    return nil
    
}

// Capability5GSM 9.11.4.1
// Iei Row, sBit, len = [], 8, 8
func (a *Capability5GSM) GetIei() (iei uint8) {
	return a.Iei
}

// Capability5GSM 9.11.4.1
// Iei Row, sBit, len = [], 8, 8
func (a *Capability5GSM) SetIei(iei uint8) {
	a.Iei = iei
}

// Capability5GSM 9.11.4.1
// Len Row, sBit, len = [], 8, 8
func (a *Capability5GSM) GetLen() (len uint8) {
	return a.Len
}

// Capability5GSM 9.11.4.1
// Len Row, sBit, len = [], 8, 8
func (a *Capability5GSM) SetLen(len uint8) {
	a.Len = len
}

// Capability5GSM 9.11.4.1
// MH6PDU Row, sBit, len = [0, 0], 2 , 1
func (a *Capability5GSM) GetMH6PDU() (mH6PDU uint8) {
	return a.Octet[0] & GetBitMask(2, 1) >> (1)
}

// Capability5GSM 9.11.4.1
// MH6PDU Row, sBit, len = [0, 0], 2 , 1
func (a *Capability5GSM) SetMH6PDU(mH6PDU uint8) {
	a.Octet[0] = (a.Octet[0] & 253) + ((mH6PDU & 1) << 1)
}

// Capability5GSM 9.11.4.1
// RqoS Row, sBit, len = [0, 0], 1 , 1
func (a *Capability5GSM) GetRqoS() (rqoS uint8) {
	return a.Octet[0] & GetBitMask(1, 0)
}

// Capability5GSM 9.11.4.1
// RqoS Row, sBit, len = [0, 0], 1 , 1
func (a *Capability5GSM) SetRqoS(rqoS uint8) {
	a.Octet[0] = (a.Octet[0] & 254) + (rqoS & 1)
}

// Capability5GSM 9.11.4.1
// Spare Row, sBit, len = [1, 12], 8 , 96
func (a *Capability5GSM) GetSpare() (spare [12]uint8) {
	copy(spare[:], a.Octet[1:13])
	return spare
}

// Capability5GSM 9.11.4.1
// Spare Row, sBit, len = [1, 12], 8 , 96
func (a *Capability5GSM) SetSpare(spare [12]uint8) {
	copy(a.Octet[1:13], spare[:])
}
