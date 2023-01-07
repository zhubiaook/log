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
