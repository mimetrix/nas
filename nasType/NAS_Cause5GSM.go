package nasType

// Cause5GSM 9.11.4.2
// CauseValue Row, sBit, len = [0, 0], 8 , 8
type Cause5GSM struct {
	Iei   uint8 `json:"-"`
	Octet uint8 `json:"-"`
	Cause string
}

var GSMCauses = map[uint8]string{
	0x04: "Operator determined barring",
	0x1a: "Insufficient resources",
	0x1b: "Missing or unknown DNN",
	0x1c: "Unknown PDU session type",
	0x1d: "User authentication or authorization failed",
	0x1f: "Request rejected, unspecified",
	0x20: "Service option not supported",
	0x21: "Requested service option not subscribed",
	0x23: "PTI already in use",
	0x24: "Regular deactivation",
	0x25: "5GS QoS not accepted",
	0x26: "Network failure",
	0x27: "Reactivation requested",
	0x29: "Semantic error in the TFT operation",
	0x2a: "Syntactical error in the TFT operation",
	0x2b: "Invalid PDU session identity",
	0x2c: "Semantic errors in packet filter(s)",
	0x2d: "Syntactical error in packet filter(s)",
	0x2e: "Out of LADN service area",
	0x2f: "PTI mismatch",
	0x32: "PDU session type IPv4 only allowed",
	0x33: "PDU session type IPv6 only allowed",
	0x36: "PDU session does not exist",
	0x39: "PDU session type IPv4v6 only allowed",
	0x3a: "PDU session type Unstructured only allowed",
	0x3b: "Unsupported 5QI value",
	0x3d: "PDU session type Ethernet only allowed",
	0x43: "Insufficient resources for specific slice and DNN",
	0x44: "Not supported SSC mode",
	0x45: "Insufficient resources for specific slice",
	0x46: "Missing or unknown DNN in a slice",
	0x51: "Invalid PTI value",
	0x52: "Maximum data rate per UE for user-plane integrity protection is too low",
	0x53: "Semantic error in the QoS operation",
	0x54: "Syntactical error in the QoS operation",
	0x55: "Invalid mapped EPS bearer identity",
	0x56: "UAS services not allowed",
	0x5f: "Semantically incorrect message",
	0x60: "Invalid mandatory information",
	0x61: "Message type non-existent or not implemented",
	0x62: "Message type not compatible with the protocol state",
	0x63: "Information element non-existent or not implemented",
	0x64: "Conditional IE error",
	0x65: "Message not compatible with the protocol state",
	0x6f: "Protocol error, unspecified",
}

func (c *Cause5GSM) DecodeNASType() error {
	c.Cause = GSMCauses[c.Octet]

	return nil
}

func NewCause5GSM(iei uint8) (cause5GSM *Cause5GSM) {
	cause5GSM = &Cause5GSM{}
	cause5GSM.SetIei(iei)
	return cause5GSM
}

// Cause5GSM 9.11.4.2
// Iei Row, sBit, len = [], 8, 8
func (a *Cause5GSM) GetIei() (iei uint8) {
	return a.Iei
}

// Cause5GSM 9.11.4.2
// Iei Row, sBit, len = [], 8, 8
func (a *Cause5GSM) SetIei(iei uint8) {
	a.Iei = iei
}

// Cause5GSM 9.11.4.2
// CauseValue Row, sBit, len = [0, 0], 8 , 8
func (a *Cause5GSM) GetCauseValue() (causeValue uint8) {
	return a.Octet
}

// Cause5GSM 9.11.4.2
// CauseValue Row, sBit, len = [0, 0], 8 , 8
func (a *Cause5GSM) SetCauseValue(causeValue uint8) {
	a.Octet = causeValue
}
