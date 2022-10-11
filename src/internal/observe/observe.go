package observe

import (
	"io"

	"github.com/schollz/progressbar/v3"
)

var writer io.Writer

// SeedWriter adds default writer to the observe package.
// Uses a noop writer until seeded.
func SeedWriter(w io.Writer) {
	writer = w
}

// ItemProgress tracks the display of an item by counting the bytes
// read through the provided readcloser, up until the byte count matches
// the totalBytes.
func ItemProgress(rc io.ReadCloser, iname string, totalBytes int64) io.ReadCloser {
	if writer == nil {
		return rc
	}

	opts := progressbar.NewOptions(
		int(totalBytes),
		progressbar.OptionSetWriter(writer),
		progressbar.OptionClearOnFinish(),
		progressbar.OptionSetDescription(" |  "+iname),
		progressbar.OptionShowDescriptionAtLineEnd(),
		progressbar.OptionSetRenderBlankState(true),
		progressbar.OptionShowCount(),
		progressbar.OptionSetPredictTime(false),
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionShowBytes(true),
		progressbar.OptionSetWidth(20),
		progressbar.OptionSetTheme(progressbar.Theme{
			Saucer:        "[green]=[reset]",
			SaucerHead:    "[green]>[reset]",
			SaucerPadding: " ",
			BarStart:      "[",
			BarEnd:        "]",
		}),
	)

	pbr := progressbar.NewReader(rc, opts)

	return &pbr
}
