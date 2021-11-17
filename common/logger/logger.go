package logger

import "log"

type Logger struct {
}

func (l Logger) Info(format string, param ...interface{}) {
	setFlags()
	log.Printf(format, param)
}

func (l Logger) Error(format string, param ...interface{}) {
	setFlags()
	log.Panic(format, param)
}

func setFlags() {
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
}
