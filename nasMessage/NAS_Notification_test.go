package nasMessage_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"

	//"fmt"
	"github.com/mimetrix/nas/logger"
	"github.com/mimetrix/nas/nasMessage"
)

type nasMessageNotificationData struct {
	inExtendedProtocolDiscriminator uint8
	inSecurityHeader                uint8
	inSpareHalfOctet1               uint8
	inNotificationMessageIdentity   uint8
	inAccessType                    uint8
	inSpareHalfOctet2               uint8
}

var nasMessageNotificationTable = []nasMessageNotificationData{
	{
		inExtendedProtocolDiscriminator: 0x01,
		inSecurityHeader:                0x08,
		inSpareHalfOctet1:               0x01,
		inNotificationMessageIdentity:   nasMessage.MsgTypeNotification,
		inAccessType:                    0x01,
		inSpareHalfOctet2:               0x01,
	},
}

func TestNasTypeNewNotification(t *testing.T) {
	a := nasMessage.NewNotification(0)
	assert.NotNil(t, a)
}

func TestNasTypeNewNotificationMessage(t *testing.T) {
	for i, table := range nasMessageNotificationTable {
		t.Logf("Test Cnt:%d", i)
		a := nasMessage.NewNotification(0)
		b := nasMessage.NewNotification(0)
		assert.NotNil(t, a)
		assert.NotNil(t, b)

		a.ExtendedProtocolDiscriminator.SetExtendedProtocolDiscriminator(table.inExtendedProtocolDiscriminator)
		a.SpareHalfOctetAndSecurityHeaderType.SetSecurityHeaderType(table.inSecurityHeader)
		a.SpareHalfOctetAndSecurityHeaderType.SetSpareHalfOctet(table.inSpareHalfOctet1)
		a.NotificationMessageIdentity.SetMessageType(table.inNotificationMessageIdentity)
		a.SpareHalfOctetAndAccessType.SetAccessType(table.inAccessType)

		buff := new(bytes.Buffer)
		a.EncodeNotification(buff)
		logger.NasMsgLog.Debugln("Encode: ", a)

		data := make([]byte, buff.Len())
		buff.Read(data)
		b.DecodeNotification(&data)
		logger.NasMsgLog.Debugln(data)
		logger.NasMsgLog.Debugln("Decode: ", b)

		// Compare the re-encoded wire form rather than the structs. Decoding
		// populates the enrichment fields (EPD, MessageType, Cause, ...) that
		// the hand-built value `a` never has, so reflect.DeepEqual(a, b) can
		// never hold. Re-encoding b and comparing bytes is the round-trip
		// property these tests actually mean to assert.
		reBuff := new(bytes.Buffer)
		if err := b.EncodeNotification(reBuff); err != nil {
			t.Fatalf("re-encode: %v", err)
		}
		if !bytes.Equal(data, reBuff.Bytes()) {
			t.Errorf("round trip mismatch:\n encoded: %v\n re-encoded: %v", data, reBuff.Bytes())
		}

	}
}
