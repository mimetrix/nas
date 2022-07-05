package nasMessage

import (
	"bytes"
	"encoding/binary"

	"github.com/free5gc/nas/nasType"
)

type PDUSessionAuthenticationComplete struct {
	nasType.ExtendedProtocolDiscriminator                   `json:"ExtendedProtocolDiscriminator,omitempty"`
	nasType.PDUSessionID                                    `json:"PDUSessionID,omitempty"`
	nasType.PTI                                             `json:"PTI,omitempty"`
	nasType.PDUSESSIONAUTHENTICATIONCOMPLETEMessageIdentity `json:"PDUSESSIONAUTHENTICATIONCOMPLETEMessageIdentity,omitempty"`
	nasType.EAPMessage                                      `json:"EAPMessage,omitempty"`
	*nasType.ExtendedProtocolConfigurationOptions           `json:"ExtendedProtocolConfigurationOptions,omitempty"`
}

func NewPDUSessionAuthenticationComplete(iei uint8) (pDUSessionAuthenticationComplete *PDUSessionAuthenticationComplete) {
	pDUSessionAuthenticationComplete = &PDUSessionAuthenticationComplete{}
	return pDUSessionAuthenticationComplete
}

const (
	PDUSessionAuthenticationCompleteExtendedProtocolConfigurationOptionsType uint8 = 0x7B
)

func (a *PDUSessionAuthenticationComplete) EncodePDUSessionAuthenticationComplete(buffer *bytes.Buffer) {
	binary.Write(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet)
	binary.Write(buffer, binary.BigEndian, &a.PDUSessionID.Octet)
	binary.Write(buffer, binary.BigEndian, &a.PTI.Octet)
	binary.Write(buffer, binary.BigEndian, &a.PDUSESSIONAUTHENTICATIONCOMPLETEMessageIdentity.Octet)
	binary.Write(buffer, binary.BigEndian, a.EAPMessage.GetLen())
	binary.Write(buffer, binary.BigEndian, &a.EAPMessage.Buffer)
	if a.ExtendedProtocolConfigurationOptions != nil {
		binary.Write(buffer, binary.BigEndian, a.ExtendedProtocolConfigurationOptions.GetIei())
		binary.Write(buffer, binary.BigEndian, a.ExtendedProtocolConfigurationOptions.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.ExtendedProtocolConfigurationOptions.Buffer)
	}
}

func (a *PDUSessionAuthenticationComplete) DecodePDUSessionAuthenticationComplete(byteArray *[]byte) {
	buffer := bytes.NewBuffer(*byteArray)
	binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet)
	binary.Read(buffer, binary.BigEndian, &a.PDUSessionID.Octet)
	binary.Read(buffer, binary.BigEndian, &a.PTI.Octet)
	binary.Read(buffer, binary.BigEndian, &a.PDUSESSIONAUTHENTICATIONCOMPLETEMessageIdentity.Octet)
	binary.Read(buffer, binary.BigEndian, &a.EAPMessage.Len)
	a.EAPMessage.SetLen(a.EAPMessage.GetLen())
	binary.Read(buffer, binary.BigEndian, &a.EAPMessage.Buffer)
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
		case PDUSessionAuthenticationCompleteExtendedProtocolConfigurationOptionsType:
			a.ExtendedProtocolConfigurationOptions = nasType.NewExtendedProtocolConfigurationOptions(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolConfigurationOptions.Len)
			a.ExtendedProtocolConfigurationOptions.SetLen(a.ExtendedProtocolConfigurationOptions.GetLen())
			binary.Read(buffer, binary.BigEndian, a.ExtendedProtocolConfigurationOptions.Buffer[:a.ExtendedProtocolConfigurationOptions.GetLen()])
		default:
		}
	}
}
