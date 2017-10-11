package chromessr

import (
	"fmt"
	"time"
)

func Capture(url string) (content string, err error) {
	fmt.Printf("Capturing url : %s\n", url)
	tabs, err := remote.TabList("")
	if err != nil {
		return
	}

	remote.NewTab(url)
	err = remote.ActivateTab(tabs[0])
	if err != nil {
		return
	}

	defer remote.CloseTab(tabs[0])

	time.Sleep(5 * time.Second)
	res, err := remote.EvaluateWrap(`
			let root = document.querySelector("#root");
			if (root) {
				return root.innerHTML;
			} else {
				return "";
			}
	`)
	if err != nil {
		return
	}

	content = res.(string)

	return
}
