package nasMessage_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mimetrix/nas/logger"
	"github.com/mimetrix/nas/nasMessage"
	"github.com/mimetrix/nas/nasType"
)

type nasMessagePDUSessionReleaseCompleteData struct {
	inExtendedProtocolDiscriminator            uint8
	inPDUSessionID                             uint8
	inPTI                                      uint8
	inPDUSESSIONRELEASECOMPLETEMessageIdentity uint8
	inCause5GSM                                nasType.Cause5GSM
	inExtendedProtocolConfigurationOptions     nasType.ExtendedProtocolConfigurationOptions
}

var nasMessagePDUSessionReleaseCompleteTable = []nasMessagePDUSessionReleaseCompleteData{
	{
		inExtendedProtocolDiscriminator: nasMessage.Epd5GSSessionManagementMessage,
		inPDUSessionID:                  0x01,
		inPTI:                           0x01,
		inPDUSESSIONRELEASECOMPLETEMessageIdentity: 0x01,
		inCause5GSM: nasType.Cause5GSM{
			Iei:   nasMessage.PDUSessionReleaseCompleteCause5GSMType,
			Octet: 0x01,
		},
		inExtendedProtocolConfigurationOptions: nasType.ExtendedProtocolConfigurationOptions{
			Iei:    nasMessage.PDUSessionReleaseCompleteExtendedProtocolConfigurationOptionsType,
			Len:    2,
			Buffer: []uint8{0x01, 0x01},
		},
	},
}

func TestNasTypeNewPDUSessionReleaseComplete(t *testing.T) {
	a := nasMessage.NewPDUSessionReleaseComplete(0)
	assert.NotNil(t, a)
}

func TestNasTypeNewPDUSessionReleaseCompleteMessage(t *testing.T) {
	for i, table := range nasMessagePDUSessionReleaseCompleteTable {
		t.Logf("Test Cnt:%d", i)
		a := nasMessage.NewPDUSessionReleaseComplete(0)
		b := nasMessage.NewPDUSessionReleaseComplete(0)
		assert.NotNil(t, a)
		assert.NotNil(t, b)

		a.ExtendedProtocolDiscriminator.SetExtendedProtocolDiscriminator(table.inExtendedProtocolDiscriminator)
		a.PDUSessionID.SetPDUSessionID(table.inPDUSessionID)
		a.PTI.SetPTI(table.inPTI)
		a.PDUSESSIONRELEASECOMPLETEMessageIdentity.SetMessageType(table.inPDUSESSIONRELEASECOMPLETEMessageIdentity)

		a.Cause5GSM = nasType.NewCause5GSM(nasMessage.PDUSessionReleaseCompleteCause5GSMType)
		a.Cause5GSM = &table.inCause5GSM

		a.ExtendedProtocolConfigurationOptions = nasType.NewExtendedProtocolConfigurationOptions(nasMessage.PDUSessionReleaseCompleteExtendedProtocolConfigurationOptionsType)
		a.ExtendedProtocolConfigurationOptions = &table.inExtendedProtocolConfigurationOptions

		buff := new(bytes.Buffer)
		a.EncodePDUSessionReleaseComplete(buff)
		logger.NasMsgLog.Debugln("Encode: ", a)

		data := make([]byte, buff.Len())
		buff.Read(data)
		logger.NasMsgLog.Debugln(data)
		b.DecodePDUSessionReleaseComplete(&data)
		logger.NasMsgLog.Debugln("Decode: ", b)
		// fmt.Println(a.Cause5GSM)
		// fmt.Println(b.Cause5GSM)
		// fmt.Println(a.ExtendedProtocolConfigurationOptions)
		// fmt.Println(b.ExtendedProtocolConfigurationOptions)
		// Compare the re-encoded wire form rather than the structs. Decoding
		// populates the enrichment fields (EPD, MessageType, Cause, ...) that
		// the hand-built value `a` never has, so reflect.DeepEqual(a, b) can
		// never hold. Re-encoding b and comparing bytes is the round-trip
		// property these tests actually mean to assert.
		reBuff := new(bytes.Buffer)
		if err := b.EncodePDUSessionReleaseComplete(reBuff); err != nil {
			t.Fatalf("re-encode: %v", err)
		}
		if !bytes.Equal(data, reBuff.Bytes()) {
			t.Errorf("round trip mismatch:\n encoded: %v\n re-encoded: %v", data, reBuff.Bytes())
		}

	}
}
