package nasType_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mimetrix/nas/nasMessage"
	"github.com/mimetrix/nas/nasType"
)

func TestNasTypeNewPDUSESSIONESTABLISHMENTREJECTMessageIdentity(t *testing.T) {
	a := nasType.NewPDUSESSIONESTABLISHMENTREJECTMessageIdentity()
	assert.NotNil(t, a)
}

type nasTypePDUSESSIONESTABLISHMENTREJECTMessageIdentityMessageType struct {
	in  uint8
	out uint8
}

var nasTypePDUSESSIONESTABLISHMENTREJECTMessageIdentityMessageTypeTable = []nasTypePDUSESSIONESTABLISHMENTREJECTMessageIdentityMessageType{
	{nasMessage.MsgTypePDUSessionEstablishmentReject, nasMessage.MsgTypePDUSessionEstablishmentReject},
}

func TestNasTypeGetSetPDUSESSIONESTABLISHMENTREJECTMessageIdentityMessageType(t *testing.T) {
	a := nasType.NewPDUSESSIONESTABLISHMENTREJECTMessageIdentity()
	for _, table := range nasTypePDUSESSIONESTABLISHMENTREJECTMessageIdentityMessageTypeTable {
		a.SetMessageType(table.in)
		assert.Equal(t, table.out, a.GetMessageType())
	}
}
