package observe

import (
	"context"
	"fmt"
	"io"
	"sync"

	"github.com/vbauerster/mpb/v8"
	"github.com/vbauerster/mpb/v8/decor"
)

var (
	wg       sync.WaitGroup
	con      context.Context
	writer   io.Writer
	progress *mpb.Progress
)

func init() {
	makeSpinFrames()
}

// SeedWriter adds default writer to the observe package.
// Uses a noop writer until seeded.
func SeedWriter(ctx context.Context, w io.Writer) {
	writer = w
	con = ctx

	if con == nil {
		con = context.Background()
	}

	progress = mpb.NewWithContext(
		con,
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

	SeedWriter(con, writer)
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
			decor.OnComplete(decor.NewPercentage("%d", decor.WC{W: 4}), ""),
			decor.OnComplete(decor.TotalKiloByte("%.1f", decor.WCSyncSpace), ""),
		),
		mpb.AppendDecorators(
			decor.OnComplete(decor.Name(iname), ""),
		),
	)

	return bar.ProxyReader(rc), waitAndCloseBar(bar)
}

var spinFrames []string

func makeSpinFrames() {
	s, l := rune('∙'), rune('●')

	line := []rune{}
	for i := 0; i < 32; i++ {
		line = append(line, s)
	}

	sl := make([]string, 0, 33)
	sl = append(sl, string(line))

	for i := 1; i < 32; i++ {
		l2 := make([]rune, len(line))
		copy(l2, line)
		l2[i] = l

		sl = append(sl, string(l2))
	}

	spinFrames = sl
}

// ItemProgress tracks the display a spinner that idles while the collection
// incrementing the count of items handled.  Each write to the provided channel
// counts as a single increment.  The caller is expected to close the channel.
func CollectionProgress(user, dirName string) (chan<- struct{}, func()) {
	if writer == nil || len(dirName) == 0 {
		return nil, func() {}
	}

	wg.Add(1)

	bar := progress.New(
		-1, // -1 to indicate an unbounded count
		mpb.SpinnerStyle(spinFrames...),
		mpb.BarFillerOnComplete(""),
		mpb.BarRemoveOnComplete(),
		// mpb.PrependDecorators(
		// ),
		mpb.AppendDecorators(
			decor.OnComplete(decor.CurrentNoUnit("%d - ", decor.WCSyncSpace), ""),
			decor.OnComplete(
				decor.Name(fmt.Sprintf("%s - %s", user, dirName)),
				""),
		),
	)

	ch := make(chan struct{})

	go func(ci <-chan struct{}) {
		for {
			select {
			case <-con.Done():
				bar.SetTotal(-1, true)
				return

			case _, ok := <-ci:
				if !ok {
					bar.SetTotal(-1, true)
					return
				}

				bar.Increment()
			}
		}
	}(ch)

	return ch, waitAndCloseBar(bar)
}

func waitAndCloseBar(bar *mpb.Bar) func() {
	return func() {
		bar.Wait()
		wg.Done()
	}
}
