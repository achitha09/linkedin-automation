package stealth

import (
	"time"

	"github.com/go-rod/rod"
)

func HumanHover(page *rod.Page, selector string) {
	el := page.MustElement(selector)
	el.MustHover()
	time.Sleep(600 * time.Millisecond)
}
