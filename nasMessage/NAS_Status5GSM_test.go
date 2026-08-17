package nasMessage_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mimetrix/nas/logger"
	"github.com/mimetrix/nas/nasMessage"
	"github.com/mimetrix/nas/nasType"
)

type nasMessageStatus5GSMData struct {
	inExtendedProtocolDiscriminator uint8
	inPDUSessionID                  nasType.PDUSessionID
	inPTI                           nasType.PTI
	inStatus5GSMMessageIdentity     uint8
	inCause5GSM                     nasType.Cause5GSM
}

var nasMessageStatus5GSMTable = []nasMessageStatus5GSMData{
	{
		inExtendedProtocolDiscriminator: nasMessage.Epd5GSSessionManagementMessage,
		inPDUSessionID: nasType.PDUSessionID{
			Octet: 0x01,
		},
		inPTI: nasType.PTI{
			Octet: 0x01,
		},
		inStatus5GSMMessageIdentity: nasMessage.MsgTypeStatus5GSM,
		inCause5GSM: nasType.Cause5GSM{
			Octet: 0x01,
		},
	},
}

func TestNasTypeNewStatus5GSM(t *testing.T) {
	a := nasMessage.NewStatus5GSM(0)
	assert.NotNil(t, a)
}

func TestNasTypeNewStatus5GSMMessage(t *testing.T) {
	for i, table := range nasMessageStatus5GSMTable {
		t.Logf("Test Cnt:%d", i)
		a := nasMessage.NewStatus5GSM(0)
		b := nasMessage.NewStatus5GSM(0)
		assert.NotNil(t, a)
		assert.NotNil(t, b)

		a.ExtendedProtocolDiscriminator.SetExtendedProtocolDiscriminator(table.inExtendedProtocolDiscriminator)
		a.PDUSessionID = table.inPDUSessionID
		a.PTI = table.inPTI

		a.STATUSMessageIdentity5GSM.SetMessageType(table.inStatus5GSMMessageIdentity)

		a.Cause5GSM = table.inCause5GSM

		buff := new(bytes.Buffer)
		a.EncodeStatus5GSM(buff)
		logger.NasMsgLog.Debugln("Encode: ", a)

		data := make([]byte, buff.Len())
		buff.Read(data)
		logger.NasMsgLog.Debugln(data)
		b.DecodeStatus5GSM(&data)
		logger.NasMsgLog.Debugln("Decode: ", b)

		// Compare the re-encoded wire form rather than the structs. Decoding
		// populates the enrichment fields (EPD, MessageType, Cause, ...) that
		// the hand-built value `a` never has, so reflect.DeepEqual(a, b) can
		// never hold. Re-encoding b and comparing bytes is the round-trip
		// property these tests actually mean to assert.
		reBuff := new(bytes.Buffer)
		if err := b.EncodeStatus5GSM(reBuff); err != nil {
			t.Fatalf("re-encode: %v", err)
		}
		if !bytes.Equal(data, reBuff.Bytes()) {
			t.Errorf("round trip mismatch:\n encoded: %v\n re-encoded: %v", data, reBuff.Bytes())
		}
	}
}
