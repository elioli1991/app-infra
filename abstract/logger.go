package abstract

type Level int8

type StdLogger interface {
	Log(level Level, vals string) error
}

type Logger interface {
	Debug(a ...any)
	Debugf(format string, a ...any)
	Debugw(keyvals ...any)
	Info(a ...any)
	Infof(format string, a ...any)
	Infow(keyvals ...any)
	Warn(a ...any)
	Warnf(format string, a ...any)
	Warnw(keyvals ...any)
	Error(a ...any)
	Errorf(format string, a ...any)
	Errorw(keyvals ...any)
	Fatal(a ...any)
	Fatalf(format string, a ...any)
	Fataw(keyvals ...any)
}
