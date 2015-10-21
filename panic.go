package httplog


import (
	"fmt"
)


func (httpLogger *internalHttpLogger) Panic(v ...interface{}) {

	msg := fmt.Sprint(v...)
	httpLogger.logs = append(httpLogger.logs, msg)

	httpLogger.Panic(v...)
}

func (httpLogger *internalHttpLogger) Panicf(format string, v ...interface{}) {

	msg := fmt.Sprintf(format, v...)
	httpLogger.logs = append(httpLogger.logs, msg)

	httpLogger.Panicf(format, v...)
}

func (httpLogger *internalHttpLogger) Panicln(v ...interface{}) {

	msg := fmt.Sprintln(v...)
	httpLogger.logs = append(httpLogger.logs, msg)

	httpLogger.Panicln(v...)
}
