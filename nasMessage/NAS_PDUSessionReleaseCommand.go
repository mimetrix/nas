package nasMessage

import (
	"bytes"
	"encoding/binary"

	"github.com/mimetrix/nas/nasType"
)

type PDUSessionReleaseCommand struct {
	nasType.ExtendedProtocolDiscriminator           `json:"ExtendedProtocolDiscriminator,omitempty"`
	nasType.PDUSessionID                            `json:"PDUSessionID,omitempty"`
	nasType.PTI                                     `json:"PTI,omitempty"`
	nasType.PDUSESSIONRELEASECOMMANDMessageIdentity `json:"PDUSESSIONRELEASECOMMANDMessageIdentity,omitempty"`
	nasType.Cause5GSM                               `json:"Cause5GSM,omitempty"`
	*nasType.BackoffTimerValue                      `json:"BackoffTimerValue,omitempty"`
	*nasType.EAPMessage                             `json:"EAPMessage,omitempty"`
	*nasType.ExtendedProtocolConfigurationOptions   `json:"ExtendedProtocolConfigurationOptions,omitempty"`
}

func NewPDUSessionReleaseCommand(iei uint8) (pDUSessionReleaseCommand *PDUSessionReleaseCommand) {
	pDUSessionReleaseCommand = &PDUSessionReleaseCommand{}
	return pDUSessionReleaseCommand
}

const (
	PDUSessionReleaseCommandBackoffTimerValueType                    uint8 = 0x37
	PDUSessionReleaseCommandEAPMessageType                           uint8 = 0x78
	PDUSessionReleaseCommandExtendedProtocolConfigurationOptionsType uint8 = 0x7B
)

func (a *PDUSessionReleaseCommand) EncodePDUSessionReleaseCommand(buffer *bytes.Buffer) {
	binary.Write(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet)
	binary.Write(buffer, binary.BigEndian, &a.PDUSessionID.Octet)
	binary.Write(buffer, binary.BigEndian, &a.PTI.Octet)
	binary.Write(buffer, binary.BigEndian, &a.PDUSESSIONRELEASECOMMANDMessageIdentity.Octet)
	binary.Write(buffer, binary.BigEndian, &a.Cause5GSM.Octet)
	if a.BackoffTimerValue != nil {
		binary.Write(buffer, binary.BigEndian, a.BackoffTimerValue.GetIei())
		binary.Write(buffer, binary.BigEndian, a.BackoffTimerValue.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.BackoffTimerValue.Octet)
	}
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

func (a *PDUSessionReleaseCommand) DecodePDUSessionReleaseCommand(byteArray *[]byte) {
	buffer := bytes.NewBuffer(*byteArray)
	binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet)
	a.ExtendedProtocolDiscriminator.DecodeNASType()

	binary.Read(buffer, binary.BigEndian, &a.PDUSessionID.Octet)
	binary.Read(buffer, binary.BigEndian, &a.PTI.Octet)
	binary.Read(buffer, binary.BigEndian, &a.PDUSESSIONRELEASECOMMANDMessageIdentity.Octet)
	a.PDUSESSIONRELEASECOMMANDMessageIdentity.DecodeNASType()

	binary.Read(buffer, binary.BigEndian, &a.Cause5GSM.Octet)
	a.Cause5GSM.DecodeNASType()

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
		case PDUSessionReleaseCommandBackoffTimerValueType:
			a.BackoffTimerValue = nasType.NewBackoffTimerValue(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.BackoffTimerValue.Len)
			a.BackoffTimerValue.SetLen(a.BackoffTimerValue.GetLen())
			binary.Read(buffer, binary.BigEndian, &a.BackoffTimerValue.Octet)
		case PDUSessionReleaseCommandEAPMessageType:
			a.EAPMessage = nasType.NewEAPMessage(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.EAPMessage.Len)
			a.EAPMessage.SetLen(a.EAPMessage.GetLen())
			binary.Read(buffer, binary.BigEndian, a.EAPMessage.Buffer[:a.EAPMessage.GetLen()])
		case PDUSessionReleaseCommandExtendedProtocolConfigurationOptionsType:
			a.ExtendedProtocolConfigurationOptions = nasType.NewExtendedProtocolConfigurationOptions(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolConfigurationOptions.Len)
			a.ExtendedProtocolConfigurationOptions.SetLen(a.ExtendedProtocolConfigurationOptions.GetLen())
			binary.Read(buffer, binary.BigEndian, a.ExtendedProtocolConfigurationOptions.Buffer[:a.ExtendedProtocolConfigurationOptions.GetLen()])
		default:
		}
	}
}
