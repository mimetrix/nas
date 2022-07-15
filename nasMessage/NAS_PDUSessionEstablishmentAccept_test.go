package nasMessage_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	"github.com/free5gc/nas"
	"github.com/free5gc/nas/logger"
	"github.com/free5gc/nas/nasMessage"
	"github.com/free5gc/nas/nasType"
	"github.com/stretchr/testify/assert"
)

type nasMessagePDUSessionEstablishmentAcceptData struct {
	inExtendedProtocolDiscriminator                uint8
	inPDUSessionID                                 uint8
	inPTI                                          uint8
	inPDUSESSIONESTABLISHMENTACCEPTMessageIdentity uint8
	inSelectedSSCModeAndSelectedPDUSessionType     nasType.SelectedSSCModeAndSelectedPDUSessionType
	inAuthorizedQosRules                           nasType.AuthorizedQosRules
	inSessionAMBR                                  nasType.SessionAMBR
	inCause5GSM                                    nasType.Cause5GSM
	inPDUAddress                                   nasType.PDUAddress
	inRQTimerValue                                 nasType.RQTimerValue
	inSNSSAI                                       nasType.SNSSAI
	inAlwaysonPDUSessionIndication                 nasType.AlwaysonPDUSessionIndication
	inMappedEPSBearerContexts                      nasType.MappedEPSBearerContexts
	inEAPMessage                                   nasType.EAPMessage
	inAuthorizedQosFlowDescriptions                nasType.AuthorizedQosFlowDescriptions
	inExtendedProtocolConfigurationOptions         nasType.ExtendedProtocolConfigurationOptions
	inDNN                                          nasType.DNN
}

var nasMessagePDUSessionEstablishmentAcceptTable = []nasMessagePDUSessionEstablishmentAcceptData{
	{
		inExtendedProtocolDiscriminator: nas.MsgTypePDUSessionEstablishmentAccept,
		inPDUSessionID:                  0x01,
		inPTI:                           0x01,
		inPDUSESSIONESTABLISHMENTACCEPTMessageIdentity: 0x01,
		inSelectedSSCModeAndSelectedPDUSessionType: nasType.SelectedSSCModeAndSelectedPDUSessionType{
			Octet: 0x01,
		},
		inAuthorizedQosRules: nasType.AuthorizedQosRules{
			Iei:    0,
			Len:    1,
			Buffer: []uint8{0x01},
		},
		inSessionAMBR: nasType.SessionAMBR{
			Iei:   0,
			Len:   6,
			Octet: [6]uint8{0x01, 0x02, 0x03, 0x04, 0x05, 0x06},
		},
		inCause5GSM: nasType.Cause5GSM{
			Iei:   nasMessage.PDUSessionEstablishmentAcceptCause5GSMType,
			Octet: 0x01,
		},
		inPDUAddress: nasType.PDUAddress{
			Iei:   nasMessage.PDUSessionEstablishmentAcceptPDUAddressType,
			Len:   13,
			Octet: [13]uint8{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0A, 0x0B, 0x0C},
		},
		inRQTimerValue: nasType.RQTimerValue{
			Iei:   nasMessage.PDUSessionEstablishmentAcceptRQTimerValueType,
			Octet: 0x01,
		},
		inSNSSAI: nasType.SNSSAI{
			Iei:   nasMessage.PDUSessionEstablishmentAcceptSNSSAIType,
			Len:   2,
			Octet: [8]uint8{0x01, 0x01},
		},
		inAlwaysonPDUSessionIndication: nasType.AlwaysonPDUSessionIndication{
			Octet: 0x80,
		},
		inMappedEPSBearerContexts: nasType.MappedEPSBearerContexts{
			Iei:    nasMessage.PDUSessionEstablishmentAcceptMappedEPSBearerContextsType,
			Len:    2,
			Buffer: []uint8{0x01, 0x01},
		},
		inEAPMessage: nasType.EAPMessage{
			Iei:    nasMessage.PDUSessionEstablishmentAcceptEAPMessageType,
			Len:    2,
			Buffer: []uint8{0x01, 0x01},
		},
		inAuthorizedQosFlowDescriptions: nasType.AuthorizedQosFlowDescriptions{
			Iei:    nasMessage.PDUSessionEstablishmentAcceptAuthorizedQosFlowDescriptionsType,
			Len:    2,
			Buffer: []uint8{0x01, 0x01},
		},
		inExtendedProtocolConfigurationOptions: nasType.ExtendedProtocolConfigurationOptions{
			Iei:    nasMessage.PDUSessionEstablishmentAcceptExtendedProtocolConfigurationOptionsType,
			Len:    2,
			Buffer: []uint8{0x01, 0x01},
		},
		inDNN: nasType.DNN{
			Iei:    nasMessage.ULNASTransportDNNType,
			Len:    2,
			Buffer: []uint8{0x01, 0x01},
		},
	},
}

