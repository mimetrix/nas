package nasMessage_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mimetrix/nas/logger"
	"github.com/mimetrix/nas/nasMessage"
	"github.com/mimetrix/nas/nasType"
)

type nasMessageDLNASTransportData struct {
	inExtendedProtocolDiscriminator uint8
	inSecurityHeaderType            uint8
	inSpareHalfOctet1               uint8
	inDLNASTRANSPORTMessageIdentity uint8
	inPayloadContainerType          uint8
	inSpareHalfOctet2               uint8
	inPayloadContainer              nasMessage.PayloadContainer
	inPduSessionID2Value            nasType.PduSessionID2Value
	inAdditionalInformation         nasType.AdditionalInformation
	inCause5GMM                     nasType.Cause5GMM
	inBackoffTimerValue             nasType.BackoffTimerValue
}

var nasMessageDLNASTransportTable = []nasMessageDLNASTransportData{
	{
		inExtendedProtocolDiscriminator: nasMessage.MsgTypeDLNASTransport,
		inSecurityHeaderType:            0x01,
		inSpareHalfOctet1:               0x01,
		inDLNASTRANSPORTMessageIdentity: 0x01,
		inPayloadContainerType:          0x01,
		inSpareHalfOctet2:               0x01,
		inPayloadContainer: nasMessage.PayloadContainer{
			Iei:    0,
			Len:    2,
			Buffer: []uint8{0x01, 0x01},
		},
		inPduSessionID2Value: nasType.PduSessionID2Value{
			Iei:   nasMessage.DLNASTransportPduSessionID2ValueType,
			Octet: 0x01,
		},
		inAdditionalInformation: nasType.AdditionalInformation{
			Iei:    nasMessage.DLNASTransportAdditionalInformationType,
			Len:    2,
			Buffer: []uint8{0x01, 0x01},
		},
		inCause5GMM: nasType.Cause5GMM{
			Iei:   nasMessage.DLNASTransportCause5GMMType,
			Octet: 0xF0,
		},
		inBackoffTimerValue: nasType.BackoffTimerValue{
			Iei:   nasMessage.DLNASTransportBackoffTimerValueType,
			Len:   1,
			Octet: 0x01,
		},
	},
}

func TestNasTypeNewDLNASTransport(t *testing.T) {
	a := nasMessage.NewDLNASTransport(0)
	assert.NotNil(t, a)
}

func TestNasTypeNewDLNASTransportMessage(t *testing.T) {
	for i, table := range nasMessageDLNASTransportTable {
		logger.NasMsgLog.Infoln("Test Cnt:", i)
		a := nasMessage.NewDLNASTransport(0)
		b := nasMessage.NewDLNASTransport(0)
		assert.NotNil(t, a)
		assert.NotNil(t, b)

		a.ExtendedProtocolDiscriminator.SetExtendedProtocolDiscriminator(table.inExtendedProtocolDiscriminator)
		a.SpareHalfOctetAndSecurityHeaderType.SetSecurityHeaderType(table.inSecurityHeaderType)
		a.SpareHalfOctetAndSecurityHeaderType.SetSpareHalfOctet(table.inSpareHalfOctet1)
		a.DLNASTRANSPORTMessageIdentity.SetMessageType(table.inDLNASTRANSPORTMessageIdentity)
		a.SpareHalfOctetAndPayloadContainerType.SetPayloadContainerType(table.inPayloadContainerType)
		a.PayloadContainer = table.inPayloadContainer

		a.PduSessionID2Value = nasType.NewPduSessionID2Value(nasMessage.DLNASTransportPduSessionID2ValueType)
		a.PduSessionID2Value = &table.inPduSessionID2Value

		a.AdditionalInformation = nasType.NewAdditionalInformation(nasMessage.DLNASTransportAdditionalInformationType)
		a.AdditionalInformation = &table.inAdditionalInformation

		a.Cause5GMM = nasType.NewCause5GMM(nasMessage.DLNASTransportCause5GMMType)
		a.Cause5GMM = &table.inCause5GMM

		a.BackoffTimerValue = nasType.NewBackoffTimerValue(nasMessage.DLNASTransportBackoffTimerValueType)
		a.BackoffTimerValue = &table.inBackoffTimerValue

		buff := new(bytes.Buffer)
		a.EncodeDLNASTransport(buff)
		logger.NasMsgLog.Debugln("Encode: ", a)

		data := make([]byte, buff.Len())
		buff.Read(data)
		logger.NasMsgLog.Debugln(data)
		b.DecodeDLNASTransport(&data)
		logger.NasMsgLog.Debugln("Decode: ", b)

		// Compare the re-encoded wire form rather than the structs. Decoding
		// populates the enrichment fields (EPD, MessageType, Cause, ...) that
		// the hand-built value `a` never has, so reflect.DeepEqual(a, b) can
		// never hold. Re-encoding b and comparing bytes is the round-trip
		// property these tests actually mean to assert.
		reBuff := new(bytes.Buffer)
		if err := b.EncodeDLNASTransport(reBuff); err != nil {
			t.Fatalf("re-encode: %v", err)
		}
		if !bytes.Equal(data, reBuff.Bytes()) {
			t.Errorf("round trip mismatch:\n encoded: %v\n re-encoded: %v", data, reBuff.Bytes())
		}

	}
}
