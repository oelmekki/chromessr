package chromessr

import (
	"time"
)

var queue []Page

func processQueue() {
	currentlyProcessing := make([]Page, 0)

	for {
		select {
		case page := <-enqueue:
			found := false
			for _, processingPage := range currentlyProcessing {
				if page.Url != processingPage.Url {
					found = true
				}
			}

			if !found {
				queue = append(queue, page)
			}

		case <-time.After(time.Second * 1):
			if len(queue) > 0 {
				page := queue[0]
				queue = queue[1:len(queue)]
				currentlyProcessing = append(currentlyProcessing, page)
				content, capturingErr := Capture(page.Url)

				pages := make([]Page, 0)
				for _, processingPage := range currentlyProcessing {
					if page.Url != processingPage.Url {
						pages = append(pages, page)
					}
				}

				currentlyProcessing = pages

				if capturingErr == nil {
					page.Save(content)
				}
			}
		}
	}
}
