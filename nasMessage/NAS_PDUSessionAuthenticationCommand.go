package nasMessage

import (
	"bytes"
	"encoding/binary"

	"github.com/free5gc/nas/nasType"
)

type PDUSessionAuthenticationCommand struct {
	nasType.ExtendedProtocolDiscriminator                  `json:"ExtendedProtocolDiscriminator,omitempty"`
	nasType.PDUSessionID                                   `json:"PDUSessionID,omitempty"`
	nasType.PTI                                            `json:"PTI,omitempty"`
	nasType.PDUSESSIONAUTHENTICATIONCOMMANDMessageIdentity `json:"PDUSESSIONAUTHENTICATIONCOMMANDMessageIdentity,omitempty"`
	nasType.EAPMessage                                     `json:"EAPMessage,omitempty"`
	*nasType.ExtendedProtocolConfigurationOptions          `json:"ExtendedProtocolConfigurationOptions,omitempty"`
}

func NewPDUSessionAuthenticationCommand(iei uint8) (pDUSessionAuthenticationCommand *PDUSessionAuthenticationCommand) {
	pDUSessionAuthenticationCommand = &PDUSessionAuthenticationCommand{}
	return pDUSessionAuthenticationCommand
}

const (
	PDUSessionAuthenticationCommandExtendedProtocolConfigurationOptionsType uint8 = 0x7B
)

func (a *PDUSessionAuthenticationCommand) EncodePDUSessionAuthenticationCommand(buffer *bytes.Buffer) {
	binary.Write(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet)
	binary.Write(buffer, binary.BigEndian, &a.PDUSessionID.Octet)
	binary.Write(buffer, binary.BigEndian, &a.PTI.Octet)
	binary.Write(buffer, binary.BigEndian, &a.PDUSESSIONAUTHENTICATIONCOMMANDMessageIdentity.Octet)
	binary.Write(buffer, binary.BigEndian, a.EAPMessage.GetLen())
	binary.Write(buffer, binary.BigEndian, &a.EAPMessage.Buffer)
	if a.ExtendedProtocolConfigurationOptions != nil {
		binary.Write(buffer, binary.BigEndian, a.ExtendedProtocolConfigurationOptions.GetIei())
		binary.Write(buffer, binary.BigEndian, a.ExtendedProtocolConfigurationOptions.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.ExtendedProtocolConfigurationOptions.Buffer)
	}
}

func (a *PDUSessionAuthenticationCommand) DecodePDUSessionAuthenticationCommand(byteArray *[]byte) {
	buffer := bytes.NewBuffer(*byteArray)
	binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet)
	binary.Read(buffer, binary.BigEndian, &a.PDUSessionID.Octet)
	binary.Read(buffer, binary.BigEndian, &a.PTI.Octet)
	binary.Read(buffer, binary.BigEndian, &a.PDUSESSIONAUTHENTICATIONCOMMANDMessageIdentity.Octet)
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
		case PDUSessionAuthenticationCommandExtendedProtocolConfigurationOptionsType:
			a.ExtendedProtocolConfigurationOptions = nasType.NewExtendedProtocolConfigurationOptions(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolConfigurationOptions.Len)
			a.ExtendedProtocolConfigurationOptions.SetLen(a.ExtendedProtocolConfigurationOptions.GetLen())
			binary.Read(buffer, binary.BigEndian, a.ExtendedProtocolConfigurationOptions.Buffer[:a.ExtendedProtocolConfigurationOptions.GetLen()])
		default:
		}
	}
}
