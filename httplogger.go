package httplog


import (
	"net/http"
)


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
