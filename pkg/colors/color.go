package colors

import "github.com/fatih/color"

func Sinfo(a ...interface{}) string {
	c := color.New(color.FgCyan).SprintFunc()
	return c(a...)
}
