package abstract

type Level int8

type StdLogger interface {
	Log(level Level, vals string) error
}

type Logger interface {
	Debug(a ...interface{})
	Debugf(format string, a ...interface{})
	Debugw(keyvals ...interface{})
	Info(a ...interface{})
	Infof(format string, a ...interface{})
	Infow(keyvals ...interface{})
	Warn(a ...interface{})
	Warnf(format string, a ...interface{})
	Warnw(keyvals ...interface{})
	Error(a ...interface{})
	Errorf(format string, a ...interface{})
	Errorw(keyvals ...interface{})
	Fatal(a ...interface{})
	Fatalf(format string, a ...interface{})
	Fataw(keyvals ...interface{})
}
