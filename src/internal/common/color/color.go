package color

import (
	"io"

	"github.com/fatih/color"
)

var (
	Red     = color.New(color.FgRed).SprintFunc()
	Blue    = color.New(color.FgBlue).SprintFunc()
	Magenta = color.New(color.FgMagenta).SprintFunc()
	Cyan    = color.New(color.FgCyan).SprintFunc()
	Green   = color.New(color.FgGreen).SprintFunc()
	Grey    = color.New(color.FgWhite).SprintFunc()
)

type colorableWriter struct {
	color  *color.Color
	writer io.Writer
}

func NewColorableWriter(color *color.Color, writer io.Writer) io.Writer {
	return &colorableWriter{color, writer}
}

func (cw *colorableWriter) Write(p []byte) (n int, err error) {
	return cw.color.Fprint(cw.writer, string(p))
}
