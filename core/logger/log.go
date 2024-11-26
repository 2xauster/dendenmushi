package logger

import (
	"fmt"
	"log"
	"path/filepath"
	"runtime"
)

const (
	reset  = "\033[0m"
	red    = "\033[31m"
	green  = "\033[32m"
	yellow = "\033[33m"
)

/* Return a string wrapped with colors for terminal */
func wrap(color, text string) string {
	return color + text + reset
}

func Info(text string) {
	pc, file, line, ok := runtime.Caller(1)
	if !ok {
		log.Printf("%s: %s", wrap(green, "[unknown] | INFO"), text)
		return
	}

	funcName := runtime.FuncForPC(pc).Name()
	shortFile := filepath.Base(file) // Get the base name of the file
	prefix := fmt.Sprintf("[%s:%d | %s] | INFO", shortFile, line, funcName)
	log.Printf("%s: %s", wrap(green, prefix), text)
}

func Err(err error) {
	pc, file, line, ok := runtime.Caller(1)
	if !ok {
		log.Printf("%s: %s", wrap(red, "[unknown] | ERROR"), err)
		return
	}

	funcName := runtime.FuncForPC(pc).Name()
	shortFile := filepath.Base(file) 
	prefix := fmt.Sprintf("[%s:%d | %s] | ERROR", shortFile, line, funcName)
	log.Printf("%s: %s", wrap(red, prefix), err)
}

func Warn(text string) {
	pc, file, line, ok := runtime.Caller(1)
	if !ok {
		log.Printf("%s: %s", wrap(yellow, "[unknown] | INFO"), text)
		return
	}

	funcName := runtime.FuncForPC(pc).Name()
	shortFile := filepath.Base(file)
	prefix := fmt.Sprintf("[%s:%d | %s] | INFO", shortFile, line, funcName)
	log.Printf("%s: %s", wrap(yellow, prefix), text)
}