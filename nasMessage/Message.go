package nasMessage

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

// Message TODO：description
type Message struct {
	SecurityHeader `json:"-"`
	*GmmMessage    `json:"GmmMessage,omitempty"`
	*GsmMessage    `json:"GsmMessage,omitempty"`
	Bytes          []byte `json:"EncryptedBytes,omitempty"`
}

// SecurityHeader TODO：description
type SecurityHeader struct {
	ProtocolDiscriminator     uint8  `json:"ProtocolDiscriminator"`
	SecurityHeaderType        uint8  `json:"SecurityHeaderType"`
	MessageAuthenticationCode uint32 `json:"MessageAuthenticationCode"`
	SequenceNumber            uint8  `json:"SequenceNumber"`
}

const (
	SecurityHeaderTypePlainNas                                                 uint8 = 0x00
	SecurityHeaderTypeIntegrityProtected                                       uint8 = 0x01
	SecurityHeaderTypeIntegrityProtectedAndCiphered                            uint8 = 0x02
	SecurityHeaderTypeIntegrityProtectedWithNew5gNasSecurityContext            uint8 = 0x03
	SecurityHeaderTypeIntegrityProtectedAndCipheredWithNew5gNasSecurityContext uint8 = 0x04
)

// NewMessage TODO:desc
func NewMessage() *Message {
	Message := &Message{}
	return Message
}

// NewGmmMessage TODO:desc
func NewGmmMessage() *GmmMessage {
	GmmMessage := &GmmMessage{}
	return GmmMessage
}

// NewGmmMessage TODO:desc
func NewGsmMessage() *GsmMessage {
	GsmMessage := &GsmMessage{}
	return GsmMessage
}

// GmmHeader Octet1 protocolDiscriminator securityHeaderType
//
// Octet2 MessageType
type GmmHeader struct {
	Octet [3]uint8 `json:"Octet,omitempty"`
}

type GsmHeader struct {
	Octet [4]uint8 `json:"Octet,omitempty"`
}

// GetMessageType 9.8
func (a *GmmHeader) GetMessageType() (messageType uint8) {
	messageType = a.Octet[2]
	return messageType
}

// GetMessageType 9.8
func (a *GmmHeader) SetMessageType(messageType uint8) {
	a.Octet[2] = messageType
}

func (a *GmmHeader) GetExtendedProtocolDiscriminator() uint8 {
	return a.Octet[0]
}

func (a *GmmHeader) SetExtendedProtocolDiscriminator(epd uint8) {
	a.Octet[0] = epd
}

func (a *GsmHeader) GetExtendedProtocolDiscriminator() uint8 {
	return a.Octet[0]
}

func (a *GsmHeader) SetExtendedProtocolDiscriminator(epd uint8) {
	a.Octet[0] = epd
}

// GetMessageType 9.8
func (a *GsmHeader) GetMessageType() (messageType uint8) {
	messageType = a.Octet[3]
	return messageType
}

// GetMessageType 9.8
func (a *GsmHeader) SetMessageType(messageType uint8) {
	a.Octet[3] = messageType
}

func GetEPD(byteArray []byte) uint8 {
	return byteArray[0]
}

func GetSecurityHeaderType(byteArray []byte) uint8 {
	return byteArray[1]
}

