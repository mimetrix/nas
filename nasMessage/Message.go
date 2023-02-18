package nasMessage

import (
	"bytes"
	"encoding/binary"
	"fmt"
	//"github.com/davecgh/go-spew/spew"
    
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
    //fmt.Println("\nNewMessage\n")
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
//           Octet2 MessageType
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
	GmmHeader                                                    `json:"-"`
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
	MsgTypeRegistrationRequest                              uint8 = 65 //0x41
	MsgTypeRegistrationAccept                               uint8 = 66 //0x42
	MsgTypeRegistrationComplete                             uint8 = 67 //0x43
	MsgTypeRegistrationReject                               uint8 = 68 //0x44
	MsgTypeDeregistrationRequestUEOriginatingDeregistration uint8 = 69 //0x45 
	MsgTypeDeregistrationAcceptUEOriginatingDeregistration  uint8 = 70 //0x46
	MsgTypeDeregistrationRequestUETerminatedDeregistration  uint8 = 71 //0x47
	MsgTypeDeregistrationAcceptUETerminatedDeregistration   uint8 = 72 //0x48
	MsgTypeServiceRequest                                   uint8 = 76 //0x4c
	MsgTypeServiceReject                                    uint8 = 77 //0x4d
	MsgTypeServiceAccept                                    uint8 = 78 //0x4e
	MsgTypeConfigurationUpdateCommand                       uint8 = 84 //0x54
	MsgTypeConfigurationUpdateComplete                      uint8 = 85 //0x55
	MsgTypeAuthenticationRequest                            uint8 = 86 //0x56
	MsgTypeAuthenticationResponse                           uint8 = 87 //0x57
	MsgTypeAuthenticationReject                             uint8 = 88 //0x58
	MsgTypeAuthenticationFailure                            uint8 = 89 //0x59
	MsgTypeAuthenticationResult                             uint8 = 90 //0x5a
	MsgTypeIdentityRequest                                  uint8 = 91 //0x5b
	MsgTypeIdentityResponse                                 uint8 = 92 //0x5c
	MsgTypeSecurityModeCommand                              uint8 = 93 //0x5d
	MsgTypeSecurityModeComplete                             uint8 = 94 //0x5e
	MsgTypeSecurityModeReject                               uint8 = 95 //0x5f 
	MsgTypeStatus5GMM                                       uint8 = 100 //0x64
	MsgTypeNotification                                     uint8 = 101 //0x65
	MsgTypeNotificationResponse                             uint8 = 102 //0x66
	MsgTypeULNASTransport                                   uint8 = 103 //0x67
	MsgTypeDLNASTransport                                   uint8 = 104 //0x67
    //0x5d
)

//placeholder
const MsgTypeSecurityProtected5GSNASMessage = 0x00

