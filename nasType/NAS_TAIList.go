package nasType

import(
    "bytes"
    "fmt"
	"github.com/davecgh/go-spew/spew"
	//"encoding/binary"
)

// TAIList 9.11.3.9
// PartialTrackingAreaIdentityList Row, sBit, len = [0, 0], 8 , INF
type TAIList struct {
	Iei    uint8   `json:"Iei,omitempty"`
	Len    uint8   `json:"Len,omitempty"`
	Buffer []uint8 `json:"Buffer,omitempty"`
}

type TAIType00 struct  {
    Type uint8
    NumElements uint8
    MCC string
    MNC string
    //LIst of TACs
}



type TAIType01 struct  {
    Type uint8
    NumElements uint8
    MCC string
    MNC string
    TAC [3]uint8
}

type TAIType10 struct  {
    Type uint8
    NumElements uint8
    //List of TAIType01
}

type TAIType interface{
    GetTAIType()
    GetNumberOfTAIElems()
}

func (t *TAIList) DecodeNASType() error {

    fmt.Println("TAIList")
    payload := bytes.NewBuffer(t.Buffer)
    for payload.Len() > 1 {
        headerByte, err := payload.ReadByte()
        if err!= nil {
            fmt.Println(err)
            return err
        }
        TAIType := (headerByte & 0x60) >> 5
        numElements := (headerByte & 0x1F)
        fmt.Printf("hdr:0x%x    type:0x%x    elems:0x%x\n",headerByte, TAIType, numElements)

        if TAIType == 0x01 {
            TAI01 := & TAIType01{}
            type01Buf := make([]byte,6)
            _, err := payload.Read(type01Buf)
            if err != nil {
                fmt.Println(err)
                return(err)
            }
            MCC := fmt.Sprintf("%d%d%d", 
                type01Buf[0] & 0xf, (type01Buf[0] &0xf0)>>4, type01Buf[1] &0xf)
            MNC := fmt.Sprintf("%d%d",
                type01Buf[2] &0xf, (type01Buf[2]&0xf0)>>4, (type01Buf[1]&0xf0) > 4 )
            TAI01.MCC = MCC
            TAI01.MNC = MNC
            spew.Dump(TAI01)

        }

    }
    //get type of 
    /*
    TODO:
    create slice of PartialTAIs 
    TAIType00
    TAIType01
    TAIType02
    */
    return nil

}

func NewTAIList(iei uint8) (tAIList *TAIList) {
	tAIList = &TAIList{}
	tAIList.SetIei(iei)
	return tAIList
}

// TAIList 9.11.3.9
// Iei Row, sBit, len = [], 8, 8
func (a *TAIList) GetIei() (iei uint8) {
	return a.Iei
}

// TAIList 9.11.3.9
// Iei Row, sBit, len = [], 8, 8
func (a *TAIList) SetIei(iei uint8) {
	a.Iei = iei
}

// TAIList 9.11.3.9
// Len Row, sBit, len = [], 8, 8
func (a *TAIList) GetLen() (len uint8) {
	return a.Len
}

// TAIList 9.11.3.9
// Len Row, sBit, len = [], 8, 8
func (a *TAIList) SetLen(len uint8) {
	a.Len = len
	a.Buffer = make([]uint8, a.Len)
}

// TAIList 9.11.3.9
// PartialTrackingAreaIdentityList Row, sBit, len = [0, 0], 8 , INF
func (a *TAIList) GetPartialTrackingAreaIdentityList() (partialTrackingAreaIdentityList []uint8) {
	partialTrackingAreaIdentityList = make([]uint8, len(a.Buffer))
	copy(partialTrackingAreaIdentityList, a.Buffer)
	return partialTrackingAreaIdentityList
}

// TAIList 9.11.3.9
// PartialTrackingAreaIdentityList Row, sBit, len = [0, 0], 8 , INF
func (a *TAIList) SetPartialTrackingAreaIdentityList(partialTrackingAreaIdentityList []uint8) {
	copy(a.Buffer, partialTrackingAreaIdentityList)
}
