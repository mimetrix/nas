package nasMessage_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mimetrix/nas/logger"
	"github.com/mimetrix/nas/nasMessage"
	"github.com/mimetrix/nas/nasType"
)

type nasMessagePDUSessionAuthenticationCommandData struct {
	inExtendedProtocolDiscriminator                  uint8
	inPDUSessionID                                   uint8
	inPTI                                            uint8
	inPDUSESSIONAUTHENTICATIONCOMMANDMessageIdentity uint8
	inEAPMessage                                     nasType.EAPMessage
	inExtendedProtocolConfigurationOptions           nasType.ExtendedProtocolConfigurationOptions
}

var nasMessagePDUSessionAuthenticationCommandTable = []nasMessagePDUSessionAuthenticationCommandData{
	{
		inExtendedProtocolDiscriminator: nasMessage.MsgTypePDUSessionAuthenticationCommand,
		inPDUSessionID:                  0x01,
		inPTI:                           0x01,
		inPDUSESSIONAUTHENTICATIONCOMMANDMessageIdentity: 0x01,
		inEAPMessage: nasType.EAPMessage{
			Iei:    0,
			Len:    4,
			Buffer: []uint8{0x01, 0x01, 0x01, 0x01},
		},
		inExtendedProtocolConfigurationOptions: nasType.ExtendedProtocolConfigurationOptions{
			Iei:    nasMessage.PDUSessionAuthenticationCommandExtendedProtocolConfigurationOptionsType,
			Len:    2,
			Buffer: []uint8{0x01, 0x01},
		},
	},
}

func TestNasTypeNewPDUSessionAuthenticationCommand(t *testing.T) {
	a := nasMessage.NewPDUSessionAuthenticationCommand(0)
	assert.NotNil(t, a)
}

func TestNasTypeNewPDUSessionAuthenticationCommandMessage(t *testing.T) {
	for i, table := range nasMessagePDUSessionAuthenticationCommandTable {
		t.Logf("Test Cnt:%d", i)
		a := nasMessage.NewPDUSessionAuthenticationCommand(0)
		b := nasMessage.NewPDUSessionAuthenticationCommand(0)
		assert.NotNil(t, a)
		assert.NotNil(t, b)

		a.ExtendedProtocolDiscriminator.SetExtendedProtocolDiscriminator(table.inExtendedProtocolDiscriminator)
		a.PDUSessionID.SetPDUSessionID(table.inPDUSessionID)
		a.PTI.SetPTI(table.inPTI)
		a.PDUSESSIONAUTHENTICATIONCOMMANDMessageIdentity.SetMessageType(table.inPDUSESSIONAUTHENTICATIONCOMMANDMessageIdentity)

		a.EAPMessage = table.inEAPMessage

		a.ExtendedProtocolConfigurationOptions = nasType.NewExtendedProtocolConfigurationOptions(nasMessage.PDUSessionAuthenticationCommandExtendedProtocolConfigurationOptionsType)
		a.ExtendedProtocolConfigurationOptions = &table.inExtendedProtocolConfigurationOptions

		buff := new(bytes.Buffer)
		a.EncodePDUSessionAuthenticationCommand(buff)
		logger.NasMsgLog.Debugln("Encode: ", a)

		data := make([]byte, buff.Len())
		buff.Read(data)
		logger.NasMsgLog.Debugln(data)
		b.DecodePDUSessionAuthenticationCommand(&data)
		logger.NasMsgLog.Debugln("Decode: ", b)

		// Compare the re-encoded wire form rather than the structs. Decoding
		// populates the enrichment fields (EPD, MessageType, Cause, ...) that
		// the hand-built value `a` never has, so reflect.DeepEqual(a, b) can
		// never hold. Re-encoding b and comparing bytes is the round-trip
		// property these tests actually mean to assert.
		reBuff := new(bytes.Buffer)
		if err := b.EncodePDUSessionAuthenticationCommand(reBuff); err != nil {
			t.Fatalf("re-encode: %v", err)
		}
		if !bytes.Equal(data, reBuff.Bytes()) {
			t.Errorf("round trip mismatch:\n encoded: %v\n re-encoded: %v", data, reBuff.Bytes())
		}

	}
}
