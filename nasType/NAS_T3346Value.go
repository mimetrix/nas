package nasType

// T3346Value 9.11.2.4
// GPRSTimer2Value Row, sBit, len = [0, 0], 8 , 8
type T3346Value struct {
	Iei   uint8 `json:"Iei,omitempty"`
	Len   uint8 `json:"Len,omitempty"`
	Octet uint8 `json:"Octet,omitempty"`
}

func NewT3346Value(iei uint8) (t3346Value *T3346Value) {
	t3346Value = &T3346Value{}
	t3346Value.SetIei(iei)
	return t3346Value
}

// T3346Value 9.11.2.4
// Iei Row, sBit, len = [], 8, 8
func (a *T3346Value) GetIei() (iei uint8) {
	return a.Iei
}

// T3346Value 9.11.2.4
// Iei Row, sBit, len = [], 8, 8
func (a *T3346Value) SetIei(iei uint8) {
	a.Iei = iei
}

// T3346Value 9.11.2.4
// Len Row, sBit, len = [], 8, 8
func (a *T3346Value) GetLen() (len uint8) {
	return a.Len
}

// T3346Value 9.11.2.4
// Len Row, sBit, len = [], 8, 8
func (a *T3346Value) SetLen(len uint8) {
	a.Len = len
}

// T3346Value 9.11.2.4
// GPRSTimer2Value Row, sBit, len = [0, 0], 8 , 8
func (a *T3346Value) GetGPRSTimer2Value() (gPRSTimer2Value uint8) {
	return a.Octet
}

// T3346Value 9.11.2.4
// GPRSTimer2Value Row, sBit, len = [0, 0], 8 , 8
func (a *T3346Value) SetGPRSTimer2Value(gPRSTimer2Value uint8) {
	a.Octet = gPRSTimer2Value
}
