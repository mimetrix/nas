package nasMessage

import (
	"bytes"
	"encoding/binary"

	"github.com/mimetrix/nas/logger"
	"github.com/mimetrix/nas/nasType"
)

type PDUSessionEstablishmentRequest struct {
	nasType.ExtendedProtocolDiscriminator                 `json:"ExtendedProtocolDiscriminator,omitempty"`
	nasType.PDUSessionID                                  `json:"PDUSessionID,omitempty"`
	nasType.PTI                                           `json:"PTI,omitempty"`
	nasType.PDUSESSIONESTABLISHMENTREQUESTMessageIdentity `json:"PDUSESSIONESTABLISHMENTREQUESTMessageIdentity,omitempty"`
	nasType.IntegrityProtectionMaximumDataRate            `json:"IntegrityProtectionMaximumDataRate,omitempty"`
	*nasType.PDUSessionType                               `json:"PDUSessionType,omitempty"`
	*nasType.SSCMode                                      `json:"SSCMode,omitempty"`
	*nasType.Capability5GSM                               `json:"Capability5GSM,omitempty"`
	*nasType.MaximumNumberOfSupportedPacketFilters        `json:"MaximumNumberOfSupportedPacketFilters,omitempty"`
	*nasType.AlwaysonPDUSessionRequested                  `json:"AlwaysonPDUSessionRequested,omitempty"`
	*nasType.SMPDUDNRequestContainer                      `json:"SMPDUDNRequestContainer,omitempty"`
	*nasType.ExtendedProtocolConfigurationOptions         `json:"ExtendedProtocolConfigurationOptions,omitempty"`
}

func NewPDUSessionEstablishmentRequest(iei uint8) (pDUSessionEstablishmentRequest *PDUSessionEstablishmentRequest) {
	pDUSessionEstablishmentRequest = &PDUSessionEstablishmentRequest{}
	return pDUSessionEstablishmentRequest
}

const (
	PDUSessionEstablishmentRequestPDUSessionTypeType                        uint8 = 0x09
	PDUSessionEstablishmentRequestSSCModeType                               uint8 = 0x0A
	PDUSessionEstablishmentRequestCapability5GSMType                        uint8 = 0x28
	PDUSessionEstablishmentRequestMaximumNumberOfSupportedPacketFiltersType uint8 = 0x55
	PDUSessionEstablishmentRequestAlwaysonPDUSessionRequestedType           uint8 = 0x0B
	PDUSessionEstablishmentRequestSMPDUDNRequestContainerType               uint8 = 0x39
	PDUSessionEstablishmentRequestExtendedProtocolConfigurationOptionsType  uint8 = 0x7B
)

func (a *PDUSessionEstablishmentRequest) EncodePDUSessionEstablishmentRequest(buffer *bytes.Buffer) {
	binary.Write(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet)
	binary.Write(buffer, binary.BigEndian, &a.PDUSessionID.Octet)
	binary.Write(buffer, binary.BigEndian, &a.PTI.Octet)
	binary.Write(buffer, binary.BigEndian, &a.PDUSESSIONESTABLISHMENTREQUESTMessageIdentity.Octet)
	binary.Write(buffer, binary.BigEndian, &a.IntegrityProtectionMaximumDataRate.Octet)
	if a.PDUSessionType != nil {
		binary.Write(buffer, binary.BigEndian, &a.PDUSessionType.Octet)
	}
	if a.SSCMode != nil {
		binary.Write(buffer, binary.BigEndian, &a.SSCMode.Octet)
	}
	if a.Capability5GSM != nil {
		binary.Write(buffer, binary.BigEndian, a.Capability5GSM.GetIei())
		binary.Write(buffer, binary.BigEndian, a.Capability5GSM.GetLen())
		binary.Write(buffer, binary.BigEndian, a.Capability5GSM.Octet[:a.Capability5GSM.GetLen()])
	}
	if a.MaximumNumberOfSupportedPacketFilters != nil {
		binary.Write(buffer, binary.BigEndian, a.MaximumNumberOfSupportedPacketFilters.GetIei())
		binary.Write(buffer, binary.BigEndian, &a.MaximumNumberOfSupportedPacketFilters.Octet)
	}
	if a.AlwaysonPDUSessionRequested != nil {
		binary.Write(buffer, binary.BigEndian, &a.AlwaysonPDUSessionRequested.Octet)
	}
	if a.SMPDUDNRequestContainer != nil {
		binary.Write(buffer, binary.BigEndian, a.SMPDUDNRequestContainer.GetIei())
		binary.Write(buffer, binary.BigEndian, a.SMPDUDNRequestContainer.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.SMPDUDNRequestContainer.Buffer)
	}
	if a.ExtendedProtocolConfigurationOptions != nil {
		logger.NasMsgLog.Infoln("Encode ExtendedProtocolConfigurationOptions in EncodePDUSessionEstablishmentRequest")
		binary.Write(buffer, binary.BigEndian, a.ExtendedProtocolConfigurationOptions.GetIei())
		binary.Write(buffer, binary.BigEndian, a.ExtendedProtocolConfigurationOptions.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.ExtendedProtocolConfigurationOptions.Buffer)
	}
}