type GmmMessage struct {
	GmmHeader                                         `json:"-"`
	*AuthenticationRequest                            `json:"AuthenticationRequest,omitempty"`
	*AuthenticationResponse                           `json:"AuthenticationResponse,omitempty"`
	*AuthenticationResult                             `json:"AuthenticationResult,omitempty"`
	*AuthenticationFailure                            `json:"AuthenticationFailure,omitempty"`
	*AuthenticationReject                             `json:"AuthenticationReject,omitempty"`
	*RegistrationRequest                              `json:"RegistrationRequest,omitempty"`
	*RegistrationAccept                               `json:"RegistrationAccept,omitempty"`
	*RegistrationComplete                             `json:"RegistrationComplete,omitempty"`
	*RegistrationReject                               `json:"RegistrationReject,omitempty"`
	*ULNASTransport                                   `json:"ULNASTransport,omitempty"`
	*DLNASTransport                                   `json:"DLNASTransport,omitempty"`
	*DeregistrationRequestUEOriginatingDeregistration `json:"DeregistrationRequestUEOriginatingDeregistration,omitempty"`
	*DeregistrationAcceptUEOriginatingDeregistration  `json:"DeregistrationAcceptUEOriginatingDeregistration,omitempty"`
	*DeregistrationRequestUETerminatedDeregistration  `json:"DeregistrationRequestUETerminatedDeregistration,omitempty"`
	*DeregistrationAcceptUETerminatedDeregistration   `json:"DeregistrationAcceptUETerminatedDeregistration,omitempty"`
	*ServiceRequest                                   `json:"ServiceRequest,omitempty"`
	*ServiceAccept                                    `json:"ServiceAccept,omitempty"`
	*ServiceReject                                    `json:"ServiceReject,omitempty"`
	*ConfigurationUpdateCommand                       `json:"ConfigurationUpdateCommand,omitempty"`
	*ConfigurationUpdateComplete                      `json:"ConfigurationUpdateComplete,omitempty"`
	*IdentityRequest                                  `json:"IdentityRequest,omitempty"`
	*IdentityResponse                                 `json:"IdentityResponse,omitempty"`
	*Notification                                     `json:"Notification,omitempty"`
	*NotificationResponse                             `json:"NotificationResponse,omitempty"`
	*SecurityModeCommand                              `json:"SecurityModeCommand,omitempty"`
	*SecurityModeComplete                             `json:"SecurityModeComplete,omitempty"`
	*SecurityModeReject                               `json:"SecurityModeReject,omitempty"`
	*SecurityProtected5GSNASMessage                   `json:"SecurityProtected5GSNASMessage,omitempty"`
	*Status5GMM                                       `json:"Status5GMM,omitempty"`
}

const (
	MsgTypeRegistrationRequest                              uint8 = 65  //0x41
	MsgTypeRegistrationAccept                               uint8 = 66  //0x42
	MsgTypeRegistrationComplete                             uint8 = 67  //0x43
	MsgTypeRegistrationReject                               uint8 = 68  //0x44
	MsgTypeDeregistrationRequestUEOriginatingDeregistration uint8 = 69  //0x45
	MsgTypeDeregistrationAcceptUEOriginatingDeregistration  uint8 = 70  //0x46
	MsgTypeDeregistrationRequestUETerminatedDeregistration  uint8 = 71  //0x47
	MsgTypeDeregistrationAcceptUETerminatedDeregistration   uint8 = 72  //0x48
	MsgTypeServiceRequest                                   uint8 = 76  //0x4c
	MsgTypeServiceReject                                    uint8 = 77  //0x4d
	MsgTypeServiceAccept                                    uint8 = 78  //0x4e
	MsgTypeConfigurationUpdateCommand                       uint8 = 84  //0x54
	MsgTypeConfigurationUpdateComplete                      uint8 = 85  //0x55
	MsgTypeAuthenticationRequest                            uint8 = 86  //0x56
	MsgTypeAuthenticationResponse                           uint8 = 87  //0x57
	MsgTypeAuthenticationReject                             uint8 = 88  //0x58
	MsgTypeAuthenticationFailure                            uint8 = 89  //0x59
	MsgTypeAuthenticationResult                             uint8 = 90  //0x5a
	MsgTypeIdentityRequest                                  uint8 = 91  //0x5b
	MsgTypeIdentityResponse                                 uint8 = 92  //0x5c
	MsgTypeSecurityModeCommand                              uint8 = 93  //0x5d
	MsgTypeSecurityModeComplete                             uint8 = 94  //0x5e
	MsgTypeSecurityModeReject                               uint8 = 95  //0x5f
	MsgTypeStatus5GMM                                       uint8 = 100 //0x64
	MsgTypeNotification                                     uint8 = 101 //0x65
	MsgTypeNotificationResponse                             uint8 = 102 //0x66
	MsgTypeULNASTransport                                   uint8 = 103 //0x67
	MsgTypeDLNASTransport                                   uint8 = 104 //0x67
	// 0x5d
)

// placeholder
const MsgTypeSecurityProtected5GSNASMessage = 0x00

func (a *Message) SecurityProtectedNasDecode(byteArray *[]byte) error {
	buffer := bytes.NewBuffer(*byteArray)
	a.GmmMessage = NewGmmMessage()
	if err := binary.Read(buffer, binary.BigEndian, &a.GmmMessage.GmmHeader); err != nil {
		return fmt.Errorf("GMM NAS decode Fail: read fail - %+v", err)
	}
	a.GmmMessage.SecurityProtected5GSNASMessage = NewSecurityProtected5GSNASMessage(MsgTypeSecurityProtected5GSNASMessage)

	if err := a.GmmMessage.DecodeSecurityProtected5GSNASMessage(byteArray); err != nil {
		return err
	}

	return nil
}