func TestNasTypeNewPDUSessionEstablishmentAccept(t *testing.T) {
	a := nasMessage.NewPDUSessionEstablishmentAccept(0)
	assert.NotNil(t, a)
}

func TestNasTypeNewPDUSessionEstablishmentAcceptMessage(t *testing.T) {
	for i, table := range nasMessagePDUSessionEstablishmentAcceptTable {
		t.Logf("Test Cnt:%d", i)
		a := nasMessage.NewPDUSessionEstablishmentAccept(0)
		b := nasMessage.NewPDUSessionEstablishmentAccept(0)
		assert.NotNil(t, a)
		assert.NotNil(t, b)

		a.ExtendedProtocolDiscriminator.SetExtendedProtocolDiscriminator(table.inExtendedProtocolDiscriminator)
		a.PDUSessionID.SetPDUSessionID(table.inPDUSessionID)
		a.PTI.SetPTI(table.inPTI)
		a.PDUSESSIONESTABLISHMENTACCEPTMessageIdentity.SetMessageType(table.inPDUSESSIONESTABLISHMENTACCEPTMessageIdentity)
		a.SelectedSSCModeAndSelectedPDUSessionType = table.inSelectedSSCModeAndSelectedPDUSessionType
		a.AuthorizedQosRules = table.inAuthorizedQosRules
		a.SessionAMBR = table.inSessionAMBR

		a.Cause5GSM = nasType.NewCause5GSM(nasMessage.PDUSessionEstablishmentAcceptCause5GSMType)
		a.Cause5GSM = &table.inCause5GSM

		a.PDUAddress = nasType.NewPDUAddress(nasMessage.PDUSessionEstablishmentAcceptPDUAddressType)
		a.PDUAddress = &table.inPDUAddress

		a.RQTimerValue = nasType.NewRQTimerValue(nasMessage.PDUSessionEstablishmentAcceptSNSSAIType)
		a.RQTimerValue = &table.inRQTimerValue

		a.SNSSAI = nasType.NewSNSSAI(nasMessage.PDUSessionEstablishmentAcceptSNSSAIType)
		a.SNSSAI = &table.inSNSSAI

		a.AlwaysonPDUSessionIndication = nasType.NewAlwaysonPDUSessionIndication(nasMessage.PDUSessionEstablishmentAcceptAlwaysonPDUSessionIndicationType)
		a.AlwaysonPDUSessionIndication = &table.inAlwaysonPDUSessionIndication

		a.MappedEPSBearerContexts = nasType.NewMappedEPSBearerContexts(nasMessage.PDUSessionEstablishmentAcceptMappedEPSBearerContextsType)
		a.MappedEPSBearerContexts = &table.inMappedEPSBearerContexts

		a.EAPMessage = nasType.NewEAPMessage(nasMessage.PDUSessionEstablishmentAcceptEAPMessageType)
		a.EAPMessage = &table.inEAPMessage

		a.AuthorizedQosFlowDescriptions = nasType.NewAuthorizedQosFlowDescriptions(nasMessage.PDUSessionEstablishmentAcceptAuthorizedQosFlowDescriptionsType)
		a.AuthorizedQosFlowDescriptions = &table.inAuthorizedQosFlowDescriptions

		a.ExtendedProtocolConfigurationOptions = nasType.NewExtendedProtocolConfigurationOptions(nasMessage.PDUSessionEstablishmentAcceptExtendedProtocolConfigurationOptionsType)
		a.ExtendedProtocolConfigurationOptions = &table.inExtendedProtocolConfigurationOptions

		a.DNN = nasType.NewDNN(nasMessage.PDUSessionEstablishmentAcceptDNNType)
		a.DNN = &table.inDNN

		buff := new(bytes.Buffer)
		a.EncodePDUSessionEstablishmentAccept(buff)
		a.PDUAddress.Parse()
		a.SNSSAI.Parse()
		a.AuthorizedQosRules.Parse()
		logger.NasMsgLog.Debugln("Encode: ", a)

		data := make([]byte, buff.Len())
		buff.Read(data)
		logger.NasMsgLog.Debugln(data)
		b.DecodePDUSessionEstablishmentAccept(&data)
		logger.NasMsgLog.Debugln("Decode: ", b)
		j, _ := json.Marshal(b)

		fmt.Println(string(j))

		if reflect.DeepEqual(a, b) != true {
			t.Errorf("Not correct")
		}

	}
}

