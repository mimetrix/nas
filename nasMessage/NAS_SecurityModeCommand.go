package nasMessage

import (
	"bytes"
	"encoding/binary"
	"github.com/mimetrix/nas/nasType"
)

type SecurityModeCommand struct {
	nasType.ExtendedProtocolDiscriminator       `json:"ExtendedProtocolDiscriminator,omitempty"`
	nasType.SpareHalfOctetAndSecurityHeaderType `json:"SpareHalfOctetAndSecurityHeaderType,omitempty"`
	nasType.SecurityModeCommandMessageIdentity  `json:"SecurityModeCommandMessageIdentity,omitempty"`
	nasType.SelectedNASSecurityAlgorithms       `json:"SelectedNASSecurityAlgorithms,omitempty"`
	nasType.SpareHalfOctetAndNgksi              `json:"SpareHalfOctetAndNgksi,omitempty"`
	nasType.ReplayedUESecurityCapabilities      `json:"ReplayedUESecurityCapabilities,omitempty"`
	*nasType.IMEISVRequest                      `json:"IMEISVRequest,omitempty"`
	*nasType.SelectedEPSNASSecurityAlgorithms   `json:"SelectedEPSNASSecurityAlgorithms,omitempty"`
	*nasType.Additional5GSecurityInformation    `json:"Additional5GSecurityInformation,omitempty"`
	*nasType.EAPMessage                         `json:"EAPMessage,omitempty"`
	*nasType.ABBA                               `json:"ABBA,omitempty"`
	*nasType.ReplayedS1UESecurityCapabilities   `json:"ReplayedS1UESecurityCapabilities,omitempty"`
}

func NewSecurityModeCommand(iei uint8) (securityModeCommand *SecurityModeCommand) {
	securityModeCommand = &SecurityModeCommand{}
	return securityModeCommand
}

const (
	SecurityModeCommandIMEISVRequestType                    uint8 = 0x0E
	SecurityModeCommandSelectedEPSNASSecurityAlgorithmsType uint8 = 0x57
	SecurityModeCommandAdditional5GSecurityInformationType  uint8 = 0x36
	SecurityModeCommandEAPMessageType                       uint8 = 0x78
	SecurityModeCommandABBAType                             uint8 = 0x38
	SecurityModeCommandReplayedS1UESecurityCapabilitiesType uint8 = 0x19
)

func (a *SecurityModeCommand) EncodeSecurityModeCommand(buffer *bytes.Buffer) {
	binary.Write(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet)
	binary.Write(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet)
	binary.Write(buffer, binary.BigEndian, &a.SecurityModeCommandMessageIdentity.Octet)
	binary.Write(buffer, binary.BigEndian, &a.SelectedNASSecurityAlgorithms.Octet)
	binary.Write(buffer, binary.BigEndian, &a.SpareHalfOctetAndNgksi.Octet)
	binary.Write(buffer, binary.BigEndian, a.ReplayedUESecurityCapabilities.GetLen())
	binary.Write(buffer, binary.BigEndian, &a.ReplayedUESecurityCapabilities.Buffer)
	if a.IMEISVRequest != nil {
		binary.Write(buffer, binary.BigEndian, &a.IMEISVRequest.Octet)
	}
	if a.SelectedEPSNASSecurityAlgorithms != nil {
		binary.Write(buffer, binary.BigEndian, a.SelectedEPSNASSecurityAlgorithms.GetIei())
		binary.Write(buffer, binary.BigEndian, &a.SelectedEPSNASSecurityAlgorithms.Octet)
	}
	if a.Additional5GSecurityInformation != nil {
		binary.Write(buffer, binary.BigEndian, a.Additional5GSecurityInformation.GetIei())
		binary.Write(buffer, binary.BigEndian, a.Additional5GSecurityInformation.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.Additional5GSecurityInformation.Octet)
	}
	if a.EAPMessage != nil {
		binary.Write(buffer, binary.BigEndian, a.EAPMessage.GetIei())
		binary.Write(buffer, binary.BigEndian, a.EAPMessage.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.EAPMessage.Buffer)
	}
	if a.ABBA != nil {
		binary.Write(buffer, binary.BigEndian, a.ABBA.GetIei())
		binary.Write(buffer, binary.BigEndian, a.ABBA.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.ABBA.Buffer)
	}
	if a.ReplayedS1UESecurityCapabilities != nil {
		binary.Write(buffer, binary.BigEndian, a.ReplayedS1UESecurityCapabilities.GetIei())
		binary.Write(buffer, binary.BigEndian, a.ReplayedS1UESecurityCapabilities.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.ReplayedS1UESecurityCapabilities.Buffer)
	}
}

