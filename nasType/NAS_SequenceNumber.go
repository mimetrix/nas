package nasType

// SequenceNumber 9.10
// SQN Row, sBit, len = [0, 0], 8 , 8
type SequenceNumber struct {
	Octet uint8 `json:"Value"`
}

func NewSequenceNumber() (sequenceNumber *SequenceNumber) {
	sequenceNumber = &SequenceNumber{}
	return sequenceNumber
}

// SequenceNumber 9.10
// SQN Row, sBit, len = [0, 0], 8 , 8
func (a *SequenceNumber) GetSQN() (sQN uint8) {
	return a.Octet
}

// SequenceNumber 9.10
// SQN Row, sBit, len = [0, 0], 8 , 8
func (a *SequenceNumber) SetSQN(sQN uint8) {
	a.Octet = sQN
}
