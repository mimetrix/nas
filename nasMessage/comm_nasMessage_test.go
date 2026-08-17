package nasMessage_test

// Table-driven helper types for the IE tests in this package.
//
// PayloadContainer and NASMessageContainer are declared in nasMessage rather
// than nasType (they recursively decode a nested NAS Message, which would be a
// circular import from nasType). Their tests moved with them, but these
// helpers live in nasType's test package and are not reachable from here, so
// they are mirrored below.

type NasTypeIeiData struct {
	in  uint8
	out uint8
}

type NasTypeLenuint8Data struct {
	in  uint8
	out uint8
}

type NasTypeLenUint16Data struct {
	in  uint16
	out uint16
}
