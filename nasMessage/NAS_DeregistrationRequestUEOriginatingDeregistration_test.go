package nasMessage_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mimetrix/nas/logger"
	"github.com/mimetrix/nas/nasMessage"
	"github.com/mimetrix/nas/nasType"
)

type nasMessageDeregistrationRequestUEOriginatingDeregistrationData struct {
	inExtendedProtocolDiscriminator        uint8
	inSecurityHeaderType                   uint8
	inSpareHalfOctet                       uint8
	inDeregistrationRequestMessageIdentity uint8
	inNgksiAndDeregistrationType           nasType.NgksiAndDeregistrationType
	inMobileIdentity5GS                    nasType.MobileIdentity5GS
}

var nasMessageDeregistrationRequestUEOriginatingDeregistrationTable = []nasMessageDeregistrationRequestUEOriginatingDeregistrationData{
	{
		inExtendedProtocolDiscriminator:        nasMessage.Epd5GSSessionManagementMessage,
		inSecurityHeaderType:                   0x01,
		inSpareHalfOctet:                       0x01,
		inDeregistrationRequestMessageIdentity: 0x01,
		inNgksiAndDeregistrationType: nasType.NgksiAndDeregistrationType{
			Octet: 0xFF,
		},
		inMobileIdentity5GS: nasType.MobileIdentity5GS{
			Iei:    0,
			Len:    4,
			Buffer: []uint8{0x01, 0x01, 0x01, 0x01},
		},
	},
}

func TestNasTypeNewDeregistrationRequestUEOriginatingDeregistration(t *testing.T) {
	a := nasMessage.NewDeregistrationRequestUEOriginatingDeregistration(0)
	assert.NotNil(t, a)
}

func TestNasTypeNewDeregistrationRequestUEOriginatingDeregistrationMessage(t *testing.T) {
	for i, table := range nasMessageDeregistrationRequestUEOriginatingDeregistrationTable {
		logger.NasMsgLog.Infoln("Test Cnt:", i)
		a := nasMessage.NewDeregistrationRequestUEOriginatingDeregistration(0)
		b := nasMessage.NewDeregistrationRequestUEOriginatingDeregistration(0)
		assert.NotNil(t, a)
		assert.NotNil(t, b)

		a.ExtendedProtocolDiscriminator.SetExtendedProtocolDiscriminator(table.inExtendedProtocolDiscriminator)
		a.SpareHalfOctetAndSecurityHeaderType.SetSecurityHeaderType(table.inSecurityHeaderType)
		a.SpareHalfOctetAndSecurityHeaderType.SetSpareHalfOctet(table.inSpareHalfOctet)
		a.DeregistrationRequestMessageIdentity.SetMessageType(table.inDeregistrationRequestMessageIdentity)

		a.NgksiAndDeregistrationType = table.inNgksiAndDeregistrationType

		a.MobileIdentity5GS = table.inMobileIdentity5GS

		buff := new(bytes.Buffer)
		a.EncodeDeregistrationRequestUEOriginatingDeregistration(buff)
		logger.NasMsgLog.Debugln("Encode: ", a)

		data := make([]byte, buff.Len())
		buff.Read(data)
		logger.NasMsgLog.Debugln(data)
		b.DecodeDeregistrationRequestUEOriginatingDeregistration(&data)
		logger.NasMsgLog.Debugln("Decode: ", b)

		// Compare the re-encoded wire form rather than the structs. Decoding
		// populates the enrichment fields (EPD, MessageType, Cause, ...) that
		// the hand-built value `a` never has, so reflect.DeepEqual(a, b) can
		// never hold. Re-encoding b and comparing bytes is the round-trip
		// property these tests actually mean to assert.
		reBuff := new(bytes.Buffer)
		if err := b.EncodeDeregistrationRequestUEOriginatingDeregistration(reBuff); err != nil {
			t.Fatalf("re-encode: %v", err)
		}
		if !bytes.Equal(data, reBuff.Bytes()) {
			t.Errorf("round trip mismatch:\n encoded: %v\n re-encoded: %v", data, reBuff.Bytes())
		}

	}
}
