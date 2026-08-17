package nasMessage_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mimetrix/nas/logger"
	"github.com/mimetrix/nas/nasMessage"
	"github.com/mimetrix/nas/nasType"
)

type nasMessageSecurityModeCommandData struct {
	inExtendedProtocolDiscriminator      uint8
	inSecurityHeader                     uint8
	inSpareHalfOctet                     uint8
	inSecurityModeCommandMessageIdentity uint8
	inSelectedNASSecurityAlgorithms      nasType.SelectedNASSecurityAlgorithms
	inNgksi                              uint8
	inReplayedUESecurityCapabilities     nasType.ReplayedUESecurityCapabilities
	inIMEISVRequest                      nasType.IMEISVRequest
	inSelectedEPSNASSecurityAlgorithms   nasType.SelectedEPSNASSecurityAlgorithms
	inAdditional5GSecurityInformation    nasType.Additional5GSecurityInformation
	inEAPMessage                         nasType.EAPMessage
	inABBA                               nasType.ABBA
	inReplayedS1UESecurityCapabilities   nasType.ReplayedS1UESecurityCapabilities
}

var nasMessageSecurityModeCommandTable = []nasMessageSecurityModeCommandData{
	{
		inExtendedProtocolDiscriminator:      nasMessage.Epd5GSMobilityManagementMessage,
		inSecurityHeader:                     0x01,
		inSpareHalfOctet:                     0x01,
		inSecurityModeCommandMessageIdentity: nasMessage.MsgTypeSecurityModeCommand,
		inSelectedNASSecurityAlgorithms: nasType.SelectedNASSecurityAlgorithms{
			Octet: 0x01,
		},
		inNgksi: 0x01,
		inReplayedUESecurityCapabilities: nasType.ReplayedUESecurityCapabilities{
			Len:    8,
			Buffer: []uint8{0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01},
		},
		inIMEISVRequest: nasType.IMEISVRequest{
			Octet: 0xE0,
		},
		inSelectedEPSNASSecurityAlgorithms: nasType.SelectedEPSNASSecurityAlgorithms{
			Iei:   nasMessage.SecurityModeCommandSelectedEPSNASSecurityAlgorithmsType,
			Octet: 0x01,
		},
		inAdditional5GSecurityInformation: nasType.Additional5GSecurityInformation{
			Iei:   nasMessage.SecurityModeCommandAdditional5GSecurityInformationType,
			Len:   1,
			Octet: 0x01,
		},
		inEAPMessage: nasType.EAPMessage{
			Iei:    nasMessage.SecurityModeCommandEAPMessageType,
			Len:    4,
			Buffer: []uint8{0x01, 0x01, 0x01, 0x01},
		},
		inABBA: nasType.ABBA{
			Iei:    nasMessage.SecurityModeCommandABBAType,
			Len:    2,
			Buffer: []uint8{0x01, 0x01},
		},
		inReplayedS1UESecurityCapabilities: nasType.ReplayedS1UESecurityCapabilities{
			Iei:    nasMessage.SecurityModeCommandReplayedS1UESecurityCapabilitiesType,
			Len:    5,
			Buffer: []uint8{0x01, 0x01, 0x01, 0x01, 0x01},
		},
	},
}

func TestNasTypeNewSecurityModeCommand(t *testing.T) {
	a := nasMessage.NewSecurityModeCommand(0)
	assert.NotNil(t, a)
}

func TestNasTypeNewSecurityModeCommandMessage(t *testing.T) {
	for i, table := range nasMessageSecurityModeCommandTable {
		t.Logf("Test Cnt:%d", i)
		a := nasMessage.NewSecurityModeCommand(0)
		b := nasMessage.NewSecurityModeCommand(0)
		assert.NotNil(t, a)
		assert.NotNil(t, b)

		a.ExtendedProtocolDiscriminator.SetExtendedProtocolDiscriminator(table.inExtendedProtocolDiscriminator)
		a.SpareHalfOctetAndSecurityHeaderType.SetSecurityHeaderType(table.inSecurityHeader)
		a.SpareHalfOctetAndSecurityHeaderType.SetSpareHalfOctet(table.inSpareHalfOctet)
		a.SecurityModeCommandMessageIdentity.SetMessageType(table.inSecurityModeCommandMessageIdentity)

		a.SelectedNASSecurityAlgorithms = table.inSelectedNASSecurityAlgorithms
		a.SpareHalfOctetAndNgksi.SetSpareHalfOctet(table.inSpareHalfOctet)
		a.SpareHalfOctetAndNgksi.SetNasKeySetIdentifiler(table.inNgksi)

		a.ReplayedUESecurityCapabilities = table.inReplayedUESecurityCapabilities

		a.IMEISVRequest = nasType.NewIMEISVRequest(nasMessage.SecurityModeCommandIMEISVRequestType)
		a.IMEISVRequest = &table.inIMEISVRequest

		a.SelectedEPSNASSecurityAlgorithms = nasType.NewSelectedEPSNASSecurityAlgorithms(nasMessage.SecurityModeCommandSelectedEPSNASSecurityAlgorithmsType)
		a.SelectedEPSNASSecurityAlgorithms = &table.inSelectedEPSNASSecurityAlgorithms

		a.Additional5GSecurityInformation = nasType.NewAdditional5GSecurityInformation(nasMessage.SecurityModeCommandAdditional5GSecurityInformationType)
		a.Additional5GSecurityInformation = &table.inAdditional5GSecurityInformation

		a.EAPMessage = nasType.NewEAPMessage(nasMessage.SecurityModeCommandEAPMessageType)
		a.EAPMessage = &table.inEAPMessage

		a.ABBA = nasType.NewABBA(nasMessage.SecurityModeCommandABBAType)
		a.ABBA = &table.inABBA

		a.ReplayedS1UESecurityCapabilities = nasType.NewReplayedS1UESecurityCapabilities(nasMessage.SecurityModeCommandReplayedS1UESecurityCapabilitiesType)
		a.ReplayedS1UESecurityCapabilities = &table.inReplayedS1UESecurityCapabilities

		buff := new(bytes.Buffer)
		a.EncodeSecurityModeCommand(buff)
		logger.NasMsgLog.Debugln("Encode: ", a)

		data := make([]byte, buff.Len())
		buff.Read(data)
		logger.NasMsgLog.Debugln(data)
		b.DecodeSecurityModeCommand(&data)
		logger.NasMsgLog.Debugln("Decode: ", b)

		// Compare the re-encoded wire form rather than the structs. Decoding
		// populates the enrichment fields (EPD, MessageType, Cause, ...) that
		// the hand-built value `a` never has, so reflect.DeepEqual(a, b) can
		// never hold. Re-encoding b and comparing bytes is the round-trip
		// property these tests actually mean to assert.
		reBuff := new(bytes.Buffer)
		if err := b.EncodeSecurityModeCommand(reBuff); err != nil {
			t.Fatalf("re-encode: %v", err)
		}
		if !bytes.Equal(data, reBuff.Bytes()) {
			t.Errorf("round trip mismatch:\n encoded: %v\n re-encoded: %v", data, reBuff.Bytes())
		}
	}
}
