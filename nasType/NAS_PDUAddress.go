package nasType

import (
	"errors"
	"net"
)

// PDUAddress 9.11.4.10
// PDUSessionTypeValue Row, sBit, len = [0, 0], 3 , 3
// PDUAddressInformation Row, sBit, len = [1, 12], 8 , 96

type PDUAddress struct {
	Iei            uint8     `json:"-"`
	Len            uint8     `json:"-"`
	Octet          [13]uint8 `json:"-"`
	PDUSessionType uint8     `json:"PDUSessionType,omitempty"`
	IPv4Address    string    `json:"IPv4Address,omitempty"`
	IPv6Address    string    `json:"IPv6Address,omitempty"`
}

func (a *PDUAddress) Parse() error {
	a.PDUSessionType = a.GetPDUSessionTypeValue()
	pduAddressInformation := a.GetPDUAddressInformation()
	var ipv4, ipv6 net.IP
	switch a.PDUSessionType {
	case PDUSessionTypeIPv4:
		ipv4 = pduAddressInformation[0:4]
		a.IPv4Address = ipv4.String()
	case PDUSessionTypeIPv6:
		ipv6 = pduAddressInformation[0:8]
		a.IPv6Address = ipv6.String()
	case PDUSessionTypeIPv4v6:
		ipv6 = pduAddressInformation[0:8]
		ipv4 = pduAddressInformation[8:12]
		a.IPv6Address = ipv6.String()
		a.IPv4Address = ipv4.String()
	default:
		return errors.New("pdu session type not valid")
	}
	return nil
}

func NewPDUAddress(iei uint8) (pDUAddress *PDUAddress) {
	pDUAddress = &PDUAddress{}
	pDUAddress.SetIei(iei)
	return pDUAddress
}

// PDUAddress 9.11.4.10
// Iei Row, sBit, len = [], 8, 8
func (a *PDUAddress) GetIei() (iei uint8) {
	return a.Iei
}

// PDUAddress 9.11.4.10
// Iei Row, sBit, len = [], 8, 8
func (a *PDUAddress) SetIei(iei uint8) {
	a.Iei = iei
}

// PDUAddress 9.11.4.10
// Len Row, sBit, len = [], 8, 8
func (a *PDUAddress) GetLen() (len uint8) {
	return a.Len
}

// PDUAddress 9.11.4.10
// Len Row, sBit, len = [], 8, 8
func (a *PDUAddress) SetLen(len uint8) {
	a.Len = len
}

// PDUAddress 9.11.4.10
// PDUSessionTypeValue Row, sBit, len = [0, 0], 3 , 3
func (a *PDUAddress) GetPDUSessionTypeValue() (pDUSessionTypeValue uint8) {
	return a.Octet[0] & GetBitMask(3, 0)
}

// PDUAddress 9.11.4.10
// PDUSessionTypeValue Row, sBit, len = [0, 0], 3 , 3
func (a *PDUAddress) SetPDUSessionTypeValue(pDUSessionTypeValue uint8) {
	a.Octet[0] = (a.Octet[0] & 248) + (pDUSessionTypeValue & 7)
}

// PDUAddress 9.11.4.10
// PDUAddressInformation Row, sBit, len = [1, 12], 8 , 96
func (a *PDUAddress) GetPDUAddressInformation() (pDUAddressInformation [12]uint8) {
	copy(pDUAddressInformation[:], a.Octet[1:13])
	return pDUAddressInformation
}

// PDUAddress 9.11.4.10
// PDUAddressInformation Row, sBit, len = [1, 12], 8 , 96
func (a *PDUAddress) SetPDUAddressInformation(pDUAddressInformation [12]uint8) {
	copy(a.Octet[1:13], pDUAddressInformation[:])
}
