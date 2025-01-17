package nasType

// PDUSESSIONRELEASECOMMANDMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
type PDUSESSIONRELEASECOMMANDMessageIdentity struct {
	Octet       uint8 `json:"0"`
	MessageType string
}

func NewPDUSESSIONRELEASECOMMANDMessageIdentity() (pDUSESSIONRELEASECOMMANDMessageIdentity *PDUSESSIONRELEASECOMMANDMessageIdentity) {
	pDUSESSIONRELEASECOMMANDMessageIdentity = &PDUSESSIONRELEASECOMMANDMessageIdentity{}
	return pDUSESSIONRELEASECOMMANDMessageIdentity
}

func (n *PDUSESSIONRELEASECOMMANDMessageIdentity) DecodeNASType() error {
	n.MessageType = MessageTypes[n.Octet]
	return nil
}

// PDUSESSIONRELEASECOMMANDMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
func (a *PDUSESSIONRELEASECOMMANDMessageIdentity) GetMessageType() (messageType uint8) {
	return a.Octet
}

// PDUSESSIONRELEASECOMMANDMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
func (a *PDUSESSIONRELEASECOMMANDMessageIdentity) SetMessageType(messageType uint8) {
	a.Octet = messageType
}
