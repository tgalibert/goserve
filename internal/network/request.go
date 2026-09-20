package network

import (
	"fmt"
	"io"
)

type Request struct {
	Raw  io.Reader
	Json string
}

func NewRequest(raw io.Reader) *Request {
	return &Request{
		Raw: raw,
	}
}

func (r *Request) Parse() {
	b := make([]byte, 8)
	var full_request = ""
	for {
		n, err :=r.Raw.Read(b)
		full_request = fmt.Sprintf("%s%s", full_request, string(b[:n]))
		if err == io.EOF {
			break
		}
	}
	fmt.Println(full_request)
}