package httplog


import (
	"fmt"
)


func (httpLogger *internalHttpLogger) Print(v ...interface{}) {

	msg := fmt.Sprint(v...)
	httpLogger.logs = append(httpLogger.logs, msg)

	httpLogger.logger.Print(v...)
}

func (httpLogger *internalHttpLogger) Printf(format string, v ...interface{}) {

	msg := fmt.Sprintf(format, v...)
	httpLogger.logs = append(httpLogger.logs, msg)

	httpLogger.logger.Printf(format, v...)
}

func (httpLogger *internalHttpLogger) Println(v ...interface{}) {

	msg := fmt.Sprintln(v...)
	httpLogger.logs = append(httpLogger.logs, msg)

	httpLogger.logger.Println(v...)
}

