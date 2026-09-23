package network

import (
	"fmt"
	"goserve/internal/constant"
	"io"
	"strconv"
)

const PROTOCOL = "HTTP/1.1"

type Response struct {
	StatusCode constant.ResponseStatusCode
	Headers    map[string]string
	Body       []byte
}

func NewResponse() *Response {
	return &Response{
		Headers: make(map[string]string),
		StatusCode: constant.StatusOK,
	}
}

func (res *Response) SetHeader(key, value string) {
	res.Headers[key] = value
}

func (res *Response) SetBody(body []byte) {
	res.Body = body
	res.Headers["Content-Length"] = strconv.Itoa(len(res.Body))
}

func (res *Response) ToBytes() []byte {
	tmp := make([]byte, 0)
	statusLine := fmt.Sprintf("%s %d %s\r\n", PROTOCOL, res.StatusCode, res.StatusCode.StatusText())
	tmp = append(tmp, []byte(statusLine)...)
	for key, value := range res.Headers {
		tmp = append(tmp, []byte(fmt.Sprintf("%s: %s\r\n", key, value))...)
	}
	tmp = append(tmp, []byte("\r\n")...)
	if len(res.Body) > 0 {
		tmp = append(tmp, res.Body...)
	}
	return tmp
}

func (res *Response) WriteTo(w io.Writer) error {
	_, err := w.Write(res.ToBytes())
	return err
}