func TestPDUSessionEstablishmentAccept_DecodePDUSessionEstablishmentAccept(t *testing.T) {
	m := []byte{
		0x2e, 0x01, 0x01, 0xc2, 0x11, 0x00, 0x09, 0x01,
		0x00, 0x06, 0x31, 0x31, 0x01, 0x01, 0xff, 0x09,
		0x06, 0x06, 0x00, 0x64, 0x06, 0x00, 0xc8, 0x29,
		0x05, 0x01, 0x0a, 0x01, 0x00, 0x02, 0x22, 0x04,
		0x01, 0x01, 0x02, 0x03, 0x79, 0x00, 0x06, 0x09,
		0x20, 0x41, 0x01, 0x01, 0x09, 0x7b, 0x00, 0x08,
		0x80, 0x00, 0x0d, 0x04, 0x08, 0x08, 0x08, 0x08,
		0x25, 0x09, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72,
		0x6e, 0x65, 0x74,
	}

	type fields struct {
		ExtendedProtocolDiscriminator                nasType.ExtendedProtocolDiscriminator
		PDUSessionID                                 nasType.PDUSessionID
		PTI                                          nasType.PTI
		PDUSESSIONESTABLISHMENTACCEPTMessageIdentity nasType.PDUSESSIONESTABLISHMENTACCEPTMessageIdentity
		SelectedSSCModeAndSelectedPDUSessionType     nasType.SelectedSSCModeAndSelectedPDUSessionType
		AuthorizedQosRules                           nasType.AuthorizedQosRules
		SessionAMBR                                  nasType.SessionAMBR
		Cause5GSM                                    *nasType.Cause5GSM
		PDUAddress                                   *nasType.PDUAddress
		RQTimerValue                                 *nasType.RQTimerValue
		SNSSAI                                       *nasType.SNSSAI
		AlwaysonPDUSessionIndication                 *nasType.AlwaysonPDUSessionIndication
		MappedEPSBearerContexts                      *nasType.MappedEPSBearerContexts
		EAPMessage                                   *nasType.EAPMessage
		AuthorizedQosFlowDescriptions                *nasType.AuthorizedQosFlowDescriptions
		ExtendedProtocolConfigurationOptions         *nasType.ExtendedProtocolConfigurationOptions
		DNN                                          *nasType.DNN
	}
	type args struct {
		byteArray *[]byte
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "pdu_session_establishement_accept_successful",
			args: args{
				byteArray: &m,
			},
			fields: fields{
				AuthorizedQosRules: nasType.AuthorizedQosRules{
					QoSRules: nasType.QoSRules{
						{
							Identifier: 1,
							PacketFilterList: nasType.PacketFilterList{
								{
									Identifier: 1,
									Direction:  3, //bidirectional
									Components: nasType.PacketFilterComponentList{
										&nasType.PacketFilterMatchAll{},
									},
								},
							},
							Precedence: 255,
							QFI:        9,
							DQR:        true,
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &nasMessage.PDUSessionEstablishmentAccept{
				// ExtendedProtocolDiscriminator: tt.fields.ExtendedProtocolDiscriminator,
				// PDUSessionID:                  tt.fields.PDUSessionID,
				// PTI:                           tt.fields.PTI,
				// PDUSESSIONESTABLISHMENTACCEPTMessageIdentity: tt.fields.PDUSESSIONESTABLISHMENTACCEPTMessageIdentity,
				// SelectedSSCModeAndSelectedPDUSessionType:     tt.fields.SelectedSSCModeAndSelectedPDUSessionType,
				AuthorizedQosRules: tt.fields.AuthorizedQosRules,
				// SessionAMBR:                                  tt.fields.SessionAMBR,
				// Cause5GSM:                                    tt.fields.Cause5GSM,
				// PDUAddress:                                   tt.fields.PDUAddress,
				// RQTimerValue:                                 tt.fields.RQTimerValue,
				// SNSSAI:                                       tt.fields.SNSSAI,
				// AlwaysonPDUSessionIndication:                 tt.fields.AlwaysonPDUSessionIndication,
				// MappedEPSBearerContexts:                      tt.fields.MappedEPSBearerContexts,
				// EAPMessage:                                   tt.fields.EAPMessage,
				// AuthorizedQosFlowDescriptions:                tt.fields.AuthorizedQosFlowDescriptions,
				// ExtendedProtocolConfigurationOptions:         tt.fields.ExtendedProtocolConfigurationOptions,
				// DNN:                                          tt.fields.DNN,
			}
			a.DecodePDUSessionEstablishmentAccept(tt.args.byteArray)
			j, _ := json.Marshal(a)
			fmt.Println(string(j))
		})
	}
}
