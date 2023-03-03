package nasType

// Cause5GMM 9.11.3.2
// CauseValue Row, sBit, len = [0, 0], 8 , 8
type Cause5GMM struct {
	Iei   uint8 `json:"-"`
	Octet uint8 `json:"-"`
	Cause string
}

var GMMCauses = map[uint8]string{
	0x03: "Illegal UE",
	0x05: "PEI not accepted",
	0x06: "Illegal ME",
	0x07: "5GS services not allowed",
	0x09: "UE identity cannot be derived by the network",
	0x0a: "Implicitly de-registered",
	0x0b: "PLMN not allowed",
	0x0c: "Tracking area not allowed",
	0x0d: "Roaming not allowed in this tracking area",
	0x0f: "No suitable cells in tracking area",
	0x14: "MAC failure",
	0x15: "Synch failure",
	0x16: "Congestion",
	0x17: "UE security capabilities mismatch",
	0x18: "Security mode rejected, unspecified",
	0x1a: "Non-5G authentication unacceptable",
	0x1b: "N1 mode not allowed",
	0x1c: "Restricted service area",
	0x1f: "Redirection to EPC required",
	0x2b: "LADN not available",
	0x3e: "No network slices available",
	0x41: "Maximum number of PDU sessions reached",
	0x43: "Insufficient resources for specific slice and DNN",
	0x45: "Insufficient resources for specific slice",
	0x47: "ngKSI already in use",
	0x48: "Non-3GPP access to 5GCN not allowed",
	0x49: "Serving network not authorized",
	0x4a: "Temporarily not authorized for this SNPN",
	0x4b: "Permanently not authorized for this SNPN",
	0x4c: "Not authorized for this CAG or authorized for CAG cells only",
	0x4d: "Wireline access area not allowed",
	0x4e: "PLMN not allowed to operate at the present UE location",
	0x4f: "UAS services not allowed",
	0x5a: "Payload was not forwarded",
	0x50: "Disaster roaming for the determined PLMN with disaster condition not allowed",
	0x5b: "DNN not supported or not subscribed in the slice",
	0x5c: "Insufficient user-plane resources for the PDU session",
	0x5d: "Onboarding services terminated",
	0x5f: "Semantically incorrect message",
	0x60: "Invalid mandatory information",
	0x61: "Message type non-existent or not implemented",
	0x62: "Message type not compatible with the protocol state",
	0x63: "Information element non-existent or not implemented",
	0x64: "Conditional IE error",
	0x65: "Message not compatible with the protocol state",
	0x6f: "Protocol error, unspecified",
}

func (c *Cause5GMM) DecodeNASType() error {

	c.Cause = GMMCauses[c.Octet]
	return nil
}

func NewCause5GMM(iei uint8) (cause5GMM *Cause5GMM) {
	cause5GMM = &Cause5GMM{}
	cause5GMM.SetIei(iei)
	return cause5GMM
}

// Cause5GMM 9.11.3.2
// Iei Row, sBit, len = [], 8, 8
func (a *Cause5GMM) GetIei() (iei uint8) {
	return a.Iei
}

// Cause5GMM 9.11.3.2
// Iei Row, sBit, len = [], 8, 8
func (a *Cause5GMM) SetIei(iei uint8) {
	a.Iei = iei
}

// Cause5GMM 9.11.3.2
// CauseValue Row, sBit, len = [0, 0], 8 , 8
func (a *Cause5GMM) GetCauseValue() (causeValue uint8) {
	return a.Octet
}

// Cause5GMM 9.11.3.2
// CauseValue Row, sBit, len = [0, 0], 8 , 8
func (a *Cause5GMM) SetCauseValue(causeValue uint8) {
	a.Octet = causeValue
}
