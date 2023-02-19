package nasType

import (
    "bytes"
    "fmt"
)
// AuthorizedQosFlowDescriptions 9.11.4.12
// QoSFlowDescriptions Row, sBit, len = [0, 0], 8 , INF
type AuthorizedQosFlowDescriptions struct {
	Iei    uint8   `json:"-"`
	Len    uint16  `json:"-"`
	Buffer []uint8 `json:"-"`
    QoSDescriptions []QoSDescription
}

type QoSDescription struct{
    QoSFlowID       string
    OpID            uint8   `json"-"`
    Operation       string 
    E               bool
    NumParameters   uint8
    Parameters      []QoSParameter
}

type QoSParameter struct {
    ParameterIdentifier uint8 `json"-"`
    ParameterLength uint8
    ParameterType string
    Bytes []uint8
}

var ParameterTYpes = map[uint8]string {
    1:"5QI",
    2:"GFBR uplink",
    3:"GFBR downlink",
    4:"MFBR uplink",
    5:"MFBR downlink",
    6:"Averaging window and",
    7:"EPS bearer identity",
}

var OperationCodes = map[uint8]string {
    1:"Create new QoS flow description",
    2:"Delete existing QoS flow description",
    3:"Modify existing QoS flow description",
}

func (a *AuthorizedQosFlowDescriptions) DecodeNASType() error {
    
    QoSBuf := bytes.NewBuffer(a.Buffer)
    
    for QoSBuf.Len() > 0{
        //Create a new QoSDescriptionstruct
        QFI,err := QoSBuf.ReadByte()
        if err != nil{
            return err
        }
        QFI = QFI & 0x3f
        
        var QFIString string
        if QFI > 0 && QFI < 64 {
            QFIString = fmt.Sprintf("QFI %d", QFI)
        } else {
            QFIString = fmt.Sprintf("Invalid QFI Value, %d", QFI)
        }

        OpCode, err := QoSBuf.ReadByte()
        if err != nil{
            return err
        }
        OpCode = (OpCode & 0xe0) >> 5
        OpName := OperationCodes[OpCode]
        
        EAndNumParams, err := QoSBuf.ReadByte()
        if err != nil {
            return err
        }
        E := ((EAndNumParams & 0x40)>>6) == 1
        NumParams := EAndNumParams & 0x3f
        
        NewQosDescription := &QoSDescription{
            QoSFlowID: QFIString,
            OpID: OpCode,
            Operation: OpName,
            E: E,
            NumParameters: NumParams,
        }


        var paramCnt uint8
        for paramCnt =0; paramCnt < NumParams; paramCnt++{
            paramID, err := QoSBuf.ReadByte()
            if err != nil {
                return err
            }

            contentLen, err := QoSBuf.ReadByte()
            if err != nil {
                return err
            }

            paramBuf := make([]uint8, contentLen)
            _, err = QoSBuf.Read(paramBuf)
            if err != nil{
                return err
            }

            newQoSParameter := &QoSParameter{
                ParameterIdentifier: paramID,
                ParameterLength: contentLen,
                ParameterType: ParameterTYpes[paramID],
                Bytes: paramBuf,
            }

            NewQosDescription.Parameters = append(NewQosDescription.Parameters, *newQoSParameter)
        }
        
        a.QoSDescriptions = append(a.QoSDescriptions, *NewQosDescription)      


    }

    return nil


}

func NewAuthorizedQosFlowDescriptions(iei uint8) (authorizedQosFlowDescriptions *AuthorizedQosFlowDescriptions) {
	authorizedQosFlowDescriptions = &AuthorizedQosFlowDescriptions{}
	authorizedQosFlowDescriptions.SetIei(iei)
	return authorizedQosFlowDescriptions
}



// AuthorizedQosFlowDescriptions 9.11.4.12
// Iei Row, sBit, len = [], 8, 8
func (a *AuthorizedQosFlowDescriptions) GetIei() (iei uint8) {
	return a.Iei
}

// AuthorizedQosFlowDescriptions 9.11.4.12
// Iei Row, sBit, len = [], 8, 8
func (a *AuthorizedQosFlowDescriptions) SetIei(iei uint8) {
	a.Iei = iei
}

// AuthorizedQosFlowDescriptions 9.11.4.12
// Len Row, sBit, len = [], 8, 16
func (a *AuthorizedQosFlowDescriptions) GetLen() (len uint16) {
	return a.Len
}

// AuthorizedQosFlowDescriptions 9.11.4.12
// Len Row, sBit, len = [], 8, 16
func (a *AuthorizedQosFlowDescriptions) SetLen(len uint16) {
	a.Len = len
	a.Buffer = make([]uint8, a.Len)
}

// AuthorizedQosFlowDescriptions 9.11.4.12
// QoSFlowDescriptions Row, sBit, len = [0, 0], 8 , INF
func (a *AuthorizedQosFlowDescriptions) GetQoSFlowDescriptions() (qoSFlowDescriptions []uint8) {
	qoSFlowDescriptions = make([]uint8, len(a.Buffer))
	copy(qoSFlowDescriptions, a.Buffer)
	return qoSFlowDescriptions
}

// AuthorizedQosFlowDescriptions 9.11.4.12
// QoSFlowDescriptions Row, sBit, len = [0, 0], 8 , INF
func (a *AuthorizedQosFlowDescriptions) SetQoSFlowDescriptions(qoSFlowDescriptions []uint8) {
	copy(a.Buffer, qoSFlowDescriptions)
}
