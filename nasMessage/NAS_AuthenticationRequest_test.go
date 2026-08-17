package nasMessage_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mimetrix/nas/logger"
	"github.com/mimetrix/nas/nasMessage"
	"github.com/mimetrix/nas/nasType"
)

type nasMessageAuthenticationRequestData struct {
	inExtendedProtocolDiscriminator        uint8
	inSecurityHeader                       uint8
	inSpareHalfOctet1                      uint8
	inAuthenticationRequestMessageIdentity uint8
	inTsc                                  uint8
	inNASKeySetIdentifier                  uint8
	inSpareHalfOctet2                      uint8
	inABBA                                 nasType.ABBA
	inAuthenticationParameterRAND          nasType.AuthenticationParameterRAND
	inAuthenticationParameterAUTN          nasType.AuthenticationParameterAUTN
	inEAPMessage                           nasType.EAPMessage
}

var nasMessageAuthenticationRequestTable = []nasMessageAuthenticationRequestData{
	{
		inExtendedProtocolDiscriminator:        0x01,
		inSecurityHeader:                       0x08,
		inSpareHalfOctet1:                      0x01,
		inAuthenticationRequestMessageIdentity: 0x01,
		inTsc:                                  0x01,
		inNASKeySetIdentifier:                  0x07,
		inSpareHalfOctet2:                      0x07,
		inABBA:                                 nasType.ABBA{Iei: 0, Len: 2, Buffer: []byte{0x00, 0x00}},
		inAuthenticationParameterRAND:          nasType.AuthenticationParameterRAND{Iei: nasMessage.AuthenticationRequestAuthenticationParameterRANDType, Octet: [16]uint8{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}},
		inAuthenticationParameterAUTN:          nasType.AuthenticationParameterAUTN{Iei: nasMessage.AuthenticationRequestAuthenticationParameterAUTNType, Len: 16, Octet: [16]uint8{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}},
		inEAPMessage:                           nasType.EAPMessage{Iei: nasMessage.AuthenticationRequestEAPMessageType, Len: 4, Buffer: []byte{0x00, 0x00, 0x00, 0x00}},
	},
}

func TestNasTypeNewAuthenticationRequest(t *testing.T) {
	a := nasMessage.NewAuthenticationRequest(0)
	assert.NotNil(t, a)
}

func TestNasTypeNewAuthenticationRequestMessage(t *testing.T) {
	logger.NasMsgLog.Infoln("---Test NAS Message: AuthenticationRequestMessage---")
	for i, table := range nasMessageAuthenticationRequestTable {
		t.Logf("Test Cnt:%d", i)
		a := nasMessage.NewAuthenticationRequest(0)
		b := nasMessage.NewAuthenticationRequest(0)
		assert.NotNil(t, a)
		assert.NotNil(t, b)

		a.ExtendedProtocolDiscriminator.SetExtendedProtocolDiscriminator(table.inExtendedProtocolDiscriminator)
		a.SpareHalfOctetAndSecurityHeaderType.SetSecurityHeaderType(table.inSecurityHeader)
		a.SpareHalfOctetAndSecurityHeaderType.SetSpareHalfOctet(table.inSpareHalfOctet1)
		a.AuthenticationRequestMessageIdentity.SetMessageType(table.inAuthenticationRequestMessageIdentity)

		a.ABBA = table.inABBA

		a.AuthenticationParameterRAND = nasType.NewAuthenticationParameterRAND(nasMessage.AuthenticationRequestAuthenticationParameterRANDType)
		a.AuthenticationParameterRAND = &table.inAuthenticationParameterRAND

		a.AuthenticationParameterAUTN = nasType.NewAuthenticationParameterAUTN(nasMessage.AuthenticationRequestAuthenticationParameterAUTNType)
		a.AuthenticationParameterAUTN = &table.inAuthenticationParameterAUTN

		a.EAPMessage = nasType.NewEAPMessage(nasMessage.AuthenticationRequestEAPMessageType)
		a.EAPMessage = &table.inEAPMessage

		buff := new(bytes.Buffer)
		a.EncodeAuthenticationRequest(buff)
		// fmt.Printf("Encode: %x\n", buff)
		logger.NasMsgLog.Debugln("Encode: ", a)

		data := make([]byte, buff.Len())
		buff.Read(data)
		b.DecodeAuthenticationRequest(&data)
		// fmt.Printf("Decode: %x\n", data)
		logger.NasMsgLog.Debugln("Decode: ", b)

		// Compare the re-encoded wire form rather than the structs. Decoding
		// populates the enrichment fields (EPD, MessageType, Cause, ...) that
		// the hand-built value `a` never has, so reflect.DeepEqual(a, b) can
		// never hold. Re-encoding b and comparing bytes is the round-trip
		// property these tests actually mean to assert.
		reBuff := new(bytes.Buffer)
		if err := b.EncodeAuthenticationRequest(reBuff); err != nil {
			t.Fatalf("re-encode: %v", err)
		}
		if !bytes.Equal(data, reBuff.Bytes()) {
			t.Errorf("round trip mismatch:\n encoded: %v\n re-encoded: %v", data, reBuff.Bytes())
		}
	}
}
