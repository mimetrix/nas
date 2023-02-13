package nasType

// Capability5GMM 9.11.3.1
// LPP  Row, sBit, len = [0, 0], 3 , 1
// HOAttach Row, sBit, len = [0, 0], 2 , 1
// S1Mode Row, sBit, len = [0, 0], 1 , 1
// Spare Row, sBit, len = [1, 12], 8 , 96
type Capability5GMM struct {
	Iei   uint8     `json:"-"`
	Len   uint8     `json:"-"`
	Octet [13]uint8 `json:"-"`
    S1Mode bool
    HOAttach bool
    LTEPositioningProtocol bool
    RestrictEC bool
    _5G_CP_IoT bool
    N3Data bool
    _5G_IPHC_CP bool
    ServingGapControl bool

    _5GSRVCC bool
    _5G_UP_CIoT bool
    V2X bool
    V2XCEPC5 bool
    V2XCNPC5 bool
    _5G_LCS bool
    NSSAA bool
    RACS bool
    
    CAG bool
    WUSA bool
    MultipleUP bool
    _5G_EHC_CP_CIoT bool
    ER_NSSAI bool
    _5G_ProSe_dd bool
    _5G_ProSe_dc bool
    _5G_ProSe_l2relay bool

    _5G_ProSe_l3relay bool
    _5G_ProSe_l2rmt bool
    _5G_ProSe_l3rmt bool
    NR_PSSI bool
    NCR bool
    PIV bool
    RPR bool
    PagingRestriction bool

    NSSRG bool
    MINT bool
    EventNotification bool
    SSNPNSI bool
    Ex_CAG bool
    NSAG bool
}

func (c *Capability5GMM ) DecodeNASType() error{

    c.S1Mode = c.Octet[0] & 0x1 == 1
    c.HOAttach = (c.Octet[0] & 0x2) >> 1 == 1
    c.LTEPositioningProtocol = (c.Octet[0] & 0x4) >> 2 == 1
    c.RestrictEC = (c.Octet[0] & 0x8) >> 3 == 1
    c._5G_CP_IoT = (c.Octet[0] & 0x10) >> 4 == 1
    c.N3Data = (c.Octet[0] & 0x20) >> 5 == 1
    c._5G_IPHC_CP = (c.Octet[0] & 0x40) >> 6 == 1
    c.ServingGapControl = (c.Octet[0] & 0x80) >> 7 == 1


    c._5GSRVCC = (c.Octet[1] & 0x1)   == 1
    c._5G_UP_CIoT = (c.Octet[1] & 0x2) >> 1  == 1
    c.V2X = (c.Octet[1] & 0x4) >>  2 == 1
    c.V2XCEPC5 = (c.Octet[1] & 0x8) >> 3  == 1
    c.V2XCNPC5 = (c.Octet[1] & 0x10) >> 4 == 1
    c._5G_LCS = (c.Octet[1] & 0x20) >> 5 == 1
    c.NSSAA = (c.Octet[1] & 0x40) >> 6 == 1
    c.RACS = (c.Octet[1] & 0x80) >> 7 == 1


    c.CAG = (c.Octet[2] & 0x1)   == 1
    c.WUSA = (c.Octet[2] & 0x2) >> 1 == 1
    c.MultipleUP = (c.Octet[2] & 0x4) >> 2 == 1
    c._5G_EHC_CP_CIoT = (c.Octet[2] & 0x8) >> 3 == 1
    c.ER_NSSAI = (c.Octet[2] & 0x10) >> 4 == 1
    c._5G_ProSe_dd = (c.Octet[2] & 0x20) >> 5 == 1
    c._5G_ProSe_dc = (c.Octet[2] & 0x40) >> 6 == 1
    c._5G_ProSe_l2relay = (c.Octet[2] & 0x80) >> 7 == 1


    c._5G_ProSe_l3relay = (c.Octet[3] & 0x1)  == 1
    c._5G_ProSe_l2rmt = (c.Octet[3] & 0x2) >> 1 == 1
    c._5G_ProSe_l3rmt = (c.Octet[3] & 0x4) >> 2 == 1
    c.NR_PSSI = (c.Octet[3] & 0x8) >> 3 == 1
    c.NCR = (c.Octet[3] & 0x10) >> 4 == 1
    c.PIV = (c.Octet[3] & 0x20) >> 5 == 1
    c.RPR = (c.Octet[3] & 0x40) >> 6 == 1
    c.PagingRestriction = (c.Octet[3] & 0x80) >> 7 == 1

    c.NSSRG = (c.Octet[4] & 0x1)   == 1
    c.MINT = (c.Octet[4] & 0x2) >> 1 == 1
    c.EventNotification = (c.Octet[4] & 0x4) >> 2 == 1
    c.SSNPNSI = (c.Octet[4] & 0x8) >> 3 == 1
    c.Ex_CAG = (c.Octet[4] & 0x10) >> 4 == 1
    c.NSAG = (c.Octet[4] & 0x20) >> 5 == 1

    return nil
}

