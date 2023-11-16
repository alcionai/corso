package color

import (
	"io"

	"github.com/fatih/color"
)

var (
	Red     = color.FgRed
	Blue    = color.FgBlue
	Magenta = color.FgMagenta
	Cyan    = color.FgCyan
	Green   = color.FgGreen
	White   = color.FgWhite

	RedOutput     = color.New(Red).SprintFunc()
	BlueOutput    = color.New(Blue).SprintFunc()
	MagentaOutput = color.New(Magenta).SprintFunc()
	CyanOutput    = color.New(Cyan).SprintFunc()
	GreenOutput   = color.New(Green).SprintFunc()
	GreyOutput    = color.New(White).SprintFunc()
)

type colorableWriter struct {
	color  color.Attribute
	writer io.Writer
}

func NewColorableWriter(clr color.Attribute, writer io.Writer) io.Writer {
	return &colorableWriter{clr, writer}
}

func (cw *colorableWriter) Write(p []byte) (n int, err error) {
	return color.New(cw.color).Fprint(cw.writer, string(p))
}
