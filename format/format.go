package format

import (
	"bufio"
	"bytes"
	"io"
	"strconv"
	"strings"
)

const MaxSerialNumber = 999

func AddSerialNumber(text string, config Config) string {
	// 初始化配置
	if err := config.init(); err != nil {
		panic(err)
	}

	levels := make([]int, MaxSerialNumber)
	empty := make([]int, MaxSerialNumber)
	newText := bytes.NewBuffer(make([]byte, 0, len(text)*2))
	reader := bufio.NewReader(strings.NewReader(text))
	skip := false
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		// 防止格式化代码注释
		b := strings.HasPrefix(strings.TrimRight(string(line), ""), "```")
		if b {
			skip = !skip
		}

		level := getLevel(string(line), config)
		var newLine string
		if !skip && level != -1 {
			levels[level] = levels[level] + 1
			copy(levels[level+1:], empty)
			// 删除旧序号
			newLine = deleteOldSerial(string(line), config)
			// 添加新序号
			newLine = addSerial(newLine, levels[:level+1], config)
		} else {
			newLine = string(line)
		}

		newLine += "\n"

		newText.WriteString(newLine)
	}

	return newText.String()
}

func deleteOldSerial(line string, config Config) string {
	line = strings.TrimRight(line, " ")
	index := strings.Index(line, " ")
	if index == -1 {
		return line
	}
	beforeStr := line[0:index]
	afterStr := line[index+1:]
	for _, r := range config.regexps {
		afterStr = r.ReplaceAllString(afterStr, "")
		afterStr = strings.TrimLeft(afterStr, " ")
	}

	return beforeStr + " " + afterStr
}

func addSerial(line string, levels []int, config Config) string {
	prefix := config.LevelsToString(levels)
	index := strings.Index(line, " ")
	str := line[:index] + " " + prefix + line[index+1:]
	return str
}

func getLevel(line string, config Config) int {
	for i := len(config.Headers) - 1; i >= 0; i-- {
		b := strings.HasPrefix(line, config.Headers[i]+" ")
		if b {
			return i
		}
	}

	return -1
}

func DefaultLevelsToString(levels []int) string {
	if len(levels) == 0 {
		return ""
	}
	buf := bytes.NewBuffer(make([]byte, 0, len(levels)*2))
	length := len(levels)
	for i := 0; i < length; i++ {
		buf.WriteString(strconv.Itoa(levels[i]))
		if i < length {
			buf.WriteString(".")
		}
		if i == length-1 {
			buf.WriteString(" ")
		}
	}
	return buf.String()
}
