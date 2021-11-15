package format

import "regexp"

// Config 所有参数都是可选
type Config struct {
	// 在哪些中添加序号
	Headers []string
	// 替换旧的序号
	ReplaceRegexps []string
	// 新的序号生成规则
	LevelsToString func(levels []int) string

	// 替换旧序号正则
	regexps []*regexp.Regexp
}

func (c *Config) init() error {
	// 哪些情况下添加序号
	if c.Headers == nil {
		c.Headers = []string{"##", "###", "####", "#####", "######", "#######"}
	}
	// 替换旧序号
	if len(c.ReplaceRegexps) == 0 {
		c.ReplaceRegexps = []string{`^[\d.]+`}
	}
	// 编译替换正则
	for _, v := range c.ReplaceRegexps {
		r, err := regexp.Compile(v)
		if err != nil {
			return err
		}
		c.regexps = append(c.regexps, r)
	}
	// 检查序列方法
	if c.LevelsToString == nil {
		// 设置默认
		c.LevelsToString = DefaultLevelsToString
	}

	return nil
}
