package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var stdinReader *bufio.Reader

func InitStdinReader() {
	stdinReader = bufio.NewReader(os.Stdin)
}

func ReadInput(msg string) string {
	if msg != "" {
		fmt.Printf("%s: ", msg)
	}

	value, _ := stdinReader.ReadString(NewLineByte)
	return strings.TrimSpace(value)
}