func (a *Message) SecurityProtectedNasDecode(byteArray *[]byte) error{
    buffer := bytes.NewBuffer(*byteArray)
	a.GmmMessage = NewGmmMessage()
	if err := binary.Read(buffer, binary.BigEndian, &a.GmmMessage.GmmHeader); err != nil {
		return fmt.Errorf("GMM NAS decode Fail: read fail - %+v", err)
	}
    a.GmmMessage.SecurityProtected5GSNASMessage = NewSecurityProtected5GSNASMessage(MsgTypeSecurityProtected5GSNASMessage)

	a.GmmMessage.DecodeSecurityProtected5GSNASMessage(byteArray)

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
		a.GmmMessage.DecodeRegistrationRequest(byteArray)
        
	case MsgTypeRegistrationAccept:
		a.GmmMessage.RegistrationAccept = NewRegistrationAccept(MsgTypeRegistrationAccept)
		a.GmmMessage.DecodeRegistrationAccept(byteArray)
	case MsgTypeRegistrationComplete:
		a.GmmMessage.RegistrationComplete = NewRegistrationComplete(MsgTypeRegistrationComplete)
		a.GmmMessage.DecodeRegistrationComplete(byteArray)
	case MsgTypeRegistrationReject:
		a.GmmMessage.RegistrationReject = NewRegistrationReject(MsgTypeRegistrationReject)
		a.GmmMessage.DecodeRegistrationReject(byteArray)
	case MsgTypeDeregistrationRequestUEOriginatingDeregistration:
		a.GmmMessage.DeregistrationRequestUEOriginatingDeregistration =
			NewDeregistrationRequestUEOriginatingDeregistration(
				MsgTypeDeregistrationRequestUEOriginatingDeregistration)
		a.GmmMessage.DecodeDeregistrationRequestUEOriginatingDeregistration(byteArray)
	case MsgTypeDeregistrationAcceptUEOriginatingDeregistration:
		a.GmmMessage.DeregistrationAcceptUEOriginatingDeregistration =
			NewDeregistrationAcceptUEOriginatingDeregistration(
				MsgTypeDeregistrationAcceptUEOriginatingDeregistration)
		a.GmmMessage.DecodeDeregistrationAcceptUEOriginatingDeregistration(byteArray)
	case MsgTypeDeregistrationRequestUETerminatedDeregistration:
		a.GmmMessage.DeregistrationRequestUETerminatedDeregistration =
			NewDeregistrationRequestUETerminatedDeregistration(
				MsgTypeDeregistrationRequestUETerminatedDeregistration)
		a.GmmMessage.DecodeDeregistrationRequestUETerminatedDeregistration(byteArray)
	case MsgTypeDeregistrationAcceptUETerminatedDeregistration:
		a.GmmMessage.DeregistrationAcceptUETerminatedDeregistration =
			NewDeregistrationAcceptUETerminatedDeregistration(
				MsgTypeDeregistrationAcceptUETerminatedDeregistration)
		a.GmmMessage.DecodeDeregistrationAcceptUETerminatedDeregistration(byteArray)
	case MsgTypeServiceRequest:
		a.GmmMessage.ServiceRequest = NewServiceRequest(MsgTypeServiceRequest)
		a.GmmMessage.DecodeServiceRequest(byteArray)
	case MsgTypeServiceReject:
		a.GmmMessage.ServiceReject = NewServiceReject(MsgTypeServiceReject)
		a.GmmMessage.DecodeServiceReject(byteArray)
	case MsgTypeServiceAccept:
		a.GmmMessage.ServiceAccept = NewServiceAccept(MsgTypeServiceAccept)
		a.GmmMessage.DecodeServiceAccept(byteArray)
	case MsgTypeConfigurationUpdateCommand:
		a.GmmMessage.ConfigurationUpdateCommand =
			NewConfigurationUpdateCommand(MsgTypeConfigurationUpdateCommand)
		a.GmmMessage.DecodeConfigurationUpdateCommand(byteArray)
	case MsgTypeConfigurationUpdateComplete:
		a.GmmMessage.ConfigurationUpdateComplete =
			NewConfigurationUpdateComplete(MsgTypeConfigurationUpdateComplete)
		a.GmmMessage.DecodeConfigurationUpdateComplete(byteArray)
	case MsgTypeAuthenticationRequest:
		a.GmmMessage.AuthenticationRequest = NewAuthenticationRequest(MsgTypeAuthenticationRequest)
		a.GmmMessage.DecodeAuthenticationRequest(byteArray)
	case MsgTypeAuthenticationResponse:
		a.GmmMessage.AuthenticationResponse = NewAuthenticationResponse(MsgTypeAuthenticationResponse)
		a.GmmMessage.DecodeAuthenticationResponse(byteArray)
	case MsgTypeAuthenticationReject:
		a.GmmMessage.AuthenticationReject = NewAuthenticationReject(MsgTypeAuthenticationReject)
		a.GmmMessage.DecodeAuthenticationReject(byteArray)
	case MsgTypeAuthenticationFailure:
		a.GmmMessage.AuthenticationFailure = NewAuthenticationFailure(MsgTypeAuthenticationFailure)
		a.GmmMessage.DecodeAuthenticationFailure(byteArray)
	case MsgTypeAuthenticationResult:
		a.GmmMessage.AuthenticationResult = NewAuthenticationResult(MsgTypeAuthenticationResult)
		a.GmmMessage.DecodeAuthenticationResult(byteArray)
	case MsgTypeIdentityRequest:
		a.GmmMessage.IdentityRequest = NewIdentityRequest(MsgTypeIdentityRequest)
		a.GmmMessage.DecodeIdentityRequest(byteArray)
	case MsgTypeIdentityResponse:
		a.GmmMessage.IdentityResponse = NewIdentityResponse(MsgTypeIdentityResponse)
		a.GmmMessage.DecodeIdentityResponse(byteArray)
	case MsgTypeSecurityModeCommand:
		a.GmmMessage.SecurityModeCommand = NewSecurityModeCommand(MsgTypeSecurityModeCommand)
		a.GmmMessage.DecodeSecurityModeCommand(byteArray)
	case MsgTypeSecurityModeComplete:
		a.GmmMessage.SecurityModeComplete = NewSecurityModeComplete(MsgTypeSecurityModeComplete)
		a.GmmMessage.DecodeSecurityModeComplete(byteArray)
	case MsgTypeSecurityModeReject:
		a.GmmMessage.SecurityModeReject = NewSecurityModeReject(MsgTypeSecurityModeReject)
		a.GmmMessage.DecodeSecurityModeReject(byteArray)
	case MsgTypeStatus5GMM:
		a.GmmMessage.Status5GMM = NewStatus5GMM(MsgTypeStatus5GMM)
		a.GmmMessage.DecodeStatus5GMM(byteArray)
	case MsgTypeNotification:
		a.GmmMessage.Notification = NewNotification(MsgTypeNotification)
		a.GmmMessage.DecodeNotification(byteArray)
	case MsgTypeNotificationResponse:
		a.GmmMessage.NotificationResponse = NewNotificationResponse(MsgTypeNotificationResponse)
		a.GmmMessage.DecodeNotificationResponse(byteArray)
	case MsgTypeULNASTransport:
		a.GmmMessage.ULNASTransport = NewULNASTransport(MsgTypeULNASTransport)
		a.GmmMessage.DecodeULNASTransport(byteArray)
	case MsgTypeDLNASTransport:
		a.GmmMessage.DLNASTransport = NewDLNASTransport(MsgTypeDLNASTransport)
		a.GmmMessage.DecodeDLNASTransport(byteArray)
    case MsgTypeSecurityProtected5GSNASMessage:  
		a.GmmMessage.SecurityProtected5GSNASMessage = NewSecurityProtected5GSNASMessage(MsgTypeSecurityProtected5GSNASMessage)
		a.GmmMessage.DecodeSecurityProtected5GSNASMessage(byteArray)
	default:
		return fmt.Errorf("NAS decode Fail: MsgType[%x] doesn't exist in GMM Message",
			a.GmmMessage.GmmHeader.GetMessageType())
	}
	return nil
}

