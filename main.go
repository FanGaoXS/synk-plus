package main

import (
	"os"
	"os/signal"

	_ "github.com/fangaoxs/synk/env"
	"github.com/fangaoxs/synk/server/browser"
	"github.com/fangaoxs/synk/server/rest"
	_ "github.com/fangaoxs/synk/utils"
)

func main() {
	// channel which listen to signal of os
	signalCh := make(chan os.Signal, 1)
	defer close(signalCh)
	signal.Notify(signalCh, os.Interrupt)
	// channel which be used in browser runtime
	browserCh := make(chan struct{}, 1)
	defer close(browserCh)

	go rest.StartGin(FS)
	go browser.StartBrowser(browserCh)

	for {
		select {
		case <-signalCh:
			// when get interrupt signal of os, notify browser runtime to shut down browser
			browserCh <- struct{}{}
		case <-browserCh:
			// when get signal of that the browser has been exited, then the App exits as well
			os.Exit(0)
		}
	}
}
