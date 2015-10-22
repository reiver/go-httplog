package httplog


import (
	"fmt"
)


func (httpLogger *internalHttpLogger) Warn(v ...interface{}) {

	msg := fmt.Sprint(v...)
	httpLogger.logs = append(httpLogger.logs, msg)

	httpLogger.logger.Warn(v...)
}

func (httpLogger *internalHttpLogger) Warnf(format string, v ...interface{}) {

	msg := fmt.Sprintf(format, v...)
	httpLogger.logs = append(httpLogger.logs, msg)

	httpLogger.logger.Warnf(format, v...)
}

func (httpLogger *internalHttpLogger) Warnln(v ...interface{}) {

	msg := fmt.Sprintln(v...)
	httpLogger.logs = append(httpLogger.logs, msg)

	httpLogger.logger.Warnln(v...)
}

