package nasMessage

import (
	"bytes"
	"encoding/binary"

	"github.com/mimetrix/nas/nasType"
)

type DLNASTransport struct {
	nasType.ExtendedProtocolDiscriminator         `json:"ExtendedProtocolDiscriminator,omitempty"`
	nasType.SpareHalfOctetAndSecurityHeaderType   `json:"SpareHalfOctetAndSecurityHeaderType,omitempty"`
	nasType.DLNASTRANSPORTMessageIdentity         `json:"DLNASTRANSPORTMessageIdentity,omitempty"`
	nasType.SpareHalfOctetAndPayloadContainerType `json:"SpareHalfOctetAndPayloadContainerType,omitempty"`
	PayloadContainer                              `json:"PayloadContainer,omitempty"`
	*nasType.PduSessionID2Value                   `json:"PduSessionID2Value,omitempty"`
	*nasType.AdditionalInformation                `json:"AdditionalInformation,omitempty"`
	*nasType.Cause5GMM                            `json:"Cause5GMM,omitempty"`
	*nasType.BackoffTimerValue                    `json:"BackoffTimerValue,omitempty"`
}

func NewDLNASTransport(iei uint8) (dLNASTransport *DLNASTransport) {
	dLNASTransport = &DLNASTransport{}
	return dLNASTransport
}

const (
	DLNASTransportPduSessionID2ValueType    uint8 = 0x12
	DLNASTransportAdditionalInformationType uint8 = 0x24
	DLNASTransportCause5GMMType             uint8 = 0x58
	DLNASTransportBackoffTimerValueType     uint8 = 0x37
)

func (a *DLNASTransport) EncodeDLNASTransport(buffer *bytes.Buffer) {
	binary.Write(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet)
	binary.Write(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet)
	binary.Write(buffer, binary.BigEndian, &a.DLNASTRANSPORTMessageIdentity.Octet)
	binary.Write(buffer, binary.BigEndian, &a.SpareHalfOctetAndPayloadContainerType.Octet)
	binary.Write(buffer, binary.BigEndian, a.PayloadContainer.GetLen())

	binary.Write(buffer, binary.BigEndian, &a.PayloadContainer.Buffer)
	if a.PduSessionID2Value != nil {
		binary.Write(buffer, binary.BigEndian, a.PduSessionID2Value.GetIei())
		binary.Write(buffer, binary.BigEndian, &a.PduSessionID2Value.Octet)
	}
	if a.AdditionalInformation != nil {
		binary.Write(buffer, binary.BigEndian, a.AdditionalInformation.GetIei())
		binary.Write(buffer, binary.BigEndian, a.AdditionalInformation.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.AdditionalInformation.Buffer)
	}
	if a.Cause5GMM != nil {
		binary.Write(buffer, binary.BigEndian, a.Cause5GMM.GetIei())
		binary.Write(buffer, binary.BigEndian, &a.Cause5GMM.Octet)
	}
	if a.BackoffTimerValue != nil {
		binary.Write(buffer, binary.BigEndian, a.BackoffTimerValue.GetIei())
		binary.Write(buffer, binary.BigEndian, a.BackoffTimerValue.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.BackoffTimerValue.Octet)
	}
}

func (a *DLNASTransport) DecodeDLNASTransport(byteArray *[]byte) {
	buffer := bytes.NewBuffer(*byteArray)
	binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet)
	a.ExtendedProtocolDiscriminator.DecodeNASType()

	a.ExtendedProtocolDiscriminator.DecodeNASType()

	binary.Read(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet)
	a.SpareHalfOctetAndSecurityHeaderType.DecodeNASType()

	binary.Read(buffer, binary.BigEndian, &a.DLNASTRANSPORTMessageIdentity.Octet)
	a.DLNASTRANSPORTMessageIdentity.DecodeNASType()

	binary.Read(buffer, binary.BigEndian, &a.SpareHalfOctetAndPayloadContainerType.Octet)
	a.SpareHalfOctetAndPayloadContainerType.DecodeNASType()

	binary.Read(buffer, binary.BigEndian, &a.PayloadContainer.Len)
	a.PayloadContainer.SetLen(a.PayloadContainer.GetLen())

	binary.Read(buffer, binary.BigEndian, &a.PayloadContainer.Buffer)
	a.PayloadContainer.DecodeNASType()

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
		case DLNASTransportPduSessionID2ValueType:
			a.PduSessionID2Value = nasType.NewPduSessionID2Value(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.PduSessionID2Value.Octet)
		case DLNASTransportAdditionalInformationType:
			a.AdditionalInformation = nasType.NewAdditionalInformation(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.AdditionalInformation.Len)
			a.AdditionalInformation.SetLen(a.AdditionalInformation.GetLen())
			binary.Read(buffer, binary.BigEndian, a.AdditionalInformation.Buffer[:a.AdditionalInformation.GetLen()])
		case DLNASTransportCause5GMMType:
			a.Cause5GMM = nasType.NewCause5GMM(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.Cause5GMM.Octet)
			a.Cause5GMM.DecodeNASType()
		case DLNASTransportBackoffTimerValueType:
			a.BackoffTimerValue = nasType.NewBackoffTimerValue(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.BackoffTimerValue.Len)
			a.BackoffTimerValue.SetLen(a.BackoffTimerValue.GetLen())
			binary.Read(buffer, binary.BigEndian, &a.BackoffTimerValue.Octet)
		default:
		}
	}
}