func (a *Message) PlainNasDecode(byteArray *[]byte) error {
	epd := GetEPD(*byteArray)

	switch epd {
	case Epd5GSMobilityManagementMessage:
		return a.GmmMessageDecode(byteArray)
	case Epd5GSSessionManagementMessage:
		return a.GsmMessageDecode(byteArray)
	default:
		a.Bytes = *byteArray
	}
	return fmt.Errorf("Extended Protocol Discriminator[%d] is not allowed in Nas Message Deocde", epd)
}

func (a *Message) PlainNasEncode() ([]byte, error) {
	data := new(bytes.Buffer)
	if a.GmmMessage != nil {
		err := a.GmmMessageEncode(data)
		return data.Bytes(), err
	} else if a.GsmMessage != nil {
		err := a.GsmMessageEncode(data)
		return data.Bytes(), err
	}
	return nil, fmt.Errorf("Gmm/Gsm Message are both empty in Nas Message Encode")
}

func (a *Message) GmmMessageDecode(byteArray *[]byte) error {
	buffer := bytes.NewBuffer(*byteArray)
	a.GmmMessage = NewGmmMessage()
	if err := binary.Read(buffer, binary.BigEndian, &a.GmmMessage.GmmHeader); err != nil {
		return fmt.Errorf("GMM NAS decode Fail: read fail - %+v", err)
	}

	switch a.GmmMessage.GmmHeader.GetMessageType() {
	case MsgTypeRegistrationRequest:
		a.GmmMessage.RegistrationRequest = NewRegistrationRequest(MsgTypeRegistrationRequest)
		if err := a.GmmMessage.DecodeRegistrationRequest(byteArray); err != nil {
			return err
		}

	case MsgTypeRegistrationAccept:
		a.GmmMessage.RegistrationAccept = NewRegistrationAccept(MsgTypeRegistrationAccept)
		if err := a.GmmMessage.DecodeRegistrationAccept(byteArray); err != nil {
			return err
		}
	case MsgTypeRegistrationComplete:
		a.GmmMessage.RegistrationComplete = NewRegistrationComplete(MsgTypeRegistrationComplete)
		if err := a.GmmMessage.DecodeRegistrationComplete(byteArray); err != nil {
			return err
		}
	case MsgTypeRegistrationReject:
		a.GmmMessage.RegistrationReject = NewRegistrationReject(MsgTypeRegistrationReject)
		if err := a.GmmMessage.DecodeRegistrationReject(byteArray); err != nil {
			return err
		}
	case MsgTypeDeregistrationRequestUEOriginatingDeregistration:
		a.GmmMessage.DeregistrationRequestUEOriginatingDeregistration = NewDeregistrationRequestUEOriginatingDeregistration(
			MsgTypeDeregistrationRequestUEOriginatingDeregistration)
		if err := a.GmmMessage.DecodeDeregistrationRequestUEOriginatingDeregistration(byteArray); err != nil {
			return err
		}
	case MsgTypeDeregistrationAcceptUEOriginatingDeregistration:
		a.GmmMessage.DeregistrationAcceptUEOriginatingDeregistration = NewDeregistrationAcceptUEOriginatingDeregistration(
			MsgTypeDeregistrationAcceptUEOriginatingDeregistration)
		if err := a.GmmMessage.DecodeDeregistrationAcceptUEOriginatingDeregistration(byteArray); err != nil {
			return err
		}
	case MsgTypeDeregistrationRequestUETerminatedDeregistration:
		a.GmmMessage.DeregistrationRequestUETerminatedDeregistration = NewDeregistrationRequestUETerminatedDeregistration(
			MsgTypeDeregistrationRequestUETerminatedDeregistration)
		if err := a.GmmMessage.DecodeDeregistrationRequestUETerminatedDeregistration(byteArray); err != nil {
			return err
		}
	case MsgTypeDeregistrationAcceptUETerminatedDeregistration:
		a.GmmMessage.DeregistrationAcceptUETerminatedDeregistration = NewDeregistrationAcceptUETerminatedDeregistration(
			MsgTypeDeregistrationAcceptUETerminatedDeregistration)
		if err := a.GmmMessage.DecodeDeregistrationAcceptUETerminatedDeregistration(byteArray); err != nil {
			return err
		}
	case MsgTypeServiceRequest:
		a.GmmMessage.ServiceRequest = NewServiceRequest(MsgTypeServiceRequest)
		if err := a.GmmMessage.DecodeServiceRequest(byteArray); err != nil {
			return err
		}
	case MsgTypeServiceReject:
		a.GmmMessage.ServiceReject = NewServiceReject(MsgTypeServiceReject)
		if err := a.GmmMessage.DecodeServiceReject(byteArray); err != nil {
			return err
		}
	case MsgTypeServiceAccept:
		a.GmmMessage.ServiceAccept = NewServiceAccept(MsgTypeServiceAccept)
		if err := a.GmmMessage.DecodeServiceAccept(byteArray); err != nil {
			return err
		}
	case MsgTypeConfigurationUpdateCommand:
		a.GmmMessage.ConfigurationUpdateCommand = NewConfigurationUpdateCommand(MsgTypeConfigurationUpdateCommand)
		if err := a.GmmMessage.DecodeConfigurationUpdateCommand(byteArray); err != nil {
			return err
		}
	case MsgTypeConfigurationUpdateComplete:
		a.GmmMessage.ConfigurationUpdateComplete = NewConfigurationUpdateComplete(MsgTypeConfigurationUpdateComplete)
		if err := a.GmmMessage.DecodeConfigurationUpdateComplete(byteArray); err != nil {
			return err
		}
	case MsgTypeAuthenticationRequest:
		a.GmmMessage.AuthenticationRequest = NewAuthenticationRequest(MsgTypeAuthenticationRequest)
		if err := a.GmmMessage.DecodeAuthenticationRequest(byteArray); err != nil {
			return err
		}
	case MsgTypeAuthenticationResponse:
		a.GmmMessage.AuthenticationResponse = NewAuthenticationResponse(MsgTypeAuthenticationResponse)
		if err := a.GmmMessage.DecodeAuthenticationResponse(byteArray); err != nil {
			return err
		}
	case MsgTypeAuthenticationReject:
		a.GmmMessage.AuthenticationReject = NewAuthenticationReject(MsgTypeAuthenticationReject)
		if err := a.GmmMessage.DecodeAuthenticationReject(byteArray); err != nil {
			return err
		}
	case MsgTypeAuthenticationFailure:
		a.GmmMessage.AuthenticationFailure = NewAuthenticationFailure(MsgTypeAuthenticationFailure)
		if err := a.GmmMessage.DecodeAuthenticationFailure(byteArray); err != nil {
			return err
		}
	case MsgTypeAuthenticationResult:
		a.GmmMessage.AuthenticationResult = NewAuthenticationResult(MsgTypeAuthenticationResult)
		if err := a.GmmMessage.DecodeAuthenticationResult(byteArray); err != nil {
			return err
		}
	case MsgTypeIdentityRequest:
		a.GmmMessage.IdentityRequest = NewIdentityRequest(MsgTypeIdentityRequest)
		if err := a.GmmMessage.DecodeIdentityRequest(byteArray); err != nil {
			return err
		}
	case MsgTypeIdentityResponse:
		a.GmmMessage.IdentityResponse = NewIdentityResponse(MsgTypeIdentityResponse)
		if err := a.GmmMessage.DecodeIdentityResponse(byteArray); err != nil {
			return err
		}
	case MsgTypeSecurityModeCommand:
		a.GmmMessage.SecurityModeCommand = NewSecurityModeCommand(MsgTypeSecurityModeCommand)
		if err := a.GmmMessage.DecodeSecurityModeCommand(byteArray); err != nil {
			return err
		}
	case MsgTypeSecurityModeComplete:
		a.GmmMessage.SecurityModeComplete = NewSecurityModeComplete(MsgTypeSecurityModeComplete)
		if err := a.GmmMessage.DecodeSecurityModeComplete(byteArray); err != nil {
			return err
		}
	case MsgTypeSecurityModeReject:
		a.GmmMessage.SecurityModeReject = NewSecurityModeReject(MsgTypeSecurityModeReject)
		if err := a.GmmMessage.DecodeSecurityModeReject(byteArray); err != nil {
			return err
		}
	case MsgTypeStatus5GMM:
		a.GmmMessage.Status5GMM = NewStatus5GMM(MsgTypeStatus5GMM)
		if err := a.GmmMessage.DecodeStatus5GMM(byteArray); err != nil {
			return err
		}
	case MsgTypeNotification:
		a.GmmMessage.Notification = NewNotification(MsgTypeNotification)
		if err := a.GmmMessage.DecodeNotification(byteArray); err != nil {
			return err
		}
	case MsgTypeNotificationResponse:
		a.GmmMessage.NotificationResponse = NewNotificationResponse(MsgTypeNotificationResponse)
		if err := a.GmmMessage.DecodeNotificationResponse(byteArray); err != nil {
			return err
		}
	case MsgTypeULNASTransport:
		a.GmmMessage.ULNASTransport = NewULNASTransport(MsgTypeULNASTransport)
		if err := a.GmmMessage.DecodeULNASTransport(byteArray); err != nil {
			return err
		}
	case MsgTypeDLNASTransport:
		a.GmmMessage.DLNASTransport = NewDLNASTransport(MsgTypeDLNASTransport)
		if err := a.GmmMessage.DecodeDLNASTransport(byteArray); err != nil {
			return err
		}
	case MsgTypeSecurityProtected5GSNASMessage:
		a.GmmMessage.SecurityProtected5GSNASMessage = NewSecurityProtected5GSNASMessage(MsgTypeSecurityProtected5GSNASMessage)
		if err := a.GmmMessage.DecodeSecurityProtected5GSNASMessage(byteArray); err != nil {
			return err
		}
	default:
		return fmt.Errorf("NAS decode Fail: MsgType[%x] doesn't exist in GMM Message",
			a.GmmMessage.GmmHeader.GetMessageType())
	}
	return nil
}

