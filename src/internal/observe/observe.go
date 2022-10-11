package observe

import (
	"io"
	"sync"

	"github.com/vbauerster/mpb/v8"
	"github.com/vbauerster/mpb/v8/decor"
)

var (
	wg       sync.WaitGroup
	writer   io.Writer
	progress *mpb.Progress
)

// SeedWriter adds default writer to the observe package.
// Uses a noop writer until seeded.
func SeedWriter(w io.Writer) {
	writer = w

	if w == nil {
		progress = nil
		return
	}

	progress = mpb.New(
		mpb.WithWidth(32),
		mpb.WithWaitGroup(&wg),
		mpb.WithOutput(writer),
	)
}

// Complete blocks until the progress finishes writing out all data.
// Afterwards, the progress instance is reset.
func Complete() {
	if progress != nil {
		progress.Wait()
	}

	SeedWriter(writer)
}

// ItemProgress tracks the display of an item by counting the bytes
// read through the provided readcloser, up until the byte count matches
// the totalBytes.
func ItemProgress(rc io.ReadCloser, iname string, totalBytes int64) (io.ReadCloser, func()) {
	if writer == nil || rc == nil || totalBytes == 0 {
		return rc, func() {}
	}

	wg.Add(1)

	bar := progress.AddBar(
		totalBytes,
		mpb.BarFillerOnComplete(""),
		mpb.BarRemoveOnComplete(),
		mpb.PrependDecorators(
			decor.OnComplete(decor.Percentage(decor.WCSyncSpace), ""),
			decor.OnComplete(decor.TotalKiloByte("% .1f", decor.WCSyncSpace), ""),
		),
		mpb.AppendDecorators(
			decor.OnComplete(decor.Name(iname), ""),
		),
	)

	return bar.ProxyReader(rc), waitAndCloseBar(iname, bar)
}

func waitAndCloseBar(n string, bar *mpb.Bar) func() {
	return func() {
		bar.Wait()
		wg.Done()
	}
}
