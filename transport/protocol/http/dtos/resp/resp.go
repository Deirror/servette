package respx

import (
	"strconv"

	"github.com/Deirror/servette/transport/dtos/resp"
)

func New(status int, key string, payload any) *respx.Resp {
	return respx.New(strconv.Itoa(status), key, payload)
}
