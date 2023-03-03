package lg

import (
	"encoding/json"
	"fmt"

	"github.com/fatih/color"
)

// const (
// 	InfoColor    = "\033[1;34m%s\033[0m"
// 	NoticeColor  = "\033[1;36m%s\033[0m"
// 	WarningColor = "\033[1;33m%s\033[0m"
// 	ErrorColor   = "\033[1;31m%s\033[0m"
// 	DebugColor   = "\033[0;36m%s\033[0m"

// 	Red     = "\033[1;31m%s\033[0m"
// 	Greens  = "\033[1;32m%s\033[0m"
// 	Black   = "\033[1;30m%s\033[0m"
// 	Yellow  = "\033[1;33m%s\033[0m"
// 	Purple  = "\033[1;34m%s\033[0m"
// 	Magenta = "\033[1;35m%s\033[0m"
// 	Teal    = "\033[1;36m%s\033[0m"
// 	White   = "\033[1;37m%s\033[0m"
// )

func Debug(y ...interface{}) string {
	//blue
	s := ""
	for _, x := range y {
		s = s + " " + fmt.Sprintf("%v", x)
	}

	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", 34, s)

}

func Info(y ...interface{}) string {
	//red
	s := ""
	for _, x := range y {
		s = s + " " + fmt.Sprintf("%v", x)
	}

	return fmt.Sprintf("\033[%d1;31m%s\033[0m", 34, s)

}

func Green(y ...interface{}) string {
	//green
	s := ""
	for _, x := range y {
		s = s + " " + fmt.Sprintf("%v", x)
	}

	return fmt.Sprintf("\033[%d1;32m%s\033[0m", 34, s)

}

func Yellow(y ...interface{}) string {
	//yellow
	s := ""
	for _, x := range y {
		s = s + " " + fmt.Sprintf("%v", x)
	}

	return fmt.Sprintf("\033[%d1;33m%s\033[0m", 34, s)

}
func Mg(y ...interface{}) string {
	//magneta
	s := ""
	for _, x := range y {
		s = s + " " + fmt.Sprintf("%v", x)
	}
	bold := color.New(color.Bold).SprintFunc()
	return bold(fmt.Sprintf("\033[%d1;35m%s\033[0m", 34, s))

}
func Dl(y ...interface{}) string {
	//magneta
	s := ""
	for _, x := range y {
		s = s + " " + fmt.Sprintf("%v", x)
	}
	bold := color.New(color.Bold).SprintFunc()
	return bold(fmt.Sprintf("\033[%d1;35m%s\033[0m", 34, s))

}

func Error(y ...interface{}) string {
	s := ""
	for _, x := range y {
		s = s + " " + fmt.Sprintf("%v", x)
	}

	return fmt.Sprintf("\033[%d1;31m%s\033[0m", 34, s)

}
func PrettyJsonString(i interface{}) string {
	s, _ := json.MarshalIndent(i, " ", "\t")
	return string(s)

}
func JsonString(i interface{}) string {

	b, err := json.Marshal(&i)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return ""
	}

	return Error(string(b))

}
