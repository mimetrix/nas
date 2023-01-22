package nasType

// ExtendedProtocolDiscriminator 9.2
// ExtendedProtocolDiscriminator Row, sBit, len = [0, 0], 8 , 8
type ExtendedProtocolDiscriminator struct {
	Octet uint8 `json:"-"`
    EPD string  `json:"EPD,omitempty"`
}

/*From 3GPP 24.007 11.2.3.1.1A ver 17.5*/
var EPDTypes = map[uint8]string {
    
    0x7: "reserved",
    0x1e: "reserved",
    0x2e: "5GS session management messages",
    0x3e: "reserved",
    0x4e: "reserved",
    0x5e: "reserved",
    0x6e: "reserved",
    0x7e: "5GS mobility management messages",
    0x8e: "reserved",
    0x9e: "reserved",
    0xae: "reserved",
    0xbe: "reserved",
    0xce: "reserved",
    0xde: "reserved",
    0xee: "reserved",
    0xfe: "reserved",
}

func (epd *ExtendedProtocolDiscriminator) DecodeNASType() error{
    epd.EPD = EPDTypes[epd.Octet]
    return nil
}

func NewExtendedProtocolDiscriminator() (extendedProtocolDiscriminator *ExtendedProtocolDiscriminator) {
	extendedProtocolDiscriminator = &ExtendedProtocolDiscriminator{}
	return extendedProtocolDiscriminator
}

// ExtendedProtocolDiscriminator 9.2
// ExtendedProtocolDiscriminator Row, sBit, len = [0, 0], 8 , 8
func (a *ExtendedProtocolDiscriminator) GetExtendedProtocolDiscriminator() (extendedProtocolDiscriminator uint8) {
	return a.Octet
}

// ExtendedProtocolDiscriminator 9.2
// ExtendedProtocolDiscriminator Row, sBit, len = [0, 0], 8 , 8
func (a *ExtendedProtocolDiscriminator) SetExtendedProtocolDiscriminator(extendedProtocolDiscriminator uint8) {
	a.Octet = extendedProtocolDiscriminator
}
