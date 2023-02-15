package nasMessage

import (
	"bytes"
	"encoding/binary"
	"github.com/mimetrix/nas/nasType"
)

type SecurityProtected5GSNASMessage struct {
	nasType.ExtendedProtocolDiscriminator       `json:"ExtendedProtocolDiscriminator,omitempty"`
	nasType.SpareHalfOctetAndSecurityHeaderType `json:"SpareHalfOctetAndSecurityHeaderType,omitempty"`
	nasType.MessageAuthenticationCode           `json:"MessageAuthenticationCode,omitempty"`
	nasType.SequenceNumber                      `json:"SequenceNumber"`
	//nasType.Plain5GSNASMessage                  `json:"Plain5GSNASMessage,omitempty"`
    PlainNASMessage *Message
}

func NewSecurityProtected5GSNASMessage(iei uint8) (securityProtected5GSNASMessage *SecurityProtected5GSNASMessage) {
	securityProtected5GSNASMessage = &SecurityProtected5GSNASMessage{}
	return securityProtected5GSNASMessage
}

func (a *SecurityProtected5GSNASMessage) EncodeSecurityProtected5GSNASMessage(buffer *bytes.Buffer) {
	binary.Write(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet)
	binary.Write(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet)
	binary.Write(buffer, binary.BigEndian, &a.MessageAuthenticationCode.Octet)
	binary.Write(buffer, binary.BigEndian, &a.SequenceNumber.Octet)
	//binary.Write(buffer, binary.BigEndian, &a.Plain5GSNASMessage)
	//binary.Write(buffer, binary.BigEndian, &a.PlainNASMessage)
    
}

func (a *SecurityProtected5GSNASMessage) DecodeSecurityProtected5GSNASMessage(byteArray *[]byte) {

	buffer := bytes.NewBuffer(*byteArray)
	binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet)
    a.ExtendedProtocolDiscriminator.DecodeNASType()
    
	binary.Read(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet)
    a.SpareHalfOctetAndSecurityHeaderType.DecodeNASType()

	binary.Read(buffer, binary.BigEndian, &a.MessageAuthenticationCode.Octet)
	binary.Read(buffer, binary.BigEndian, &a.SequenceNumber.Octet)
    
    //spew.Dump(buffer.Bytes())
    b := buffer.Bytes()
    a.PlainNASMessage = &Message{}
    a.PlainNASMessage.PlainNasDecode(&b)

	for buffer.Len() > 0 {
		var ieiN uint8
		var tmpIeiN uint8
		binary.Read(buffer, binary.BigEndian, &ieiN)
		// fmt.Println(ieiN)
		if ieiN >= 0x80 {
			tmpIeiN = (ieiN & 0xf0) >> 4
		} else {
			tmpIeiN = ieiN
		}
		// fmt.Println("type", tmpIeiN)
		switch tmpIeiN {
		default:
		}
	}
}
