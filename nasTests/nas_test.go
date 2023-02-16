package nasTests

import (
    "os"
    //"fmt"
	"flag"
    "github.com/davecgh/go-spew/spew"
	"github.com/mimetrix/nas/nasMessage"
	"github.com/mimetrix/nas"
    "path/filepath"
	"testing"
	"encoding/json"
)

var hexString = "7e00560102000021e440b883d63a9f9c56b3703217152eba2010068f241c77748000b2180e54a9760068"
var printJSON = flag.Bool("printjson", false, "Dump unmarshalled json to stdout")
var saveJSON = flag.String("savejson", "", "Write json to directory")


/*
================================================
        MAIN
================================================
*/

func TestMain(m *testing.M){
    flag.Parse()
    os.Exit(m.Run())
}

/*
================================================
        TEST INDIVIDUAL MESSAGES
================================================
*/

func TestRegistrationRequest(t *testing.T){
    decodedJSON, _:= decodeNASMsg(t, RegistrationRequest)
    outputJSON(t, "RegistrationRequest.json", decodedJSON) 
    //spew.Dump(nasMsg)
}

func TestAuthenticationRequest(t *testing.T){
    decodedJSON, nasMsg := decodeNASMsg(t, AuthenticationRequest)
    outputJSON(t, "AuthenticationRequest.json", decodedJSON) 
    spew.Dump(nasMsg)
}

func TestAuthenticationResponse(t *testing.T){
    decodedJSON, nasMsg := decodeNASMsg(t, AuthenticationResponse)
    outputJSON(t, "AuthenticationResponse.json", decodedJSON) 
    spew.Dump(nasMsg)
}

func TestSecProtectedSecurityModeCommand(t *testing.T){
    decodedJSON, nasMsg := decodeNASMsg(t, SecProtectedSecurityModeCommand)
    outputJSON(t, "SecProtectedSecurityModeCommand.json", decodedJSON) 
    spew.Dump(nasMsg)
}

func TestSecProtectedSecurityModeComplete(t *testing.T){
    decodedJSON, nasMsg := decodeNASMsg(t, SecProtectedSecurityModeComplete)
    outputJSON(t, "SecProtectedSecurityModeComplete.json", decodedJSON) 
    spew.Dump(nasMsg)
}

func TestSecProtectedRegistrationAccept(t *testing.T){
    decodedJSON, nasMsg := decodeNASMsg(t, SecProtectedRegistrationAccept)
    outputJSON(t, "SecProtectedRegistrationAccept.json", decodedJSON) 
    spew.Dump(nasMsg)
}

func TestSecProtectedRegistrationComplete(t *testing.T){
    decodedJSON, nasMsg := decodeNASMsg(t, SecProtectedRegistrationComplete)
    outputJSON(t, "SecProtectedRegistrationComplete.json", decodedJSON) 
    spew.Dump(nasMsg)
}

func TestSecProtectedConfigurationUpdateCommand(t *testing.T){
    decodedJSON, nasMsg := decodeNASMsg(t, SecProtectedConfigurationUpdateCommand)
    outputJSON(t, "SecProtectedConfigurationUpdateCommand.json", decodedJSON) 
    spew.Dump(nasMsg)
}

func TestSecProtectedULNASTransport(t *testing.T){
    decodedJSON, nasMsg := decodeNASMsg(t, SecProtectedULNASTransport)
    outputJSON(t, "SecProtectedULNASTransport.json", decodedJSON) 
    spew.Dump(nasMsg)
}

func TestSecProtectedDLNASTransport(t *testing.T){
    decodedJSON, nasMsg := decodeNASMsg(t, SecProtectedDLNASTransport  )
    outputJSON(t, "SecProtectedDLNASTransport.json", decodedJSON) 
    spew.Dump(nasMsg)
}


/*
================================================
        AUXILIARY METHODS 
================================================
*/
func outputJSON(t *testing.T, fname string, jsonString string){
	if *printJSON {
		t.Log(jsonString)
	}
    if *saveJSON != ""{
        err := writeJSONToFile(fname, jsonString)
        if err!= nil{
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
        //spew.Dump(nasMsg)
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

func writeJSONToFile(fname string, jsonString string) error{
    mode := int(0777)
    fileMode := os.FileMode(mode)
    
    _ , err := os.Stat(*saveJSON)
    if os.IsNotExist(err) {
        err := os.Mkdir(*saveJSON, os.FileMode(mode))
        if err != nil {
            return err
        }
    }
    fpath := filepath.Join(*saveJSON, fname )
    err = os.WriteFile(fpath,[]byte(jsonString),fileMode)
    if err != nil {
        return err
    }
    return nil
}



