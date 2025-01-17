package nasType

// RegistrationAcceptMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
type RegistrationAcceptMessageIdentity struct {
	Octet uint8 `json:"-"`
    MessageType string `json:",omitempty"`
}

func (r *RegistrationAcceptMessageIdentity ) DecodeNASType() error {
    r.MessageType = MessageTypes[r.Octet]
    return nil
}



func NewRegistrationAcceptMessageIdentity() (registrationAcceptMessageIdentity *RegistrationAcceptMessageIdentity) {
	registrationAcceptMessageIdentity = &RegistrationAcceptMessageIdentity{}
	return registrationAcceptMessageIdentity
}

// RegistrationAcceptMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
func (a *RegistrationAcceptMessageIdentity) GetMessageType() (messageType uint8) {
	return a.Octet
}

// RegistrationAcceptMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
func (a *RegistrationAcceptMessageIdentity) SetMessageType(messageType uint8) {
	a.Octet = messageType
}
