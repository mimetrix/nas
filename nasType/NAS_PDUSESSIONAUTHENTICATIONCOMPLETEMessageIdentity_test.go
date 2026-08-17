package nasType_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mimetrix/nas/nasMessage"
	"github.com/mimetrix/nas/nasType"
)

func TestNasTypeNewPDUSESSIONAUTHENTICATIONCOMPLETEMessageIdentity(t *testing.T) {
	a := nasType.NewPDUSESSIONAUTHENTICATIONCOMPLETEMessageIdentity()
	assert.NotNil(t, a)
}

type nasTypePDUSESSIONAUTHENTICATIONCOMPLETEMessageIdentityMessageType struct {
	in  uint8
	out uint8
}

var nasTypePDUSESSIONAUTHENTICATIONCOMPLETEMessageIdentityMessageTypeTable = []nasTypePDUSESSIONAUTHENTICATIONCOMPLETEMessageIdentityMessageType{
	{nasMessage.MsgTypePDUSessionAuthenticationComplete, nasMessage.MsgTypePDUSessionAuthenticationComplete},
}

func TestNasTypeGetSetPDUSESSIONAUTHENTICATIONCOMPLETEMessageIdentityMessageType(t *testing.T) {
	a := nasType.NewPDUSESSIONAUTHENTICATIONCOMPLETEMessageIdentity()
	for _, table := range nasTypePDUSESSIONAUTHENTICATIONCOMPLETEMessageIdentityMessageTypeTable {
		a.SetMessageType(table.in)
		assert.Equal(t, table.out, a.GetMessageType())
	}
}
