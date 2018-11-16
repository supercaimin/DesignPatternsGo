package bridge

import (
	"fmt"
)

type Formatter interface {
	Format(text string) string
}

type PainTextFormatter struct {
}

func (self *PainTextFormatter) Format(text string) string {
	return text
}

type HtmlFormatter struct {
}

func (self *HtmlFormatter) Format(text string) string {
	return fmt.Sprintf("<p>%s</p>", text)
}
