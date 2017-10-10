package chromessr

import (
	"fmt"
	"regexp"
)

func Retrieve(url, placeholder string) (content string, err error) {
	fmt.Printf("asked to retrieve %s\n", url)
	page := retrieveFromCache(url)
	if !page.Exists() || page.Outdated() {
		enqueue <- page
	}

	if page.Exists() {
		content = loadContent(page, placeholder)
	} else {
		content = placeholder
	}

	return
}

func retrieveFromCache(url string) (page Page) {
	page = Page{Url: url}
	page.Load()
	return
}

func loadContent(page Page, base string) (newContent string) {
	bodyMatcher := regexp.MustCompile(`<div id="root"></div>`)
	newContent = bodyMatcher.ReplaceAllString(base, `<div id="root">`+page.Content+`</div>`)
	return
}
