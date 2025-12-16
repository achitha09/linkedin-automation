package main

import (
	"fmt"
	"time"

	"linkedin-automation/stealth"

	"linkedin-automation/storage"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

func main() {
	if !stealth.IsBusinessHours() {
		fmt.Println("Outside business hours. Exiting to simulate human behavior.")
		return
	}
	rateLimiter := stealth.NewRateLimiter(20) // max 20 actions per day

	if !rateLimiter.Allow() {
		fmt.Println("Daily limit reached. Stopping automation.")
		return
	}

	fmt.Println("Program started successfully")

	url := launcher.New().
		Leakless(false).
		Headless(false).
		Set("binary", "C:\\Program Files\\Google\\Chrome\\Application\\chrome.exe").
		Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 Chrome/120.0.0.0 Safari/537.36").
		MustLaunch()

	browser := rod.New().ControlURL(url).MustConnect()

	page := browser.MustPage("https://www.linkedin.com/login")
	stealth.RandomDelay(2000, 5000)
	page.MustEval(`() => window.scrollBy(0, 300)`)
	stealth.RandomDelay(1000, 2000)
	page.MustEval(`() => window.scrollBy(0, -150)`)
	// Hover like a human
	stealth.HumanHover(page, "input[name='session_key']")

	// Small pause
	stealth.RandomDelay(500, 1200)

	// Click input box
	emailInput := page.MustElement("input[name='session_key']")
	emailInput.MustClick()

	// Pause again
	stealth.RandomDelay(800, 1500)

	// Type like human
	stealth.HumanType(page, "input[name='session_key']", "test@example.com")
	storage.SaveState(storage.State{SentMessages: 1})

	fmt.Println("Chrome opened LinkedIn login page")

	page.MustWaitLoad()

	fmt.Println("Waiting 10 seconds...")
	time.Sleep(10 * time.Second)
}
