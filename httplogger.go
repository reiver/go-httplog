package httplog


import (
	"net/http"
)


// HttpLogger is the interface that represents what an 'http logger' looks like.
//
// The New func returns an HttpLogger interface (rather than the underlying
// struct).
//
// The HttpLogger interface includes 9 methods that are exactly the same as Go's
// built-in "log" library. Namely: Fatal, Fatalf, Fatalln, Panic, Panicf, Panicln,
// Print, Printf and Println.
//
// The HttpLogger interface also includes 25 methods (that the Go's built-in log
// library does NOT have) that makes it easier to return an 4xx or 5xx HTTP response.
//
// For example, a "404 Not Found" corresponds to the NotFound method, a "500 Internal
// Server Error" corresponds to the InternalServerError method, and a "408 Request
// Timeout" corresponds to the RequestTimeout. Etc.
type HttpLogger interface {
	Fatal(...interface{})
	Fatalf(string, ...interface{})
	Fatalln(...interface{})

	Panic(...interface{})
	Panicf(string, ...interface{})
	Panicln(...interface{})

	Print(...interface{})
	Printf(string, ...interface{})
	Println(...interface{})

	BadRequest(http.ResponseWriter)
	Unauthorized(http.ResponseWriter)
	PaymentRequired(http.ResponseWriter)
	Forbidden(http.ResponseWriter)
	NotFound(http.ResponseWriter)
	MethodNotAllowed(http.ResponseWriter)
	NotAcceptable(http.ResponseWriter)
	ProxyAuthRequired(http.ResponseWriter)
	RequestTimeout(http.ResponseWriter)
	Conflict(http.ResponseWriter)
	Gone(http.ResponseWriter)
	LengthRequired(http.ResponseWriter)
	PreconditionFailed(http.ResponseWriter)
	RequestEntityTooLarge(http.ResponseWriter)
	RequestURITooLong(http.ResponseWriter)
	UnsupportedMediaType(http.ResponseWriter)
	RequestedRangeNotSatisfiable(http.ResponseWriter)
	ExpectationFailed(http.ResponseWriter)
	Teapot(http.ResponseWriter)

	InternalServerError(http.ResponseWriter)
	NotImplemented(http.ResponseWriter)
	BadGateway(http.ResponseWriter)
	ServiceUnavailable(http.ResponseWriter)
	GatewayTimeout(http.ResponseWriter)
	HTTPVersionNotSupported(http.ResponseWriter)
}