func (a *Message) GmmMessageEncode(buffer *bytes.Buffer) error {
	switch a.GmmMessage.GmmHeader.GetMessageType() {
	case MsgTypeRegistrationRequest:
		if err := a.GmmMessage.EncodeRegistrationRequest(buffer); err != nil {
			return err
		}
	case MsgTypeRegistrationAccept:
		if err := a.GmmMessage.EncodeRegistrationAccept(buffer); err != nil {
			return err
		}
	case MsgTypeRegistrationComplete:
		if err := a.GmmMessage.EncodeRegistrationComplete(buffer); err != nil {
			return err
		}
	case MsgTypeRegistrationReject:
		if err := a.GmmMessage.EncodeRegistrationReject(buffer); err != nil {
			return err
		}
	case MsgTypeDeregistrationRequestUEOriginatingDeregistration:
		if err := a.GmmMessage.EncodeDeregistrationRequestUEOriginatingDeregistration(buffer); err != nil {
			return err
		}
	case MsgTypeDeregistrationAcceptUEOriginatingDeregistration:
		if err := a.GmmMessage.EncodeDeregistrationAcceptUEOriginatingDeregistration(buffer); err != nil {
			return err
		}
	case MsgTypeDeregistrationRequestUETerminatedDeregistration:
		if err := a.GmmMessage.EncodeDeregistrationRequestUETerminatedDeregistration(buffer); err != nil {
			return err
		}
	case MsgTypeDeregistrationAcceptUETerminatedDeregistration:
		if err := a.GmmMessage.EncodeDeregistrationAcceptUETerminatedDeregistration(buffer); err != nil {
			return err
		}
	case MsgTypeServiceRequest:
		if err := a.GmmMessage.EncodeServiceRequest(buffer); err != nil {
			return err
		}
	case MsgTypeServiceReject:
		if err := a.GmmMessage.EncodeServiceReject(buffer); err != nil {
			return err
		}
	case MsgTypeServiceAccept:
		if err := a.GmmMessage.EncodeServiceAccept(buffer); err != nil {
			return err
		}
	case MsgTypeConfigurationUpdateCommand:
		if err := a.GmmMessage.EncodeConfigurationUpdateCommand(buffer); err != nil {
			return err
		}
	case MsgTypeConfigurationUpdateComplete:
		if err := a.GmmMessage.EncodeConfigurationUpdateComplete(buffer); err != nil {
			return err
		}
	case MsgTypeAuthenticationRequest:
		if err := a.GmmMessage.EncodeAuthenticationRequest(buffer); err != nil {
			return err
		}
	case MsgTypeAuthenticationResponse:
		if err := a.GmmMessage.EncodeAuthenticationResponse(buffer); err != nil {
			return err
		}
	case MsgTypeAuthenticationReject:
		if err := a.GmmMessage.EncodeAuthenticationReject(buffer); err != nil {
			return err
		}
	case MsgTypeAuthenticationFailure:
		if err := a.GmmMessage.EncodeAuthenticationFailure(buffer); err != nil {
			return err
		}
	case MsgTypeAuthenticationResult:
		if err := a.GmmMessage.EncodeAuthenticationResult(buffer); err != nil {
			return err
		}
	case MsgTypeIdentityRequest:
		if err := a.GmmMessage.EncodeIdentityRequest(buffer); err != nil {
			return err
		}
	case MsgTypeIdentityResponse:
		if err := a.GmmMessage.EncodeIdentityResponse(buffer); err != nil {
			return err
		}
	case MsgTypeSecurityModeCommand:
		if err := a.GmmMessage.EncodeSecurityModeCommand(buffer); err != nil {
			return err
		}
	case MsgTypeSecurityModeComplete:
		if err := a.GmmMessage.EncodeSecurityModeComplete(buffer); err != nil {
			return err
		}
	case MsgTypeSecurityModeReject:
		if err := a.GmmMessage.EncodeSecurityModeReject(buffer); err != nil {
			return err
		}
	case MsgTypeStatus5GMM:
		if err := a.GmmMessage.EncodeStatus5GMM(buffer); err != nil {
			return err
		}
	case MsgTypeNotification:
		if err := a.GmmMessage.EncodeNotification(buffer); err != nil {
			return err
		}
	case MsgTypeNotificationResponse:
		if err := a.GmmMessage.EncodeNotificationResponse(buffer); err != nil {
			return err
		}
	case MsgTypeULNASTransport:
		if err := a.GmmMessage.EncodeULNASTransport(buffer); err != nil {
			return err
		}
	case MsgTypeDLNASTransport:
		if err := a.GmmMessage.EncodeDLNASTransport(buffer); err != nil {
			return err
		}
	default:
		return fmt.Errorf("NAS Encode Fail: MsgType[%d] doesn't exist in GMM Message",
			a.GmmMessage.GmmHeader.GetMessageType())
	}
	return nil
}

