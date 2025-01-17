package nasTests

import (
	"flag"
	"fmt"
	"os"

	//"github.com/davecgh/go-spew/spew"
	"encoding/json"
	"path/filepath"
	"testing"

	"github.com/mimetrix/nas"

	//"nas"

	"github.com/mimetrix/nas/nasMessage"
)

var hexString = "7e00560102000021e440b883d63a9f9c56b3703217152eba2010068f241c77748000b2180e54a9760068"
var printJSON = flag.Bool("printjson", false, "Dump unmarshalled json to stdout")
var saveJSON = flag.String("savejson", "", "Write json to directory")

/*
================================================
        MAIN
================================================
*/

func TestMain(m *testing.M) {
	flag.Parse()
	os.Exit(m.Run())
}

/*
================================================
        TEST INDIVIDUAL MESSAGES
================================================
*/

func TestRegistrationRequest(t *testing.T) {
	decodedJSON, nasMsg := decodeNASMsg(t, RegistrationRequest)
	outputJSON(t, "RegistrationRequest.json", decodedJSON)

	RegistrationRequest := nasMsg.GmmMessage.RegistrationRequest
	EPD := RegistrationRequest.ExtendedProtocolDiscriminator.EPD
	SUCI := RegistrationRequest.MobileIdentity5GS.SUCI
	S1Mode := fmt.Sprintf("%t", RegistrationRequest.Capability5GMM.S1Mode)
	HOAttach := fmt.Sprintf("%t", RegistrationRequest.Capability5GMM.HOAttach)

	var expectedList = map[string]string{
		EPD:      "5GS mobility management messages",
		SUCI:     "suci-0-001-01-0000-0-0-0123456780",
		S1Mode:   "true",
		HOAttach: "true",
	}
	compareResults(t, expectedList)

}

func TestAuthenticationRequest(t *testing.T) {
	decodedJSON, nasMsg := decodeNASMsg(t, AuthenticationRequest)
	outputJSON(t, "AuthenticationRequest.json", decodedJSON)

	AuthenticationRequest := nasMsg.GmmMessage.AuthenticationRequest
	MessageType := AuthenticationRequest.AuthenticationRequestMessageIdentity.MessageType
	RAND := AuthenticationRequest.AuthenticationParameterRAND.Rand
	AUTN := AuthenticationRequest.AuthenticationParameterAUTN.Autn

	var expectedList = map[string]string{
		MessageType: "Authentication request",
		RAND:        "d00408bc739df5113e9d4a6e89c1be31",
		AUTN:        "3452632de0a88000f05f4041f98d5a5d",
	}
	compareResults(t, expectedList)
}

func TestAuthenticationResponse(t *testing.T) {
	decodedJSON, nasMsg := decodeNASMsg(t, AuthenticationResponse)
	outputJSON(t, "AuthenticationResponse.json", decodedJSON)

	AuthenticationResponse := nasMsg.GmmMessage.AuthenticationResponse
	MessageType := AuthenticationResponse.AuthenticationResponseMessageIdentity.MessageType
	RES := AuthenticationResponse.AuthenticationResponseParameter.RES

	var expectedList = map[string]string{
		MessageType: "Authentication response",
		RES:         "920c41828fefd1208e4d06d117949b17",
	}
	compareResults(t, expectedList)
}

func TestSecProtectedSecurityModeCommand(t *testing.T) {
	decodedJSON, nasMsg := decodeNASMsg(t, SecProtectedSecurityModeCommand)
	outputJSON(t, "SecProtectedSecurityModeCommand.json", decodedJSON)

	secModeCommand := nasMsg.GmmMessage.PlainNASMessage.GmmMessage.SecurityModeCommand
	MessageType := secModeCommand.SecurityModeCommandMessageIdentity.MessageType
	integrityAlgorithm := secModeCommand.SelectedNASSecurityAlgorithms.IntegrityAlgorithm
	IMEISVReq := secModeCommand.IMEISVRequest.RequestValue

	var expectedList = map[string]string{
		MessageType:        "Security mode command",
		integrityAlgorithm: "5G integrity algorithm 128-5G-IA2",
		IMEISVReq:          "IMEISV requested",
	}
	compareResults(t, expectedList)
}

