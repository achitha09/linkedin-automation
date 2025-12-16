package stealth

import "time"

func IsBusinessHours() bool {
	now := time.Now()
	hour := now.Hour()

	// Business hours: 9 AM to 6 PM
	return hour >= 9 && hour <= 18
}
