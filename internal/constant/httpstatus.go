package constant

type ResponseStatusCode int

const (
	// 1xx Informational
	StatusContinue           ResponseStatusCode = 100
	StatusSwitchingProtocols ResponseStatusCode = 101
	StatusProcessing         ResponseStatusCode = 102
	StatusEarlyHints         ResponseStatusCode = 103

	// 2xx Success
	StatusOK                   ResponseStatusCode = 200
	StatusCreated              ResponseStatusCode = 201
	StatusAccepted             ResponseStatusCode = 202
	StatusNonAuthoritativeInfo ResponseStatusCode = 203
	StatusNoContent            ResponseStatusCode = 204
	StatusResetContent         ResponseStatusCode = 205
	StatusPartialContent       ResponseStatusCode = 206
	StatusMultiStatus          ResponseStatusCode = 207
	StatusAlreadyReported      ResponseStatusCode = 208
	StatusIMUsed               ResponseStatusCode = 226

	// 3xx Redirection
	StatusMultipleChoices   ResponseStatusCode = 300
	StatusMovedPermanently  ResponseStatusCode = 301
	StatusFound             ResponseStatusCode = 302
	StatusSeeOther          ResponseStatusCode = 303
	StatusNotModified       ResponseStatusCode = 304
	StatusUseProxy          ResponseStatusCode = 305
	StatusTemporaryRedirect ResponseStatusCode = 307
	StatusPermanentRedirect ResponseStatusCode = 308

	// 4xx Client Error
	StatusBadRequest                  ResponseStatusCode = 400
	StatusUnauthorized                ResponseStatusCode = 401
	StatusPaymentRequired             ResponseStatusCode = 402
	StatusForbidden                   ResponseStatusCode = 403
	StatusNotFound                    ResponseStatusCode = 404
	StatusMethodNotAllowed            ResponseStatusCode = 405
	StatusNotAcceptable               ResponseStatusCode = 406
	StatusProxyAuthRequired           ResponseStatusCode = 407
	StatusRequestTimeout              ResponseStatusCode = 408
	StatusConflict                    ResponseStatusCode = 409
	StatusGone                        ResponseStatusCode = 410
	StatusLengthRequired              ResponseStatusCode = 411
	StatusPreconditionFailed          ResponseStatusCode = 412
	StatusPayloadTooLarge             ResponseStatusCode = 413
	StatusURITooLong                  ResponseStatusCode = 414
	StatusUnsupportedMediaType        ResponseStatusCode = 415
	StatusRangeNotSatisfiable         ResponseStatusCode = 416
	StatusExpectationFailed           ResponseStatusCode = 417
	StatusTeapot                      ResponseStatusCode = 418
	StatusMisdirectedRequest          ResponseStatusCode = 421
	StatusUnprocessableEntity         ResponseStatusCode = 422
	StatusLocked                      ResponseStatusCode = 423
	StatusFailedDependency            ResponseStatusCode = 424
	StatusTooEarly                    ResponseStatusCode = 425
	StatusUpgradeRequired             ResponseStatusCode = 426
	StatusPreconditionRequired        ResponseStatusCode = 428
	StatusTooManyRequests             ResponseStatusCode = 429
	StatusRequestHeaderFieldsTooLarge ResponseStatusCode = 431
	StatusUnavailableForLegalReasons  ResponseStatusCode = 451

	// 5xx Server Error
	StatusInternalServerError           ResponseStatusCode = 500
	StatusNotImplemented                ResponseStatusCode = 501
	StatusBadGateway                    ResponseStatusCode = 502
	StatusServiceUnavailable            ResponseStatusCode = 503
	StatusGatewayTimeout                ResponseStatusCode = 504
	StatusHTTPVersionNotSupported       ResponseStatusCode = 505
	StatusVariantAlsoNegotiates         ResponseStatusCode = 506
	StatusInsufficientStorage           ResponseStatusCode = 507
	StatusLoopDetected                  ResponseStatusCode = 508
	StatusNotExtended                   ResponseStatusCode = 510
	StatusNetworkAuthenticationRequired ResponseStatusCode = 511
)

