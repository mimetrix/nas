package nasMessage

import (
	"bytes"
	"encoding/binary"
    "fmt"

	"github.com/mimetrix/nas/nasType"
)

type AuthenticationRequest struct {
	nasType.ExtendedProtocolDiscriminator        `json:"ExtendedProtocolDiscriminator,omitempty"`
	nasType.SpareHalfOctetAndSecurityHeaderType  `json:"SpareHalfOctetAndSecurityHeaderType,omitempty"`
	nasType.AuthenticationRequestMessageIdentity `json:"AuthenticationRequestMessageIdentity,omitempty"`
	nasType.SpareHalfOctetAndNgksi               `json:"SpareHalfOctetAndNgksi,omitempty"`
	nasType.ABBA                                 `json:"ABBA,omitempty"`
	*nasType.AuthenticationParameterRAND         `json:"AuthenticationParameterRAND,omitempty"`
	*nasType.AuthenticationParameterAUTN         `json:"AuthenticationParameterAUTN,omitempty"`
	*nasType.EAPMessage                          `json:"EAPMessage,omitempty"`
}

func NewAuthenticationRequest(iei uint8) (authenticationRequest *AuthenticationRequest) {
	authenticationRequest = &AuthenticationRequest{}
	return authenticationRequest
}

const (
	AuthenticationRequestAuthenticationParameterRANDType uint8 = 0x21
	AuthenticationRequestAuthenticationParameterAUTNType uint8 = 0x20
	AuthenticationRequestEAPMessageType                  uint8 = 0x78
)

func (a *AuthenticationRequest) EncodeAuthenticationRequest(buffer *bytes.Buffer) {
	binary.Write(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet)
	binary.Write(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet)
	binary.Write(buffer, binary.BigEndian, &a.AuthenticationRequestMessageIdentity.Octet)

	binary.Write(buffer, binary.BigEndian, &a.SpareHalfOctetAndNgksi.Octet)

	binary.Write(buffer, binary.BigEndian, a.ABBA.GetLen())
	binary.Write(buffer, binary.BigEndian, &a.ABBA.Buffer)
	if a.AuthenticationParameterRAND != nil {
		binary.Write(buffer, binary.BigEndian, a.AuthenticationParameterRAND.GetIei())
		binary.Write(buffer, binary.BigEndian, &a.AuthenticationParameterRAND.Octet)
	}
	if a.AuthenticationParameterAUTN != nil {
		binary.Write(buffer, binary.BigEndian, a.AuthenticationParameterAUTN.GetIei())
		binary.Write(buffer, binary.BigEndian, a.AuthenticationParameterAUTN.GetLen())
		binary.Write(buffer, binary.BigEndian, a.AuthenticationParameterAUTN.Octet[:a.AuthenticationParameterAUTN.GetLen()])
	}
	if a.EAPMessage != nil {
		binary.Write(buffer, binary.BigEndian, a.EAPMessage.GetIei())
		binary.Write(buffer, binary.BigEndian, a.EAPMessage.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.EAPMessage.Buffer)
	}
}

func (a *AuthenticationRequest) DecodeAuthenticationRequest(byteArray *[]byte) {
	buffer := bytes.NewBuffer(*byteArray)
	binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet)
    a.ExtendedProtocolDiscriminator.DecodeNASType()

	binary.Read(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet)
    a.SpareHalfOctetAndSecurityHeaderType.DecodeNASType()

	binary.Read(buffer, binary.BigEndian, &a.AuthenticationRequestMessageIdentity.Octet)
    a.AuthenticationRequestMessageIdentity.DecodeNASType()

	binary.Read(buffer, binary.BigEndian, &a.SpareHalfOctetAndNgksi.Octet)
    err := a.SpareHalfOctetAndNgksi.DecodeNASType()
    if err != nil {
        fmt.Println(err)
    }

	binary.Read(buffer, binary.BigEndian, &a.ABBA.Len)
	a.ABBA.SetLen(a.ABBA.GetLen())
	binary.Read(buffer, binary.BigEndian, &a.ABBA.Buffer)
    a.ABBA.DecodeNASType()
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
		case AuthenticationRequestAuthenticationParameterRANDType:
			a.AuthenticationParameterRAND = nasType.NewAuthenticationParameterRAND(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.AuthenticationParameterRAND.Octet)
            a.AuthenticationParameterRAND.DecodeNASType() 
		case AuthenticationRequestAuthenticationParameterAUTNType:
			a.AuthenticationParameterAUTN = nasType.NewAuthenticationParameterAUTN(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.AuthenticationParameterAUTN.Len)
			a.AuthenticationParameterAUTN.SetLen(a.AuthenticationParameterAUTN.GetLen())
			binary.Read(buffer, binary.BigEndian, a.AuthenticationParameterAUTN.Octet[:a.AuthenticationParameterAUTN.GetLen()])
            a.AuthenticationParameterAUTN.DecodeNASType()
		case AuthenticationRequestEAPMessageType:
			a.EAPMessage = nasType.NewEAPMessage(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.EAPMessage.Len)
			a.EAPMessage.SetLen(a.EAPMessage.GetLen())
			binary.Read(buffer, binary.BigEndian, a.EAPMessage.Buffer[:a.EAPMessage.GetLen()])
		default:
		}
	}
}
