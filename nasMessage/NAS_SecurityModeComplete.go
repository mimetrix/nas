package nasMessage

import (
	"bytes"
	"encoding/binary"

	"github.com/mimetrix/nas/nasType"
)

type SecurityModeComplete struct {
	nasType.ExtendedProtocolDiscriminator       `json:"ExtendedProtocolDiscriminator,omitempty"`
	nasType.SpareHalfOctetAndSecurityHeaderType `json:"SpareHalfOctetAndSecurityHeaderType,omitempty"`
	nasType.SecurityModeCompleteMessageIdentity `json:"SecurityModeCompleteMessageIdentity,omitempty"`
	*nasType.IMEISV                             `json:"IMEISV,omitempty"`
	*NASMessageContainer                `json:"NASMessageContainer,omitempty"`
}

func NewSecurityModeComplete(iei uint8) (securityModeComplete *SecurityModeComplete) {
	securityModeComplete = &SecurityModeComplete{}
	return securityModeComplete
}

const (
	SecurityModeCompleteIMEISVType              uint8 = 0x77
	SecurityModeCompleteNASMessageContainerType uint8 = 0x71
)

func (a *SecurityModeComplete) EncodeSecurityModeComplete(buffer *bytes.Buffer) {
	binary.Write(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet)
	binary.Write(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet)
	binary.Write(buffer, binary.BigEndian, &a.SecurityModeCompleteMessageIdentity.Octet)
	if a.IMEISV != nil {
		binary.Write(buffer, binary.BigEndian, a.IMEISV.GetIei())
		binary.Write(buffer, binary.BigEndian, a.IMEISV.GetLen())
		binary.Write(buffer, binary.BigEndian, a.IMEISV.Octet[:a.IMEISV.GetLen()])
	}
	if a.NASMessageContainer != nil {
		binary.Write(buffer, binary.BigEndian, a.NASMessageContainer.GetIei())
		binary.Write(buffer, binary.BigEndian, a.NASMessageContainer.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.NASMessageContainer.Buffer)
	}
}

func (a *SecurityModeComplete) DecodeSecurityModeComplete(byteArray *[]byte) {
	buffer := bytes.NewBuffer(*byteArray)
	binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet)
    a.ExtendedProtocolDiscriminator.DecodeNASType()
	
    binary.Read(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet)
    a.SpareHalfOctetAndSecurityHeaderType.DecodeNASType()

	binary.Read(buffer, binary.BigEndian, &a.SecurityModeCompleteMessageIdentity.Octet)
    a.SecurityModeCompleteMessageIdentity.DecodeNASType()
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
		case SecurityModeCompleteIMEISVType:
			a.IMEISV = nasType.NewIMEISV(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.IMEISV.Len)
			a.IMEISV.SetLen(a.IMEISV.GetLen())
			binary.Read(buffer, binary.BigEndian, a.IMEISV.Octet[:a.IMEISV.GetLen()])
            a.IMEISV.DecodeNASType()
		case SecurityModeCompleteNASMessageContainerType:
			a.NASMessageContainer = NewNASMessageContainer(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.NASMessageContainer.Len)
			a.NASMessageContainer.SetLen(a.NASMessageContainer.GetLen())
			binary.Read(buffer, binary.BigEndian, a.NASMessageContainer.Buffer[:a.NASMessageContainer.GetLen()])

            a.NASMessageContainer.DecodeNASType()
		default:
		}
	}
}
