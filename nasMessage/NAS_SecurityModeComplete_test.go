package nasMessage_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mimetrix/nas/logger"
	"github.com/mimetrix/nas/nasMessage"
	"github.com/mimetrix/nas/nasType"
)

type nasMessageSecurityModeCompleteData struct {
	inExtendedProtocolDiscriminator       uint8
	inSecurityHeader                      uint8
	inSpareHalfOctet                      uint8
	inSecurityModeCompleteMessageIdentity uint8
	inIMEISV                              nasType.IMEISV
	inNASMessageContainer                 nasMessage.NASMessageContainer
}

var nasMessageSecurityModeCompleteTable = []nasMessageSecurityModeCompleteData{
	{
		inExtendedProtocolDiscriminator:       nasMessage.Epd5GSMobilityManagementMessage,
		inSecurityHeader:                      0x01,
		inSpareHalfOctet:                      0x01,
		inSecurityModeCompleteMessageIdentity: nasMessage.MsgTypeSecurityModeComplete,
		inIMEISV: nasType.IMEISV{
			Iei:   nasMessage.SecurityModeCompleteIMEISVType,
			Len:   9,
			Octet: [9]uint8{0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01},
		},
		inNASMessageContainer: nasMessage.NASMessageContainer{
			Iei:    nasMessage.SecurityModeCompleteNASMessageContainerType,
			Len:    2,
			Buffer: []uint8{0x01, 0x01},
		},
	},
}

func TestNasTypeNewSecurityModeComplete(t *testing.T) {
	a := nasMessage.NewSecurityModeComplete(0)
	assert.NotNil(t, a)
}

func TestNasTypeNewSecurityModeCompleteMessage(t *testing.T) {
	for i, table := range nasMessageSecurityModeCompleteTable {
		t.Logf("Test Cnt:%d", i)
		a := nasMessage.NewSecurityModeComplete(0)
		b := nasMessage.NewSecurityModeComplete(0)
		assert.NotNil(t, a)
		assert.NotNil(t, b)

		a.ExtendedProtocolDiscriminator.SetExtendedProtocolDiscriminator(table.inExtendedProtocolDiscriminator)
		a.SpareHalfOctetAndSecurityHeaderType.SetSecurityHeaderType(table.inSecurityHeader)
		a.SpareHalfOctetAndSecurityHeaderType.SetSpareHalfOctet(table.inSpareHalfOctet)
		a.SecurityModeCompleteMessageIdentity.SetMessageType(table.inSecurityModeCompleteMessageIdentity)

		a.IMEISV = nasType.NewIMEISV(nasMessage.SecurityModeCompleteIMEISVType)
		a.IMEISV = &table.inIMEISV

		a.NASMessageContainer = nasMessage.NewNASMessageContainer(nasMessage.SecurityModeCompleteNASMessageContainerType)
		a.NASMessageContainer = &table.inNASMessageContainer

		buff := new(bytes.Buffer)
		a.EncodeSecurityModeComplete(buff)
		logger.NasMsgLog.Debugln("Encode: ", a)

		data := make([]byte, buff.Len())
		buff.Read(data)
		logger.NasMsgLog.Debugln(data)
		b.DecodeSecurityModeComplete(&data)
		logger.NasMsgLog.Debugln("Decode: ", b)

		// Compare the re-encoded wire form rather than the structs. Decoding
		// populates the enrichment fields (EPD, MessageType, Cause, ...) that
		// the hand-built value `a` never has, so reflect.DeepEqual(a, b) can
		// never hold. Re-encoding b and comparing bytes is the round-trip
		// property these tests actually mean to assert.
		reBuff := new(bytes.Buffer)
		if err := b.EncodeSecurityModeComplete(reBuff); err != nil {
			t.Fatalf("re-encode: %v", err)
		}
		if !bytes.Equal(data, reBuff.Bytes()) {
			t.Errorf("round trip mismatch:\n encoded: %v\n re-encoded: %v", data, reBuff.Bytes())
		}
	}
}
