package nasMessage_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mimetrix/nas/logger"

	//"fmt"
	"github.com/mimetrix/nas/nasMessage"
	"github.com/mimetrix/nas/nasType"
)

type nasMessageAuthenticationResponseData struct {
	inExtendedProtocolDiscriminator         uint8
	inSecurityHeader                        uint8
	inSpareHalfOctet                        uint8
	inAuthenticationResponseMessageIdentity uint8
	inAuthenticationResponseParameter       nasType.AuthenticationResponseParameter
	inEAPMessage                            nasType.EAPMessage
}

var nasMessageAuthenticationResponseTable = []nasMessageAuthenticationResponseData{
	{
		inExtendedProtocolDiscriminator:         0x01,
		inSecurityHeader:                        0x08,
		inSpareHalfOctet:                        0x01,
		inAuthenticationResponseMessageIdentity: 0x01,
		inAuthenticationResponseParameter:       nasType.AuthenticationResponseParameter{Iei: nasMessage.AuthenticationResponseAuthenticationResponseParameterType, Len: 16, Octet: [16]uint8{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}},
		inEAPMessage:                            nasType.EAPMessage{Iei: nasMessage.AuthenticationResponseEAPMessageType, Len: 4, Buffer: []uint8{0x01, 0x01, 0x01, 0x01}},
	},
}

func TestNasTypeNewAuthenticationResponse(t *testing.T) {
	a := nasMessage.NewAuthenticationResponse(0)
	assert.NotNil(t, a)
}

func TestNasTypeNewAuthenticationResponseMessage(t *testing.T) {
	logger.NasMsgLog.Infoln("---Test NAS Message: AuthenticationResponseMessage---")
	for i, table := range nasMessageAuthenticationResponseTable {
		logger.NasMsgLog.Infoln("Test Cnt:", i)
		a := nasMessage.NewAuthenticationResponse(0)
		b := nasMessage.NewAuthenticationResponse(0)
		assert.NotNil(t, a)
		assert.NotNil(t, b)

		a.ExtendedProtocolDiscriminator.SetExtendedProtocolDiscriminator(table.inExtendedProtocolDiscriminator)
		a.SpareHalfOctetAndSecurityHeaderType.SetSecurityHeaderType(table.inSecurityHeader)
		a.SpareHalfOctetAndSecurityHeaderType.SetSpareHalfOctet(table.inSpareHalfOctet)
		a.AuthenticationResponseMessageIdentity.SetMessageType(table.inAuthenticationResponseMessageIdentity)

		a.AuthenticationResponseParameter = nasType.NewAuthenticationResponseParameter(nasMessage.AuthenticationResponseAuthenticationResponseParameterType)
		a.AuthenticationResponseParameter = &table.inAuthenticationResponseParameter

		a.EAPMessage = nasType.NewEAPMessage(nasMessage.AuthenticationResponseEAPMessageType)
		a.EAPMessage = &table.inEAPMessage

		buff := new(bytes.Buffer)
		a.EncodeAuthenticationResponse(buff)
		logger.NasMsgLog.Debugln("Encode: ", a)

		data := make([]byte, buff.Len())
		buff.Read(data)
		b.DecodeAuthenticationResponse(&data)
		logger.NasMsgLog.Debugln(data)
		logger.NasMsgLog.Debugln("Decode: ", b)

		// Compare the re-encoded wire form rather than the structs. Decoding
		// populates the enrichment fields (EPD, MessageType, Cause, ...) that
		// the hand-built value `a` never has, so reflect.DeepEqual(a, b) can
		// never hold. Re-encoding b and comparing bytes is the round-trip
		// property these tests actually mean to assert.
		reBuff := new(bytes.Buffer)
		if err := b.EncodeAuthenticationResponse(reBuff); err != nil {
			t.Fatalf("re-encode: %v", err)
		}
		if !bytes.Equal(data, reBuff.Bytes()) {
			t.Errorf("round trip mismatch:\n encoded: %v\n re-encoded: %v", data, reBuff.Bytes())
		}

	}
}
