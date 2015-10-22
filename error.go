package httplog


import (
	"fmt"
)


func (httpLogger *internalHttpLogger) Error(v ...interface{}) {

	msg := fmt.Sprint(v...)
	httpLogger.logs = append(httpLogger.logs, msg)

	httpLogger.logger.Error(v...)
}

func (httpLogger *internalHttpLogger) Errorf(format string, v ...interface{}) {

	msg := fmt.Sprintf(format, v...)
	httpLogger.logs = append(httpLogger.logs, msg)

	httpLogger.logger.Errorf(format, v...)
}

func (httpLogger *internalHttpLogger) Errorln(v ...interface{}) {

	msg := fmt.Sprintln(v...)
	httpLogger.logs = append(httpLogger.logs, msg)

	httpLogger.logger.Errorln(v...)
}

