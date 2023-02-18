package nasMessage

import (
	//"github.com/mimetrix/nas/nasType"
)
// PayloadContainer 9.11.3.39
// PayloadContainerContents Row, sBit, len = [0, 0], 8 , INF

/*
From 24.501 section 9.11.3.39
If the payload container type is set to "N1 SM information" and is included in the UL
NAS TRANSPORT or DL NAS TRANSPORT message, the payload container contents
contain a 5GSM message as defined in subclause 8.3.


This structure is technically a NAS "Type", as defined in the specs and therefore
technically belongs in the nasType package/directory. However, making it a part of
the nasMessage package/directory solves the circular dependency problem without
introducing a third package

*/
type PayloadContainer struct {
	Iei    uint8   `json:"-"`
	Len    uint16  `json:"-"`
	Buffer []uint8 `json:"-"`
    NASMessage  *Message
}

func (p *PayloadContainer) DecodeNASType() error{
    p.NASMessage = NewMessage()
    err := p.NASMessage.PlainNasDecode(&p.Buffer)
    if err != nil {
        return err
    }
 
    return nil
}

func NewPayloadContainer(iei uint8) (payloadContainer *PayloadContainer) {
	payloadContainer = &PayloadContainer{}
	payloadContainer.SetIei(iei)
	return payloadContainer
}

// PayloadContainer 9.11.3.39
// Iei Row, sBit, len = [], 8, 8
func (a *PayloadContainer) GetIei() (iei uint8) {
	return a.Iei
}

// PayloadContainer 9.11.3.39
// Iei Row, sBit, len = [], 8, 8
func (a *PayloadContainer) SetIei(iei uint8) {
	a.Iei = iei
}

// PayloadContainer 9.11.3.39
// Len Row, sBit, len = [], 8, 16
func (a *PayloadContainer) GetLen() (len uint16) {
	return a.Len
}

// PayloadContainer 9.11.3.39
// Len Row, sBit, len = [], 8, 16
func (a *PayloadContainer) SetLen(len uint16) {
	a.Len = len
	a.Buffer = make([]uint8, a.Len)
}

// PayloadContainer 9.11.3.39
// PayloadContainerContents Row, sBit, len = [0, 0], 8 , INF
func (a *PayloadContainer) GetPayloadContainerContents() (payloadContainerContents []uint8) {
	payloadContainerContents = make([]uint8, len(a.Buffer))
	copy(payloadContainerContents, a.Buffer)
	return payloadContainerContents
}

// PayloadContainer 9.11.3.39
// PayloadContainerContents Row, sBit, len = [0, 0], 8 , INF
func (a *PayloadContainer) SetPayloadContainerContents(payloadContainerContents []uint8) {
	copy(a.Buffer, payloadContainerContents)
}
