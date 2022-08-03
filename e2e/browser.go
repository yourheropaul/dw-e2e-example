package e2e

import (
	"context"
	"fmt"
	"io"
	"strings"

	"github.com/chromedp/cdproto/runtime"
	"github.com/chromedp/chromedp"
)

type Browser struct {
	ctx    context.Context
	cancel func()
	logger *ServiceLogger
}

func NewBrowser(logger *ServiceLogger) *Browser {
	ctx, cancel := chromedp.NewContext(context.Background())

	handleChromeDpLogs(ctx, logger)

	return &Browser{
		ctx:    ctx,
		cancel: cancel,
		logger: logger,
	}
}

func (b *Browser) Close() {
	b.cancel()
}

func (b *Browser) Visit(url string) error {
	return chromedp.Run(b.ctx, chromedp.Navigate(url))
}

func (b *Browser) Text(selector string) (string, error) {
	var res string
	err := chromedp.Run(b.ctx, chromedp.Text(selector, &res, chromedp.NodeVisible))
	return res, err
}

func (b *Browser) FullScreenshot() ([]byte, error) {
	var buf []byte

	if err := chromedp.Run(b.ctx, chromedp.FullScreenshot(&buf, 100)); err != nil {
		return []byte{}, err
	}

	return buf, nil
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
