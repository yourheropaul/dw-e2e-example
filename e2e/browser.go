package e2e

import (
	"context"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/chromedp/cdproto/runtime"
	"github.com/chromedp/chromedp"
)

type browser struct {
	ctx    context.Context
	cancel func()
	logger *ServiceLogger
}

func NewBrowser(logger *ServiceLogger) *browser {
	ctx, cancel := chromedp.NewContext(context.Background())

	handleChromeDpLogs(ctx, logger)

	return &browser{
		ctx:    ctx,
		cancel: cancel,
		logger: logger,
	}
}

func (b *browser) Close() {
	b.cancel()
}

func (b *browser) Visit(url string) error {
	return chromedp.Run(b.ctx, chromedp.Navigate(url))
}

func (b *browser) Text(selector string) (string, error) {
	var res string
	err := chromedp.Run(
		b.ctx,
		runWithTimeOut(
			&b.ctx,
			1*time.Second,
			chromedp.Tasks{
				chromedp.Text(selector, &res, chromedp.NodeVisible),
			},
		),
	)
	return res, err
}

func (b *browser) FullScreenshot() ([]byte, error) {
	var buf []byte

	if err := chromedp.Run(b.ctx, chromedp.FullScreenshot(&buf, 100)); err != nil {
		return []byte{}, err
	}

	return buf, nil
}

func runWithTimeOut(ctx *context.Context, timeout time.Duration, tasks chromedp.Tasks) chromedp.ActionFunc {
	return func(ctx context.Context) error {
		timeoutContext, cancel := context.WithTimeout(ctx, timeout)
		defer cancel()
		return tasks.Do(timeoutContext)
	}
}

func handleChromeDpLogs(ctx context.Context, logger io.Writer) {
	chromedp.ListenTarget(ctx, func(ev interface{}) {
		switch ev := ev.(type) {
		case *runtime.EventConsoleAPICalled:
			io.WriteString(logger, fmt.Sprintf("* console.%s call:\n", ev.Type))
			for _, arg := range ev.Args {
				io.WriteString(logger, fmt.Sprintf("%s - %s\n", arg.Type, arg.Value))
			}

		case *runtime.EventExceptionThrown:
			s := ev.ExceptionDetails.Error()
			// V8 has changed the error messages for property access on null/undefined in version 9.3.310.
			// see: https://chromium.googlesource.com/v8/v8/+/c0fd89c3c089e888c4f4e8582e56db7066fa779b
			//      https://github.com/chromium/chromium/commit/1735cbf94c98c70ff7554a1e9e01bb9a4f91beb6
			// The message is normalized to make it compatible with the versions before this change.
			s = strings.ReplaceAll(s, "Cannot read property 'throwsException' of null", "Cannot read properties of null (reading 'throwsException')")
			io.WriteString(logger, fmt.Sprintf("* %s\n", s))
		}
	})
}