func TestSecProtectedSecurityModeComplete(t *testing.T) {
	decodedJSON, nasMsg := decodeNASMsg(t, SecProtectedSecurityModeComplete)
	outputJSON(t, "SecProtectedSecurityModeComplete.json", decodedJSON)

	secModeComplete := nasMsg.GmmMessage.PlainNASMessage.GmmMessage.SecurityModeComplete
	MAC := nasMsg.GmmMessage.MessageAuthenticationCode.MAC
	MessageType := secModeComplete.SecurityModeCompleteMessageIdentity.MessageType
	var expectedList = map[string]string{
		MessageType: "Security mode complete",
		MAC:         "f5f20c34",
	}
	compareResults(t, expectedList)
}

func TestSecProtectedRegistrationAccept(t *testing.T) {
	decodedJSON, nasMsg := decodeNASMsg(t, SecProtectedRegistrationAccept)
	outputJSON(t, "SecProtectedRegistrationAccept.json", decodedJSON)

	regAccept := nasMsg.GmmMessage.PlainNASMessage.GmmMessage.RegistrationAccept
	MessageType := regAccept.RegistrationAcceptMessageIdentity.MessageType
	MAC := nasMsg.GmmMessage.MessageAuthenticationCode.MAC
	TMSI := regAccept.GUTI5G.TMSI

	var expectedList = map[string]string{
		MessageType: "Registration accept",
		MAC:         "bf53cc9a",
		TMSI:        "f70006fb",
	}
	compareResults(t, expectedList)
}

func TestSecProtectedRegistrationComplete(t *testing.T) {
	decodedJSON, nasMsg := decodeNASMsg(t, SecProtectedRegistrationComplete)
	outputJSON(t, "SecProtectedRegistrationComplete.json", decodedJSON)

	regComplete := nasMsg.GmmMessage.PlainNASMessage.GmmMessage.RegistrationComplete
	MAC := nasMsg.GmmMessage.MessageAuthenticationCode.MAC
	MessageType := regComplete.RegistrationCompleteMessageIdentity.MessageType

	var expectedList = map[string]string{
		MessageType: "Registration complete",
		MAC:         "ce215fba",
	}
	compareResults(t, expectedList)
}

func TestSecProtectedConfigurationUpdateCommand(t *testing.T) {
	decodedJSON, nasMsg := decodeNASMsg(t, SecProtectedConfigurationUpdateCommand)
	outputJSON(t, "SecProtectedConfigurationUpdateCommand.json", decodedJSON)

	updateCmd := nasMsg.GmmMessage.PlainNASMessage.GmmMessage.ConfigurationUpdateCommand
	MAC := nasMsg.GmmMessage.MessageAuthenticationCode.MAC
	networkFullName := updateCmd.FullNameForNetwork.NetworkName
	networkShortName := updateCmd.ShortNameForNetwork.NetworkName
	MessageType := updateCmd.ConfigurationUpdateCommandMessageIdentity.MessageType
	var expectedList = map[string]string{
		MessageType:      "Configuration update command",
		MAC:              "af2e0089",
		networkFullName:  "DougCo 5G Network Emulation",
		networkShortName: "Open5GS",
		MAC:              "af2e0089",
	}
	compareResults(t, expectedList)
}

func TestSecProtectedULNASTransport1(t *testing.T) {
	decodedJSON, nasMsg := decodeNASMsg(t, SecProtectedULNASTransport1)
	outputJSON(t, "SecProtectedULNASTransport1.json", decodedJSON)

	ULNASTrans := nasMsg.GmmMessage.PlainNASMessage.GmmMessage.ULNASTransport
	MessageType := ULNASTrans.ULNASTRANSPORTMessageIdentity.MessageType
	MAC := nasMsg.GmmMessage.MessageAuthenticationCode.MAC
	SDBytes := ULNASTrans.SNSSAI.SDBytes
	FQDN := ULNASTrans.DNN.FQDN

	var expectedList = map[string]string{
		MessageType: "UL NAS transport",
		MAC:         "fa507e7e",
		SDBytes:     "ffffff",
		FQDN:        "internet",
	}
	compareResults(t, expectedList)
}

func TestSecProtectedULNASTransport2(t *testing.T) {
	decodedJSON, _ := decodeNASMsg(t, SecProtectedULNASTransport2)
	outputJSON(t, "SecProtectedULNASTransport2.json", decodedJSON)
}

