package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {

		lineOut := s.Text()

		lineOut = colorWord(lineOut, "trace", CyanU)
		lineOut = colorWord(lineOut, "debug", GreenU)
		lineOut = colorWord(lineOut, "info", YellowU)
		lineOut = colorWord(lineOut, "warn", MagentaU)
		lineOut = colorWord(lineOut, "error", RedU)
		lineOut = colorWord(lineOut, "fatal", RedU)
		lineOut = colorWord(lineOut, "exception", RedU)

		fmt.Println(lineOut)
	}

}

func colorWord(source, word string, color func(...interface{}) string) string {

	first := word

	rawchar := string(first[0])
	capchar := strings.ToUpper(rawchar)
	new := bytes.Replace([]byte(first), []byte(rawchar), []byte(capchar), 1)
	first = string(new)

	upper := strings.ToUpper(word)
	lower := strings.ToLower(word)

	source = strings.Replace(source, first, color(first), -1)
	source = strings.Replace(source, upper, color(upper), -1)
	source = strings.Replace(source, lower, color(lower), -1)
	return source
}

// normal colors
var Red = color.New(color.FgRed).SprintFunc()
var Green = color.New(color.FgGreen).SprintFunc()
var Yellow = color.New(color.FgYellow).SprintFunc()
var Blue = color.New(color.FgBlue).SprintFunc()
var Magenta = color.New(color.FgMagenta).SprintFunc()
var Cyan = color.New(color.FgCyan).SprintFunc()
var White = color.New(color.FgWhite).SprintFunc()
var Black = color.New(color.FgBlack).SprintFunc()

// normal with underline
var RedU = color.New(color.FgRed, color.Underline).SprintFunc()
var GreenU = color.New(color.FgGreen, color.Underline).SprintFunc()
var YellowU = color.New(color.FgYellow, color.Underline).SprintFunc()
var BlueU = color.New(color.FgBlue, color.Underline).SprintFunc()
var MagentaU = color.New(color.FgMagenta, color.Underline).SprintFunc()
var CyanU = color.New(color.FgCyan, color.Underline).SprintFunc()
var WhiteU = color.New(color.FgWhite, color.Underline).SprintFunc()
