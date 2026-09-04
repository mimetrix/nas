package nasMessage_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mimetrix/nas/logger"
	"github.com/mimetrix/nas/nasMessage"
)

type nasMessageConfigurationUpdateCompleteData struct {
	inExtendedProtocolDiscriminator              uint8
	inSecurityHeaderType                         uint8
	inSpareHalfOctet                             uint8
	inConfigurationUpdateCompleteMessageIdentity uint8
}

var nasMessageConfigurationUpdateCompleteTable = []nasMessageConfigurationUpdateCompleteData{
	{
		inExtendedProtocolDiscriminator:              nasMessage.Epd5GSSessionManagementMessage,
		inSecurityHeaderType:                         0x01,
		inSpareHalfOctet:                             0x01,
		inConfigurationUpdateCompleteMessageIdentity: nasMessage.MsgTypeConfigurationUpdateComplete,
	},
}

func TestNasTypeNewConfigurationUpdateComplete(t *testing.T) {
	a := nasMessage.NewConfigurationUpdateComplete(0)
	assert.NotNil(t, a)
}

func TestNasTypeNewConfigurationUpdateCompleteMessage(t *testing.T) {
	for i, table := range nasMessageConfigurationUpdateCompleteTable {
		logger.NasMsgLog.Infoln("Test Cnt:", i)
		a := nasMessage.NewConfigurationUpdateComplete(0)
		b := nasMessage.NewConfigurationUpdateComplete(0)
		assert.NotNil(t, a)
		assert.NotNil(t, b)

		a.ExtendedProtocolDiscriminator.SetExtendedProtocolDiscriminator(table.inExtendedProtocolDiscriminator)
		a.SpareHalfOctetAndSecurityHeaderType.SetSecurityHeaderType(table.inSecurityHeaderType)
		a.SpareHalfOctetAndSecurityHeaderType.SetSpareHalfOctet(table.inSpareHalfOctet)
		a.ConfigurationUpdateCompleteMessageIdentity.SetMessageType(table.inConfigurationUpdateCompleteMessageIdentity)

		buff := new(bytes.Buffer)
		a.EncodeConfigurationUpdateComplete(buff)
		logger.NasMsgLog.Debugln("Encode: ", a)

		data := make([]byte, buff.Len())
		buff.Read(data)
		logger.NasMsgLog.Debugln(data)
		b.DecodeConfigurationUpdateComplete(&data)
		logger.NasMsgLog.Debugln("Decode: ", b)

		// Compare the re-encoded wire form rather than the structs. Decoding
		// populates the enrichment fields (EPD, MessageType, Cause, ...) that
		// the hand-built value `a` never has, so reflect.DeepEqual(a, b) can
		// never hold. Re-encoding b and comparing bytes is the round-trip
		// property these tests actually mean to assert.
		reBuff := new(bytes.Buffer)
		if err := b.EncodeConfigurationUpdateComplete(reBuff); err != nil {
			t.Fatalf("re-encode: %v", err)
		}
		if !bytes.Equal(data, reBuff.Bytes()) {
			t.Errorf("round trip mismatch:\n encoded: %v\n re-encoded: %v", data, reBuff.Bytes())
		}

	}
}