func (a *Message) GmmMessageEncode(buffer *bytes.Buffer) error {
	switch a.GmmMessage.GmmHeader.GetMessageType() {
	case MsgTypeRegistrationRequest:
		a.GmmMessage.EncodeRegistrationRequest(buffer)
	case MsgTypeRegistrationAccept:
		a.GmmMessage.EncodeRegistrationAccept(buffer)
	case MsgTypeRegistrationComplete:
		a.GmmMessage.EncodeRegistrationComplete(buffer)
	case MsgTypeRegistrationReject:
		a.GmmMessage.EncodeRegistrationReject(buffer)
	case MsgTypeDeregistrationRequestUEOriginatingDeregistration:
		a.GmmMessage.EncodeDeregistrationRequestUEOriginatingDeregistration(buffer)
	case MsgTypeDeregistrationAcceptUEOriginatingDeregistration:
		a.GmmMessage.EncodeDeregistrationAcceptUEOriginatingDeregistration(buffer)
	case MsgTypeDeregistrationRequestUETerminatedDeregistration:
		a.GmmMessage.EncodeDeregistrationRequestUETerminatedDeregistration(buffer)
	case MsgTypeDeregistrationAcceptUETerminatedDeregistration:
		a.GmmMessage.EncodeDeregistrationAcceptUETerminatedDeregistration(buffer)
	case MsgTypeServiceRequest:
		a.GmmMessage.EncodeServiceRequest(buffer)
	case MsgTypeServiceReject:
		a.GmmMessage.EncodeServiceReject(buffer)
	case MsgTypeServiceAccept:
		a.GmmMessage.EncodeServiceAccept(buffer)
	case MsgTypeConfigurationUpdateCommand:
		a.GmmMessage.EncodeConfigurationUpdateCommand(buffer)
	case MsgTypeConfigurationUpdateComplete:
		a.GmmMessage.EncodeConfigurationUpdateComplete(buffer)
	case MsgTypeAuthenticationRequest:
		a.GmmMessage.EncodeAuthenticationRequest(buffer)
	case MsgTypeAuthenticationResponse:
		a.GmmMessage.EncodeAuthenticationResponse(buffer)
	case MsgTypeAuthenticationReject:
		a.GmmMessage.EncodeAuthenticationReject(buffer)
	case MsgTypeAuthenticationFailure:
		a.GmmMessage.EncodeAuthenticationFailure(buffer)
	case MsgTypeAuthenticationResult:
		a.GmmMessage.EncodeAuthenticationResult(buffer)
	case MsgTypeIdentityRequest:
		a.GmmMessage.EncodeIdentityRequest(buffer)
	case MsgTypeIdentityResponse:
		a.GmmMessage.EncodeIdentityResponse(buffer)
	case MsgTypeSecurityModeCommand:
		a.GmmMessage.EncodeSecurityModeCommand(buffer)
	case MsgTypeSecurityModeComplete:
		a.GmmMessage.EncodeSecurityModeComplete(buffer)
	case MsgTypeSecurityModeReject:
		a.GmmMessage.EncodeSecurityModeReject(buffer)
	case MsgTypeStatus5GMM:
		a.GmmMessage.EncodeStatus5GMM(buffer)
	case MsgTypeNotification:
		a.GmmMessage.EncodeNotification(buffer)
	case MsgTypeNotificationResponse:
		a.GmmMessage.EncodeNotificationResponse(buffer)
	case MsgTypeULNASTransport:
		a.GmmMessage.EncodeULNASTransport(buffer)
	case MsgTypeDLNASTransport:
		a.GmmMessage.EncodeDLNASTransport(buffer)
	default:
		return fmt.Errorf("NAS Encode Fail: MsgType[%d] doesn't exist in GMM Message",
			a.GmmMessage.GmmHeader.GetMessageType())
	}
	return nil
}

