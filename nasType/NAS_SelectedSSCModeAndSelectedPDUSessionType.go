package nasType

import "errors"

// SelectedSSCModeAndSelectedPDUSessionType 9.11.4.11 9.11.4.16
// SSCMode Row, sBit, len = [0, 0], 7 , 3
// PDUSessionType  Row, sBit, len = [0, 0], 3 , 3
const (
	SSCMode1  = 0b001
	SSCMode2  = 0b010
	SSCMode3  = 0b011
	SSCMode1U = 0b100
	SSCMode2U = 0b101
	SSCMode3U = 0b110
	// TODO create a const file to put all the values.
	// PDUSessionTypeIPv4         = 0b001
	// PDUSessionTypeIPv6         = 0b010
	// PDUSessionTypeIPv4v6       = 0b011
	PDUSessionTypeUnstructered = 0b100
	PDUSessionTypeEthernet     = 0b101
	PDUSessionTypeReserved     = 0b111
)

type SelectedSSCModeAndSelectedPDUSessionType struct {
	Octet          uint8  `json:"-"`
	SSCMode        string `json:"SSCMode,omitempty"`
	PDUSessionType string `json:"PDUSessionType,omitempty"`
}

func (a *SelectedSSCModeAndSelectedPDUSessionType) Parse() error {
	switch a.GetSSCMode() {
	case SSCMode1:
		a.SSCMode = "SSCMode1"
	case SSCMode2:
		a.SSCMode = "SSCMode2"
	case SSCMode3:
		a.SSCMode = "SSCMode3"
	case SSCMode1U:
		a.SSCMode = "SSCMode1"
	case SSCMode2U:
		a.SSCMode = "SSCMode2"
	case SSCMode3U:
		a.SSCMode = "SSCMode3"
	default:
		return errors.New("invalid scc mode")
	}

	switch a.GetPDUSessionType() {
	case PDUSessionTypeIPv4:
		a.PDUSessionType = "IPv4"
	case PDUSessionTypeIPv6:
		a.PDUSessionType = "IPv6"
	case PDUSessionTypeIPv4v6:
		a.PDUSessionType = "IPv4v6"
	case PDUSessionTypeUnstructered:
		a.PDUSessionType = "Unstructered"
	case PDUSessionTypeEthernet:
		a.PDUSessionType = "Ethernet"
	case PDUSessionTypeReserved:
		a.PDUSessionType = "Reserved"
	default:
		a.PDUSessionType = "IPv4v6"
	}

	return nil
}

func NewSelectedSSCModeAndSelectedPDUSessionType() (selectedSSCModeAndSelectedPDUSessionType *SelectedSSCModeAndSelectedPDUSessionType) {
	selectedSSCModeAndSelectedPDUSessionType = &SelectedSSCModeAndSelectedPDUSessionType{}
	return selectedSSCModeAndSelectedPDUSessionType
}

// SelectedSSCModeAndSelectedPDUSessionType 9.11.4.11 9.11.4.16
// SSCMode Row, sBit, len = [0, 0], 7 , 3
func (a *SelectedSSCModeAndSelectedPDUSessionType) GetSSCMode() (sSCMode uint8) {
	return a.Octet & GetBitMask(7, 4) >> (4)
}

// SelectedSSCModeAndSelectedPDUSessionType 9.11.4.11 9.11.4.16
// SSCMode Row, sBit, len = [0, 0], 7 , 3
func (a *SelectedSSCModeAndSelectedPDUSessionType) SetSSCMode(sSCMode uint8) {
	a.Octet = (a.Octet & 143) + ((sSCMode & 7) << 4)
}

// SelectedSSCModeAndSelectedPDUSessionType 9.11.4.11 9.11.4.16
// PDUSessionType Row, sBit, len = [0, 0], 3 , 3
func (a *SelectedSSCModeAndSelectedPDUSessionType) GetPDUSessionType() (pDUSessionType uint8) {
	return a.Octet & GetBitMask(3, 0)
}

// SelectedSSCModeAndSelectedPDUSessionType 9.11.4.11 9.11.4.16
// PDUSessionType Row, sBit, len = [0, 0], 3 , 3
func (a *SelectedSSCModeAndSelectedPDUSessionType) SetPDUSessionType(pDUSessionType uint8) {
	a.Octet = (a.Octet & 248) + (pDUSessionType & 7)
}
