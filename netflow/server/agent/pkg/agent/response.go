package agent

type Response struct {
	Resp   interface{} `json:"response"`
	Errors []string    `json:"errors"`
}

func NewResponse() *Response {
	return &Response{
		Errors: []string{},
	}
}