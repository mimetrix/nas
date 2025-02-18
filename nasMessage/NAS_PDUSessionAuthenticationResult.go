package nasMessage

import (
	"bytes"
	"encoding/binary"

	"github.com/mimetrix/nas/nasType"
)

type PDUSessionAuthenticationResult struct {
	nasType.ExtendedProtocolDiscriminator                 `json:"ExtendedProtocolDiscriminator,omitempty"`
	nasType.PDUSessionID                                  `json:"PDUSessionID,omitempty"`
	nasType.PTI                                           `json:"PTI,omitempty"`
	nasType.PDUSESSIONAUTHENTICATIONRESULTMessageIdentity `json:"PDUSESSIONAUTHENTICATIONRESULTMessageIdentity,omitempty"`
	*nasType.EAPMessage                                   `json:"EAPMessage,omitempty"`
	*nasType.ExtendedProtocolConfigurationOptions         `json:"ExtendedProtocolConfigurationOptions,omitempty"`
}

func NewPDUSessionAuthenticationResult(iei uint8) (pDUSessionAuthenticationResult *PDUSessionAuthenticationResult) {
	pDUSessionAuthenticationResult = &PDUSessionAuthenticationResult{}
	return pDUSessionAuthenticationResult
}

const (
	PDUSessionAuthenticationResultEAPMessageType                           uint8 = 0x78
	PDUSessionAuthenticationResultExtendedProtocolConfigurationOptionsType uint8 = 0x7B
)

func (a *PDUSessionAuthenticationResult) EncodePDUSessionAuthenticationResult(buffer *bytes.Buffer) {
	binary.Write(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet)
	binary.Write(buffer, binary.BigEndian, &a.PDUSessionID.Octet)
	binary.Write(buffer, binary.BigEndian, &a.PTI.Octet)
	binary.Write(buffer, binary.BigEndian, &a.PDUSESSIONAUTHENTICATIONRESULTMessageIdentity.Octet)
	if a.EAPMessage != nil {
		binary.Write(buffer, binary.BigEndian, a.EAPMessage.GetIei())
		binary.Write(buffer, binary.BigEndian, a.EAPMessage.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.EAPMessage.Buffer)
	}
	if a.ExtendedProtocolConfigurationOptions != nil {
		binary.Write(buffer, binary.BigEndian, a.ExtendedProtocolConfigurationOptions.GetIei())
		binary.Write(buffer, binary.BigEndian, a.ExtendedProtocolConfigurationOptions.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.ExtendedProtocolConfigurationOptions.Buffer)
	}
}

func (a *PDUSessionAuthenticationResult) DecodePDUSessionAuthenticationResult(byteArray *[]byte) {
	buffer := bytes.NewBuffer(*byteArray)
	binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet)
	binary.Read(buffer, binary.BigEndian, &a.PDUSessionID.Octet)
	binary.Read(buffer, binary.BigEndian, &a.PTI.Octet)
	binary.Read(buffer, binary.BigEndian, &a.PDUSESSIONAUTHENTICATIONRESULTMessageIdentity.Octet)
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
		case PDUSessionAuthenticationResultEAPMessageType:
			a.EAPMessage = nasType.NewEAPMessage(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.EAPMessage.Len)
			a.EAPMessage.SetLen(a.EAPMessage.GetLen())
			binary.Read(buffer, binary.BigEndian, a.EAPMessage.Buffer[:a.EAPMessage.GetLen()])
		case PDUSessionAuthenticationResultExtendedProtocolConfigurationOptionsType:
			a.ExtendedProtocolConfigurationOptions = nasType.NewExtendedProtocolConfigurationOptions(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolConfigurationOptions.Len)
			a.ExtendedProtocolConfigurationOptions.SetLen(a.ExtendedProtocolConfigurationOptions.GetLen())
			binary.Read(buffer, binary.BigEndian, a.ExtendedProtocolConfigurationOptions.Buffer[:a.ExtendedProtocolConfigurationOptions.GetLen()])
		default:
		}
	}
}
