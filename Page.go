package chromessr

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"regexp"
	"time"
)

type Page struct {
	Url       string
	Timestamp int64
	Content   string
	Found     bool
}

func (page *Page) Outdated() bool {
	return page.Timestamp+int64((time.Hour*24).Seconds()) < time.Now().Unix()
}

func (page *Page) Exists() bool {
	return page.Found
}

func (page *Page) Load() {
	content, err := ioutil.ReadFile(page.Path())
	if err != nil {
		return
	}

	json.Unmarshal(content, &page)
	page.Found = true

	return
}

func (page *Page) Save(content string) {
	page.Timestamp = time.Now().Unix()
	page.Content = content
	page.Found = true
	body, _ := json.Marshal(page)
	ioutil.WriteFile(page.Path(), body, 0600)
}

func (page *Page) Path() string {
	return os.Getenv("SSR_CACHE_PATH") + "/" + page.fsName()
}

func (page *Page) fsName() string {
	mehMatcher := regexp.MustCompile(`[^\w\d]`)
	return mehMatcher.ReplaceAllString(page.Url, "_")
}
