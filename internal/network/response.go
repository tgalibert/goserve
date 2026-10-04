package network

import (
	"fmt"
	"goserve/internal/constant"
	"io"
	"strconv"
	"time"
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
	res.loadGenericHeaders()
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

func (res *Response) loadGenericHeaders() {
	if _, exists := res.Headers["Content-Type"]; !exists && len(res.Body) > 0 {
		res.Headers["Content-Type"] = string(constant.TEXT)
	}
	if _, exists := res.Headers["Connection"]; !exists {
		res.Headers["Connection"] = "close"
	}
	res.Headers["Date"] = time.Now().UTC().Format("Mon, 02 Jan 2006 15:04:05 GMT")
}

func (res *Response) SetJSON(body []byte) {
	res.SetBody(body)
	res.Headers["Content-Type"] = string(constant.JSON)
}

func (res *Response) SetHTML(body []byte) {
	res.SetBody(body)
	res.Headers["Content-Type"] = string(constant.HTML)
}

func (res *Response) Preflight() {
	res.StatusCode = constant.StatusNoContent
}

func (res *Response) AddCorsHeaders(origin string) {
	res.SetHeader("Access-Control-Allow-Origin", origin)
	res.SetHeader("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH, HEAD")
	res.SetHeader("Access-Control-Allow-Headers", "Content-Type, Authorization")
	res.SetHeader("Access-Control-Max-Age", "86400")
	if origin != "*" {
		res.SetHeader("Vary", "Origin")
	}
}


func (res *Response) WriteTo(w io.Writer) error {
	_, err := w.Write(res.ToBytes())
	return err
}