package main

import (
	"context"
	"fmt"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/emulation"
	"log"
	"os"
	"sync"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/vincent-petithory/dataurl"
)

type Chrome struct {
	ctx    context.Context
	cancel context.CancelFunc
	once   sync.Once
}

func (c *Chrome) Start() {
	log.Printf("starting")
	opts := []chromedp.ExecAllocatorOption{
		chromedp.NoFirstRun,
		chromedp.NoDefaultBrowserCheck,
		chromedp.CombinedOutput(os.Stderr), // stdout and stderr from the browser
		chromedp.Headless,                  // comment out to see chrome in foreground

		chromedp.Flag("disable-background-networking", true),
		chromedp.Flag("enable-features", "NetworkService,NetworkServiceInProcess"),
		chromedp.Flag("disable-background-timer-throttling", true),
		chromedp.Flag("disable-backgrounding-occluded-windows", true),
		chromedp.Flag("disable-breakpad", true),
		chromedp.Flag("disable-client-side-phishing-detection", true),
		chromedp.Flag("disable-default-apps", true),
		chromedp.Flag("disable-dev-shm-usage", true),
		chromedp.Flag("disable-extensions", true),
		chromedp.Flag("disable-features", "site-per-process,Translate,BlinkGenPropertyTrees"),
		chromedp.Flag("disable-hang-monitor", true),
		chromedp.Flag("disable-ipc-flooding-protection", true),
		chromedp.Flag("disable-popup-blocking", true),
		chromedp.Flag("disable-prompt-on-repost", true),
		chromedp.Flag("disable-renderer-backgrounding", true),
		chromedp.Flag("disable-sync", true),
		chromedp.Flag("force-color-profile", "srgb"),
		chromedp.Flag("metrics-recording-only", true),
		chromedp.Flag("safebrowsing-disable-auto-update", true),
		chromedp.Flag("enable-automation", true),
		chromedp.Flag("password-store", "basic"),
		chromedp.Flag("use-mock-keychain", true),
	}

	c.ctx, c.cancel = chromedp.NewExecAllocator(context.Background(), opts...)
}

func (c *Chrome) Run(filename string, in []byte) ([]byte, error) {
	c.once.Do(c.Start)

	// set a parent timeout so we bound our total time in case something hangs
	tctx, cancel := context.WithTimeout(c.ctx, time.Minute)
	defer cancel()

	ctx, cancel := chromedp.NewContext(tctx)

	defer func() {
		// bound cleanup time
		// https://pkg.go.dev/github.com/chromedp/chromedp#Cancel
		tctx, tcancel := context.WithTimeout(ctx, 10*time.Second)
		defer tcancel()
		err := chromedp.Cancel(tctx)
		if err != nil {
			log.Printf("cleanup error %s", err)
		}
	}()

	imagedataURI := dataurl.New(in, "image/svg+xml").String()
	body := fmt.Sprintf(`<html><body><img src="%s" width="512"></body></html>`, imagedataURI)
	dataURI := dataurl.New([]byte(body), "text/html").String()
	log.Printf("loading %s...", dataURI[:50])
	var buf []byte
	if err := chromedp.Run(ctx,
		chromedp.Navigate(dataURI),
		chromedp.ActionFunc(func(ctx context.Context) error {
			err := emulation.SetDefaultBackgroundColorOverride().WithColor(&cdp.RGBA{0, 0, 0, 0}).Do(ctx)
			if err != nil {
				return err
			}
			return chromedp.Screenshot("img", &buf, chromedp.NodeVisible).Do(ctx)
		}),
	); err != nil {
		return nil, err
	}
	log.Printf("done")
	return buf, nil
}
