package nasType

import(
    //"bytes"
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
    //TAC
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

func (t *TAIList) DecodeNASType() {
    /*
    TODO:
    create slice of PartialTAIs 
    TAIType00
    TAIType01
    TAIType02
    */

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
