package errutil

import (
	"testing"

	"github.com/save95/xerror"
	"github.com/stretchr/testify/assert"
)

func TestKratos_XError(t *testing.T) {
	xe := xerror.WithCode(1003, "to kratos test")
	ke := Kratos(xe)

	assert.Equal(t, ke.Reason, "XCODE_1003")
	assert.Equal(t, ke.Message, "to kratos test")

	xe2 := XError(ke)
	assert.True(t, xerror.IsErrorCode(xe2, 1003))
	assert.Equal(t, xe2.Error(), "to kratos test")
}
