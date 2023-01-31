package nasType

// AuthenticationRequestMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
type AuthenticationRequestMessageIdentity struct {
	Octet uint8         `json:"-"`
    MessageType string `json:",omitempty"`
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
