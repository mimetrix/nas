package nasMessage

import (
	"bytes"
	"encoding/binary"
	"github.com/mimetrix/nas/nasType"
)

type AuthenticationFailure struct {
	nasType.ExtendedProtocolDiscriminator        `json:"ExtendedProtocolDiscriminator,omitempty"`
	nasType.SpareHalfOctetAndSecurityHeaderType  `json:"SpareHalfOctetAndSecurityHeaderType,omitempty"`
	nasType.AuthenticationFailureMessageIdentity `json:"AuthenticationFailureMessageIdentity,omitempty"`
	nasType.Cause5GMM                            `json:"Cause5GMM,omitempty"`
	*nasType.AuthenticationFailureParameter      `json:"AuthenticationFailureParameter,omitempty"`
}

func NewAuthenticationFailure(iei uint8) (authenticationFailure *AuthenticationFailure) {
	authenticationFailure = &AuthenticationFailure{}
	return authenticationFailure
}

const (
	AuthenticationFailureAuthenticationFailureParameterType uint8 = 0x30
)

func (a *AuthenticationFailure) EncodeAuthenticationFailure(buffer *bytes.Buffer) {
	binary.Write(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet)
	binary.Write(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet)
	binary.Write(buffer, binary.BigEndian, &a.AuthenticationFailureMessageIdentity.Octet)
	binary.Write(buffer, binary.BigEndian, &a.Cause5GMM.Octet)
	if a.AuthenticationFailureParameter != nil {
		binary.Write(buffer, binary.BigEndian, a.AuthenticationFailureParameter.GetIei())
		binary.Write(buffer, binary.BigEndian, a.AuthenticationFailureParameter.GetLen())
		binary.Write(buffer, binary.BigEndian, a.AuthenticationFailureParameter.Octet[:a.AuthenticationFailureParameter.GetLen()])
	}
}

func (a *AuthenticationFailure) DecodeAuthenticationFailure(byteArray *[]byte) {
	buffer := bytes.NewBuffer(*byteArray)
	binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet)
	binary.Read(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet)
	binary.Read(buffer, binary.BigEndian, &a.AuthenticationFailureMessageIdentity.Octet)
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
		case AuthenticationFailureAuthenticationFailureParameterType:
			a.AuthenticationFailureParameter = nasType.NewAuthenticationFailureParameter(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.AuthenticationFailureParameter.Len)
			a.AuthenticationFailureParameter.SetLen(a.AuthenticationFailureParameter.GetLen())
			binary.Read(buffer, binary.BigEndian, a.AuthenticationFailureParameter.Octet[:a.AuthenticationFailureParameter.GetLen()])
		default:
		}
	}
}
