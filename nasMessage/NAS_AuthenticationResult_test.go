package nasMessage_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mimetrix/nas/logger"
	"github.com/mimetrix/nas/nasMessage"
	"github.com/mimetrix/nas/nasType"
)

type nasMessageAuthenticationResultData struct {
	inExtendedProtocolDiscriminator uint8
	inSecurityHeaderType            uint8
	inMessageType                   uint8
	inTsc                           uint8
	inNASKeySetIdentifier           uint8
	inEAPLen                        uint16
	inEAPMessage                    []uint8
	inABBA                          nasType.ABBA
}

var aBBATestData = []nasType.ABBA{
	{Iei: nasMessage.AuthenticationResultABBAType, Len: 2, Buffer: []byte{0x00, 0x00}},
	//{Iei: 0x81, Len: 2, Buffer: []byte{0x00, 0x00}},
}

var nasMessageAuthenticationResultTable = []nasMessageAuthenticationResultData{
	{
		inExtendedProtocolDiscriminator: nasMessage.Epd5GSSessionManagementMessage,
		inSecurityHeaderType:            0x01,
		inMessageType:                   nasMessage.MsgTypeAuthenticationResult,
		inTsc:                           0x01,
		inNASKeySetIdentifier:           0x01,
		inEAPLen:                        0x04,
		inEAPMessage:                    []uint8{0x10, 0x11, 0x10, 0x11},
		inABBA:                          aBBATestData[0],
	},
	/*{inExtendedProtocolDiscriminator: nasMessage.Epd5GSSessionManagementMessage,
	inSecurityHeaderType:  0x01,
	inMessageType:         nasMessage.MsgTypeAuthenticationResult,
	inTsc:                 0x01,
	inNASKeySetIdentifier: 0x01,
	inEAPLen:              0x02,
	inEAPMessage:          []uint8{0x10, 0x11},
	inABBA:                aBBATestData[1]},*/
}

func TestNasTypeNewAuthenticationResult(t *testing.T) {
	a := nasMessage.NewAuthenticationResult(0)
	assert.NotNil(t, a)
}

func TestNasTypeNewAuthenticationResultMessage(t *testing.T) {
	for i, table := range nasMessageAuthenticationResultTable {
		logger.NasMsgLog.Infoln("Test Cnt:", i)
		a := nasMessage.NewAuthenticationResult(0)
		b := nasMessage.NewAuthenticationResult(0)
		assert.NotNil(t, a)
		assert.NotNil(t, b)

		a.ExtendedProtocolDiscriminator.SetExtendedProtocolDiscriminator(table.inExtendedProtocolDiscriminator)
		a.SpareHalfOctetAndSecurityHeaderType.SetSecurityHeaderType(table.inSecurityHeaderType)
		a.AuthenticationResultMessageIdentity.SetMessageType(table.inMessageType)
		a.SpareHalfOctetAndNgksi.SetTSC(table.inTsc)
		a.SpareHalfOctetAndNgksi.SetNasKeySetIdentifiler(table.inNASKeySetIdentifier)
		a.EAPMessage.SetLen(table.inEAPLen)
		a.EAPMessage.SetEAPMessage(table.inEAPMessage)

		a.ABBA = nasType.NewABBA(nasMessage.AuthenticationResultABBAType)
		a.ABBA = &table.inABBA

		buff := new(bytes.Buffer)
		a.EncodeAuthenticationResult(buff)
		logger.NasMsgLog.Debugln(buff)

		data := make([]byte, buff.Len())
		buff.Read(data)
		b.DecodeAuthenticationResult(&data)
		logger.NasMsgLog.Debugln(data)
		logger.NasMsgLog.Debugln("Decode: ", b)

		// Compare the re-encoded wire form rather than the structs. Decoding
		// populates the enrichment fields (EPD, MessageType, Cause, ...) that
		// the hand-built value `a` never has, so reflect.DeepEqual(a, b) can
		// never hold. Re-encoding b and comparing bytes is the round-trip
		// property these tests actually mean to assert.
		reBuff := new(bytes.Buffer)
		if err := b.EncodeAuthenticationResult(reBuff); err != nil {
			t.Fatalf("re-encode: %v", err)
		}
		if !bytes.Equal(data, reBuff.Bytes()) {
			t.Errorf("round trip mismatch:\n encoded: %v\n re-encoded: %v", data, reBuff.Bytes())
		}
	}
}
