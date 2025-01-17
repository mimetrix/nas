package nasType

// AlwaysonPDUSessionRequested 9.11.4.4
// Iei Row, sBit, len = [0, 0], 8 , 4
// APSR Row, sBit, len = [0, 0], 1 , 1
type AlwaysonPDUSessionRequested struct {
	Octet uint8 `json:"Octet,omitempty"`
}

func NewAlwaysonPDUSessionRequested(iei uint8) (alwaysonPDUSessionRequested *AlwaysonPDUSessionRequested) {
	alwaysonPDUSessionRequested = &AlwaysonPDUSessionRequested{}
	alwaysonPDUSessionRequested.SetIei(iei)
	return alwaysonPDUSessionRequested
}

// AlwaysonPDUSessionRequested 9.11.4.4
// Iei Row, sBit, len = [0, 0], 8 , 4
func (a *AlwaysonPDUSessionRequested) GetIei() (iei uint8) {
	return a.Octet & GetBitMask(8, 4) >> (4)
}

// AlwaysonPDUSessionRequested 9.11.4.4
// Iei Row, sBit, len = [0, 0], 8 , 4
func (a *AlwaysonPDUSessionRequested) SetIei(iei uint8) {
	a.Octet = (a.Octet & 15) + ((iei & 15) << 4)
}

// AlwaysonPDUSessionRequested 9.11.4.4
// APSR Row, sBit, len = [0, 0], 1 , 1
func (a *AlwaysonPDUSessionRequested) GetAPSR() (aPSR uint8) {
	return a.Octet & GetBitMask(1, 0)
}

// AlwaysonPDUSessionRequested 9.11.4.4
// APSR Row, sBit, len = [0, 0], 1 , 1
func (a *AlwaysonPDUSessionRequested) SetAPSR(aPSR uint8) {
	a.Octet = (a.Octet & 254) + (aPSR & 1)
}
