package nasType

// AuthenticationRequestMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
type AuthenticationRequestMessageIdentity struct {
	Octet uint8         `json:"-"`
    MessageType string `json:",omitempty"`
}

var MessageTypes = map[uint8]string {
    0x41: "Registration request",
    0x42: "Registration accept",
    0x43: "Registration complete",
    0x44: "Registration reject",
    0x45: "Deregistration request (UE originating)",
    0x46: "Deregistration accept (UE originating)",
    0x47: "Deregistration request (UE terminated)",
    0x48: "Deregistration accept (UE terminated)",
    0x4c: "1Service request",
    0x4d: "Service reject",
    0x4e: "Service accept",
    0x4f: "Control plane service request",
    0x50: "Network slice-specific authentication command",
    0x51: "Network slice-specific authentication complete",
    0x52: "Network slice-specific authentication result",
    0x54: "Configuration update command",
    0x55: "Configuration update complete",
    0x56: "Authentication request",
    0x57: "Authentication response",
    0x58: "Authentication reject",
    0x59: "Authentication failure",
    0x5a: "Authentication result",
    0x5b: "Identity request",
    0x5c: "Identity response",
    0x5d: "Security mode command",
    0x5e: "Security mode complete",
    0x5f: "Security mode reject",
    0x64: "5GMM status",
    0x65: "Notification",
    0x66: "Notification response",
    0x67: "UL NAS transport",
    0x68: "DL NAS transport",
    0x69: "Relay key request",
    0x6a: "Relay key accept",
    0x6b: "Relay key reject",
    0x6c: "Relay authentication request",
    0x6d: "Relay authentication response",
    0x61: "PDU session establishment request",
    0x62: "PDU session establishment accept",
    0x63: "PDU session establishment reject",
    0xc5: "PDU session authentication command",
    0xc6: "PDU session authentication complete",
    0xc7: "PDU session authentication result",
    0xc9: "PDU session modification request",
    0xca: "PDU session modification reject",
    0xcb: "PDU session modification command",
    0xcc: "PDU session modification complete",
    0xcd: "PDU session modification command reject",
    0xd1: "PDU session release request",
    0xd2: "PDU session release reject",
    0xd3: "PDU session release command",
    0xd4: "PDU session release complete",
    0xd6: "5GSM status",
    0xd8: "Service-level authentication command",
    0xd9: "Service-level authentication complete",
    0xda: "Remote UE report",
    0xdb: "Remote UE report response",
}

func (a *AuthenticationRequestMessageIdentity ) DecodeNASType() error{
    a.MessageType = MessageTypes[a.GetMessageType()]

    return nil
}

func NewAuthenticationRequestMessageIdentity() (authenticationRequestMessageIdentity *AuthenticationRequestMessageIdentity) {
	authenticationRequestMessageIdentity = &AuthenticationRequestMessageIdentity{}
	return authenticationRequestMessageIdentity
}

// AuthenticationRequestMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
func (a *AuthenticationRequestMessageIdentity) GetMessageType() (messageType uint8) {
	return a.Octet
}

// AuthenticationRequestMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
func (a *AuthenticationRequestMessageIdentity) SetMessageType(messageType uint8) {
	a.Octet = messageType
}
