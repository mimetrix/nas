package nasType

// RequestType 9.11.3.47
// Iei Row, sBit, len = [0, 0], 8 , 4
// RequestTypeValue Row, sBit, len = [0, 0], 3 , 3
type RequestType struct {
	Octet uint8 `json:"-"`
    Type string
}

var RequestTypes = map[uint8]string {
    1:"initial request",
    2:"existing PDU session",
    3:"initial emergency request",
    4:"existing emergency PDU session",
    5:"modification request",
    6:"MA PDU request",
    7:"reserved",
}

func (r *RequestType) DecodeNASType() error {

    r.Type = RequestTypes[r.GetRequestTypeValue()]
    return nil
}

func NewRequestType(iei uint8) (requestType *RequestType) {
	requestType = &RequestType{}
	requestType.SetIei(iei)
	return requestType
}

// RequestType 9.11.3.47
// Iei Row, sBit, len = [0, 0], 8 , 4
func (a *RequestType) GetIei() (iei uint8) {
	return a.Octet & GetBitMask(8, 4) >> (4)
}

// RequestType 9.11.3.47
// Iei Row, sBit, len = [0, 0], 8 , 4
func (a *RequestType) SetIei(iei uint8) {
	a.Octet = (a.Octet & 15) + ((iei & 15) << 4)
}

// RequestType 9.11.3.47
// RequestTypeValue Row, sBit, len = [0, 0], 3 , 3
func (a *RequestType) GetRequestTypeValue() (requestTypeValue uint8) {
	return a.Octet & GetBitMask(3, 0)
}

// RequestType 9.11.3.47
// RequestTypeValue Row, sBit, len = [0, 0], 3 , 3
func (a *RequestType) SetRequestTypeValue(requestTypeValue uint8) {
	a.Octet = (a.Octet & 248) + (requestTypeValue & 7)
}
