package httplog


import (
	"fmt"
	"io"
	"os"
)


type internalWriterLogger struct {
	writer io.Writer
}


func newWriterLogger(writer io.Writer) extendedLogger {

	logger := internalWriterLogger{
		writer: writer,
	}

	return &logger
}



func (logger *internalWriterLogger) Print(v ...interface{}) {

	msg := fmt.Sprint(v...)
	fmt.Fprintln(logger.writer, msg)
}

func (logger *internalWriterLogger) Printf(format string, v ...interface{}) {

	msg := fmt.Sprintf(format, v...)
	fmt.Fprintln(logger.writer, msg)
}

func (logger *internalWriterLogger) Println(v ...interface{}) {

	msg := fmt.Sprintln(v...)
	fmt.Fprintln(logger.writer, msg)
}



func (logger *internalWriterLogger) Debug(v ...interface{}) {

	logger.Print(v...)
}

func (logger *internalWriterLogger) Debugf(format string, v ...interface{}) {

	logger.Printf(format, v...)
}

func (logger *internalWriterLogger) Debugln(v ...interface{}) {

	logger.Println(v...)
}



func (logger *internalWriterLogger) Error(v ...interface{}) {

	logger.Print(v...)
}

func (logger *internalWriterLogger) Errorf(format string, v ...interface{}) {

	logger.Printf(format, v...)
}

func (logger *internalWriterLogger) Errorln(v ...interface{}) {

	logger.Println(v...)
}



func (logger *internalWriterLogger) Fatal(v ...interface{}) {

	logger.Print(v...)
	os.Exit(1)
}

func (logger *internalWriterLogger) Fatalf(format string, v ...interface{}) {

	logger.Printf(format, v...)
	os.Exit(1)
}

func (logger *internalWriterLogger) Fatalln(v ...interface{}) {

	logger.Println(v...)
	os.Exit(1)
}



func (logger *internalWriterLogger) Panic(v ...interface{}) {

	logger.Print(v...)
	panic(fmt.Sprint(v...))
}

func (logger *internalWriterLogger) Panicf(format string, v ...interface{}) {

	logger.Printf(format, v...)
	panic(fmt.Sprintf(format, v...))
}

func (logger *internalWriterLogger) Panicln(v ...interface{}) {

	logger.Println(v...)
	panic(fmt.Sprintln(v...))
}



func (logger *internalWriterLogger) Warn(v ...interface{}) {

	logger.Print(v...)
}

func (logger *internalWriterLogger) Warnf(format string, v ...interface{}) {

	logger.Printf(format, v...)
}

func (logger *internalWriterLogger) Warnln(v ...interface{}) {

	logger.Println(v...)
}



func (logger *internalWriterLogger) Trace(v ...interface{}) {

	logger.Print(v...)
}

func (logger *internalWriterLogger) Tracef(format string, v ...interface{}) {

	logger.Printf(format, v...)
}

func (logger *internalWriterLogger) Traceln(v ...interface{}) {

	logger.Println(v...)
}

