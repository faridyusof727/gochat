package gochat

import "fmt"

type FormatterImpl struct {
	content string
}

// Text implements Formatter.
func (f *FormatterImpl) Text(text string) Formatter {
	f.content += text
	return f
}

// NewLine implements Formatter.
func (f *FormatterImpl) NewLine() Formatter {
	f.content += "\n"
	return f
}

// ToString implements Formatter.
func (f *FormatterImpl) ToString() string {
	return f.content
}

// Bold implements Formatter.
func (f *FormatterImpl) Bold(text string) Formatter {
	f.content += fmt.Sprintf("*%s*", text)
	return f
}

// Code implements Formatter.
func (f *FormatterImpl) Code(text string) Formatter {
	f.content += fmt.Sprintf("```\n%s\n```", text)
	return f
}

// Italic implements Formatter.
func (f *FormatterImpl) Italic(text string) Formatter {
	f.content += fmt.Sprintf("_%s_", text)
	return f
}

// List implements Formatter.
func (f *FormatterImpl) List(list []string) Formatter {
	for _, item := range list {
		f.content += fmt.Sprintf("- %s\n", item)
	}
	return f
}

// Monospace implements Formatter.
func (f *FormatterImpl) Monospace(text string) Formatter {
	f.content += fmt.Sprintf("`%s`", text)
	return f
}

// StrikeThrough implements Formatter.
func (f *FormatterImpl) StrikeThrough(text string) Formatter {
	f.content += fmt.Sprintf("~%s~", text)
	return f
}

func NewFormatter() Formatter {
	return &FormatterImpl{}
}
