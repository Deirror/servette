package respx

type Resp struct {
	Code    string `json:"code"`
	MsgKey  string `json:"message_key"`
	Payload any    `json:"payload,omitempty"`
}

func New(code, key string, payload any) *Resp {
	return &Resp{
		Code:    code,
		MsgKey:  key,
		Payload: payload,
	}
}
