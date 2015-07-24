
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
// The HttpLogger interface also includes 26 methods (that the Go's built-in log
// library does NOT have) that makes it easier to return an 2xx, 4xx or 5xx HTTP response.
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

	OK(http.ResponseWriter, ...interface{})

	BadRequest(http.ResponseWriter, ...interface{})
	Unauthorized(http.ResponseWriter, ...interface{})
	PaymentRequired(http.ResponseWriter, ...interface{})
	Forbidden(http.ResponseWriter, ...interface{})
	NotFound(http.ResponseWriter, ...interface{})
	MethodNotAllowed(http.ResponseWriter, ...interface{})
	NotAcceptable(http.ResponseWriter, ...interface{})
	ProxyAuthRequired(http.ResponseWriter, ...interface{})
	RequestTimeout(http.ResponseWriter, ...interface{})
	Conflict(http.ResponseWriter, ...interface{})
	Gone(http.ResponseWriter, ...interface{})
	LengthRequired(http.ResponseWriter, ...interface{})
	PreconditionFailed(http.ResponseWriter, ...interface{})
	RequestEntityTooLarge(http.ResponseWriter, ...interface{})
	RequestURITooLong(http.ResponseWriter, ...interface{})
	UnsupportedMediaType(http.ResponseWriter, ...interface{})
	RequestedRangeNotSatisfiable(http.ResponseWriter, ...interface{})
	ExpectationFailed(http.ResponseWriter, ...interface{})
	Teapot(http.ResponseWriter, ...interface{})

	InternalServerError(http.ResponseWriter, ...interface{})
	NotImplemented(http.ResponseWriter, ...interface{})
	BadGateway(http.ResponseWriter, ...interface{})
	ServiceUnavailable(http.ResponseWriter, ...interface{})
	GatewayTimeout(http.ResponseWriter, ...interface{})
	HTTPVersionNotSupported(http.ResponseWriter, ...interface{})
}
