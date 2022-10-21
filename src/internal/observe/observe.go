package observe

import (
	"context"
	"fmt"
	"io"
	"sync"

	"github.com/vbauerster/mpb/v8"
	"github.com/vbauerster/mpb/v8/decor"
)

const progressBarWidth = 32

var (
	wg       sync.WaitGroup
	con      context.Context
	writer   io.Writer
	progress *mpb.Progress
)

func init() {
	makeSpinFrames(progressBarWidth)
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
		mpb.WithWidth(progressBarWidth),
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

const (
	ItemBackupMsg  = "Backing up item:"
	ItemRestoreMsg = "Restoring item:"
	ItemQueueMsg   = "Queuing items:"
	// Use the longest message
	dynamicBarMsgLength = len(ItemBackupMsg)
)

// ---------------------------------------------------------------------------
// Progress for Known Quantities
// ---------------------------------------------------------------------------

// ItemProgress tracks the display of an item in a folder by counting the bytes
// read through the provided readcloser, up until the byte count matches
// the totalBytes.
func ItemProgress(rc io.ReadCloser, header, iname string, totalBytes int64) (io.ReadCloser, func()) {
	if writer == nil || rc == nil || totalBytes == 0 {
		return rc, func() {}
	}

	wg.Add(1)

	bar := progress.New(
		totalBytes,
		mpb.NopStyle(),
		mpb.BarRemoveOnComplete(),
		mpb.PrependDecorators(
			decor.Name(header, decor.WCSyncSpaceR),
			decor.Name(iname, decor.WCSyncSpaceR),
			decor.CountersKibiByte(" %.1f/%.1f ", decor.WC{W: 8}),
			decor.NewPercentage("%d ", decor.WC{W: 4}),
		),
	)

	return bar.ProxyReader(rc), waitAndCloseBar(bar)
}

// Progress is used to display progress with a message
func Progress(message string) {
	if writer == nil {
		return
	}

	wg.Add(1)

	bar := progress.New(
		-1,
		mpb.NopStyle(),
		mpb.PrependDecorators(
			decor.Name(message, decor.WC{W: len(message) + 1, C: decor.DidentRight}),
		),
	)

	// Complete the bar immediately
	bar.SetTotal(-1, true)

	waitAndCloseBar(bar)()
}

// ProgressWithCompletion is used to display progress with a spinner
// that switches to "done" when the completion channel is signalled
func ProgressWithCompletion(message string) (chan<- struct{}, func()) {
	completionCh := make(chan struct{}, 1)

	if writer == nil {
		return completionCh, func() {}
	}

	wg.Add(1)
	frames := []string{"∙∙∙", "●∙∙", "∙●∙", "∙∙●", "∙∙∙"}
	bar := progress.New(
		-1,
		mpb.SpinnerStyle(frames...).PositionLeft(),
		mpb.PrependDecorators(
			decor.Name(message),
		),
		mpb.BarFillerOnComplete("done"),
	)

	go func(ci <-chan struct{}) {
		for {
			select {
			case <-con.Done():
				bar.SetTotal(-1, true)
			case <-ci:
				// We don't care whether the channel was signalled or closed
				// Use either one as an indication that the bar is done
				bar.SetTotal(-1, true)
			}
		}
	}(completionCh)

	return completionCh, waitAndCloseBar(bar)
}

// ProgressWithCount tracks the display of a bar that tracks the completion
// of the specified count.
// Each write to the provided channel counts as a single increment.
// The caller is expected to close the channel.
func ProgressWithCount(header, message string, count int64) (chan<- struct{}, func()) {
	progressCh := make(chan struct{})

	if writer == nil {

		go func(ci <-chan struct{}) {
			for {
				_, ok := <-ci
				if !ok {
					return
				}
			}
		}(progressCh)

		return progressCh, func() {}
	}

	wg.Add(1)

	bar := progress.New(
		count,
		mpb.NopStyle(),
		mpb.BarRemoveOnComplete(),
		mpb.PrependDecorators(
			decor.Name(header, decor.WCSyncSpaceR),
			decor.Counters(0, " %d/%d "),
			decor.Name(message),
		),
	)

	ch := make(chan struct{})

	go func(ci <-chan struct{}) {
		for {
			select {
			case <-con.Done():
				bar.Abort(true)
				return

			case _, ok := <-ci:
				if !ok {
					bar.Abort(true)
					return
				}

				bar.Increment()
			}
		}
	}(ch)

	return ch, waitAndCloseBar(bar)
}

// ---------------------------------------------------------------------------
// Progress for Unknown Quantities
// ---------------------------------------------------------------------------

var spinFrames []string

// The bar width is set to a static 32 characters.  The default spinner is only
// one char wide, which puts a lot of white space between it and the useful text.
// This builds a custom spinner animation to fill up that whitespace for a cleaner
// display.
func makeSpinFrames(barWidth int) {
	s, l := rune('∙'), rune('●')

	line := []rune{}
	for i := 0; i < barWidth; i++ {
		line = append(line, s)
	}

	sl := make([]string, 0, barWidth+1)
	sl = append(sl, string(line))

	for i := 1; i < barWidth; i++ {
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
func CollectionProgress(user, category, dirName string) (chan<- struct{}, func()) {
	if writer == nil || len(user) == 0 || len(dirName) == 0 {
		ch := make(chan struct{})

		go func(ci <-chan struct{}) {
			for {
				_, ok := <-ci
				if !ok {
					return
				}
			}
		}(ch)

		return ch, func() {}
	}

	wg.Add(1)

	bar := progress.New(
		-1, // -1 to indicate an unbounded count
		mpb.SpinnerStyle(spinFrames...),
		mpb.BarFillerOnComplete(""),
		mpb.BarRemoveOnComplete(),
		mpb.PrependDecorators(
			decor.OnComplete(decor.Name(category), ""),
		),
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
