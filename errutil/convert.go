package errutil

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/save95/xerror"
)

func Kratos(err error) *errors.Error {
	if kerr, ok := err.(*errors.Error); ok {
		return kerr
	}

	if xe, ok := err.(xerror.XError); ok {
		return errors.New(xe.HttpStatus(), fmt.Sprintf("XCODE_%d", xe.ErrorCode()), xe.Error())
	}

	return errors.New(http.StatusInternalServerError, err.Error(), err.Error())
}

func XError(err error) xerror.XError {
	if xe, ok := err.(xerror.XError); ok {
		return xe
	}

	if kerr, ok := err.(*errors.Error); ok {
		xec, err := strconv.Atoi(strings.ReplaceAll(kerr.Reason, "XCODE_", ""))
		if nil == err {
			msg := kerr.Message

			if bs, err := json.Marshal(kerr.GetMetadata()); len(kerr.Metadata) > 0 && err == nil {
				msg += ", metadata: " + string(bs)
			}

			return xerror.WithCode(xec, msg)
		}
	}

	return xerror.Wrap(err, err.Error())
}
