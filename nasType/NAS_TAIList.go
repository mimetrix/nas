package nasType


import(
    "bytes"
    "fmt"
	//"github.com/davecgh/go-spew/spew"
	//"encoding/binary"
)

// TAIList 9.11.3.9
// PartialTrackingAreaIdentityList Row, sBit, len = [0, 0], 8 , INF
type TAIList struct {
	Iei    uint8   `json:"-"`
	Len    uint8   `json:"-"`
	Buffer []uint8 `json:"-"`
    TAIs   []TAIType
}

type TAIType00 struct  {
    Type uint8
    NumElements uint8
    MCC string
    MNC string
    TACs []uint32//LIst of TACs
}



type TAIType01 struct  {
    Type uint8
    NumElements uint8
    MCC string
    MNC string
    TAC uint32
}

type TAIType10 struct  {
    Type uint8
    NumElements uint8
    TAIs []TAIType01
}

type TAIType interface{
    GetTAIType() uint8
    GetNumberOfTAIElems() uint8
}

func (t *TAIType01) GetTAIType() uint8{
    return t.Type
}


func (t *TAIType01) GetNumberOfTAIElems() uint8 {
    return t.NumElements
}

func NewTAI01(numElements uint8, buf [6]byte) *TAIType01{
     

    MCC1 := buf[0] & 0xf
    MCC2 := (buf[0] &0xf0) >> 4
    MCC3 := buf[1] &0xf
    MCC := fmt.Sprintf("%d%d%d", MCC1, MCC2, MCC3)
    
    MNC1 := buf[2] &0xf
    MNC2 := (buf[2]&0xf0) >> 4
    MNC3 := (buf[1]&0xf0) >> 4 

    var MNC string
    if MNC3 == 0xf{
        MNC = fmt.Sprintf("%d%d", MNC1, MNC2)
    } else {
        MNC = fmt.Sprintf("%d%d%d", MNC1, MNC2, MNC3)
    }

    TAC1 := buf[3]
    TAC2 := buf[4]
    TAC3 := buf[5]
    TAC := uint32(TAC1 << 16 | TAC2 << 8 | TAC3)

    return &TAIType01{1, numElements, MCC, MNC, TAC}
    

}

func (t *TAIList) DecodeNASType() error {
    
    t.TAIs = make([]TAIType,0)

    payload := bytes.NewBuffer(t.Buffer)
    for payload.Len() > 1 {
        headerByte, err := payload.ReadByte()
        if err!= nil {
            fmt.Println(err)
            return err
        }

        TAIType := (headerByte & 0x60) >> 5
        
        numElements := (headerByte & 0x1F) + 1

        if TAIType == 0x01 {

            var type01Buf [6]byte
            _, err := payload.Read(type01Buf[:])
            if err != nil {
                fmt.Println(err)
                return(err)
            }

            TAI01 := NewTAI01(numElements, type01Buf)
            t.TAIs = append(t.TAIs, TAI01)


        } else if TAIType == 0x00 {
            //TODO
        } else if TAIType == 0x10 {
            //TODO
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
