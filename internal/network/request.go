package network

import (
	"fmt"
	"strings"
)

type Request struct {
	RequestLines	[]string
}

type RequestMethod string

const (
	GET	RequestMethod = "GET"
	POST RequestMethod = "POST"
	PUT RequestMethod = "PUT"
	DELETE RequestMethod = "DELETE"
	OPTION RequestMethod = "OPTION"
	PATCH RequestMethod = "PATCH"
	UNKNOW RequestMethod = ""
)

func NewRequest(requestLines []string) *Request {
	return &Request{
		RequestLines: requestLines,
	}
}

func (r *Request) Parse() {
	method := r.retrieveMethod()
	fmt.Println(method)
}

func (r *Request) retrieveMethod() RequestMethod {
	if len(r.RequestLines)  < 1 {
		return UNKNOW
	}
	first := r.RequestLines[0]
	splitted := strings.SplitN(first, " ", -1)
	if len(splitted) < 1 {
		return UNKNOW
	}
	return RequestMethod(splitted[0])
}