func NewCapability5GMM(iei uint8) (capability5GMM *Capability5GMM) {
	capability5GMM = &Capability5GMM{}
	capability5GMM.SetIei(iei)
	return capability5GMM
}

// Capability5GMM 9.11.3.1
// Iei Row, sBit, len = [], 8, 8
func (a *Capability5GMM) GetIei() (iei uint8) {
	return a.Iei
}

// Capability5GMM 9.11.3.1
// Iei Row, sBit, len = [], 8, 8
func (a *Capability5GMM) SetIei(iei uint8) {
	a.Iei = iei
}

// Capability5GMM 9.11.3.1
// Len Row, sBit, len = [], 8, 8
func (a *Capability5GMM) GetLen() (len uint8) {
	return a.Len
}

// Capability5GMM 9.11.3.1
// Len Row, sBit, len = [], 8, 8
func (a *Capability5GMM) SetLen(len uint8) {
	a.Len = len
}

// Capability5GMM 9.11.3.1
// LPP Row, sBit, len = [0, 0], 3 , 1
func (a *Capability5GMM) GetLPP() (lPP uint8) {
	return a.Octet[0] & GetBitMask(3, 2) >> (2)
}

// Capability5GMM 9.11.3.1
// LPP Row, sBit, len = [0, 0], 3 , 1
func (a *Capability5GMM) SetLPP(lPP uint8) {
	a.Octet[0] = (a.Octet[0] & 251) + ((lPP & 1) << 2)
}

// Capability5GMM 9.11.3.1
// HOAttach Row, sBit, len = [0, 0], 2 , 1
func (a *Capability5GMM) GetHOAttach() (hOAttach uint8) {
	return a.Octet[0] & GetBitMask(2, 1) >> (1)
}

// Capability5GMM 9.11.3.1
// HOAttach Row, sBit, len = [0, 0], 2 , 1
func (a *Capability5GMM) SetHOAttach(hOAttach uint8) {
	a.Octet[0] = (a.Octet[0] & 253) + ((hOAttach & 1) << 1)
}

// Capability5GMM 9.11.3.1
// S1Mode Row, sBit, len = [0, 0], 1 , 1
func (a *Capability5GMM) GetS1Mode() (s1Mode uint8) {
	return a.Octet[0] & GetBitMask(1, 0)
}

// Capability5GMM 9.11.3.1
// S1Mode Row, sBit, len = [0, 0], 1 , 1
func (a *Capability5GMM) SetS1Mode(s1Mode uint8) {
	a.Octet[0] = (a.Octet[0] & 254) + (s1Mode & 1)
}

// Capability5GMM 9.11.3.1
// Spare Row, sBit, len = [1, 12], 8 , 96
func (a *Capability5GMM) GetSpare() (spare [12]uint8) {
	copy(spare[:], a.Octet[1:13])
	return spare
}

// Capability5GMM 9.11.3.1
// Spare Row, sBit, len = [1, 12], 8 , 96
func (a *Capability5GMM) SetSpare(spare [12]uint8) {
	copy(a.Octet[1:13], spare[:])
}
