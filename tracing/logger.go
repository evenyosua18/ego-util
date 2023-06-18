package tracing

import "log"

type DefaultLogger struct{}

func (l *DefaultLogger) Println(args ...interface{}) {
	log.Println(args...)
}

func (l *DefaultLogger) Printf(name string, obj ...interface{}) {
	log.Printf("%s: %v", name, obj)
}
