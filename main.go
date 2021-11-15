package main

import (
	"serial-number/format"

	"github.com/atotto/clipboard"
)

func main() {
	text, _ := clipboard.ReadAll()
	newText := format.AddSerialNumber(text)
	clipboard.WriteAll(newText)
}
