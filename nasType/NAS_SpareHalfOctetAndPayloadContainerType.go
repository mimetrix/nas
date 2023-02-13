package nasType

// SpareHalfOctetAndPayloadContainerType 9.11.3.40 9.5
// PayloadContainerType Row, sBit, len = [0, 0], 4 , 4
type SpareHalfOctetAndPayloadContainerType struct {
	Octet uint8 `json:"-"`
    PayloadContainerType string
}

var ContainerTypes  = map[uint8]string {
    1: "N1 SM information",
    2: "SMS",
    3: "LTE Positioning Protocol (LPP) message container",
    4: "SOR transparent container",
    5: "UE policy container",
    6: "UE parameters update transparent container",
    7: "Location services message container (see 3GPP TS 23.273 [6B])",
    8: "CIoT user data container",
    9: "Service-level-AA container",
    10: "Event notification",
    11: "Reserved",
    12: "Reserved",
    13: "Reserved",
    14: "Reserved",
    15: "Reserved",
    16: "Multiple payloads",
}

func (s *SpareHalfOctetAndPayloadContainerType) DecodeNASType() error{
    s.PayloadContainerType = ContainerTypes[s.Octet]
    return nil
}

func NewSpareHalfOctetAndPayloadContainerType() (spareHalfOctetAndPayloadContainerType *SpareHalfOctetAndPayloadContainerType) {
	spareHalfOctetAndPayloadContainerType = &SpareHalfOctetAndPayloadContainerType{}
	return spareHalfOctetAndPayloadContainerType
}

// SpareHalfOctetAndPayloadContainerType 9.11.3.40 9.5
// PayloadContainerType Row, sBit, len = [0, 0], 4 , 4
func (a *SpareHalfOctetAndPayloadContainerType) GetPayloadContainerType() (payloadContainerType uint8) {
	return a.Octet & GetBitMask(4, 0)
}

// SpareHalfOctetAndPayloadContainerType 9.11.3.40 9.5
// PayloadContainerType Row, sBit, len = [0, 0], 4 , 4
func (a *SpareHalfOctetAndPayloadContainerType) SetPayloadContainerType(payloadContainerType uint8) {
	a.Octet = (a.Octet & 240) + (payloadContainerType & 15)
}