type GsmMessage struct {
	GsmHeader                            `json:"-"`
	*PDUSessionEstablishmentRequest      `json:"PDUSessionEstablishmentRequest,omitempty"`
	*PDUSessionEstablishmentAccept       `json:"PDUSessionEstablishmentAccept,omitempty"`
	*PDUSessionEstablishmentReject       `json:"PDUSessionEstablishmentReject,omitempty"`
	*PDUSessionAuthenticationCommand     `json:"PDUSessionAuthenticationCommand,omitempty"`
	*PDUSessionAuthenticationComplete    `json:"PDUSessionAuthenticationComplete,omitempty"`
	*PDUSessionAuthenticationResult      `json:"PDUSessionAuthenticationResult,omitempty"`
	*PDUSessionModificationRequest       `json:"PDUSessionModificationRequest,omitempty"`
	*PDUSessionModificationReject        `json:"PDUSessionModificationReject,omitempty"`
	*PDUSessionModificationCommand       `json:"PDUSessionModificationCommand,omitempty"`
	*PDUSessionModificationComplete      `json:"PDUSessionModificationComplete,omitempty"`
	*PDUSessionModificationCommandReject `json:"PDUSessionModificationCommandReject,omitempty"`
	*PDUSessionReleaseRequest            `json:"PDUSessionReleaseRequest,omitempty"`
	*PDUSessionReleaseReject             `json:"PDUSessionReleaseReject,omitempty"`
	*PDUSessionReleaseCommand            `json:"PDUSessionReleaseCommand,omitempty"`
	*PDUSessionReleaseComplete           `json:"PDUSessionReleaseComplete,omitempty"`
	*Status5GSM                          `json:"Status5GSM,omitempty"`
}