func (a *SecurityModeCommand) DecodeSecurityModeCommand(byteArray *[]byte) {

    //fmt.Println("\nDecodeSecurityModeCommand\n")
	buffer := bytes.NewBuffer(*byteArray)
	binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet)
    a.ExtendedProtocolDiscriminator.DecodeNASType()

	binary.Read(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet)
    a.SpareHalfOctetAndSecurityHeaderType.DecodeNASType()

	binary.Read(buffer, binary.BigEndian, &a.SecurityModeCommandMessageIdentity.Octet)
    a.SecurityModeCommandMessageIdentity.DecodeNASType()

	binary.Read(buffer, binary.BigEndian, &a.SelectedNASSecurityAlgorithms.Octet)
    a.SelectedNASSecurityAlgorithms.DecodeNASType()

	binary.Read(buffer, binary.BigEndian, &a.SpareHalfOctetAndNgksi.Octet)
    a.SpareHalfOctetAndNgksi.DecodeNASType()

	binary.Read(buffer, binary.BigEndian, &a.ReplayedUESecurityCapabilities.Len)
	a.ReplayedUESecurityCapabilities.SetLen(a.ReplayedUESecurityCapabilities.GetLen())
	binary.Read(buffer, binary.BigEndian, &a.ReplayedUESecurityCapabilities.Buffer)
    a.ReplayedUESecurityCapabilities.DecodeNASType()

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
		case SecurityModeCommandIMEISVRequestType:
			a.IMEISVRequest = nasType.NewIMEISVRequest(ieiN)
			a.IMEISVRequest.Octet = ieiN
			a.IMEISVRequest.DecodeNASType()
		case SecurityModeCommandSelectedEPSNASSecurityAlgorithmsType:
			a.SelectedEPSNASSecurityAlgorithms = nasType.NewSelectedEPSNASSecurityAlgorithms(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.SelectedEPSNASSecurityAlgorithms.Octet)
		case SecurityModeCommandAdditional5GSecurityInformationType:
			a.Additional5GSecurityInformation = nasType.NewAdditional5GSecurityInformation(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.Additional5GSecurityInformation.Len)
			a.Additional5GSecurityInformation.SetLen(a.Additional5GSecurityInformation.GetLen())
			binary.Read(buffer, binary.BigEndian, &a.Additional5GSecurityInformation.Octet)
            a.Additional5GSecurityInformation.DecodeNASType()
		case SecurityModeCommandEAPMessageType:
			a.EAPMessage = nasType.NewEAPMessage(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.EAPMessage.Len)
			a.EAPMessage.SetLen(a.EAPMessage.GetLen())
			binary.Read(buffer, binary.BigEndian, a.EAPMessage.Buffer[:a.EAPMessage.GetLen()])
		case SecurityModeCommandABBAType:
			a.ABBA = nasType.NewABBA(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.ABBA.Len)
			a.ABBA.SetLen(a.ABBA.GetLen())
			binary.Read(buffer, binary.BigEndian, a.ABBA.Buffer[:a.ABBA.GetLen()])
		case SecurityModeCommandReplayedS1UESecurityCapabilitiesType:
			a.ReplayedS1UESecurityCapabilities = nasType.NewReplayedS1UESecurityCapabilities(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.ReplayedS1UESecurityCapabilities.Len)
			a.ReplayedS1UESecurityCapabilities.SetLen(a.ReplayedS1UESecurityCapabilities.GetLen())
			binary.Read(buffer, binary.BigEndian, a.ReplayedS1UESecurityCapabilities.Buffer[:a.ReplayedS1UESecurityCapabilities.GetLen()])
		default:
		}
	}
}
