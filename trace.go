package httplog


import (
	"fmt"
)


func (httpLogger *internalHttpLogger) Trace(v ...interface{}) {

	msg := fmt.Sprint(v...)
	httpLogger.logs = append(httpLogger.logs, msg)

	httpLogger.logger.Trace(v...)
}

func (httpLogger *internalHttpLogger) Tracef(format string, v ...interface{}) {

	msg := fmt.Sprintf(format, v...)
	httpLogger.logs = append(httpLogger.logs, msg)

	httpLogger.logger.Tracef(format, v...)
}

func (httpLogger *internalHttpLogger) Traceln(v ...interface{}) {

	msg := fmt.Sprintln(v...)
	httpLogger.logs = append(httpLogger.logs, msg)

	httpLogger.logger.Traceln(v...)
}

