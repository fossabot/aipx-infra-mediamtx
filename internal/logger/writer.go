package logger

// Writer is an object that provides a log method.
type Writer interface {
	Log(Level, string, ...interface{})
}

type WriterWithConfSource interface {
	Log(Level, string, ...interface{})
}