const (
	MsgTypePDUSessionEstablishmentRequest      uint8 = 193
	MsgTypePDUSessionEstablishmentAccept       uint8 = 194
	MsgTypePDUSessionEstablishmentReject       uint8 = 195
	MsgTypePDUSessionAuthenticationCommand     uint8 = 197
	MsgTypePDUSessionAuthenticationComplete    uint8 = 198
	MsgTypePDUSessionAuthenticationResult      uint8 = 199
	MsgTypePDUSessionModificationRequest       uint8 = 201
	MsgTypePDUSessionModificationReject        uint8 = 202
	MsgTypePDUSessionModificationCommand       uint8 = 203
	MsgTypePDUSessionModificationComplete      uint8 = 204
	MsgTypePDUSessionModificationCommandReject uint8 = 205
	MsgTypePDUSessionReleaseRequest            uint8 = 209
	MsgTypePDUSessionReleaseReject             uint8 = 210
	MsgTypePDUSessionReleaseCommand            uint8 = 211
	MsgTypePDUSessionReleaseComplete           uint8 = 212
	MsgTypeStatus5GSM                          uint8 = 214
)

func (a *Message) GsmMessageDecode(byteArray *[]byte) error {
	buffer := bytes.NewBuffer(*byteArray)
	a.GsmMessage = NewGsmMessage()
	if err := binary.Read(buffer, binary.BigEndian, &a.GsmMessage.GsmHeader); err != nil {
		return fmt.Errorf("GSM NAS decode Fail: read fail - %+v", err)
	}

	switch a.GsmMessage.GsmHeader.GetMessageType() {
	case MsgTypePDUSessionEstablishmentRequest:
		a.GsmMessage.PDUSessionEstablishmentRequest = NewPDUSessionEstablishmentRequest(MsgTypePDUSessionEstablishmentRequest)
		if err := a.GsmMessage.DecodePDUSessionEstablishmentRequest(byteArray); err != nil {
			return err
		}
	case MsgTypePDUSessionEstablishmentAccept:
		a.GsmMessage.PDUSessionEstablishmentAccept = NewPDUSessionEstablishmentAccept(MsgTypePDUSessionEstablishmentAccept)
		if err := a.GsmMessage.DecodePDUSessionEstablishmentAccept(byteArray); err != nil {
			return err
		}
	case MsgTypePDUSessionEstablishmentReject:
		a.GsmMessage.PDUSessionEstablishmentReject = NewPDUSessionEstablishmentReject(MsgTypePDUSessionEstablishmentReject)
		if err := a.GsmMessage.DecodePDUSessionEstablishmentReject(byteArray); err != nil {
			return err
		}
	case MsgTypePDUSessionAuthenticationCommand:
		a.GsmMessage.PDUSessionAuthenticationCommand = NewPDUSessionAuthenticationCommand(MsgTypePDUSessionAuthenticationCommand)
		if err := a.GsmMessage.DecodePDUSessionAuthenticationCommand(byteArray); err != nil {
			return err
		}
	case MsgTypePDUSessionAuthenticationComplete:
		a.GsmMessage.PDUSessionAuthenticationComplete = NewPDUSessionAuthenticationComplete(MsgTypePDUSessionAuthenticationComplete)
		if err := a.GsmMessage.DecodePDUSessionAuthenticationComplete(byteArray); err != nil {
			return err
		}
	case MsgTypePDUSessionAuthenticationResult:
		a.GsmMessage.PDUSessionAuthenticationResult = NewPDUSessionAuthenticationResult(MsgTypePDUSessionAuthenticationResult)
		if err := a.GsmMessage.DecodePDUSessionAuthenticationResult(byteArray); err != nil {
			return err
		}
	case MsgTypePDUSessionModificationRequest:
		a.GsmMessage.PDUSessionModificationRequest = NewPDUSessionModificationRequest(MsgTypePDUSessionModificationRequest)
		if err := a.GsmMessage.DecodePDUSessionModificationRequest(byteArray); err != nil {
			return err
		}
	case MsgTypePDUSessionModificationReject:
		a.GsmMessage.PDUSessionModificationReject = NewPDUSessionModificationReject(MsgTypePDUSessionModificationReject)
		if err := a.GsmMessage.DecodePDUSessionModificationReject(byteArray); err != nil {
			return err
		}
	case MsgTypePDUSessionModificationCommand:
		a.GsmMessage.PDUSessionModificationCommand = NewPDUSessionModificationCommand(MsgTypePDUSessionModificationCommand)
		if err := a.GsmMessage.DecodePDUSessionModificationCommand(byteArray); err != nil {
			return err
		}
	case MsgTypePDUSessionModificationComplete:
		a.GsmMessage.PDUSessionModificationComplete = NewPDUSessionModificationComplete(MsgTypePDUSessionModificationComplete)
		if err := a.GsmMessage.DecodePDUSessionModificationComplete(byteArray); err != nil {
			return err
		}
	case MsgTypePDUSessionModificationCommandReject:
		a.GsmMessage.PDUSessionModificationCommandReject = NewPDUSessionModificationCommandReject(MsgTypePDUSessionModificationCommandReject)
		if err := a.GsmMessage.DecodePDUSessionModificationCommandReject(byteArray); err != nil {
			return err
		}
	case MsgTypePDUSessionReleaseRequest:
		a.GsmMessage.PDUSessionReleaseRequest = NewPDUSessionReleaseRequest(MsgTypePDUSessionReleaseRequest)
		if err := a.GsmMessage.DecodePDUSessionReleaseRequest(byteArray); err != nil {
			return err
		}
	case MsgTypePDUSessionReleaseReject:
		a.GsmMessage.PDUSessionReleaseReject = NewPDUSessionReleaseReject(MsgTypePDUSessionReleaseReject)
		if err := a.GsmMessage.DecodePDUSessionReleaseReject(byteArray); err != nil {
			return err
		}
	case MsgTypePDUSessionReleaseCommand:
		a.GsmMessage.PDUSessionReleaseCommand = NewPDUSessionReleaseCommand(MsgTypePDUSessionReleaseCommand)
		if err := a.GsmMessage.DecodePDUSessionReleaseCommand(byteArray); err != nil {
			return err
		}
	case MsgTypePDUSessionReleaseComplete:
		a.GsmMessage.PDUSessionReleaseComplete = NewPDUSessionReleaseComplete(MsgTypePDUSessionReleaseComplete)
		if err := a.GsmMessage.DecodePDUSessionReleaseComplete(byteArray); err != nil {
			return err
		}
	case MsgTypeStatus5GSM:
		a.GsmMessage.Status5GSM = NewStatus5GSM(MsgTypeStatus5GSM)
		if err := a.GsmMessage.DecodeStatus5GSM(byteArray); err != nil {
			return err
		}
	default:
		return fmt.Errorf("NAS Decode Fail: MsgType[%d] doesn't exist in GSM Message",
			a.GsmMessage.GsmHeader.GetMessageType())
	}
	return nil
}

