package nasMessage_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mimetrix/nas/logger"
	"github.com/mimetrix/nas/nasMessage"
	"github.com/mimetrix/nas/nasType"
)

type nasMessageServiceAcceptData struct {
	inExtendedProtocolDiscriminator          uint8
	inSecurityHeader                         uint8
	inSpareHalfOctet                         uint8
	inServiceAcceptMessageIdentity           uint8
	inPDUSessionStatus                       nasType.PDUSessionStatus
	inPDUSessionReactivationResult           nasType.PDUSessionReactivationResult
	inPDUSessionReactivationResultErrorCause nasType.PDUSessionReactivationResultErrorCause
	inEAPMessage                             nasType.EAPMessage
}

var nasMessageServiceAcceptTable = []nasMessageServiceAcceptData{
	{
		inExtendedProtocolDiscriminator: nasMessage.Epd5GSMobilityManagementMessage,
		inSecurityHeader:                0x01,
		inSpareHalfOctet:                0x01,
		inServiceAcceptMessageIdentity:  nasMessage.MsgTypeServiceAccept,
		inPDUSessionStatus: nasType.PDUSessionStatus{
			Iei:    nasMessage.ServiceAcceptPDUSessionStatusType,
			Len:    2,
			Buffer: []uint8{0x01, 0x01},
		},
		inPDUSessionReactivationResult: nasType.PDUSessionReactivationResult{
			Iei:    nasMessage.ServiceAcceptPDUSessionReactivationResultType,
			Len:    2,
			Buffer: []uint8{0x01, 0x01},
		},
		inPDUSessionReactivationResultErrorCause: nasType.PDUSessionReactivationResultErrorCause{
			Iei:    nasMessage.ServiceAcceptPDUSessionReactivationResultErrorCauseType,
			Len:    2,
			Buffer: []uint8{0x01, 0x01},
		},
		inEAPMessage: nasType.EAPMessage{
			Iei:    nasMessage.ServiceAcceptEAPMessageType,
			Len:    4,
			Buffer: []uint8{0x01, 0x01, 0x01, 0x01},
		},
	},
}

func TestNasTypeNewServiceAccept(t *testing.T) {
	a := nasMessage.NewServiceAccept(0)
	assert.NotNil(t, a)
}

func TestNasTypeNewServiceAcceptMessage(t *testing.T) {
	for i, table := range nasMessageServiceAcceptTable {
		t.Logf("Test Cnt:%d", i)
		a := nasMessage.NewServiceAccept(0)
		b := nasMessage.NewServiceAccept(0)
		assert.NotNil(t, a)
		assert.NotNil(t, b)

		a.ExtendedProtocolDiscriminator.SetExtendedProtocolDiscriminator(table.inExtendedProtocolDiscriminator)
		a.SpareHalfOctetAndSecurityHeaderType.SetSecurityHeaderType(table.inSecurityHeader)
		a.SpareHalfOctetAndSecurityHeaderType.SetSpareHalfOctet(table.inSpareHalfOctet)
		a.ServiceAcceptMessageIdentity.SetMessageType(table.inServiceAcceptMessageIdentity)

		a.PDUSessionStatus = nasType.NewPDUSessionStatus(nasMessage.ServiceAcceptPDUSessionStatusType)
		a.PDUSessionStatus = &table.inPDUSessionStatus

		a.PDUSessionReactivationResult = nasType.NewPDUSessionReactivationResult(nasMessage.ServiceAcceptPDUSessionReactivationResultType)
		a.PDUSessionReactivationResult = &table.inPDUSessionReactivationResult

		a.PDUSessionReactivationResultErrorCause = nasType.NewPDUSessionReactivationResultErrorCause(nasMessage.ServiceAcceptPDUSessionReactivationResultErrorCauseType)
		a.PDUSessionReactivationResultErrorCause = &table.inPDUSessionReactivationResultErrorCause

		a.EAPMessage = nasType.NewEAPMessage(nasMessage.ServiceAcceptEAPMessageType)
		a.EAPMessage = &table.inEAPMessage

		buff := new(bytes.Buffer)
		a.EncodeServiceAccept(buff)
		logger.NasMsgLog.Debugln("Encode: ", a)

		data := make([]byte, buff.Len())
		buff.Read(data)
		logger.NasMsgLog.Debugln(data)
		b.DecodeServiceAccept(&data)
		logger.NasMsgLog.Debugln("Dncode: ", b)

		// Compare the re-encoded wire form rather than the structs. Decoding
		// populates the enrichment fields (EPD, MessageType, Cause, ...) that
		// the hand-built value `a` never has, so reflect.DeepEqual(a, b) can
		// never hold. Re-encoding b and comparing bytes is the round-trip
		// property these tests actually mean to assert.
		reBuff := new(bytes.Buffer)
		if err := b.EncodeServiceAccept(reBuff); err != nil {
			t.Fatalf("re-encode: %v", err)
		}
		if !bytes.Equal(data, reBuff.Bytes()) {
			t.Errorf("round trip mismatch:\n encoded: %v\n re-encoded: %v", data, reBuff.Bytes())
		}
	}
}
