package nasMessage

import (
	"bytes"
	"encoding/binary"

	"github.com/free5gc/nas/nasType"
)

type Status5GMM struct {
	nasType.ExtendedProtocolDiscriminator       `json:"ExtendedProtocolDiscriminator,omitempty"`
	nasType.SpareHalfOctetAndSecurityHeaderType `json:"SpareHalfOctetAndSecurityHeaderType,omitempty"`
	nasType.STATUSMessageIdentity5GMM           `json:"STATUSMessageIdentity5GMM,omitempty"`
	nasType.Cause5GMM                           `json:"Cause5GMM,omitempty"`
}

func NewStatus5GMM(iei uint8) (status5GMM *Status5GMM) {
	status5GMM = &Status5GMM{}
	return status5GMM
}

func (a *Status5GMM) EncodeStatus5GMM(buffer *bytes.Buffer) {
	binary.Write(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet)
	binary.Write(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet)
	binary.Write(buffer, binary.BigEndian, &a.STATUSMessageIdentity5GMM.Octet)
	binary.Write(buffer, binary.BigEndian, &a.Cause5GMM.Octet)
}

func (a *Status5GMM) DecodeStatus5GMM(byteArray *[]byte) {
	buffer := bytes.NewBuffer(*byteArray)
	binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet)
	binary.Read(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet)
	binary.Read(buffer, binary.BigEndian, &a.STATUSMessageIdentity5GMM.Octet)
	binary.Read(buffer, binary.BigEndian, &a.Cause5GMM.Octet)
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
