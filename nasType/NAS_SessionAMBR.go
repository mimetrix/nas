package nasType

// SessionAMBR 9.11.4.14
// UnitForSessionAMBRForDownlink Row, sBit, len = [0, 0], 8 , 8
// SessionAMBRForDownlink Row, sBit, len = [1, 2], 8 , 16
// UnitForSessionAMBRForUplink Row, sBit, len = [3, 3], 8 , 8
// SessionAMBRForUplink Row, sBit, len = [4, 5], 8 , 16
type SessionAMBR struct {
	Iei   uint8    `json:"-"`
	Len   uint8    `json:"-"`
	Octet [6]uint8 `json:"-"`
    DownlinkUnits string
    DownlinkAMBR uint16 
    UplinkUnits string
    UplinkAMBR uint16
}

var UnitMap = map[uint8]string{
    1:"1Kbps",
    2:"4Kbps",
    3:"16Kbps",
    4:"64Kbps",
    5:"256Kbps",
    6:"1Mbps",
    7:"4Mbps",
    8:"16Mbps",
    9:"64Mbps",
    10:"256Mbps",
    11:"1Gbps",
    12:"4Gbps",
    13:"16Gbps",
    14:"64Gbps",
    15:"256Gbps",
    16:"1Tbps",
    17:"4Tbps",
    18:"16Tbps",
    19:"64Tbps",
    20:"256Tbps",
    21:"1Pbps",
    22:"4Pbps",
    23:"16Pbps",
    24:"64Pbps",
    25:"256Pbps",
}


func (s *SessionAMBR) DecodeNASType() error{
    s.DownlinkUnits = UnitMap[s.GetUnitForSessionAMBRForDownlink()]
    s.UplinkUnits = UnitMap[s.GetUnitForSessionAMBRForUplink()]

    DLAMBR := s.GetSessionAMBRForDownlink()
    ULAMBR := s.GetSessionAMBRForUplink()
    
    s.DownlinkAMBR = calculateAMBR(DLAMBR)
    s.UplinkAMBR = calculateAMBR(ULAMBR)
    return nil
}

func calculateAMBR(buf [2]uint8) uint16{
    AMBR := ( uint16(buf[0] & 0xf0) >> 4) * 1000
    AMBR = AMBR + uint16(buf[0] & 0xf) * 100
    AMBR = ( uint16(buf[1] & 0xf0) >> 4) * 1000
    AMBR = AMBR + uint16(buf[1] & 0xf) 
    return AMBR
}

func NewSessionAMBR(iei uint8) (sessionAMBR *SessionAMBR) {
	sessionAMBR = &SessionAMBR{}
	sessionAMBR.SetIei(iei)
	return sessionAMBR
}

// SessionAMBR 9.11.4.14
// Iei Row, sBit, len = [], 8, 8
func (a *SessionAMBR) GetIei() (iei uint8) {
	return a.Iei
}

// SessionAMBR 9.11.4.14
// Iei Row, sBit, len = [], 8, 8
func (a *SessionAMBR) SetIei(iei uint8) {
	a.Iei = iei
}

// SessionAMBR 9.11.4.14
// Len Row, sBit, len = [], 8, 8
func (a *SessionAMBR) GetLen() (len uint8) {
	return a.Len
}

// SessionAMBR 9.11.4.14
// Len Row, sBit, len = [], 8, 8
func (a *SessionAMBR) SetLen(len uint8) {
	a.Len = len
}

// SessionAMBR 9.11.4.14
// UnitForSessionAMBRForDownlink Row, sBit, len = [0, 0], 8 , 8
func (a *SessionAMBR) GetUnitForSessionAMBRForDownlink() (unitForSessionAMBRForDownlink uint8) {
	return a.Octet[0]
}

// SessionAMBR 9.11.4.14
// UnitForSessionAMBRForDownlink Row, sBit, len = [0, 0], 8 , 8
func (a *SessionAMBR) SetUnitForSessionAMBRForDownlink(unitForSessionAMBRForDownlink uint8) {
	a.Octet[0] = unitForSessionAMBRForDownlink
}

// SessionAMBR 9.11.4.14
// SessionAMBRForDownlink Row, sBit, len = [1, 2], 8 , 16
func (a *SessionAMBR) GetSessionAMBRForDownlink() (sessionAMBRForDownlink [2]uint8) {
	copy(sessionAMBRForDownlink[:], a.Octet[1:3])
	return sessionAMBRForDownlink
}

// SessionAMBR 9.11.4.14
// SessionAMBRForDownlink Row, sBit, len = [1, 2], 8 , 16
func (a *SessionAMBR) SetSessionAMBRForDownlink(sessionAMBRForDownlink [2]uint8) {
	copy(a.Octet[1:3], sessionAMBRForDownlink[:])
}

// SessionAMBR 9.11.4.14
// UnitForSessionAMBRForUplink Row, sBit, len = [3, 3], 8 , 8
func (a *SessionAMBR) GetUnitForSessionAMBRForUplink() (unitForSessionAMBRForUplink uint8) {
	return a.Octet[3]
}

// SessionAMBR 9.11.4.14
// UnitForSessionAMBRForUplink Row, sBit, len = [3, 3], 8 , 8
func (a *SessionAMBR) SetUnitForSessionAMBRForUplink(unitForSessionAMBRForUplink uint8) {
	a.Octet[3] = unitForSessionAMBRForUplink
}

// SessionAMBR 9.11.4.14
// SessionAMBRForUplink Row, sBit, len = [4, 5], 8 , 16
func (a *SessionAMBR) GetSessionAMBRForUplink() (sessionAMBRForUplink [2]uint8) {
	copy(sessionAMBRForUplink[:], a.Octet[4:6])
	return sessionAMBRForUplink
}

// SessionAMBR 9.11.4.14
// SessionAMBRForUplink Row, sBit, len = [4, 5], 8 , 16
func (a *SessionAMBR) SetSessionAMBRForUplink(sessionAMBRForUplink [2]uint8) {
	copy(a.Octet[4:6], sessionAMBRForUplink[:])
}
