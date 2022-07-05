package nasMessage

import (
	"bytes"
	"encoding/binary"

	"github.com/free5gc/nas/nasType"
)

type Status5GSM struct {
	nasType.ExtendedProtocolDiscriminator `json:"ExtendedProtocolDiscriminator,omitempty"`
	nasType.PDUSessionID                  `json:"PDUSessionID,omitempty"`
	nasType.PTI                           `json:"PTI,omitempty"`
	nasType.STATUSMessageIdentity5GSM     `json:"STATUSMessageIdentity5GSM,omitempty"`
	nasType.Cause5GSM                     `json:"Cause5GSM,omitempty"`
}

func NewStatus5GSM(iei uint8) (status5GSM *Status5GSM) {
	status5GSM = &Status5GSM{}
	return status5GSM
}

func (a *Status5GSM) EncodeStatus5GSM(buffer *bytes.Buffer) {
	binary.Write(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet)
	binary.Write(buffer, binary.BigEndian, &a.PDUSessionID.Octet)
	binary.Write(buffer, binary.BigEndian, &a.PTI.Octet)
	binary.Write(buffer, binary.BigEndian, &a.STATUSMessageIdentity5GSM.Octet)
	binary.Write(buffer, binary.BigEndian, &a.Cause5GSM.Octet)
}

func (a *Status5GSM) DecodeStatus5GSM(byteArray *[]byte) {
	buffer := bytes.NewBuffer(*byteArray)
	binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet)
	binary.Read(buffer, binary.BigEndian, &a.PDUSessionID.Octet)
	binary.Read(buffer, binary.BigEndian, &a.PTI.Octet)
	binary.Read(buffer, binary.BigEndian, &a.STATUSMessageIdentity5GSM.Octet)
	binary.Read(buffer, binary.BigEndian, &a.Cause5GSM.Octet)
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
