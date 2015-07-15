package httplog


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

	BadRequest()
	Unauthorized()
	PaymentRequired()
	Forbidden()
	NotFound()
	MethodNotAllowed()
	NotAcceptable()
	ProxyAuthRequired()
	RequestTimeout()
	Conflict()
	Gone()
	LengthRequired()
	PreconditionFailed()
	RequestEntityTooLarge()
	RequestURITooLong()
	UnsupportedMediaType()
	RequestedRangeNotSatisfiable()
	ExpectationFailed()
	Teapot()

	InternalServerError()
	StatusNotImplemented()
	StatusBadGateway()
	ServiceUnavailable()
	GatewayTimeout()
	HTTPVersionNotSupported()
}
