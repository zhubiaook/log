package log

import "github.com/fatih/color"

func Info(msg string) {
	color.Cyan(msg)
}

func Warning(msg string) {
	color.Yellow(msg)
}

func Error(msg string) {
	color.Red(msg)
}

func Infof(format string, a ...interface{}) {
	c := color.New(color.FgCyan).PrintfFunc()
	c(format, a...)
}

func Warningf(format string, a ...interface{}) {
	c := color.New(color.FgYellow).PrintfFunc()
	c(format, a...)
}
