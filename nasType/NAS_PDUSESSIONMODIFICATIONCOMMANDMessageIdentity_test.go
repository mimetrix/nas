package nasType_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mimetrix/nas/nasMessage"
	"github.com/mimetrix/nas/nasType"
)

func TestNasTypeNewPDUSESSIONMODIFICATIONCOMMANDMessageIdentity(t *testing.T) {
	a := nasType.NewPDUSESSIONMODIFICATIONCOMMANDMessageIdentity()
	assert.NotNil(t, a)
}

type nasTypePDUSESSIONMODIFICATIONCOMMANDMessageIdentityMessageType struct {
	in  uint8
	out uint8
}

var nasTypePDUSESSIONMODIFICATIONCOMMANDMessageIdentityMessageTypeTable = []nasTypePDUSESSIONMODIFICATIONCOMMANDMessageIdentityMessageType{
	{nasMessage.MsgTypePDUSessionModificationCommand, nasMessage.MsgTypePDUSessionModificationCommand},
}

func TestNasTypeGetSetPDUSESSIONMODIFICATIONCOMMANDMessageIdentityMessageType(t *testing.T) {
	a := nasType.NewPDUSESSIONMODIFICATIONCOMMANDMessageIdentity()
	for _, table := range nasTypePDUSESSIONMODIFICATIONCOMMANDMessageIdentityMessageTypeTable {
		a.SetMessageType(table.in)
		assert.Equal(t, table.out, a.GetMessageType())
	}
}
