package httplog


import (
	"io"
)


type internalHttpLogger struct {
	logger   extendedLogger
	logs   []string
	dumpLogs bool
}


// New creates a new 'http logger'.
//
// New takes 2 parameters: 'writer' and 'dumpLogs',
//
// The 'writer' parameter tells New where the logs should
// be written to, as part of its normal logging operations.
//
// The 'dumpLogs' parameter tells New whether the 26 methods for
// generating 2xx, 4xx and 5xx HTTP responses should dump the (past) logs
// that the 'http logger' has seen.
//
// One would NOT want to dump these log in a 'production' deployment
// environment. And in that case 'dumpLogs' should be given a value
// of 'false'.
//
// However, in a 'staging' or 'development' deployment environment one would
// indeed want to dump these logs (to make debugging easier). And in that case
// 'dumpLogs' should be given a value of 'true'.
func New(writer io.Writer, dumpLogs bool) HttpLogger {

	logger := newWriterLogger(writer)

	httpLogger := Wrap(logger, dumpLogs)

	return httpLogger
}



func Wrap(logger Logger, dumpLogs bool) HttpLogger {
	logs := make([]string, 0, 16)

	extLogger, ok := logger.(extendedLogger)
	if !ok {
		extLogger = newAdaptor(logger)
	}

	httpLogger := internalHttpLogger{
		logger:   extLogger,
		logs:     logs,
		dumpLogs: dumpLogs,
	}

	return &httpLogger
}
