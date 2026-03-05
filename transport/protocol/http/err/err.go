package errx

import (
	"fmt"
	"strconv"

	"github.com/Deirror/servette/transport/err"
)

func New(status int, key, errMsg string, srvErr error) *errx.Err {
	srvErrMsg := errMsg
	if srvErr != nil {
		srvErrMsg = fmt.Sprintf("%s: %v", errMsg, srvErr)	
	}

	return errx.New(strconv.Itoa(status), key, srvErrMsg)
}
