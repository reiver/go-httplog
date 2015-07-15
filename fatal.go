package httplog


import (
	"os"
)


func (httpLogger *internalHttpLogger) Fatal(v ...interface{}) {

	httpLogger.Print(v...)
	os.Exit(1)
}

func (httpLogger *internalHttpLogger) Fatalf(format string, v ...interface{}) {

	httpLogger.Printf(format, v...)
	os.Exit(1)
}

func (httpLogger *internalHttpLogger) Fatalln(v ...interface{}) {

	httpLogger.Println(v...)
	os.Exit(1)
}

