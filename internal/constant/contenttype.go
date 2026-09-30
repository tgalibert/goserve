package constant

type ContentType string

const (
	JSON  ContentType = "application/json"
	TEXT  ContentType = "text/plain; charset=utf-8"
	HTML  ContentType = "text/html; charset=utf-8"
	OCTET ContentType = "application/octet-stream"
)