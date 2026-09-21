package network

import (
	"strings"
)

type Request struct {
	RequestLines	[]string
	Method	RequestMethod
	path	string
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
	r.retrieveMethod()
	r.retrieveRoutePath()
}

func (r *Request) retrieveMethod() {
	if len(r.RequestLines)  < 1 {
		r.Method = UNKNOW
	}
	first := r.RequestLines[0]
	splitted := strings.SplitN(first, " ", -1)
	if len(splitted) < 1 {
		r.Method = UNKNOW
	}
	r.Method = RequestMethod(splitted[0])
}

func (r *Request) retrieveRoutePath() string {
	if len(r.RequestLines)  < 1 {
		r.path = ""
	}
	first := r.RequestLines[0]
	splitted := strings.SplitN(first, " ", -1)
	if len(splitted) < 1 {
		r.path = ""
	}
		r.path = splitted[1]
}