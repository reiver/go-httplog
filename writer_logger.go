package httplog


import (
	"fmt"
	"io"
	"os"
)


type internalWriterLogger struct {
	writer io.Writer
}


func newWriterLogger(writer io.Writer) Logger {

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

