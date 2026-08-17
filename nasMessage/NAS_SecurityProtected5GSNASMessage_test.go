package nasMessage_test

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mimetrix/nas/logger"
	"github.com/mimetrix/nas/nasMessage"
	"github.com/mimetrix/nas/nasType"
)

type nasMessageSecurityProtected5GSNASMessageData struct {
	inExtendedProtocolDiscriminator uint8
	inSecurityHeader                uint8
	inSpareHalfOctet                uint8
	inMessageAuthenticationCode     nasType.MessageAuthenticationCode
	inSequenceNumber                nasType.SequenceNumber
	inPlainNASMessage               []uint8
}

var nasMessageSecurityProtected5GSNASMessageTable = []nasMessageSecurityProtected5GSNASMessageData{
	{
		inExtendedProtocolDiscriminator: nasMessage.Epd5GSMobilityManagementMessage,
		inSecurityHeader:                0x01,
		inSpareHalfOctet:                0x01,
		inMessageAuthenticationCode: nasType.MessageAuthenticationCode{
			Octet: [4]uint8{0x01, 0x01, 0x01, 0x01},
		},
		inSequenceNumber: nasType.SequenceNumber{
			Octet: 0x01,
		},
		// A minimal plain 5GMM Registration complete as the protected payload.
		inPlainNASMessage: []uint8{0x7e, 0x00, 0x43},
	},
}

func TestNasTypeNewSecurityProtected5GSNASMessage(t *testing.T) {
	a := nasMessage.NewSecurityProtected5GSNASMessage(0)
	assert.NotNil(t, a)
}

func TestNasTypeNewSecurityProtected5GSNASMessageMessage(t *testing.T) {
	for i, table := range nasMessageSecurityProtected5GSNASMessageTable {
		t.Logf("Test Cnt:%d", i)
		a := nasMessage.NewSecurityProtected5GSNASMessage(0)
		b := nasMessage.NewSecurityProtected5GSNASMessage(0)
		assert.NotNil(t, a)
		assert.NotNil(t, b)

		a.ExtendedProtocolDiscriminator.SetExtendedProtocolDiscriminator(table.inExtendedProtocolDiscriminator)
		a.SpareHalfOctetAndSecurityHeaderType.SetSecurityHeaderType(table.inSecurityHeader)
		a.SpareHalfOctetAndSecurityHeaderType.SetSpareHalfOctet(table.inSpareHalfOctet)

		a.MessageAuthenticationCode = table.inMessageAuthenticationCode
		a.SequenceNumber = table.inSequenceNumber

		// Build the wire form: 6-octet security header + protected payload.
		buff := new(bytes.Buffer)
		if err := a.EncodeSecurityProtected5GSNASMessage(buff); err != nil {
			t.Fatalf("encode: %v", err)
		}
		buff.Write(table.inPlainNASMessage)
		logger.NasMsgLog.Debugln("Encode: ", a)

		data := make([]byte, buff.Len())
		buff.Read(data)
		logger.NasMsgLog.Debugln(data)
		if err := b.DecodeSecurityProtected5GSNASMessage(&data); err != nil {
			t.Fatalf("decode: %v", err)
		}
		logger.NasMsgLog.Debugln("Decode: ", b)

		// The security header must survive the round trip. Compare the raw
		// octets only: decoding also populates the enrichment fields (here,
		// the hex MAC string), which the hand-built value `a` never has.
		if !reflect.DeepEqual(a.MessageAuthenticationCode.Octet, b.MessageAuthenticationCode.Octet) {
			t.Errorf("MessageAuthenticationCode mismatch: %v != %v",
				a.MessageAuthenticationCode.Octet, b.MessageAuthenticationCode.Octet)
		}
		if !reflect.DeepEqual(a.SequenceNumber.Octet, b.SequenceNumber.Octet) {
			t.Errorf("SequenceNumber mismatch: %v != %v",
				a.SequenceNumber.Octet, b.SequenceNumber.Octet)
		}

		// Unlike upstream, this fork recursively decodes the protected payload
		// instead of holding it as an opaque blob.
		if b.PlainNASMessage == nil {
			t.Errorf("PlainNASMessage was not decoded")
		}
	}
}
