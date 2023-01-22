package abstract

type Level int8

type StdLogger interface {
	Log(level Level, keyvals ...interface{}) error
}

type Logger interface {
	Info(a ...interface{})
	Infof(format string, a ...interface{})
	Infow(keyvals ...interface{})
	Error(a ...interface{})
	Errorf(format string, a ...interface{})
	Errorw(keyvals ...interface{})
	Fatal(a ...interface{})
	Fatalf(format string, a ...interface{})
	Fataw(keyvals ...interface{})
}