type GsmMessage struct {
	GsmHeader                                       `json:"-"`
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
		a.GsmMessage.PDUSessionEstablishmentRequest =
			NewPDUSessionEstablishmentRequest(MsgTypePDUSessionEstablishmentRequest)
		a.GsmMessage.DecodePDUSessionEstablishmentRequest(byteArray)
	case MsgTypePDUSessionEstablishmentAccept:
		a.GsmMessage.PDUSessionEstablishmentAccept =
			NewPDUSessionEstablishmentAccept(MsgTypePDUSessionEstablishmentAccept)
		a.GsmMessage.DecodePDUSessionEstablishmentAccept(byteArray)
	case MsgTypePDUSessionEstablishmentReject:
		a.GsmMessage.PDUSessionEstablishmentReject =
			NewPDUSessionEstablishmentReject(MsgTypePDUSessionEstablishmentReject)
		a.GsmMessage.DecodePDUSessionEstablishmentReject(byteArray)
	case MsgTypePDUSessionAuthenticationCommand:
		a.GsmMessage.PDUSessionAuthenticationCommand =
			NewPDUSessionAuthenticationCommand(MsgTypePDUSessionAuthenticationCommand)
		a.GsmMessage.DecodePDUSessionAuthenticationCommand(byteArray)
	case MsgTypePDUSessionAuthenticationComplete:
		a.GsmMessage.PDUSessionAuthenticationComplete =
			NewPDUSessionAuthenticationComplete(MsgTypePDUSessionAuthenticationComplete)
		a.GsmMessage.DecodePDUSessionAuthenticationComplete(byteArray)
	case MsgTypePDUSessionAuthenticationResult:
		a.GsmMessage.PDUSessionAuthenticationResult =
			NewPDUSessionAuthenticationResult(MsgTypePDUSessionAuthenticationResult)
		a.GsmMessage.DecodePDUSessionAuthenticationResult(byteArray)
	case MsgTypePDUSessionModificationRequest:
		a.GsmMessage.PDUSessionModificationRequest =
			NewPDUSessionModificationRequest(MsgTypePDUSessionModificationRequest)
		a.GsmMessage.DecodePDUSessionModificationRequest(byteArray)
	case MsgTypePDUSessionModificationReject:
		a.GsmMessage.PDUSessionModificationReject =
			NewPDUSessionModificationReject(MsgTypePDUSessionModificationReject)
		a.GsmMessage.DecodePDUSessionModificationReject(byteArray)
	case MsgTypePDUSessionModificationCommand:
		a.GsmMessage.PDUSessionModificationCommand =
			NewPDUSessionModificationCommand(MsgTypePDUSessionModificationCommand)
		a.GsmMessage.DecodePDUSessionModificationCommand(byteArray)
	case MsgTypePDUSessionModificationComplete:
		a.GsmMessage.PDUSessionModificationComplete =
			NewPDUSessionModificationComplete(MsgTypePDUSessionModificationComplete)
		a.GsmMessage.DecodePDUSessionModificationComplete(byteArray)
	case MsgTypePDUSessionModificationCommandReject:
		a.GsmMessage.PDUSessionModificationCommandReject =
			NewPDUSessionModificationCommandReject(MsgTypePDUSessionModificationCommandReject)
		a.GsmMessage.DecodePDUSessionModificationCommandReject(byteArray)
	case MsgTypePDUSessionReleaseRequest:
		a.GsmMessage.PDUSessionReleaseRequest = NewPDUSessionReleaseRequest(MsgTypePDUSessionReleaseRequest)
		a.GsmMessage.DecodePDUSessionReleaseRequest(byteArray)
	case MsgTypePDUSessionReleaseReject:
		a.GsmMessage.PDUSessionReleaseReject = NewPDUSessionReleaseReject(MsgTypePDUSessionReleaseReject)
		a.GsmMessage.DecodePDUSessionReleaseReject(byteArray)
	case MsgTypePDUSessionReleaseCommand:
		a.GsmMessage.PDUSessionReleaseCommand = NewPDUSessionReleaseCommand(MsgTypePDUSessionReleaseCommand)
		a.GsmMessage.DecodePDUSessionReleaseCommand(byteArray)
	case MsgTypePDUSessionReleaseComplete:
		a.GsmMessage.PDUSessionReleaseComplete = NewPDUSessionReleaseComplete(MsgTypePDUSessionReleaseComplete)
		a.GsmMessage.DecodePDUSessionReleaseComplete(byteArray)
	case MsgTypeStatus5GSM:
		a.GsmMessage.Status5GSM = NewStatus5GSM(MsgTypeStatus5GSM)
		a.GsmMessage.DecodeStatus5GSM(byteArray)
	default:
		return fmt.Errorf("NAS Decode Fail: MsgType[%d] doesn't exist in GSM Message",
			a.GsmMessage.GsmHeader.GetMessageType())
	}
	return nil
}

