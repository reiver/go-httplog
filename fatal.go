package httplog


import (
	"fmt"
)


func (httpLogger *internalHttpLogger) Fatal(v ...interface{}) {

	msg := fmt.Sprint(v...)
	httpLogger.logs = append(httpLogger.logs, msg)

	httpLogger.Fatal(v...)
}

func (httpLogger *internalHttpLogger) Fatalf(format string, v ...interface{}) {

	msg := fmt.Sprintf(format, v...)
	httpLogger.logs = append(httpLogger.logs, msg)

	httpLogger.Fatalf(format, v...)
}

func (httpLogger *internalHttpLogger) Fatalln(v ...interface{}) {

	msg := fmt.Sprintln(v...)
	httpLogger.logs = append(httpLogger.logs, msg)

	httpLogger.Fatalln(v...)
}
