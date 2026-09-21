package network

import (
	"fmt"
	"strings"
)

type Request struct {
	RequestLines	[]string
	Method	RequestMethod
	Path	string
	QueryParameters	map[string]string
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
	r.retrieveQueryParameters()
	fmt.Printf("%s%s \n", r.Path, r.Method)
	fmt.Println(r.QueryParameters)
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

func (r *Request) retrieveRoutePath() {
	if len(r.RequestLines)  < 1 {
		r.Path = ""
	}
	first := r.RequestLines[0]
	splitted := strings.SplitN(first, " ", -1)
	if len(splitted) < 1 {
		r.Path = ""
	}
	r.Path = splitted[1]
}

func (r *Request) retrieveQueryParameters() {
	if r.Path == "" || !strings.Contains(r.Path, "?") {
		r.QueryParameters = make(map[string]string)
	}
	pstring := strings.Split(r.Path, "?")
	if len(pstring) < 1 {
		r.QueryParameters = make(map[string]string)
	}
	parameters := strings.Split(pstring[len(pstring) - 1], "&")
	parametersSet := make(map[string]string)

	for _, parameter := range parameters {
		splitted := strings.Split(parameter, "=")
		if len(splitted) == 2 {
			parametersSet[splitted[0]] = splitted[1]
		}
	}
	r.QueryParameters = parametersSet
}