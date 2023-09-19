package nas

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"

	"github.com/mimetrix/nas/nasMessage"
)

func NASDecode(byteArray *[]byte) (*nasMessage.Message, error) {

	nasMsg := nasMessage.NewMessage()

	epd := nasMessage.GetEPD(*byteArray)
	securityHdrType := (*byteArray)[1]

	switch epd {
	case nasMessage.Epd5GSMobilityManagementMessage:
		if securityHdrType == 0 {
			err := nasMsg.PlainNasDecode(byteArray)
			return nasMsg, err
		} else if securityHdrType > 0 && securityHdrType <= 4 {
			err := nasMsg.SecurityProtectedNasDecode(byteArray)
			return nasMsg, err
		} else {
			err := fmt.Errorf("Unrecognized Security Header Type: 0x%x", securityHdrType)
			return nasMsg, err
		}
	case nasMessage.Epd5GSSessionManagementMessage:
		err := nasMsg.PlainNasDecode(byteArray)
		return nasMsg, err
	default:
		err := fmt.Errorf("Unrecognized Extended Protocol Discriminator: 0x%x", epd)
		return nasMsg, err

	}

}

type mimeJSONPart struct {
	SUPI               string
	PEI                string
	GPSI               string
	PDUSessionID       uint32
	DNN                string
	SNSSAI             json.RawMessage
	ServingNfId        string
	GUAMI              json.RawMessage
	ServingNetwork     json.RawMessage
	N1SmMsg            *N1SMMSG
	ANType             string
	ContextStatusUri   string
	N1MessageContainer *N1MessageContainer
}

type N1SMMSG struct {
	ContentID string
}

type N1MessageContainer struct {
	N1MessageClass   string
	N1MessageContent N1SMMSG
}

func SecurityProtectedNasDecode(byteArray *[]byte) error {
	buffer := bytes.NewBuffer(*byteArray)

	GmmMsg := nasMessage.NewGmmMessage()
	if err := binary.Read(buffer, binary.BigEndian, &GmmMsg.GmmHeader); err != nil {
		return fmt.Errorf("GMM NAS decode Fail: read fail - %+v", err)
	}
	GmmMsg.SecurityProtected5GSNASMessage = nasMessage.NewSecurityProtected5GSNASMessage(nasMessage.MsgTypeSecurityProtected5GSNASMessage)
	GmmMsg.DecodeSecurityProtected5GSNASMessage(byteArray)

	return nil
}

func getNASMultipartType(jsonblob []byte) (msgType string, err error) {

	var jp mimeJSONPart

	json.Unmarshal([]byte(jsonblob), &jp)

	switch {
	case jp.N1MessageContainer != nil:
		msgType = jp.N1MessageContainer.N1MessageContent.ContentID
	case jp.N1SmMsg != nil:
		msgType = jp.N1SmMsg.ContentID
	default:
		return "", fmt.Errorf("currently not supporting NGAP type found in %s", jsonblob)
	}

	if msgType == "" {
		err = fmt.Errorf("Didn't find NAS message type in MIME Multipart JSON")
	}

	return

}

func MultipartDecoder(jsonblob []byte, data []byte) (nasMsg *nasMessage.Message, err error) {

	msgType, err := getNASMultipartType(jsonblob)
	if err != nil {
		return nil, err
	}

	switch msgType {
	case "n1SmMsg":
		nasMsg, err = NASDecode(&data)
		//nasMsg, err = SecurityProtectedNasDecode(&data)
	case "GSM_NAS":
		nasMsg, err = NASDecode(&data)
	}

	return

}
