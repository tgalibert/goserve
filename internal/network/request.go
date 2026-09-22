package network

import (
	"errors"
	"goserve/internal/network/lib"
	"strconv"
	"strings"
)

type Request struct {
	parser *lib.IoParser
	Method	RequestMethod
	Path	string
	QueryParameters	map[string]string
	Headers	map[string]string
	ProtocolVersion	RequestProtocol
	Body	[]byte
}

type RequestMethod string
type RequestProtocol string

const (
	GET	RequestMethod = "GET"
	POST RequestMethod = "POST"
	PUT RequestMethod = "PUT"
	DELETE RequestMethod = "DELETE"
	OPTIONS RequestMethod = "OPTIONS"
	PATCH RequestMethod = "PATCH"
	UNKNOWN RequestMethod = ""
)

const (
	HTTP11 RequestProtocol = "HTTP/1.1"
	HTTP1 RequestProtocol = "HTTP/1.0"
)

func (p RequestProtocol) IsValid() bool {
	switch p {
		case HTTP1, HTTP11:
			return true
		default:
			return false
	}
}

func NewRequest(parser *lib.IoParser) *Request {
	return &Request{
		parser: parser,
		Headers: make(map[string]string),
	}
}

func (r *Request) Parse() error {
	err := r.retrieveRequestInformations()
	if err != nil {
		return err
	}
	err = r.parseHeaders()
	if err != nil {
		return err
	}
	err = r.parseBody()
	if err != nil {
		return err
	}
	return nil
}

func (r *Request) retrieveRequestInformations() error {
	first, err := r.parser.ReadLine("\r\n")
	if err != nil {
		return err
	}
	splitted := strings.SplitN(first, " ", -1)
	if len(splitted) < 3 {
		return errors.New("Bad request format.")
	}
	r.retrieveMethod(splitted)
	r.retrieveRoutePath(splitted)
	r.retrieveQueryParameters(splitted)
	errProto := r.retrieveProtocolVersion(splitted)
	if errProto != nil {
		return errProto
	}
	return nil
}

func (r *Request) retrieveMethod(splitted []string) {
	r.Method = RequestMethod(splitted[0])
}

func (r *Request) retrieveRoutePath(splitted []string) {
	path,_,_ := strings.Cut(splitted[1], "?")
	r.Path = path
}

func (r *Request) retrieveQueryParameters(splitted []string) {
	if !strings.Contains(splitted[1], "?") {
		r.QueryParameters = make(map[string]string)
		return
	}
	_, after, found := strings.Cut(splitted[1], "?")
	if !found {
		r.QueryParameters = make(map[string]string)
		return
	}
	parameters := strings.Split(after, "&")
	parametersSet := make(map[string]string)

	for _, parameter := range parameters {
		splitted := strings.Split(parameter, "=")
		if len(splitted) == 2 {
			parametersSet[splitted[0]] = splitted[1]
		}
	}
	r.QueryParameters = parametersSet
}

func (r *Request) retrieveProtocolVersion(splitted []string) error {
	proto := RequestProtocol(splitted[2])
	if !proto.IsValid() {
		return errors.New("Invalid protocol")
	}
	r.ProtocolVersion = proto
	return nil
}

func (r *Request) parseHeaders() error {
	for {
		line, err := r.parser.ReadLine("\r\n")
		if err != nil {
			return err
		}
		if line == "" {
			break
		}
		before, after, found := strings.Cut(line, ":")
		if found {
			r.Headers[strings.ToLower(strings.TrimSpace(before))] = strings.TrimSpace(after)
		} else {
			return errors.New("Malformed header")
		}
	}
	return nil
}

func (r *Request) parseBody() error {
	_, exist := r.Headers["content-length"]
	if !exist {
		return nil
	}
	length, err := strconv.Atoi(r.Headers["content-length"])
	if err != nil {
		return err
	}
	temp, err := r.parser.ReadBytes(length)
	if err != nil {
		return err
	}
	r.Body = temp
	return nil
}