func TestSecProtectedDLNASTransport1(t *testing.T) {
	decodedJSON, nasMsg := decodeNASMsg(t, SecProtectedDLNASTransport1)
	outputJSON(t, "SecProtectedDLNASTransport.json", decodedJSON)
	PlainMessage := nasMsg.GmmMessage.SecurityProtected5GSNASMessage.PlainNASMessage
	DLNASTrans := PlainMessage.GmmMessage.DLNASTransport

	MessageType := DLNASTrans.DLNASTRANSPORTMessageIdentity.MessageType
	MAC := nasMsg.GmmMessage.MessageAuthenticationCode.MAC
	payloadContainer := DLNASTrans.SpareHalfOctetAndPayloadContainerType.PayloadContainerType
	var expectedList = map[string]string{
		MessageType:      "DL NAS transport",
		MAC:              "86ee320d",
		payloadContainer: "N1 SM information",
	}
	compareResults(t, expectedList)
}

func TestSecProtectedDLNASTransport2(t *testing.T) {
	decodedJSON, nasMsg := decodeNASMsg(t, SecProtectedDLNASTransport2)
	outputJSON(t, "SecProtectedDLNASTransport.json", decodedJSON)
	PlainMessage := nasMsg.GmmMessage.SecurityProtected5GSNASMessage.PlainNASMessage
	DLNASTrans := PlainMessage.GmmMessage.DLNASTransport
	Cause := DLNASTrans.Cause5GMM.Cause

	MessageType := DLNASTrans.DLNASTRANSPORTMessageIdentity.MessageType
	MAC := nasMsg.GmmMessage.MessageAuthenticationCode.MAC
	payloadContainer := DLNASTrans.SpareHalfOctetAndPayloadContainerType.PayloadContainerType
	var expectedList = map[string]string{
		MessageType:      "DL NAS transport",
		MAC:              "d940a158",
		payloadContainer: "N1 SM information",
		Cause:            "Payload was not forwarded",
	}
	compareResults(t, expectedList)
}

func TestSecProtectedDLNASTransport3(t *testing.T) {
	decodedJSON, nasMsg := decodeNASMsg(t, SecProtectedDLNASTransport3)
	outputJSON(t, "SecProtectedDLNASTransport.json", decodedJSON)
	PlainMessage := nasMsg.GmmMessage.SecurityProtected5GSNASMessage.PlainNASMessage
	DLNASTrans := PlainMessage.GmmMessage.DLNASTransport
	gsmMessage := DLNASTrans.PayloadContainer.NASMessage.GsmMessage
	Cause := gsmMessage.PDUSessionReleaseCommand.Cause5GSM.Cause

	MessageType := DLNASTrans.DLNASTRANSPORTMessageIdentity.MessageType
	MAC := nasMsg.GmmMessage.MessageAuthenticationCode.MAC
	payloadContainer := DLNASTrans.SpareHalfOctetAndPayloadContainerType.PayloadContainerType
	var expectedList = map[string]string{
		MessageType:      "DL NAS transport",
		MAC:              "b80ee9d9",
		payloadContainer: "N1 SM information",
		Cause:            "Regular deactivation",
	}
	compareResults(t, expectedList)
}

/*
================================================
       MIME MULTIPART TESTS
================================================
*/

func TestMIME_N1SMSG(t *testing.T) {
	nasMsg, err := nas.MultipartDecoder(N1SMMSG, N1SMMSGBytes)
	if err != nil {
		t.Log(err)
		t.Error("Failed to Decode Message")
	}
	imJSON, err := json.MarshalIndent(nasMsg, "", "    ")
	jsonStr := string(imJSON)
	if err != nil {
		t.Log(err)
		t.Error("Failed to Marshal JSON")
	}

	//spew.Dump(PlainMessage.GsmMessage)
	pser := nasMsg.GsmMessage.PDUSessionEstablishmentRequest

	EPD := pser.ExtendedProtocolDiscriminator.EPD
	uplinkRate := pser.IntegrityProtectionMaximumDataRate.UplinkDataRate
	SessionType := pser.PDUSessionType.SessionType

	var expectedList = map[string]string{
		EPD:         "5GS session management messages",
		uplinkRate:  "Full data rate",
		SessionType: "IPv4",
	}
	compareResults(t, expectedList)

	outputJSON(t, "MIME_N1SMSG.json", jsonStr)

}