func (a *Message) GsmMessageEncode(buffer *bytes.Buffer) error {
	switch a.GsmMessage.GsmHeader.GetMessageType() {
	case MsgTypePDUSessionEstablishmentRequest:
		a.GsmMessage.EncodePDUSessionEstablishmentRequest(buffer)
	case MsgTypePDUSessionEstablishmentAccept:
		a.GsmMessage.EncodePDUSessionEstablishmentAccept(buffer)
	case MsgTypePDUSessionEstablishmentReject:
		a.GsmMessage.EncodePDUSessionEstablishmentReject(buffer)
	case MsgTypePDUSessionAuthenticationCommand:
		a.GsmMessage.EncodePDUSessionAuthenticationCommand(buffer)
	case MsgTypePDUSessionAuthenticationComplete:
		a.GsmMessage.EncodePDUSessionAuthenticationComplete(buffer)
	case MsgTypePDUSessionAuthenticationResult:
		a.GsmMessage.EncodePDUSessionAuthenticationResult(buffer)
	case MsgTypePDUSessionModificationRequest:
		a.GsmMessage.EncodePDUSessionModificationRequest(buffer)
	case MsgTypePDUSessionModificationReject:
		a.GsmMessage.EncodePDUSessionModificationReject(buffer)
	case MsgTypePDUSessionModificationCommand:
		a.GsmMessage.EncodePDUSessionModificationCommand(buffer)
	case MsgTypePDUSessionModificationComplete:
		a.GsmMessage.EncodePDUSessionModificationComplete(buffer)
	case MsgTypePDUSessionModificationCommandReject:
		a.GsmMessage.EncodePDUSessionModificationCommandReject(buffer)
	case MsgTypePDUSessionReleaseRequest:
		a.GsmMessage.EncodePDUSessionReleaseRequest(buffer)
	case MsgTypePDUSessionReleaseReject:
		a.GsmMessage.EncodePDUSessionReleaseReject(buffer)
	case MsgTypePDUSessionReleaseCommand:
		a.GsmMessage.EncodePDUSessionReleaseCommand(buffer)
	case MsgTypePDUSessionReleaseComplete:
		a.GsmMessage.EncodePDUSessionReleaseComplete(buffer)
	case MsgTypeStatus5GSM:
		a.GsmMessage.EncodeStatus5GSM(buffer)
	default:
		return fmt.Errorf("NAS Encode Fail: MsgType[%d] doesn't exist in GSM Message",
			a.GsmMessage.GsmHeader.GetMessageType())
	}
	return nil
}

