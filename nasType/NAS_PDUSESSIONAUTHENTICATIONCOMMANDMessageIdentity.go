package nasType

// PDUSESSIONAUTHENTICATIONCOMMANDMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
type PDUSESSIONAUTHENTICATIONCOMMANDMessageIdentity struct {
	Octet uint8 `json:"Octet,omitempty"`
    MessageType string
}

func (p *PDUSESSIONAUTHENTICATIONCOMMANDMessageIdentity) DecodeNASType() error{
    p.MessageType = MessageTypes[p.GetMessageType()]
    return nil
}

func NewPDUSESSIONAUTHENTICATIONCOMMANDMessageIdentity() (pDUSESSIONAUTHENTICATIONCOMMANDMessageIdentity *PDUSESSIONAUTHENTICATIONCOMMANDMessageIdentity) {
	pDUSESSIONAUTHENTICATIONCOMMANDMessageIdentity = &PDUSESSIONAUTHENTICATIONCOMMANDMessageIdentity{}
	return pDUSESSIONAUTHENTICATIONCOMMANDMessageIdentity
}

// PDUSESSIONAUTHENTICATIONCOMMANDMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
func (a *PDUSESSIONAUTHENTICATIONCOMMANDMessageIdentity) GetMessageType() (messageType uint8) {
	return a.Octet
}

// PDUSESSIONAUTHENTICATIONCOMMANDMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
func (a *PDUSESSIONAUTHENTICATIONCOMMANDMessageIdentity) SetMessageType(messageType uint8) {
	a.Octet = messageType
}
