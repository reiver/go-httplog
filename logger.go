
package httplog


// Logger is the interface that represents what an 'logger' looks like.
//
// The Wrap func takes one of these as an argument.
//
// The Logger interface includes 9 methods that are exactly the same as Go's
// built-in "log" library. Namely: Fatal, Fatalf, Fatalln, Panic, Panicf, Panicln,
// Print, Printf and Println.
type Logger interface {
	Fatal(...interface{})
	Fatalf(string, ...interface{})
	Fatalln(...interface{})

	Panic(...interface{})
	Panicf(string, ...interface{})
	Panicln(...interface{})

	Print(...interface{})
	Printf(string, ...interface{})
	Println(...interface{})
}



type extendedLogger interface {
	Logger

	Debug(...interface{})
	Debugf(string, ...interface{})
	Debugln(...interface{})

	Error(...interface{})
	Errorf(string, ...interface{})
	Errorln(...interface{})

	Trace(...interface{})
	Tracef(string, ...interface{})
	Traceln(...interface{})

	Warn(...interface{})
	Warnf(string, ...interface{})
	Warnln(...interface{})
}
