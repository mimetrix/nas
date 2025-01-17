package nasType

// STATUSMessageIdentity5GMM 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
type STATUSMessageIdentity5GMM struct {
	Octet uint8 `json:"Octet,omitempty"`
}

func NewSTATUSMessageIdentity5GMM() (sTATUSMessageIdentity5GMM *STATUSMessageIdentity5GMM) {
	sTATUSMessageIdentity5GMM = &STATUSMessageIdentity5GMM{}
	return sTATUSMessageIdentity5GMM
}

// STATUSMessageIdentity5GMM 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
func (a *STATUSMessageIdentity5GMM) GetMessageType() (messageType uint8) {
	return a.Octet
}

// STATUSMessageIdentity5GMM 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
func (a *STATUSMessageIdentity5GMM) SetMessageType(messageType uint8) {
	a.Octet = messageType
}