func (s ResponseStatusCode) StatusText() string {
	switch s {
	case StatusContinue:
		return "Continue"
	case StatusSwitchingProtocols:
		return "Switching Protocols"
	case StatusProcessing:
		return "Processing"
	case StatusEarlyHints:
		return "Early Hints"
	case StatusOK:
		return "OK"
	case StatusCreated:
		return "Created"
	case StatusAccepted:
		return "Accepted"
	case StatusNonAuthoritativeInfo:
		return "Non-Authoritative Information"
	case StatusNoContent:
		return "No Content"
	case StatusResetContent:
		return "Reset Content"
	case StatusPartialContent:
		return "Partial Content"
	case StatusMultiStatus:
		return "Multi-Status"
	case StatusAlreadyReported:
		return "Already Reported"
	case StatusIMUsed:
		return "IM Used"
	case StatusMultipleChoices:
		return "Multiple Choices"
	case StatusMovedPermanently:
		return "Moved Permanently"
	case StatusFound:
		return "Found"
	case StatusSeeOther:
		return "See Other"
	case StatusNotModified:
		return "Not Modified"
	case StatusUseProxy:
		return "Use Proxy"
	case StatusTemporaryRedirect:
		return "Temporary Redirect"
	case StatusPermanentRedirect:
		return "Permanent Redirect"
	case StatusBadRequest:
		return "Bad Request"
	case StatusUnauthorized:
		return "Unauthorized"
	case StatusPaymentRequired:
		return "Payment Required"
	case StatusForbidden:
		return "Forbidden"
	case StatusNotFound:
		return "Not Found"
	case StatusMethodNotAllowed:
		return "Method Not Allowed"
	case StatusNotAcceptable:
		return "Not Acceptable"
	case StatusProxyAuthRequired:
		return "Proxy Authentication Required"
	case StatusRequestTimeout:
		return "Request Timeout"
	case StatusConflict:
		return "Conflict"
	case StatusGone:
		return "Gone"
	case StatusLengthRequired:
		return "Length Required"
	case StatusPreconditionFailed:
		return "Precondition Failed"
	case StatusPayloadTooLarge:
		return "Payload Too Large"
	case StatusURITooLong:
		return "URI Too Long"
	case StatusUnsupportedMediaType:
		return "Unsupported Media Type"
	case StatusRangeNotSatisfiable:
		return "Range Not Satisfiable"
	case StatusExpectationFailed:
		return "Expectation Failed"
	case StatusTeapot:
		return "I'm a teapot"
	case StatusMisdirectedRequest:
		return "Misdirected Request"
	case StatusUnprocessableEntity:
		return "Unprocessable Entity"
	case StatusLocked:
		return "Locked"
	case StatusFailedDependency:
		return "Failed Dependency"
	case StatusTooEarly:
		return "Too Early"
	case StatusUpgradeRequired:
		return "Upgrade Required"
	case StatusPreconditionRequired:
		return "Precondition Required"
	case StatusTooManyRequests:
		return "Too Many Requests"
	case StatusRequestHeaderFieldsTooLarge:
		return "Request Header Fields Too Large"
	case StatusUnavailableForLegalReasons:
		return "Unavailable For Legal Reasons"
	case StatusInternalServerError:
		return "Internal Server Error"
	case StatusNotImplemented:
		return "Not Implemented"
	case StatusBadGateway:
		return "Bad Gateway"
	case StatusServiceUnavailable:
		return "Service Unavailable"
	case StatusGatewayTimeout:
		return "Gateway Timeout"
	case StatusHTTPVersionNotSupported:
		return "HTTP Version Not Supported"
	case StatusVariantAlsoNegotiates:
		return "Variant Also Negotiates"
	case StatusInsufficientStorage:
		return "Insufficient Storage"
	case StatusLoopDetected:
		return "Loop Detected"
	case StatusNotExtended:
		return "Not Extended"
	case StatusNetworkAuthenticationRequired:
		return "Network Authentication Required"
	default:
		return ""
	}
}