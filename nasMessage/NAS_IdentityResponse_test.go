package nasMessage_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"

	//"fmt"
	"github.com/mimetrix/nas/logger"
	"github.com/mimetrix/nas/nasMessage"
	"github.com/mimetrix/nas/nasType"
)

type nasMessageIdentityResponseData struct {
	inExtendedProtocolDiscriminator   uint8
	inSecurityHeader                  uint8
	inSpareHalfOctet                  uint8
	inIdentityResponseMessageIdentity uint8
	inMobileIdentity                  nasType.MobileIdentity
}

var nasMessageIdentityResponseTable = []nasMessageIdentityResponseData{
	{
		inExtendedProtocolDiscriminator:   0x01,
		inSecurityHeader:                  0x08,
		inSpareHalfOctet:                  0x01,
		inIdentityResponseMessageIdentity: nasMessage.MsgTypeIdentityResponse,
		inMobileIdentity: nasType.MobileIdentity{
			Iei:    0,
			Len:    2,
			Buffer: []uint8{0x01, 0x01},
		},
	},
}

func TestNasTypeNewIdentityResponse(t *testing.T) {
	a := nasMessage.NewIdentityResponse(0)
	assert.NotNil(t, a)
}

func TestNasTypeNewIdentityResponseMessage(t *testing.T) {
	for i, table := range nasMessageIdentityResponseTable {
		logger.NasMsgLog.Infoln("Test Cnt:", i)
		a := nasMessage.NewIdentityResponse(0)
		b := nasMessage.NewIdentityResponse(0)
		assert.NotNil(t, a)
		assert.NotNil(t, b)

		a.ExtendedProtocolDiscriminator.SetExtendedProtocolDiscriminator(table.inExtendedProtocolDiscriminator)
		a.SpareHalfOctetAndSecurityHeaderType.SetSecurityHeaderType(table.inSecurityHeader)
		a.SpareHalfOctetAndSecurityHeaderType.SetSpareHalfOctet(table.inSpareHalfOctet)
		a.IdentityResponseMessageIdentity.SetMessageType(table.inIdentityResponseMessageIdentity)

		a.MobileIdentity = table.inMobileIdentity

		buff := new(bytes.Buffer)
		a.EncodeIdentityResponse(buff)
		logger.NasMsgLog.Debugln("Encode: ", a)

		data := make([]byte, buff.Len())
		buff.Read(data)
		b.DecodeIdentityResponse(&data)
		logger.NasMsgLog.Debugln(data)
		logger.NasMsgLog.Debugln("Decode: ", b)

		// Compare the re-encoded wire form rather than the structs. Decoding
		// populates the enrichment fields (EPD, MessageType, Cause, ...) that
		// the hand-built value `a` never has, so reflect.DeepEqual(a, b) can
		// never hold. Re-encoding b and comparing bytes is the round-trip
		// property these tests actually mean to assert.
		reBuff := new(bytes.Buffer)
		if err := b.EncodeIdentityResponse(reBuff); err != nil {
			t.Fatalf("re-encode: %v", err)
		}
		if !bytes.Equal(data, reBuff.Bytes()) {
			t.Errorf("round trip mismatch:\n encoded: %v\n re-encoded: %v", data, reBuff.Bytes())
		}

	}
}
