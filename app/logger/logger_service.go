package logger

import (
	"fmt"
	"time"

	"github.com/goravel/framework/support/color"
)

type Logger struct {
	name string
}

func NewLogger(name ...string) Logger {
	var n string
	if name[0] != "" {
		n = name[0]
	}
	formatStr := fmt.Sprintf(
		"[%v]\t[%v]\t",
		time.Now().Format(time.TimeOnly),
		n,
	)
	return Logger{name: formatStr}
}

func (l *Logger) Info(format string, args ...any) {
	println()
	color.Infof(l.name+format, args...)
	println()
}