func (a *PDUSessionEstablishmentRequest) DecodePDUSessionEstablishmentRequest(byteArray *[]byte) {
	buffer := bytes.NewBuffer(*byteArray)
	binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet)
    a.ExtendedProtocolDiscriminator.DecodeNASType()
	binary.Read(buffer, binary.BigEndian, &a.PDUSessionID.Octet)
	binary.Read(buffer, binary.BigEndian, &a.PTI.Octet)
	binary.Read(buffer, binary.BigEndian, &a.PDUSESSIONESTABLISHMENTREQUESTMessageIdentity.Octet)

    a.PDUSESSIONESTABLISHMENTREQUESTMessageIdentity.DecodeNASType()
	binary.Read(buffer, binary.BigEndian, &a.IntegrityProtectionMaximumDataRate.Octet)
    a.IntegrityProtectionMaximumDataRate.DecodeNASType()
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
		case PDUSessionEstablishmentRequestPDUSessionTypeType:
			a.PDUSessionType = nasType.NewPDUSessionType(ieiN)
			a.PDUSessionType.Octet = ieiN
            a.PDUSessionType.DecodeNASType()
		case PDUSessionEstablishmentRequestSSCModeType:
			a.SSCMode = nasType.NewSSCMode(ieiN)
			a.SSCMode.Octet = ieiN
		case PDUSessionEstablishmentRequestCapability5GSMType:
			a.Capability5GSM = nasType.NewCapability5GSM(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.Capability5GSM.Len)
			a.Capability5GSM.SetLen(a.Capability5GSM.GetLen())
			binary.Read(buffer, binary.BigEndian, a.Capability5GSM.Octet[:a.Capability5GSM.GetLen()])
            a.Capability5GSM.DecodeNASType()

		case PDUSessionEstablishmentRequestMaximumNumberOfSupportedPacketFiltersType:
			a.MaximumNumberOfSupportedPacketFilters = nasType.NewMaximumNumberOfSupportedPacketFilters(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.MaximumNumberOfSupportedPacketFilters.Octet)
		case PDUSessionEstablishmentRequestAlwaysonPDUSessionRequestedType:
			a.AlwaysonPDUSessionRequested = nasType.NewAlwaysonPDUSessionRequested(ieiN)
			a.AlwaysonPDUSessionRequested.Octet = ieiN
		case PDUSessionEstablishmentRequestSMPDUDNRequestContainerType:
			a.SMPDUDNRequestContainer = nasType.NewSMPDUDNRequestContainer(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.SMPDUDNRequestContainer.Len)
			a.SMPDUDNRequestContainer.SetLen(a.SMPDUDNRequestContainer.GetLen())
			binary.Read(buffer, binary.BigEndian, a.SMPDUDNRequestContainer.Buffer[:a.SMPDUDNRequestContainer.GetLen()])
		case PDUSessionEstablishmentRequestExtendedProtocolConfigurationOptionsType:
			a.ExtendedProtocolConfigurationOptions = nasType.NewExtendedProtocolConfigurationOptions(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolConfigurationOptions.Len)
			a.ExtendedProtocolConfigurationOptions.SetLen(a.ExtendedProtocolConfigurationOptions.GetLen())
			binary.Read(buffer, binary.BigEndian, a.ExtendedProtocolConfigurationOptions.Buffer[:a.ExtendedProtocolConfigurationOptions.GetLen()])
            a.ExtendedProtocolConfigurationOptions.DecodeNASType()
		default:
		}
	}
}