/*
TODO: fix SessionAMBR message, values are not correct
TODO: ExtendedProtocolConfigurationOptions not decoding at all
TODO: SSCMode not decoding at all
*/
func TestMIME_GSMNAS(t *testing.T) {
	nasMsg, err := nas.MultipartDecoder(GSMNAS, GSMNASBytes)
	if err != nil {
		t.Log(err)
		t.Error("Failed to Decode Message")
	}
	imJSON, err := json.MarshalIndent(nasMsg, "", "    ")
	jsonStr := string(imJSON)
	if err != nil {
		t.Log(err)
		t.Error("Failed to Marshal JSON")
	}

	gsmMsg := nasMsg.GsmMessage

	EPD := gsmMsg.PDUSessionEstablishmentAccept.ExtendedProtocolDiscriminator.EPD
	IPv4Address := gsmMsg.PDUAddress.IPv4Address
	FQDN := gsmMsg.DNN.FQDN

	var expectedList = map[string]string{
		EPD:         "5GS session management messages",
		IPv4Address: "10.1.0.5",
		FQDN:        "internet",
	}

	compareResults(t, expectedList)

	outputJSON(t, "MIME_GSMNAS.json", jsonStr)

}

func TestMobilityMessage(t *testing.T) {
	nasMsg, err := nas.NASDecode(&MobilityMgmtMsg)
	if err != nil {
		t.Log(err)
		t.Error("Failed to Decode Message")
	}
	imJSON, err := json.MarshalIndent(nasMsg, "", "    ")
	jsonStr := string(imJSON)
	if err != nil {
		t.Log(err)
		t.Error("Failed to Marshal JSON")
	}
	MAC := nasMsg.GmmMessage.SecurityProtected5GSNASMessage.MessageAuthenticationCode.MAC
	extGUTI := nasMsg.GmmMessage.SecurityProtected5GSNASMessage.PlainNASMessage.
		GmmMessage.RegistrationRequest.MobileIdentity5GS.FiveGGUTI
	intGUTI := nasMsg.GmmMessage.SecurityProtected5GSNASMessage.PlainNASMessage.
		GmmMessage.RegistrationRequest.NASMessageContainer.NASMessage.GmmMessage.RegistrationRequest.MobileIdentity5GS.FiveGGUTI

	var expectedList = map[string]string{
		MAC:     "aaa94d1e",
		extGUTI: "310012010041c2c00005",
		intGUTI: "310012010041c2c00005",
	}
	compareResults(t, expectedList)

	outputJSON(t, "Mobility_Message.json", jsonStr)
}

/*
================================================
        AUXILIARY METHODS
================================================
*/

func compareResults(t *testing.T, results map[string]string) {
	for got := range results {
		if got != results[got] {
			t.Errorf("Expected: %s, got: %s", results[got], got)
		}
	}
}

func outputJSON(t *testing.T, fname string, jsonString string) {
	if *printJSON {
		t.Log(jsonString)
	}
	if *saveJSON != "" {
		err := writeJSONToFile(fname, jsonString)
		if err != nil {
			t.Error(err)
		}
	}
}

func decodeNASMsg(t *testing.T, bytes []byte) (jsonStr string, nasMsg *nasMessage.Message) {
	nasMsg, err := nas.NASDecode(&bytes)
	if err != nil {
		t.Log(err)
		t.Error("Failed to Marshal aper")
		t.FailNow()
	} else {
		imJSON, err := json.MarshalIndent(nasMsg, "", "    ")
		jsonStr = string(imJSON)
		if err != nil {
			t.Log(err)
			t.Error("Failed to Marshal JSON")
		} else {
			return
		}
	}
	return
}

func writeJSONToFile(fname string, jsonString string) error {
	mode := int(0777)
	fileMode := os.FileMode(mode)

	_, err := os.Stat(*saveJSON)
	if os.IsNotExist(err) {
		err := os.Mkdir(*saveJSON, os.FileMode(mode))
		if err != nil {
			return err
		}
	}
	fpath := filepath.Join(*saveJSON, fname)
	err = os.WriteFile(fpath, []byte(jsonString), fileMode)
	if err != nil {
		return err
	}
	return nil
}
