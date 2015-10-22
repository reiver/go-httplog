package httplog


import (
	"fmt"
)


func (httpLogger *internalHttpLogger) Debug(v ...interface{}) {

	msg := fmt.Sprint(v...)
	httpLogger.logs = append(httpLogger.logs, msg)

	httpLogger.logger.Debug(v...)
}

func (httpLogger *internalHttpLogger) Debugf(format string, v ...interface{}) {

	msg := fmt.Sprintf(format, v...)
	httpLogger.logs = append(httpLogger.logs, msg)

	httpLogger.logger.Debugf(format, v...)
}

func (httpLogger *internalHttpLogger) Debugln(v ...interface{}) {

	msg := fmt.Sprintln(v...)
	httpLogger.logs = append(httpLogger.logs, msg)

	httpLogger.logger.Debugln(v...)
}

