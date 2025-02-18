package nasType

// ServiceRequestMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
type ServiceRequestMessageIdentity struct {
	Octet uint8 `json:"Octet,omitempty"`
}

func NewServiceRequestMessageIdentity() (serviceRequestMessageIdentity *ServiceRequestMessageIdentity) {
	serviceRequestMessageIdentity = &ServiceRequestMessageIdentity{}
	return serviceRequestMessageIdentity
}

// ServiceRequestMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
func (a *ServiceRequestMessageIdentity) GetMessageType() (messageType uint8) {
	return a.Octet
}

// ServiceRequestMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
func (a *ServiceRequestMessageIdentity) SetMessageType(messageType uint8) {
	a.Octet = messageType
}
