package nasMessage

import (
	"fmt"

	"github.com/mimetrix/nas/nasType"
)

// NASMessageContainer 9.11.3.33
// NASMessageContainerContents Row, sBit, len = [0, 0], 8 , INF
/*
NAS message container contents (octet 4 to octet n); Max value of 65535 octets
This IE can contain
    a REGISTRATION REQUEST message as defined in subclause 5.5.1 (pg 286), or
    a SERVICE REQUEST message as defined in subclause 5.6.1, or
    non-cleartext IEs of a CONTROL PLANE SERVICE REQUEST message as defined in subclause 5.6.1.

    This structure is technically a NAS "Type", as defined in the specs and therefore
    technically belongs in the nasType package/directory. However, making it a part of
    the nasMessage package/directory solves the circular dependency problem without
    introducing a third package
*/
type NASMessageContainer struct {
	Iei         uint8   `json:"-"`
	Len         uint16  `json:"-"`
	Buffer      []uint8 `json:"-"`
	MessageID   uint8   `json:"-"`
	MessageType string
	NASMessage  *Message
}

func (n *NASMessageContainer) DecodeNASType() error {
	messageID, err := n.GetMessageID()
	if err != nil {
		return err
	}

	messageType, err := n.GetMessageName()
	if err != nil {
		return err
	}

	n.MessageID = messageID
	n.MessageType = messageType

	n.NASMessage = NewMessage()
	err = n.NASMessage.PlainNasDecode(&n.Buffer)
	if err != nil {
		return err
	}

	return nil

}

func (n *NASMessageContainer) GetMessageID() (uint8, error) {

	if len(n.Buffer) > 3 {
		return n.Buffer[2], nil
	} else {
		return 0, fmt.Errorf("Buffer in NASMessageContainer is too short\n")
	}
}

func (n *NASMessageContainer) GetMessageName() (string, error) {

	msgID, err := n.GetMessageID()
	if err != nil {
		return "", err
	}

	//return message type if valid to be inside NAS Message Container
	msgName, ok := nasType.MessageTypes[msgID]
	if ok {
		switch msgID {
		case MsgTypeRegistrationRequest, MsgTypeServiceRequest /*, 0x4f*/ :
			return msgName, nil
		default:
			return "", fmt.Errorf("NS Message Container can't contain %s.", msgName)
		}
	} else {
		return "", fmt.Errorf("Message type, %d, does not exist", msgID)
	}

}

func NewNASMessageContainer(iei uint8) (nASMessageContainer *NASMessageContainer) {
	nASMessageContainer = &NASMessageContainer{}
	nASMessageContainer.SetIei(iei)
	return nASMessageContainer
}

// NASMessageContainer 9.11.3.33
// Iei Row, sBit, len = [], 8, 8
func (a *NASMessageContainer) GetIei() (iei uint8) {
	return a.Iei
}

// NASMessageContainer 9.11.3.33
// Iei Row, sBit, len = [], 8, 8
func (a *NASMessageContainer) SetIei(iei uint8) {
	a.Iei = iei
}

// NASMessageContainer 9.11.3.33
// Len Row, sBit, len = [], 8, 16
func (a *NASMessageContainer) GetLen() (len uint16) {
	return a.Len
}

// NASMessageContainer 9.11.3.33
// Len Row, sBit, len = [], 8, 16
func (a *NASMessageContainer) SetLen(len uint16) {
	a.Len = len
	a.Buffer = make([]uint8, a.Len)
}

// NASMessageContainer 9.11.3.33
// NASMessageContainerContents Row, sBit, len = [0, 0], 8 , INF
func (a *NASMessageContainer) GetNASMessageContainerContents() (nASMessageContainerContents []uint8) {
	nASMessageContainerContents = make([]uint8, len(a.Buffer))
	copy(nASMessageContainerContents, a.Buffer)
	return nASMessageContainerContents
}

// NASMessageContainer 9.11.3.33
// NASMessageContainerContents Row, sBit, len = [0, 0], 8 , INF
func (a *NASMessageContainer) SetNASMessageContainerContents(nASMessageContainerContents []uint8) {
	copy(a.Buffer, nASMessageContainerContents)
}
