package stealth

import (
	"time"

	"github.com/go-rod/rod"
)

func HumanType(page *rod.Page, selector string, text string) {
	el := page.MustElement(selector)

	for _, char := range text {
		el.MustInput(string(char))
		time.Sleep(time.Duration(100+time.Now().UnixNano()%200) * time.Millisecond)
	}
}
