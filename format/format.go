package format

import (
	"bufio"
	"bytes"
	"io"
	"strconv"
	"strings"
)

const MaxSerialNumber = 999

var headers []string = []string{"##", "###", "####", "#####", "######", "#######"}

func AddSerialNumber(text string) string {
	levels := make([]int, MaxSerialNumber)
	newText := bytes.NewBuffer(make([]byte, 0, len(text)*2))
	reader := bufio.NewReader(strings.NewReader(text))
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		level := getLevel(string(line))
		var newLine string
		if level != -1 {
			levels[level] = levels[level] + 1
			newLine = addSerial(deleteOldSerial(string(line)), levels[:level+1])
		} else {
			newLine = string(line)
		}

		newLine += "\n"

		newText.WriteString(newLine)
	}

	return newText.String()
}

func deleteOldSerial(line string) string {
	strs := strings.Split(line, " ")
	if len(strs) <= 2 {
		return line
	}

	return strs[0] + " " +  strs[2]
}

func addSerial(line string, levels []int) string {
	prefix := levelsToString(levels)
	index := strings.Index(line, " ")
	str := line[:index] + " " + prefix + line[index+1:]
	return str
}

func getLevel(line string) int {
	for i := len(headers) - 1; i >= 0; i-- {
		b := strings.HasPrefix(line, headers[i])
		if b {
			return i
		}
	}

	return -1
}

func levelsToString(levels []int) string {
	if len(levels) == 0 {
		return ""
	}
	buf := bytes.NewBuffer(make([]byte, 0, len(levels)*2))
	length := len(levels)
	for i := 0; i < length; i++ {
		buf.WriteString(strconv.Itoa(levels[i]))
		if i < length-1 {
			buf.WriteString(".")
		} else {
			buf.WriteString(" ")
		}
	}
	return buf.String()
}
