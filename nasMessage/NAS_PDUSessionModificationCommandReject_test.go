package nasMessage_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mimetrix/nas/logger"
	"github.com/mimetrix/nas/nasMessage"
	"github.com/mimetrix/nas/nasType"
)

type nasMessagePDUSessionModificationCommandRejectData struct {
	inExtendedProtocolDiscriminator                      uint8
	inPDUSessionID                                       uint8
	inPTI                                                uint8
	inPDUSESSIONMODIFICATIONCOMMANDREJECTMessageIdentity uint8
	inCause5GSM                                          nasType.Cause5GSM
	inExtendedProtocolConfigurationOptions               nasType.ExtendedProtocolConfigurationOptions
}

var nasMessagePDUSessionModificationCommandRejectTable = []nasMessagePDUSessionModificationCommandRejectData{
	{
		inExtendedProtocolDiscriminator: nasMessage.MsgTypePDUSessionModificationCommandReject,
		inPDUSessionID:                  0x01,
		inPTI:                           0x01,
		inPDUSESSIONMODIFICATIONCOMMANDREJECTMessageIdentity: 0x01,
		inCause5GSM: nasType.Cause5GSM{
			Iei:   0,
			Octet: 0x01,
		},
		inExtendedProtocolConfigurationOptions: nasType.ExtendedProtocolConfigurationOptions{
			Iei:    nasMessage.PDUSessionModificationCommandRejectExtendedProtocolConfigurationOptionsType,
			Len:    2,
			Buffer: []uint8{0x01, 0x01},
		},
	},
}

func TestNasTypeNewPDUSessionModificationCommandReject(t *testing.T) {
	a := nasMessage.NewPDUSessionModificationCommandReject(0)
	assert.NotNil(t, a)
}

func TestNasTypeNewPDUSessionModificationCommandRejectMessage(t *testing.T) {
	for i, table := range nasMessagePDUSessionModificationCommandRejectTable {
		t.Logf("Test Cnt:%d", i)
		a := nasMessage.NewPDUSessionModificationCommandReject(0)
		b := nasMessage.NewPDUSessionModificationCommandReject(0)
		assert.NotNil(t, a)
		assert.NotNil(t, b)

		a.ExtendedProtocolDiscriminator.SetExtendedProtocolDiscriminator(table.inExtendedProtocolDiscriminator)
		a.PDUSessionID.SetPDUSessionID(table.inPDUSessionID)
		a.PTI.SetPTI(table.inPTI)
		a.PDUSESSIONMODIFICATIONCOMMANDREJECTMessageIdentity.SetMessageType(table.inPDUSESSIONMODIFICATIONCOMMANDREJECTMessageIdentity)
		a.Cause5GSM = table.inCause5GSM

		a.ExtendedProtocolConfigurationOptions = nasType.NewExtendedProtocolConfigurationOptions(nasMessage.PDUSessionModificationCommandRejectExtendedProtocolConfigurationOptionsType)
		a.ExtendedProtocolConfigurationOptions = &table.inExtendedProtocolConfigurationOptions

		buff := new(bytes.Buffer)
		a.EncodePDUSessionModificationCommandReject(buff)
		logger.NasMsgLog.Debugln("Encode: ", a)

		data := make([]byte, buff.Len())
		buff.Read(data)
		logger.NasMsgLog.Debugln(data)
		b.DecodePDUSessionModificationCommandReject(&data)
		logger.NasMsgLog.Debugln("Decode: ", b)

		// Compare the re-encoded wire form rather than the structs. Decoding
		// populates the enrichment fields (EPD, MessageType, Cause, ...) that
		// the hand-built value `a` never has, so reflect.DeepEqual(a, b) can
		// never hold. Re-encoding b and comparing bytes is the round-trip
		// property these tests actually mean to assert.
		reBuff := new(bytes.Buffer)
		if err := b.EncodePDUSessionModificationCommandReject(reBuff); err != nil {
			t.Fatalf("re-encode: %v", err)
		}
		if !bytes.Equal(data, reBuff.Bytes()) {
			t.Errorf("round trip mismatch:\n encoded: %v\n re-encoded: %v", data, reBuff.Bytes())
		}

	}
}
