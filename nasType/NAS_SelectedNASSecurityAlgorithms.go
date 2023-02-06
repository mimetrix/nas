package nasType

import(
    "fmt"
)

// SelectedNASSecurityAlgorithms 9.11.3.34
// TypeOfCipheringAlgorithm Row, sBit, len = [0, 0], 8 , 4
// TypeOfIntegrityProtectionAlgorithm Row, sBit, len = [0, 0], 4 , 4
type SelectedNASSecurityAlgorithms struct {
	Iei   uint8 `json:"-"`
	Octet uint8 `json:"-"`
    CipheringAlgorithm string
    IntegrityAlgorithm string
}

var IntegrityAlgorithms = map[uint8]string  {
    0:"5G integrity algorithm 5G-IA0 (null integrity protection algorithm)",
    1:"5G integrity algorithm 128-5G-IA1",
    2:"5G integrity algorithm 128-5G-IA2",
    3:"5G integrity algorithm 128-5G-IA3",
    4:"5G integrity algorithm 5G-IA4",
    5:"5G integrity algorithm 5G-IA5",
    6:"5G integrity algorithm 5G-IA6",
    7:"5G integrity algorithm 5G-IA7",
}

var CipheringAlgorithms = map[uint8]string  {
    0:"5G encryption algorithm 5G-EA0 (null ciphering algorithm)",
    1:"5G encryption algorithm 128-5G-EA1",
    2:"5G encryption algorithm 128-5G-EA2",
    3:"5G encryption algorithm 128-5G-EA3",
    4:"5G encryption algorithm 5G-EA4",
    5:"5G encryption algorithm 5G-EA5",
    6:"5G encryption algorithm 5G-EA6",
    7:"5G encryption algorithm 5G-EA7",
}


func (s *SelectedNASSecurityAlgorithms) DecodeNASType() error{
    cipherID := s.GetTypeOfCipheringAlgorithm()
    integID := s.GetTypeOfIntegrityProtectionAlgorithm()

    ca, ok := CipheringAlgorithms[cipherID]
    
    if !ok {
        return fmt.Errorf("\nUndefined Cipher Type: %x", cipherID)
    }
    s.CipheringAlgorithm = ca

    ia , ok:= IntegrityAlgorithms[integID]
    if !ok {
        return fmt.Errorf("\nUndefined Integrity Type: %x", cipherID)
    }
    s.IntegrityAlgorithm = ia
    
    return nil
    
}


func NewSelectedNASSecurityAlgorithms(iei uint8) (selectedNASSecurityAlgorithms *SelectedNASSecurityAlgorithms) {
	selectedNASSecurityAlgorithms = &SelectedNASSecurityAlgorithms{}
	selectedNASSecurityAlgorithms.SetIei(iei)
	return selectedNASSecurityAlgorithms
}

// SelectedNASSecurityAlgorithms 9.11.3.34
// Iei Row, sBit, len = [], 8, 8
func (a *SelectedNASSecurityAlgorithms) GetIei() (iei uint8) {
	return a.Iei
}

// SelectedNASSecurityAlgorithms 9.11.3.34
// Iei Row, sBit, len = [], 8, 8
func (a *SelectedNASSecurityAlgorithms) SetIei(iei uint8) {
	a.Iei = iei
}

// SelectedNASSecurityAlgorithms 9.11.3.34
// TypeOfCipheringAlgorithm Row, sBit, len = [0, 0], 8 , 4
func (a *SelectedNASSecurityAlgorithms) GetTypeOfCipheringAlgorithm() (typeOfCipheringAlgorithm uint8) {
	return a.Octet & GetBitMask(8, 4) >> (4)
}

// SelectedNASSecurityAlgorithms 9.11.3.34
// TypeOfCipheringAlgorithm Row, sBit, len = [0, 0], 8 , 4
func (a *SelectedNASSecurityAlgorithms) SetTypeOfCipheringAlgorithm(typeOfCipheringAlgorithm uint8) {
	a.Octet = (a.Octet & 15) + ((typeOfCipheringAlgorithm & 15) << 4)
}

// SelectedNASSecurityAlgorithms 9.11.3.34
// TypeOfIntegrityProtectionAlgorithm Row, sBit, len = [0, 0], 4 , 4
func (a *SelectedNASSecurityAlgorithms) GetTypeOfIntegrityProtectionAlgorithm() (typeOfIntegrityProtectionAlgorithm uint8) {
	return a.Octet & GetBitMask(4, 0)
}

// SelectedNASSecurityAlgorithms 9.11.3.34
// TypeOfIntegrityProtectionAlgorithm Row, sBit, len = [0, 0], 4 , 4
func (a *SelectedNASSecurityAlgorithms) SetTypeOfIntegrityProtectionAlgorithm(typeOfIntegrityProtectionAlgorithm uint8) {
	a.Octet = (a.Octet & 240) + (typeOfIntegrityProtectionAlgorithm & 15)
}
