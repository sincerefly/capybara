package log

// wrapper

func Info(msg interface{}, keyvals ...interface{}) {
	GetDefaultLogger().Info(msg, keyvals...)
}

func Infof(format string, args ...interface{}) {
	GetDefaultLogger().Infof(format, args...)
}

func Warn(msg interface{}, keyvals ...interface{}) {
	GetDefaultLogger().Warn(msg, keyvals...)
}

func Warnf(format string, args ...interface{}) {
	GetDefaultLogger().Warnf(format, args...)
}

func Debug(msg interface{}, keyvals ...interface{}) {
	GetDefaultLogger().Debug(msg, keyvals...)
}

func Debugf(format string, args ...interface{}) {
	GetDefaultLogger().Debugf(format, args...)
}

func Error(msg interface{}, keyvals ...interface{}) {
	GetDefaultLogger().Error(msg, keyvals...)
}

func Errorf(format string, args ...interface{}) {
	GetDefaultLogger().Errorf(format, args...)
}

func Fatal(msg interface{}, keyvals ...interface{}) {
	GetDefaultLogger().Fatal(msg, keyvals...)
}

func Fatalf(format string, args ...interface{}) {
	GetDefaultLogger().Fatalf(format, args...)
}