func (a *Message) GsmMessageEncode(buffer *bytes.Buffer) error {
	switch a.GsmMessage.GsmHeader.GetMessageType() {
	case MsgTypePDUSessionEstablishmentRequest:
		if err := a.GsmMessage.EncodePDUSessionEstablishmentRequest(buffer); err != nil {
			return err
		}
	case MsgTypePDUSessionEstablishmentAccept:
		if err := a.GsmMessage.EncodePDUSessionEstablishmentAccept(buffer); err != nil {
			return err
		}
	case MsgTypePDUSessionEstablishmentReject:
		if err := a.GsmMessage.EncodePDUSessionEstablishmentReject(buffer); err != nil {
			return err
		}
	case MsgTypePDUSessionAuthenticationCommand:
		if err := a.GsmMessage.EncodePDUSessionAuthenticationCommand(buffer); err != nil {
			return err
		}
	case MsgTypePDUSessionAuthenticationComplete:
		if err := a.GsmMessage.EncodePDUSessionAuthenticationComplete(buffer); err != nil {
			return err
		}
	case MsgTypePDUSessionAuthenticationResult:
		if err := a.GsmMessage.EncodePDUSessionAuthenticationResult(buffer); err != nil {
			return err
		}
	case MsgTypePDUSessionModificationRequest:
		if err := a.GsmMessage.EncodePDUSessionModificationRequest(buffer); err != nil {
			return err
		}
	case MsgTypePDUSessionModificationReject:
		if err := a.GsmMessage.EncodePDUSessionModificationReject(buffer); err != nil {
			return err
		}
	case MsgTypePDUSessionModificationCommand:
		if err := a.GsmMessage.EncodePDUSessionModificationCommand(buffer); err != nil {
			return err
		}
	case MsgTypePDUSessionModificationComplete:
		if err := a.GsmMessage.EncodePDUSessionModificationComplete(buffer); err != nil {
			return err
		}
	case MsgTypePDUSessionModificationCommandReject:
		if err := a.GsmMessage.EncodePDUSessionModificationCommandReject(buffer); err != nil {
			return err
		}
	case MsgTypePDUSessionReleaseRequest:
		if err := a.GsmMessage.EncodePDUSessionReleaseRequest(buffer); err != nil {
			return err
		}
	case MsgTypePDUSessionReleaseReject:
		if err := a.GsmMessage.EncodePDUSessionReleaseReject(buffer); err != nil {
			return err
		}
	case MsgTypePDUSessionReleaseCommand:
		if err := a.GsmMessage.EncodePDUSessionReleaseCommand(buffer); err != nil {
			return err
		}
	case MsgTypePDUSessionReleaseComplete:
		if err := a.GsmMessage.EncodePDUSessionReleaseComplete(buffer); err != nil {
			return err
		}
	case MsgTypeStatus5GSM:
		if err := a.GsmMessage.EncodeStatus5GSM(buffer); err != nil {
			return err
		}
	default:
		return fmt.Errorf("NAS Encode Fail: MsgType[%d] doesn't exist in GSM Message",
			a.GsmMessage.GsmHeader.GetMessageType())
	}
	return nil
}
