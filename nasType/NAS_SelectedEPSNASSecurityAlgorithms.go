package nasType

// SelectedEPSNASSecurityAlgorithms 9.11.3.25
// TypeOfCipheringAlgorithm Row, sBit, len = [0, 0], 7 , 3
// TypeOfIntegrityProtectionAlgorithm Row, sBit, len = [0, 0], 3 , 3
type SelectedEPSNASSecurityAlgorithms struct {
	Iei   uint8 `json:"Iei,omitempty"`
	Octet uint8 `json:"Octet,omitempty"`
}

func NewSelectedEPSNASSecurityAlgorithms(iei uint8) (selectedEPSNASSecurityAlgorithms *SelectedEPSNASSecurityAlgorithms) {
	selectedEPSNASSecurityAlgorithms = &SelectedEPSNASSecurityAlgorithms{}
	selectedEPSNASSecurityAlgorithms.SetIei(iei)
	return selectedEPSNASSecurityAlgorithms
}

// SelectedEPSNASSecurityAlgorithms 9.11.3.25
// Iei Row, sBit, len = [], 8, 8
func (a *SelectedEPSNASSecurityAlgorithms) GetIei() (iei uint8) {
	return a.Iei
}

// SelectedEPSNASSecurityAlgorithms 9.11.3.25
// Iei Row, sBit, len = [], 8, 8
func (a *SelectedEPSNASSecurityAlgorithms) SetIei(iei uint8) {
	a.Iei = iei
}

// SelectedEPSNASSecurityAlgorithms 9.11.3.25
// TypeOfCipheringAlgorithm Row, sBit, len = [0, 0], 7 , 3
func (a *SelectedEPSNASSecurityAlgorithms) GetTypeOfCipheringAlgorithm() (typeOfCipheringAlgorithm uint8) {
	return a.Octet & GetBitMask(7, 4) >> (4)
}

// SelectedEPSNASSecurityAlgorithms 9.11.3.25
// TypeOfCipheringAlgorithm Row, sBit, len = [0, 0], 7 , 3
func (a *SelectedEPSNASSecurityAlgorithms) SetTypeOfCipheringAlgorithm(typeOfCipheringAlgorithm uint8) {
	a.Octet = (a.Octet & 143) + ((typeOfCipheringAlgorithm & 7) << 4)
}

// SelectedEPSNASSecurityAlgorithms 9.11.3.25
// TypeOfIntegrityProtectionAlgorithm Row, sBit, len = [0, 0], 3 , 3
func (a *SelectedEPSNASSecurityAlgorithms) GetTypeOfIntegrityProtectionAlgorithm() (typeOfIntegrityProtectionAlgorithm uint8) {
	return a.Octet & GetBitMask(3, 0)
}

// SelectedEPSNASSecurityAlgorithms 9.11.3.25
// TypeOfIntegrityProtectionAlgorithm Row, sBit, len = [0, 0], 3 , 3
func (a *SelectedEPSNASSecurityAlgorithms) SetTypeOfIntegrityProtectionAlgorithm(typeOfIntegrityProtectionAlgorithm uint8) {
	a.Octet = (a.Octet & 248) + (typeOfIntegrityProtectionAlgorithm & 7)
}
