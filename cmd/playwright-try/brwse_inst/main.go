package main

import (
	"fmt"
	"log"

	"github.com/mxschmitt/playwright-go"
)

func main() {
	pw := initPlaywright()
	defer clsPlaywright(pw)

	brwse, pg := browserPage(pw)
	defer clsBrowser(brwse)

	scrape(pg)
}

func initPlaywright() *playwright.Playwright {
	pw, err := playwright.Run(
		&playwright.RunOptions{
			DriverDirectory: "/home/siuyin/go",
			Browsers:        []string{"chromium"}},
	)
	if err != nil {
		log.Fatalf("could not start playwright: %v", err)
	}
	return pw
}

func browserPage(pw *playwright.Playwright) (playwright.Browser, playwright.Page) {
	opt := playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(true),
		Args:     []string{"--disable-gpu"},
	}
	browser, err := pw.Chromium.Launch(opt)
	if err != nil {
		log.Fatalf("could not launch browser: %v", err)
	}
	fmt.Printf("Browser version: %s %s\n", pw.Chromium.Name(), browser.Version())

	ctx, err := browser.NewContext()
	if err != nil {
		log.Fatalf("could not create context: %v", err)
	}

	page, err := ctx.NewPage()
	if err != nil {
		log.Fatalf("could not create page: %v", err)
	}
	if _, err = page.Goto("https://news.ycombinator.com"); err != nil {
		log.Fatalf("could not goto: %v", err)
	}
	return browser, page
}

func scrape(pg playwright.Page) {
	entries, err := pg.QuerySelectorAll(".athing")
	if err != nil {
		log.Fatalf("could not get entries: %v", err)
	}
	for i, entry := range entries {
		titleElement, err := entry.QuerySelector("td.title > a")
		if err != nil {
			log.Fatalf("could not get title element: %v", err)
		}
		title, err := titleElement.TextContent()
		if err != nil {
			log.Fatalf("could not get text content: %v", err)
		}
		fmt.Printf("%d: %s\n", i+1, title)
	}
}

func clsBrowser(browser playwright.Browser) {
	if err := browser.Close(); err != nil {
		log.Fatalf("could not close browser: %v", err)
	}
}

func clsPlaywright(pw *playwright.Playwright) {
	if err := pw.Stop(); err != nil {
		log.Fatalf("could not stop Playwright: %v", err)
	}
}
