package nasMessage_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mimetrix/nas/logger"
	"github.com/mimetrix/nas/nasMessage"
	"github.com/mimetrix/nas/nasType"
)

type nasMessageServiceRequestData struct {
	inExtendedProtocolDiscriminator uint8
	inSecurityHeader                uint8
	inSpareHalfOctet                uint8
	inServiceRequestMessageIdentity uint8
	inTMSI5GS                       nasType.TMSI5GS
	inUplinkDataStatus              nasType.UplinkDataStatus
	inPDUSessionStatus              nasType.PDUSessionStatus
	inAllowedPDUSessionStatus       nasType.AllowedPDUSessionStatus
	inNASMessageContainer           nasMessage.NASMessageContainer
}

var nasMessageServiceRequestTable = []nasMessageServiceRequestData{
	{
		inExtendedProtocolDiscriminator: nasMessage.Epd5GSMobilityManagementMessage,
		inSecurityHeader:                0x01,
		inSpareHalfOctet:                0x01,
		inServiceRequestMessageIdentity: nasMessage.MsgTypeServiceRequest,
		inTMSI5GS: nasType.TMSI5GS{
			Len:   7,
			Octet: [7]uint8{0x01, 0x01},
		},
		inUplinkDataStatus: nasType.UplinkDataStatus{
			Iei:    nasMessage.ServiceRequestUplinkDataStatusType,
			Len:    2,
			Buffer: []uint8{0x01, 0x01},
		},
		inPDUSessionStatus: nasType.PDUSessionStatus{
			Iei:    nasMessage.ServiceRequestPDUSessionStatusType,
			Len:    2,
			Buffer: []uint8{0x01, 0x01},
		},
		inAllowedPDUSessionStatus: nasType.AllowedPDUSessionStatus{
			Iei:    nasMessage.ServiceRequestAllowedPDUSessionStatusType,
			Len:    2,
			Buffer: []uint8{0x01, 0x01},
		},
		inNASMessageContainer: nasMessage.NASMessageContainer{
			Iei:    nasMessage.ServiceRequestNASMessageContainerType,
			Len:    2,
			Buffer: []uint8{0x01, 0x01},
		},
	},
}

func TestNasTypeNewServiceRequest(t *testing.T) {
	a := nasMessage.NewServiceRequest(0)
	assert.NotNil(t, a)
}

func TestNasTypeNewServiceRequestMessage(t *testing.T) {
	for i, table := range nasMessageServiceRequestTable {
		t.Logf("Test Cnt:%d", i)
		a := nasMessage.NewServiceRequest(0)
		b := nasMessage.NewServiceRequest(0)
		assert.NotNil(t, a)
		assert.NotNil(t, b)

		a.ExtendedProtocolDiscriminator.SetExtendedProtocolDiscriminator(table.inExtendedProtocolDiscriminator)
		a.SpareHalfOctetAndSecurityHeaderType.SetSecurityHeaderType(table.inSecurityHeader)
		a.SpareHalfOctetAndSecurityHeaderType.SetSpareHalfOctet(table.inSpareHalfOctet)
		a.ServiceRequestMessageIdentity.SetMessageType(table.inServiceRequestMessageIdentity)

		a.TMSI5GS = table.inTMSI5GS

		a.UplinkDataStatus = nasType.NewUplinkDataStatus(nasMessage.ServiceRequestUplinkDataStatusType)
		a.UplinkDataStatus = &table.inUplinkDataStatus

		a.PDUSessionStatus = nasType.NewPDUSessionStatus(nasMessage.ServiceRequestPDUSessionStatusType)
		a.PDUSessionStatus = &table.inPDUSessionStatus

		a.AllowedPDUSessionStatus = nasType.NewAllowedPDUSessionStatus(nasMessage.ServiceRequestAllowedPDUSessionStatusType)
		a.AllowedPDUSessionStatus = &table.inAllowedPDUSessionStatus

		a.NASMessageContainer = nasMessage.NewNASMessageContainer(nasMessage.ServiceRequestNASMessageContainerType)
		a.NASMessageContainer = &table.inNASMessageContainer

		buff := new(bytes.Buffer)
		a.EncodeServiceRequest(buff)
		logger.NasMsgLog.Debugln("Encode: ", a)

		data := make([]byte, buff.Len())
		buff.Read(data)
		logger.NasMsgLog.Debugln(data)
		b.DecodeServiceRequest(&data)
		logger.NasMsgLog.Debugln("Decode: ", b)

		// Compare the re-encoded wire form rather than the structs. Decoding
		// populates the enrichment fields (EPD, MessageType, Cause, ...) that
		// the hand-built value `a` never has, so reflect.DeepEqual(a, b) can
		// never hold. Re-encoding b and comparing bytes is the round-trip
		// property these tests actually mean to assert.
		reBuff := new(bytes.Buffer)
		if err := b.EncodeServiceRequest(reBuff); err != nil {
			t.Fatalf("re-encode: %v", err)
		}
		if !bytes.Equal(data, reBuff.Bytes()) {
			t.Errorf("round trip mismatch:\n encoded: %v\n re-encoded: %v", data, reBuff.Bytes())
		}
	}
}
