package nasType_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mimetrix/nas/nasMessage"
	"github.com/mimetrix/nas/nasType"
)

func TestNasTypeNewPDUSESSIONAUTHENTICATIONRESULTMessageIdentity(t *testing.T) {
	a := nasType.NewPDUSESSIONAUTHENTICATIONRESULTMessageIdentity()
	assert.NotNil(t, a)
}

type nasTypePDUSESSIONAUTHENTICATIONRESULTMessageIdentityMessageType struct {
	in  uint8
	out uint8
}

var nasTypePDUSESSIONAUTHENTICATIONRESULTMessageIdentityMessageTypeTable = []nasTypePDUSESSIONAUTHENTICATIONRESULTMessageIdentityMessageType{
	{nasMessage.MsgTypePDUSessionAuthenticationResult, nasMessage.MsgTypePDUSessionAuthenticationResult},
}

func TestNasTypeGetSetPDUSESSIONAUTHENTICATIONRESULTMessageIdentityMessageType(t *testing.T) {
	a := nasType.NewPDUSESSIONAUTHENTICATIONRESULTMessageIdentity()
	for _, table := range nasTypePDUSESSIONAUTHENTICATIONRESULTMessageIdentityMessageTypeTable {
		a.SetMessageType(table.in)
		assert.Equal(t, table.out, a.GetMessageType())
	}
}
