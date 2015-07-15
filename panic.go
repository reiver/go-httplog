package httplog


import (
	"fmt"
)


func (httpLogger *internalHttpLogger) Panic(v ...interface{}) {

	httpLogger.Print(v...)
	panic(fmt.Sprint(v...))
}

func (httpLogger *internalHttpLogger) Panicf(format string, v ...interface{}) {

	httpLogger.Printf(format, v...)
	panic(fmt.Sprintf(format, v...))
}

func (httpLogger *internalHttpLogger) Panicln(v ...interface{}) {

	httpLogger.Println(v...)
	panic(fmt.Sprintln(v...))
}

