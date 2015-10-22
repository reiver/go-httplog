package httplog


type adaptor struct {
	logger Logger
}


func newAdaptor(logger Logger) extendedLogger {
	extLogger := adaptor{
		logger:logger,
	}

	return &extLogger
}



func (lg *adaptor) Debug(v ...interface{}) {
	lg.logger.Print(v)
}

func (lg *adaptor) Debugf(format string, v ...interface{}) {
	lg.logger.Printf(format, v)
}

func (lg *adaptor) Debugln(v ...interface{}) {
	lg.logger.Println(v)
}



func (lg *adaptor) Error(v ...interface{}) {
	lg.logger.Print(v)
}

func (lg *adaptor) Errorf(format string, v ...interface{}) {
	lg.logger.Printf(format, v)
}

func (lg *adaptor) Errorln(v ...interface{}) {
	lg.logger.Println(v)
}



func (lg *adaptor) Fatal(v ...interface{}) {
	lg.logger.Fatal(v)
}

func (lg *adaptor) Fatalf(format string, v ...interface{}) {
	lg.logger.Fatalf(format, v)
}

func (lg *adaptor) Fatalln(v ...interface{}) {
	lg.logger.Fatalln(v)
}



func (lg *adaptor) Panic(v ...interface{}) {
	lg.logger.Panic(v)
}

func (lg *adaptor) Panicf(format string, v ...interface{}) {
	lg.logger.Panicf(format, v)
}
func (lg *adaptor) Panicln(v ...interface{}) {
	lg.logger.Panicln(v)
}



func (lg *adaptor) Print(v ...interface{}) {
	lg.logger.Print(v)
}

func (lg *adaptor) Printf(format string, v ...interface{}) {
	lg.logger.Printf(format, v)
}

func (lg *adaptor) Println(v ...interface{}) {
	lg.logger.Println(v)
}



func (lg *adaptor) Trace(v ...interface{}) {
	lg.logger.Print(v)
}

func (lg *adaptor) Tracef(format string, v ...interface{}) {
	lg.logger.Printf(format, v)
}

func (lg *adaptor) Traceln(v ...interface{}) {
	lg.logger.Println(v)
}



func (lg *adaptor) Warn(v ...interface{}) {
	lg.logger.Print(v)
}

func (lg *adaptor) Warnf(format string, v ...interface{}) {
	lg.logger.Printf(format, v)
}

func (lg *adaptor) Warnln(v ...interface{}) {
	lg.logger.Println(v)
}
